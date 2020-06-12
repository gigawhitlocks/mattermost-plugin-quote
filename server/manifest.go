// This file is automatically generated. Do not modify it manually.

package main

import (
	"strings"

	"github.com/mattermost/mattermost-server/v5/model"
)

var manifest *model.Manifest

const manifestStr = `
{
  "id": "business.silly.quote",
  "name": "Permalink Quote Expander",
  "description": "This plugin turns links to scrollback into pretty quotes",
  "version": "0.3.3",
  "min_server_version": "5.18.0",
  "server": {
    "executables": {
      "linux-amd64": "server/dist/plugin-linux-amd64"
    },
    "executable": ""
  },
  "webapp": {
    "bundle_path": "webapp/dist/main.js"
  },
  "settings_schema": {
    "header": "",
    "footer": "",
    "settings": []
  }
}
`

func init() {
	manifest = model.ManifestFromJson(strings.NewReader(manifestStr))
}
