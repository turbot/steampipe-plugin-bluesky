package bluesky

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/bluesky-social/indigo/api/bsky"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableBlueskyUserFollowing(ctx context.Context) *plugin.Table {

	return &plugin.Table{
		Name:        "bluesky_user_following",
		Description: "List of users that the specified user follows. To query by handle, use a join with the bluesky_user table.",
		List: &plugin.ListConfig{
			Hydrate: listUserFollowing,
			KeyColumns: plugin.KeyColumnSlice{
				{
					Name:    "target_did",
					Require: plugin.Required,
				},
			},
		},
		DefaultTransform: transform.FromGo().NullIfZero(),
		Columns:          userColumns("target_did"),
	}
}

func processFollowing(ctx context.Context, d *plugin.QueryData, following *bsky.ActorDefs_ProfileView, targetDid string) error {
	logger := plugin.Logger(ctx)

	// Get the full profile for the following
	client, err := connect(ctx, d)
	if err != nil {
		logger.Error("listUserFollowing: Error connecting", "error", err)
		return err
	}

	profile, err := bsky.ActorGetProfile(ctx, client, following.Did)
	if err != nil {
		logger.Error("listUserFollowing: Error getting profile", "error", err, "did", following.Did)
		return fmt.Errorf("failed to get profile for did %s: %v", following.Did, err)
	}

	item := map[string]interface{}{
		"did":             profile.Did,
		"target_did":      targetDid,
		"handle":          profile.Handle,
		"display_name":    derefString(profile.DisplayName),
		"description":     derefString(profile.Description),
		"indexed_at":      derefString(profile.IndexedAt),
		"follower_count":  derefInt64(profile.FollowersCount),
		"following_count": derefInt64(profile.FollowsCount),
		"post_count":      derefInt64(profile.PostsCount),
		"avatar":          derefString(profile.Avatar),
		"banner":          derefString(profile.Banner),
	}

	logger.Debug("listUserFollowing: Created item map with keys", "keys", getMapKeys(item))
	logger.Debug("listUserFollowing: Item map values", "item", item)

	return nil
}

func listUserFollowing(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	targetDid := d.EqualsQuals["target_did"].GetStringValue()
	if targetDid == "" {
		logger.Error("listUserFollowing: No target_did specified")
		return nil, fmt.Errorf("target_did must be specified")
	}

	if !strings.HasPrefix(targetDid, "did:") {
		logger.Error("listUserFollowing: Invalid DID format", "did", targetDid)
		return nil, fmt.Errorf("invalid DID format: %s", targetDid)
	}

	// Get the connection
	client, err := connect(ctx, d)
	if err != nil {
		logger.Error("listUserFollowing: Error connecting", "error", err)
		return nil, err
	}

	// Get following

	following, err := bsky.GraphGetFollows(ctx, client, targetDid, "", 100)
	if err != nil {
		logger.Error("listUserFollowing: Error getting following", "error", err, "did", targetDid)
		return nil, fmt.Errorf("failed to get following for %s: %v", targetDid, err)
	}
	if following == nil {
		logger.Error("listUserFollowing: Empty response from GraphGetFollows", "did", targetDid)
		return nil, fmt.Errorf("received empty response from GraphGetFollows for %s", targetDid)
	}

	// Process each following
	for _, following := range following.Follows {
		if err := processFollowing(ctx, d, following, targetDid); err != nil {
			logger.Error("listUserFollowing: Error processing following", "error", err)
			return nil, err
		}
		// Add a small delay to avoid rate limiting
		time.Sleep(100 * time.Millisecond)
	}

	// Handle pagination
	cursor := following.Cursor
	for cursor != nil {

		// Add a small delay to avoid rate limiting
		time.Sleep(100 * time.Millisecond)

		nextFollowing, err := bsky.GraphGetFollows(ctx, client, targetDid, *cursor, 100)
		if err != nil {
			logger.Error("listUserFollowing: Error getting next page", "error", err)
			return nil, fmt.Errorf("failed to get next page: %v", err)
		}

		for _, following := range nextFollowing.Follows {
			if err := processFollowing(ctx, d, following, targetDid); err != nil {
				logger.Error("listUserFollowing: Error processing following", "error", err)
				return nil, err
			}
			// Add a small delay to avoid rate limiting
			time.Sleep(100 * time.Millisecond)
		}

		cursor = nextFollowing.Cursor
	}

	return nil, nil
}
