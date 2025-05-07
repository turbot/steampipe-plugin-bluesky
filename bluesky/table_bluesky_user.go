package bluesky

import (
	"context"
	"fmt"
	"strings"

	"github.com/bluesky-social/indigo/api/atproto"
	"github.com/bluesky-social/indigo/api/bsky"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableBlueskyUser(ctx context.Context) *plugin.Table {

	return &plugin.Table{
		Name:        "bluesky_user",
		Description: "Get information about a Bluesky user",
		List: &plugin.ListConfig{
			Hydrate: listUser,
			KeyColumns: plugin.KeyColumnSlice{
				{
					Name:    "did",
					Require: plugin.Optional,
				},
				{
					Name:    "handle",
					Require: plugin.Optional,
				},
			},
		},
		DefaultTransform: transform.FromGo().NullIfZero(),
		Columns:          userColumns(),
	}
}

func listUser(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	client, err := connect(ctx, d)
	if err != nil {
		logger.Error("listUser: Connection error", "error", err)
		return nil, fmt.Errorf("connection error: %v", err)
	}

	var did string
	var handle string

	// Check if DID is provided
	if d.EqualsQuals["did"] != nil {
		did = d.EqualsQuals["did"].GetStringValue()
		if did != "" {
			if !strings.HasPrefix(did, "did:") {
				logger.Error("listUser: Invalid DID format", "did", did)
				return nil, fmt.Errorf("invalid did format: must start with 'did:'")
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
	if did == "" && handle == "" {
		logger.Error("listUser: No DID or handle specified")
		return nil, fmt.Errorf("either did or handle must be specified")
	}

	// If handle is provided but DID is not, resolve handle to DID
	if did == "" && handle != "" {
		resp, err := atproto.IdentityResolveHandle(ctx, client, handle)
		if err != nil {
			logger.Error("listUser: Error resolving handle", "error", err, "handle", handle)
			return nil, fmt.Errorf("failed to resolve handle %s: %v", handle, err)
		}
		did = resp.Did
	}

	profile, err := bsky.ActorGetProfile(ctx, client, did)
	if err != nil {
		logger.Error("listUser: Error getting profile", "error", err, "did", did)
		return nil, fmt.Errorf("failed to get profile for did %s: %v", did, err)
	}

	item := map[string]interface{}{
		"did":             profile.Did,
		"handle":          profile.Handle,
		"display_name":    derefString(profile.DisplayName),
		"description":     derefString(profile.Description),
		"indexed_at":      profile.IndexedAt,
		"follower_count":  derefInt64(profile.FollowersCount),
		"following_count": derefInt64(profile.FollowsCount),
		"post_count":      derefInt64(profile.PostsCount),
		"avatar":          derefString(profile.Avatar),
		"banner":          derefString(profile.Banner),
	}

	logger.Debug("listUser: Streaming item", "did", did, "handle", profile.Handle)
	d.StreamListItem(ctx, item)
	return nil, nil
}
