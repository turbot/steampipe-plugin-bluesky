package main

import (
	"turbot/steampipe-plugin-bluesky/bluesky"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: bluesky.Plugin})
}
