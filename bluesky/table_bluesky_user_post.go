package bluesky

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/bluesky-social/indigo/api/atproto"
	"github.com/bluesky-social/indigo/api/bsky"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableBlueskyUserPost(ctx context.Context) *plugin.Table {

	return &plugin.Table{
		Name:        "bluesky_user_post",
		Description: "List of posts published by a specific Bluesky user.",
		List: &plugin.ListConfig{
			Hydrate: listUserPosts,
			KeyColumns: plugin.KeyColumnSlice{
				{
					Name:    "target_did",
					Require: plugin.Optional,
				},
				{
					Name:    "handle",
					Require: plugin.Optional,
				},
			},
		},
		DefaultTransform: transform.FromGo().NullIfZero(),
		Columns:          postColumns("target_did", "handle"),
	}
}

func listUserPosts(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	var targetDid string
	var handle string

	// Check if DID is provided
	if d.EqualsQuals["target_did"] != nil {
		targetDid = d.EqualsQuals["target_did"].GetStringValue()
		if targetDid != "" {
			if !strings.HasPrefix(targetDid, "did:") {
				logger.Error("listUserPosts: Invalid DID format", "did", targetDid)
				return nil, fmt.Errorf("invalid DID format: %s", targetDid)
			}
		}
	}

	// Check if handle is provided
	if d.EqualsQuals["handle"] != nil {
		handle = d.EqualsQuals["handle"].GetStringValue()
		if handle != "" {
			// Remove @ prefix if present
			handle = strings.TrimPrefix(handle, "@")
		}
	}

	// If neither DID nor handle is provided, return error
	if targetDid == "" && handle == "" {
		logger.Error("listUserPosts: No target_did or handle specified")
		return nil, fmt.Errorf("either target_did or handle must be specified")
	}

	client, err := connect(ctx, d)
	if err != nil {
		logger.Error("listUserPosts: Failed to connect", "error", err)
		return nil, fmt.Errorf("failed to connect: %w", err)
	}

	// If handle is provided but DID is not, resolve handle to DID
	if targetDid == "" && handle != "" {
		resp, err := atproto.IdentityResolveHandle(ctx, client, handle)
		if err != nil {
			logger.Error("listUserPosts: Error resolving handle", "error", err, "handle", handle)
			return nil, fmt.Errorf("failed to resolve handle %s: %v", handle, err)
		}
		targetDid = resp.Did
	}

	// Get the user's feed
	feed, err := bsky.FeedGetAuthorFeed(ctx, client, targetDid, "", "", false, 100)
	if err != nil {
		logger.Error("listUserPosts: Failed to get author feed", "error", err)
		return nil, fmt.Errorf("failed to get author feed: %w", err)
	}

	for _, item := range feed.Feed {
		feedPost := item.Post.Record.Val.(*bsky.FeedPost)
		metadata := extractPostMetadata(feedPost)
		mentionedDIDs := metadata["mentioned_handles"].([]string)
		mentionedHandles := resolveDIDsToHandles(ctx, client, mentionedDIDs)

		d.StreamListItem(ctx, map[string]interface{}{
			"uri":                     item.Post.Uri,
			"http_url":                convertToHttpUrl(item.Post.Uri),
			"cid":                     item.Post.Cid,
			"text":                    feedPost.Text,
			"author":                  item.Post.Author.Handle,
			"created_at":              feedPost.CreatedAt,
			"indexed_at":              item.Post.IndexedAt,
			"like_count":              item.Post.LikeCount,
			"repost_count":            item.Post.RepostCount,
			"reply_root":              getReplyRoot(feedPost),
			"reply_parent":            getReplyParent(feedPost),
			"has_external_links":      metadata["has_external_links"],
			"image_count":             metadata["image_count"],
			"hashtags":                metadata["hashtags"],
			"mentioned_handles":       metadata["mentioned_handles"],
			"mentioned_handles_names": mentionedHandles,
			"external_links":          metadata["external_links"],
			"target_did":              targetDid,
			"handle":                  handle,
		})
	}

	// Handle pagination
	cursor := feed.Cursor
	for cursor != nil {
		// Add a small delay to avoid rate limiting
		time.Sleep(100 * time.Millisecond)

		nextFeed, err := bsky.FeedGetAuthorFeed(ctx, client, targetDid, *cursor, "", false, 100)
		if err != nil {
			logger.Error("listUserPosts: Failed to fetch next page", "error", err)
			return nil, fmt.Errorf("failed to fetch next page: %w", err)
		}

		for _, item := range nextFeed.Feed {
			feedPost := item.Post.Record.Val.(*bsky.FeedPost)
			metadata := extractPostMetadata(feedPost)
			mentionedDIDs := metadata["mentioned_handles"].([]string)
			mentionedHandles := resolveDIDsToHandles(ctx, client, mentionedDIDs)

			d.StreamListItem(ctx, map[string]interface{}{
				"uri":                     item.Post.Uri,
				"http_url":                convertToHttpUrl(item.Post.Uri),
				"cid":                     item.Post.Cid,
				"text":                    feedPost.Text,
				"author":                  item.Post.Author.Handle,
				"created_at":              feedPost.CreatedAt,
				"indexed_at":              item.Post.IndexedAt,
				"like_count":              item.Post.LikeCount,
				"repost_count":            item.Post.RepostCount,
				"reply_root":              getReplyRoot(feedPost),
				"reply_parent":            getReplyParent(feedPost),
				"has_external_links":      metadata["has_external_links"],
				"image_count":             metadata["image_count"],
				"hashtags":                metadata["hashtags"],
				"mentioned_handles":       metadata["mentioned_handles"],
				"mentioned_handles_names": mentionedHandles,
				"external_links":          metadata["external_links"],
				"target_did":              targetDid,
				"handle":                  handle,
			})
		}

		cursor = nextFeed.Cursor
	}

	return nil, nil
}
