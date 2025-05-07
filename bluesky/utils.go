package bluesky

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/bluesky-social/indigo/api/atproto"
	"github.com/bluesky-social/indigo/api/bsky"
	"github.com/bluesky-social/indigo/xrpc"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

// Global map to hold XRPC clients keyed by connection name
var (
	xrpcClients   = make(map[string]*xrpc.Client)
	xrpcClientsMu sync.Mutex
)

// connect ensures an authenticated XRPC client is available for the connection.
// It handles reuse and creation of clients.
func connect(ctx context.Context, d *plugin.QueryData) (*xrpc.Client, error) {
	logger := plugin.Logger(ctx)

	connName := d.Connection.Name
	xrpcClientsMu.Lock()
	defer xrpcClientsMu.Unlock()

	if client, ok := xrpcClients[connName]; ok && client != nil {
		return client, nil
	}

	blueskyConfig, err := GetConfig(d.Connection)
	if err != nil {
		logger.Error("connect: Failed to get config", "error", err)
		return nil, fmt.Errorf("failed to get config: %v", err)
	}

	// Validate required configuration
	if blueskyConfig.Handle == nil || *blueskyConfig.Handle == "" {
		logger.Error("connect: handle is empty")
		return nil, fmt.Errorf("handle is required")
	}

	if blueskyConfig.AppPassword == nil || *blueskyConfig.AppPassword == "" {
		logger.Error("connect: app_password is empty")
		return nil, fmt.Errorf("app_password is required")
	}

	pdsHost := "https://bsky.social"
	if blueskyConfig.PdsHost != nil && *blueskyConfig.PdsHost != "" {
		pdsHost = *blueskyConfig.PdsHost
	}

	c := &xrpc.Client{
		Host: pdsHost,
	}

	sessResp, err := atproto.ServerCreateSession(ctx, c, &atproto.ServerCreateSession_Input{
		Identifier: *blueskyConfig.Handle,
		Password:   *blueskyConfig.AppPassword,
	})
	if err != nil {
		logger.Error("connect: Authentication failed", "error", err, "handle", *blueskyConfig.Handle)
		return nil, fmt.Errorf("authentication failed for handle '%s': %w", *blueskyConfig.Handle, err)
	}

	c.Auth = &xrpc.AuthInfo{
		AccessJwt:  sessResp.AccessJwt,
		RefreshJwt: sessResp.RefreshJwt,
		Handle:     sessResp.Handle,
		Did:        sessResp.Did,
	}

	xrpcClients[connName] = c
	return c, nil
}

func isDID(s string) bool {
	return strings.HasPrefix(s, "did:")
}

// resolveToDID resolves a handle to a DID using the Bluesky API
func resolveToDID(ctx context.Context, client *xrpc.Client, handle string) (string, error) {
	resolved, err := atproto.IdentityResolveHandle(ctx, client, handle)
	if err != nil {
		return "", fmt.Errorf("failed to resolve handle: %v", err)
	}
	return resolved.Did, nil
}

// resolveDIDsToHandles resolves a list of DIDs to their corresponding handles
func resolveDIDsToHandles(ctx context.Context, client *xrpc.Client, dids []string) []string {
	handles := make([]string, 0, len(dids))
	for _, did := range dids {
		// Skip if it's already a handle
		if !strings.HasPrefix(did, "did:") {
			handles = append(handles, did)
			continue
		}

		// Try to get the profile
		resp, err := bsky.ActorGetProfile(ctx, client, did)
		if err != nil {
			// If we can't resolve it, keep the DID
			handles = append(handles, did)
			continue
		}

		// Extract handle from the profile
		handles = append(handles, resp.Handle)
	}
	return handles
}

