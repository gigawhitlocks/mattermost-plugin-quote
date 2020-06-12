package main

import (
	"fmt"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
)

// Plugin implements the interface expected by the Mattermost server to communicate between the server and plugin processes.
type Plugin struct {
	plugin.MattermostPlugin

	// configurationLock synchronizes access to the configuration.
	configurationLock sync.RWMutex

	// configuration is the active plugin configuration. Consult getConfiguration and
	// setConfiguration for usage.
	configuration *configuration
}

func (p *Plugin) MessageWillBePosted(c *plugin.Context, post *model.Post) (*model.Post, string) {
	siteURL := p.API.GetConfig().ServiceSettings.SiteURL
	channel, apiErr := p.API.GetChannel(post.ChannelId)
	if apiErr != nil {
		return post, apiErr.Message
	}

	if channel.Type != model.CHANNEL_OPEN {
		return post, ""
	}

	team, apiErr := p.API.GetTeam(channel.TeamId)
	if apiErr != nil {
		return post, apiErr.Message
	}

	selfLink := fmt.Sprintf("%s/%s", *siteURL, team.Name)
	selfLinkPattern, er := regexp.Compile(fmt.Sprintf("%s%s", selfLink, `/[\w/]+`))
	if er != nil {
		return post, er.Error()
	}

	for _, match := range selfLinkPattern.FindAllString(post.Message, -1) {
		markdownLink, err := regexp.Compile(fmt.Sprintf("\\]\\(%s\\)", match))
		if err != nil {
			continue
		}

		// skip permalinks inside markdown links
		if markdownLink.FindStringIndex(post.Message) != nil {
			continue
		}

		separated := strings.Split(match, "/")
		postId := separated[len(separated)-1]
		oldPost, apiErr := p.API.GetPost(postId)
		if apiErr != nil {
			return post, apiErr.Message
		}

		postUser, apiErr := p.API.GetUser(oldPost.UserId)
		if apiErr != nil {
			return post, apiErr.Message
		}

		postedAt := time.Unix(oldPost.CreateAt/1000, 0)
		quote := fmt.Sprintf("##### @%s at [%s](%s) in ~%s said:\n",
			postUser.Username,
			postedAt.Format("Mon Jan 2 08:04PM -0700 MST 2006"),
			match,
			channel.Name,
		)

		messageLines := strings.Split(oldPost.Message, "\n")
		for _, line := range messageLines {
			quote = fmt.Sprintf("%s\n> %s\n", quote, line)
		}
		post.Message = strings.Replace(post.Message, match, quote, 1)
	}

	return post, ""
}
