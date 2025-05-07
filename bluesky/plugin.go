/*
Package bluesky implements a Steampipe plugin for Bluesky.
This plugin provides data that Steampipe uses to present foreign
tables that represent Bluesky posts, users, and other resources.
*/

package bluesky

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {

	p := &plugin.Plugin{
		Name: "steampipe-plugin-bluesky",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
		},
		DefaultTransform: transform.FromGo().NullIfZero(),
		TableMap: map[string]*plugin.Table{
			"bluesky_post":           tableBlueskyPost(ctx),
			"bluesky_search_recent":  tableBlueskySearchRecent(ctx),
			"bluesky_user":           tableBlueskyUser(ctx),
			"bluesky_user_follower":  tableBlueskyUserFollower(ctx),
			"bluesky_user_following": tableBlueskyUserFollowing(ctx),
			"bluesky_user_mention":   tableBlueskyUserMention(ctx),
			"bluesky_user_post":      tableBlueskyUserPost(ctx),
		},
	}
	return p
}