func postColumns(optionalCols ...string) []*plugin.Column {
	cols := []*plugin.Column{
		{Name: "uri", Type: proto.ColumnType_STRING, Description: "The URI of the post.", Transform: transform.FromField("uri")},
		{Name: "http_url", Type: proto.ColumnType_STRING, Description: "The HTTP URL for the post on bsky.app.", Transform: transform.FromField("http_url")},
		{Name: "cid", Type: proto.ColumnType_STRING, Description: "The CID of the post.", Transform: transform.FromField("cid")},
		{Name: "author", Type: proto.ColumnType_STRING, Description: "The handle of the post author.", Transform: transform.FromField("author")},
		{Name: "text", Type: proto.ColumnType_STRING, Description: "The text content of the post.", Transform: transform.FromField("text")},
		{Name: "reply_root", Type: proto.ColumnType_STRING, Description: "The URI of the root post if this is a reply.", Transform: transform.FromField("reply_root")},
		{Name: "reply_parent", Type: proto.ColumnType_STRING, Description: "The URI of the parent post if this is a reply.", Transform: transform.FromField("reply_parent")},
		{Name: "created_at", Type: proto.ColumnType_STRING, Description: "When the post was created.", Transform: transform.FromField("created_at")},
		{Name: "indexed_at", Type: proto.ColumnType_STRING, Description: "When the post was indexed.", Transform: transform.FromField("indexed_at")},
		{Name: "like_count", Type: proto.ColumnType_INT, Description: "Number of likes on the post.", Transform: transform.FromField("like_count")},
		{Name: "repost_count", Type: proto.ColumnType_INT, Description: "Number of reposts of the post.", Transform: transform.FromField("repost_count")},
		{Name: "has_external_links", Type: proto.ColumnType_BOOL, Description: "Whether the post contains external links.", Transform: transform.FromField("has_external_links")},
		{Name: "image_count", Type: proto.ColumnType_INT, Description: "Number of images in the post.", Transform: transform.FromField("image_count")},
		{Name: "hashtags", Type: proto.ColumnType_JSON, Description: "List of hashtags in the post.", Transform: transform.FromField("hashtags")},
		{Name: "mentioned_handles", Type: proto.ColumnType_JSON, Description: "List of handles mentioned in the post.", Transform: transform.FromField("mentioned_handles")},
		{Name: "mentioned_handles_names", Type: proto.ColumnType_JSON, Description: "List of handle names (not DIDs) mentioned in the post.", Transform: transform.FromField("mentioned_handles_names")},
		{Name: "external_links", Type: proto.ColumnType_JSON, Description: "List of external links in the post.", Transform: transform.FromField("external_links")},
	}

	// Add optional columns
	for _, col := range optionalCols {
		switch col {
		case "target_did":
			cols = append(cols, &plugin.Column{
				Name:        "target_did",
				Type:        proto.ColumnType_STRING,
				Description: "The DID of the target user.",
				Transform:   transform.FromField("target_did"),
			})
		case "handle":
			cols = append(cols, &plugin.Column{
				Name:        "handle",
				Type:        proto.ColumnType_STRING,
				Description: "The handle of the target user.",
				Transform:   transform.FromField("handle"),
			})
		case "query":
			cols = append(cols, &plugin.Column{
				Name:        "query",
				Type:        proto.ColumnType_STRING,
				Description: "The search query used to find this post.",
				Transform:   transform.FromField("query"),
			})
		case "limit":
			cols = append(cols, &plugin.Column{
				Name:        "limit",
				Type:        proto.ColumnType_INT,
				Description: "The maximum number of results to return.",
				Transform:   transform.FromField("limit"),
			})
		}
	}

	return cols
}

func userColumns(optionalCols ...string) []*plugin.Column {
	cols := []*plugin.Column{
		{Name: "did", Type: proto.ColumnType_STRING, Description: "The DID of the user.", Transform: transform.FromField("did")},
		{Name: "handle", Type: proto.ColumnType_STRING, Description: "The handle of the user.", Transform: transform.FromField("handle")},
		{Name: "display_name", Type: proto.ColumnType_STRING, Description: "The display name of the user.", Transform: transform.FromField("display_name")},
		{Name: "description", Type: proto.ColumnType_STRING, Description: "The user's bio/description.", Transform: transform.FromField("description")},
		{Name: "indexed_at", Type: proto.ColumnType_STRING, Description: "When the user was indexed.", Transform: transform.FromField("indexed_at")},
		{Name: "follower_count", Type: proto.ColumnType_INT, Description: "Number of followers.", Transform: transform.FromField("follower_count")},
		{Name: "following_count", Type: proto.ColumnType_INT, Description: "Number of users being followed.", Transform: transform.FromField("following_count")},
		{Name: "post_count", Type: proto.ColumnType_INT, Description: "Number of posts by the user.", Transform: transform.FromField("post_count")},
		{Name: "avatar", Type: proto.ColumnType_STRING, Description: "URL of the user's avatar image.", Transform: transform.FromField("avatar")},
		{Name: "banner", Type: proto.ColumnType_STRING, Description: "URL of the user's banner image.", Transform: transform.FromField("banner")},
	}

	for _, col := range optionalCols {
		switch col {
		case "target_did":
			cols = append(cols, &plugin.Column{
				Name:        "target_did",
				Type:        proto.ColumnType_STRING,
				Description: "The DID of the target user.",
				Transform:   transform.FromField("target_did"),
			})
		case "handle":
			cols = append(cols, &plugin.Column{
				Name:        "handle",
				Type:        proto.ColumnType_STRING,
				Description: "The handle of the target user.",
				Transform:   transform.FromField("handle"),
			})
		}
	}
	return cols
}

func targetDidString(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.EqualsQuals
	did := quals["target_did"].GetStringValue()
	return did, nil
}

func didString(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.EqualsQuals
	did := quals["did"].GetStringValue()
	return did, nil
}

// Helper functions for safe dereferencing
func derefString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func derefInt64(i *int64) int64 {
	if i == nil {
		return 0
	}
	return *i
}

// convertToHttpUrl converts an at:// URI to a https://bsky.app URL
func convertToHttpUrl(uri string) string {
	if uri == "" {
		return ""
	}
	// Example: at://did:plc:vipregezugaizr3kfcjijzrv/app.bsky.feed.post/3k2m6q5dpl42g
	// to https://bsky.app/profile/matty.wtf/post/3k2m6q5dpl42g
	parts := strings.Split(uri, "/")
	if len(parts) != 5 {
		return uri
	}
	did := parts[2]
	rkey := parts[4]
	return fmt.Sprintf("https://bsky.app/profile/%s/post/%s", did, rkey)
}
