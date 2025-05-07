package bluesky

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/bluesky-social/indigo/api/bsky"
	"github.com/bluesky-social/indigo/xrpc"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func userFollowerColumns() []*plugin.Column {
	return []*plugin.Column{
		{Name: "did", Type: proto.ColumnType_STRING, Description: "The DID of the follower."},
		{Name: "target_did", Type: proto.ColumnType_STRING, Description: "The DID of the user being followed."},
		{Name: "handle", Type: proto.ColumnType_STRING, Description: "The handle of the follower."},
		{Name: "display_name", Type: proto.ColumnType_STRING, Description: "The display name of the follower."},
		{Name: "description", Type: proto.ColumnType_STRING, Description: "The description/bio of the follower."},
		{Name: "indexed_at", Type: proto.ColumnType_STRING, Description: "When the follower was indexed."},
		{Name: "follower_count", Type: proto.ColumnType_INT, Description: "Number of followers the user has."},
		{Name: "following_count", Type: proto.ColumnType_INT, Description: "Number of users the follower is following."},
		{Name: "post_count", Type: proto.ColumnType_INT, Description: "Number of posts the follower has made."},
		{Name: "avatar", Type: proto.ColumnType_STRING, Description: "URL of the follower's avatar."},
		{Name: "banner", Type: proto.ColumnType_STRING, Description: "URL of the follower's banner image."},
	}
}

func tableBlueskyUserFollower(ctx context.Context) *plugin.Table {
	logger := plugin.Logger(ctx)
	logger.Debug("tableBlueskyUserFollower: Creating table definition")

	return &plugin.Table{
		Name:        "bluesky_user_follower",
		Description: "List of users who follow the specified user. To query by handle, use a join with the bluesky_user table.",
		List: &plugin.ListConfig{
			Hydrate: listUserFollower,
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

func processFollower(ctx context.Context, d *plugin.QueryData, follower *bsky.ActorDefs_ProfileView, targetDid string) error {
	logger := plugin.Logger(ctx)

	// Get the full profile for the follower
	client, err := connect(ctx, d)
	if err != nil {
		logger.Error("listUserFollower: Error connecting", "error", err)
		return err
	}

	profile, err := bsky.ActorGetProfile(ctx, client, follower.Did)
	if err != nil {
		logger.Error("listUserFollower: Error getting profile", "error", err, "did", follower.Did)
		return fmt.Errorf("failed to get profile for did %s: %v", follower.Did, err)
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

	logger.Debug("listUserFollower: Created item map with keys", "keys", getMapKeys(item))
	logger.Debug("listUserFollower: Item map values", "item", item)

	return nil
}

func getFollowersWithRetry(ctx context.Context, client *xrpc.Client, did string, cursor string, maxRetries int) (*bsky.GraphGetFollowers_Output, error) {
	logger := plugin.Logger(ctx)
	var lastErr error

	for i := 0; i < maxRetries; i++ {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			if i > 0 {
				// Exponential backoff
				backoff := time.Duration(1<<uint(i)) * time.Second
				logger.Warn("listUserFollower: Retrying after error", "attempt", i+1, "backoff", backoff)
				time.Sleep(backoff)
			}

			// Create a new context with a timeout for the followers fetch
			followersCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
			defer cancel()

			followers, err := bsky.GraphGetFollowers(followersCtx, client, did, cursor, 100)
			if err == nil {
				return followers, nil
			}

			lastErr = err
			logger.Error("listUserFollower: Error getting followers", "error", err, "attempt", i+1)
		}
	}

	return nil, fmt.Errorf("failed to get followers after %d attempts: %v", maxRetries, lastErr)
}

func listUserFollower(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	logger.Debug("listUserFollower: Starting follower lookup")

	targetDid := d.EqualsQuals["target_did"].GetStringValue()
	if targetDid == "" {
		logger.Error("listUserFollower: No target_did specified")
		return nil, fmt.Errorf("target_did must be specified")
	}

	if !strings.HasPrefix(targetDid, "did:") {
		logger.Error("listUserFollower: Invalid DID format", "did", targetDid)
		return nil, fmt.Errorf("invalid DID format: %s", targetDid)
	}

	// Get the connection
	client, err := connect(ctx, d)
	if err != nil {
		logger.Error("listUserFollower: Error connecting", "error", err)
		return nil, err
	}

	// Get followers

	followers, err := bsky.GraphGetFollowers(ctx, client, targetDid, "", 100)
	if err != nil {
		logger.Error("listUserFollower: Error getting followers", "error", err, "did", targetDid)
		return nil, fmt.Errorf("failed to get followers for %s: %v", targetDid, err)
	}
	if followers == nil {
		logger.Error("listUserFollower: Empty response from GraphGetFollowers", "did", targetDid)
		return nil, fmt.Errorf("received empty response from GraphGetFollowers for %s", targetDid)
	}

	// Process each follower
	for _, follower := range followers.Followers {
		if err := processFollower(ctx, d, follower, targetDid); err != nil {
			logger.Error("listUserFollower: Error processing follower", "error", err)
			return nil, err
		}
		// Add a small delay to avoid rate limiting
		time.Sleep(100 * time.Millisecond)
	}

	// Handle pagination
	cursor := followers.Cursor
	for cursor != nil {

		// Add a small delay to avoid rate limiting
		time.Sleep(100 * time.Millisecond)

		nextFollowers, err := bsky.GraphGetFollowers(ctx, client, targetDid, *cursor, 100)
		if err != nil {
			logger.Error("listUserFollower: Error getting next page", "error", err)
			return nil, fmt.Errorf("failed to get next page: %v", err)
		}

		for _, follower := range nextFollowers.Followers {
			if err := processFollower(ctx, d, follower, targetDid); err != nil {
				logger.Error("listUserFollower: Error processing follower", "error", err)
				return nil, err
			}
			// Add a small delay to avoid rate limiting
			time.Sleep(100 * time.Millisecond)
		}

		cursor = nextFollowers.Cursor
	}

	return nil, nil
}
