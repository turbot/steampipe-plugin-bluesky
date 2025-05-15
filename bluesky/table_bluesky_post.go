package bluesky

import (
	"context"
	"fmt"
	"strings"

	comatproto "github.com/bluesky-social/indigo/api/atproto"
	"github.com/bluesky-social/indigo/api/bsky"
	"github.com/bluesky-social/indigo/xrpc"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableBlueskyPost(ctx context.Context) *plugin.Table {

	return &plugin.Table{
		Name:        "bluesky_post",
		Description: "Lookup a specific post by URI or HTTP URL.",
		List: &plugin.ListConfig{
			Hydrate: listPost,
			KeyColumns: plugin.KeyColumnSlice{
				{
					Name:    "uri",
					Require: plugin.Optional,
				},
				{
					Name:    "http_url",
					Require: plugin.Optional,
				},
			},
		},
		Columns: postColumns(),
	}
}

// convertToAtURI converts a web URL to an at-uri format
func convertToAtURI(ctx context.Context, client *xrpc.Client, uri string) (string, error) {

	// Remove @ prefix if present
	uri = strings.TrimPrefix(uri, "@")

	// If it's already an at-uri, return it
	if strings.HasPrefix(uri, "at://") {

		return uri, nil
	}

	// Handle bsky.app URLs
	if strings.Contains(uri, "bsky.app") {

		// Remove any query parameters or fragments
		uri = strings.Split(uri, "?")[0]
		uri = strings.Split(uri, "#")[0]

		// Ensure the URL ends with a slash for consistent parsing
		if !strings.HasSuffix(uri, "/") {
			uri = uri + "/"
		}

		// Split the URL into parts and remove empty parts
		var parts []string
		for _, part := range strings.Split(uri, "/") {
			if part != "" {
				parts = append(parts, part)
			}
		}

		// Find the profile and post segments
		var identifier, postID string
		for i := 0; i < len(parts)-1; i++ {

			if parts[i] == "profile" && i+1 < len(parts) {
				identifier = parts[i+1]
			}
			if parts[i] == "post" && i+1 < len(parts) {
				postID = parts[i+1]
			}
		}

		if identifier == "" || postID == "" {
			return "", fmt.Errorf("invalid bsky.app URL format: could not find identifier or post ID")
		}

		// Clean up the identifier
		identifier = strings.TrimSpace(identifier)

		// If the identifier is already a DID, use it directly
		if strings.HasPrefix(identifier, "did:") {
			atURI := fmt.Sprintf("at://%s/app.bsky.feed.post/%s", identifier, postID)
			return atURI, nil
		}

		// Otherwise, try to resolve it as a handle
		didResp, err := comatproto.IdentityResolveHandle(ctx, client, identifier)
		if err != nil {
			return "", fmt.Errorf("failed to resolve identifier '%s' to DID: %v", identifier, err)
		}

		atURI := fmt.Sprintf("at://%s/app.bsky.feed.post/%s", didResp.Did, postID)
		return atURI, nil
	}

	return "", fmt.Errorf("unsupported URI format")
}

func listPost(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	conn, err := connect(ctx, d)
	if err != nil {
		logger.Error("listPost: Connection error", "error", err)
		return nil, err
	}

	var uri string
	if d.EqualsQuals["uri"] != nil {
		uri = d.EqualsQuals["uri"].GetStringValue()
	} else if d.EqualsQuals["http_url"] != nil {
		httpUrl := d.EqualsQuals["http_url"].GetStringValue()
		uri, err = convertToAtURI(ctx, conn, httpUrl)
		if err != nil {
			logger.Error("listPost: Error converting HTTP URL to URI", "error", err)
			return nil, fmt.Errorf("failed to convert HTTP URL to URI: %w", err)
		}
	}

	if uri == "" {
		logger.Error("listPost: URI is empty")
		return nil, fmt.Errorf("either uri or http_url must be specified")
	}

	thread, err := bsky.FeedGetPostThread(ctx, conn, 0, 0, uri)
	if err != nil {
		logger.Error("listPost: Error getting thread", "error", err)
		return nil, err
	}

	if thread.Thread.FeedDefs_ThreadViewPost == nil {
		logger.Debug("listPost: No post view found")
		return nil, nil
	}

	post := thread.Thread.FeedDefs_ThreadViewPost.Post
	feedPost, ok := post.Record.Val.(*bsky.FeedPost)
	if !ok {
		logger.Error("listPost: Could not convert to FeedPost")
		return nil, nil
	}

	metadata := extractPostMetadata(feedPost)
	mentionedDIDs := metadata["mentioned_handles"].([]string)
	mentionedHandles := resolveDIDsToHandles(ctx, conn, mentionedDIDs)

	item := map[string]interface{}{
		"uri":                     post.Uri,
		"http_url":                convertToHttpUrl(post.Uri),
		"cid":                     post.Cid,
		"author":                  post.Author.Handle,
		"text":                    feedPost.Text,
		"reply_root":              getReplyRoot(feedPost),
		"reply_parent":            getReplyParent(feedPost),
		"created_at":              feedPost.CreatedAt,
		"indexed_at":              post.IndexedAt,
		"like_count":              post.LikeCount,
		"repost_count":            post.RepostCount,
		"has_external_links":      metadata["has_external_links"],
		"image_count":             metadata["image_count"],
		"hashtags":                metadata["hashtags"],
		"mentioned_handles":       metadata["mentioned_handles"],
		"mentioned_handles_names": mentionedHandles,
		"external_links":          metadata["external_links"],
	}

	d.StreamListItem(ctx, item)
	return nil, nil
}

// getReplyRoot safely extracts the reply root URI if it exists
func getReplyRoot(post *bsky.FeedPost) string {
	if post.Reply != nil && post.Reply.Root != nil {
		return post.Reply.Root.Uri
	}
	return ""
}

// getReplyParent safely extracts the reply parent URI if it exists
func getReplyParent(post *bsky.FeedPost) string {
	if post.Reply != nil && post.Reply.Parent != nil {
		return post.Reply.Parent.Uri
	}
	return ""
}

// Helper function to get map keys for debugging
func getMapKeys(m map[string]interface{}) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func extractPostMetadata(post *bsky.FeedPost) map[string]interface{} {
	metadata := make(map[string]interface{})

	// Extract hashtags and mentions from facets
	var hashtags []string
	var mentionedHandles []string
	var externalLinks []string

	if post.Facets != nil {
		for _, facet := range post.Facets {
			for _, feature := range facet.Features {
				switch {
				case feature.RichtextFacet_Tag != nil:
					hashtags = append(hashtags, feature.RichtextFacet_Tag.Tag)
				case feature.RichtextFacet_Mention != nil:
					mentionedHandles = append(mentionedHandles, feature.RichtextFacet_Mention.Did)
				case feature.RichtextFacet_Link != nil:
					externalLinks = append(externalLinks, feature.RichtextFacet_Link.Uri)
				}
			}
		}
	}

	metadata["hashtags"] = hashtags
	metadata["mentioned_handles"] = mentionedHandles
	metadata["external_links"] = externalLinks
	metadata["has_external_links"] = len(externalLinks) > 0

	// Count images
	imageCount := 0
	if post.Embed != nil {
		// Access embed data directly
		if post.Embed.EmbedImages != nil {
			imageCount = len(post.Embed.EmbedImages.Images)
		}
	}
	metadata["image_count"] = imageCount

	return metadata
}
