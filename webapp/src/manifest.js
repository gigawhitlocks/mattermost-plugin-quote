// This file is automatically generated. Do not modify it manually.

const manifest = JSON.parse(`
{
    "id": "business.silly.quote",
    "name": "Quote (like Rocket Chat) Plugin",
    "description": "This plugin turns links to scrollback into pretty quotes",
    "version": "0.2.0",
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
`);

export default manifest;
export const id = manifest.id;
export const version = manifest.version;
