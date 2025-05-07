package bluesky

import (
	"context"
	"fmt"
	"strings"

	"github.com/bluesky-social/indigo/api/bsky"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableBlueskyUserMention(ctx context.Context) *plugin.Table {

	return &plugin.Table{
		Name:        "bluesky_user_mention",
		Description: "List of posts that mention the specified user. To query by handle, use a join with the bluesky_user table.",
		List: &plugin.ListConfig{
			Hydrate: listUserMentions,
			KeyColumns: plugin.KeyColumnSlice{
				{
					Name:    "target_did",
					Require: plugin.Required,
				},
			},
		},
		DefaultTransform: transform.FromGo().NullIfZero(),
		Columns:          postColumns("target_did"),
	}
}

func listUserMentions(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	targetDid := d.EqualsQuals["target_did"].GetStringValue()
	if targetDid == "" {
		logger.Error("listUserMentions: No target_did specified")
		return nil, fmt.Errorf("target_did must be specified")
	}

	if !strings.HasPrefix(targetDid, "did:") {
		logger.Error("listUserMentions: Invalid DID format", "did", targetDid)
		return nil, fmt.Errorf("invalid DID format: %s", targetDid)
	}

	client, err := connect(ctx, d)
	if err != nil {
		logger.Error("listUserMentions: Failed to connect", "error", err)
		return nil, fmt.Errorf("failed to connect: %w", err)
	}

	// Get the user's handle
	profile, err := bsky.ActorGetProfile(ctx, client, targetDid)
	if err != nil {
		logger.Error("listUserMentions: Error getting profile", "error", err)
		return nil, fmt.Errorf("failed to get profile: %v", err)
	}

	searchQuery := fmt.Sprintf("@%s", profile.Handle)

	searchResults, err := bsky.FeedSearchPosts(ctx, client, "", "", "", "", 100, "", searchQuery, "", "", nil, "", "")
	if err != nil {
		logger.Error("listUserMentions: Failed to search posts", "error", err)
		return nil, fmt.Errorf("failed to search posts: %w", err)
	}

	for _, post := range searchResults.Posts {
		feedPost := post.Record.Val.(*bsky.FeedPost)
		metadata := extractPostMetadata(feedPost)
		mentionedDIDs := metadata["mentioned_handles"].([]string)
		mentionedHandles := resolveDIDsToHandles(ctx, client, mentionedDIDs)

		d.StreamListItem(ctx, map[string]interface{}{
			"uri":                     post.Uri,
			"http_url":                convertToHttpUrl(post.Uri),
			"cid":                     post.Cid,
			"text":                    feedPost.Text,
			"author":                  post.Author.Handle,
			"created_at":              feedPost.CreatedAt,
			"indexed_at":              post.IndexedAt,
			"like_count":              post.LikeCount,
			"repost_count":            post.RepostCount,
			"reply_root":              getReplyRoot(feedPost),
			"reply_parent":            getReplyParent(feedPost),
			"has_external_links":      metadata["has_external_links"],
			"image_count":             metadata["image_count"],
			"hashtags":                metadata["hashtags"],
			"mentioned_handles":       metadata["mentioned_handles"],
			"mentioned_handles_names": mentionedHandles,
			"external_links":          metadata["external_links"],
			"target_did":              targetDid,
		})
	}

	// Handle pagination
	cursor := searchResults.Cursor
	for cursor != nil {
		nextResults, err := bsky.FeedSearchPosts(ctx, client, "", *cursor, "", "", 100, "", searchQuery, "", "", nil, "", "")
		if err != nil {
			logger.Error("listUserMentions: Failed to fetch next page", "error", err)
			return nil, fmt.Errorf("failed to fetch next page: %w", err)
		}

		for _, post := range nextResults.Posts {
			feedPost := post.Record.Val.(*bsky.FeedPost)
			metadata := extractPostMetadata(feedPost)
			mentionedDIDs := metadata["mentioned_handles"].([]string)
			mentionedHandles := resolveDIDsToHandles(ctx, client, mentionedDIDs)

			d.StreamListItem(ctx, map[string]interface{}{
				"uri":                     post.Uri,
				"http_url":                convertToHttpUrl(post.Uri),
				"cid":                     post.Cid,
				"text":                    feedPost.Text,
				"author":                  post.Author.Handle,
				"created_at":              feedPost.CreatedAt,
				"indexed_at":              post.IndexedAt,
				"like_count":              post.LikeCount,
				"repost_count":            post.RepostCount,
				"reply_root":              getReplyRoot(feedPost),
				"reply_parent":            getReplyParent(feedPost),
				"has_external_links":      metadata["has_external_links"],
				"image_count":             metadata["image_count"],
				"hashtags":                metadata["hashtags"],
				"mentioned_handles":       metadata["mentioned_handles"],
				"mentioned_handles_names": mentionedHandles,
				"external_links":          metadata["external_links"],
				"target_did":              targetDid,
			})
		}

		cursor = nextResults.Cursor
	}

	return nil, nil
}
