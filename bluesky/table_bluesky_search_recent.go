package bluesky

import (
	"context"
	"fmt"
	"time"

	"github.com/bluesky-social/indigo/api/bsky"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableBlueskySearchRecent(ctx context.Context) *plugin.Table {

	return &plugin.Table{
		Name:        "bluesky_search_recent",
		Description: "Search for recent posts on Bluesky. Results are limited to 100 by default, but can be changed using the limit parameter.",
		List: &plugin.ListConfig{
			Hydrate: listSearchRecent,
			KeyColumns: plugin.KeyColumnSlice{
				{
					Name:    "query",
					Require: plugin.Required,
				},
				{
					Name:    "limit",
					Require: plugin.Optional,
				},
			},
		},
		DefaultTransform: transform.FromGo().NullIfZero(),
		Columns:          postColumns("query", "limit"),
	}
}

func listSearchRecent(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Get the search query
	query := d.EqualsQuals["query"].GetStringValue()
	if query == "" {
		logger.Error("listSearchRecent: No query specified")
		return nil, fmt.Errorf("query must be specified")
	}

	// Get the limit, default to 100 if not specified
	limit := int64(100)
	if d.EqualsQuals["limit"] != nil {
		limit = d.EqualsQuals["limit"].GetInt64Value()
		if limit <= 0 {
			limit = 100
		}
	}

	// Get the connection
	client, err := connect(ctx, d)
	if err != nil {
		logger.Error("listSearchRecent: Error connecting", "error", err)
		return nil, err
	}

	// Calculate how many results to fetch per page
	// Use min(limit, 100) since the API has a max of 100 per page
	perPage := limit
	if perPage > 100 {
		perPage = 100
	}

	// Get initial search results
	searchResults, err := bsky.FeedSearchPosts(ctx, client, "", "", "", "", perPage, "", query, "", "", nil, "", "")
	if err != nil {
		logger.Error("listSearchRecent: Failed to search posts", "error", err)
		return nil, fmt.Errorf("failed to search posts: %w", err)
	}

	// Track total results returned
	totalReturned := int64(0)

	// Process each post
	for _, post := range searchResults.Posts {
		if totalReturned >= limit {
			break
		}

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
			"query":                   query,
			"limit":                   limit,
		})

		totalReturned++

		// Add a small delay to avoid rate limiting
		time.Sleep(100 * time.Millisecond)
	}

	// Handle pagination if we haven't reached the limit
	cursor := searchResults.Cursor
	for cursor != nil && totalReturned < limit {
		// Add a small delay to avoid rate limiting
		time.Sleep(100 * time.Millisecond)

		// Calculate how many results to fetch in this page
		remaining := limit - totalReturned
		if remaining > 100 {
			remaining = 100
		}

		nextResults, err := bsky.FeedSearchPosts(ctx, client, "", *cursor, "", "", remaining, "", query, "", "", nil, "", "")
		if err != nil {
			logger.Error("listSearchRecent: Failed to fetch next page", "error", err)
			return nil, fmt.Errorf("failed to fetch next page: %w", err)
		}

		for _, post := range nextResults.Posts {
			if totalReturned >= limit {
				break
			}

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
				"query":                   query,
				"limit":                   limit,
			})

			totalReturned++

			// Add a small delay to avoid rate limiting
			time.Sleep(100 * time.Millisecond)
		}

		cursor = nextResults.Cursor
	}

	return nil, nil
}
