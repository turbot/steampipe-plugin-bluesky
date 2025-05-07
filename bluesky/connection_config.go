package bluesky

import (
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type blueskyConfig struct {
	AppPassword *string `hcl:"app_password"`
	Handle      *string `hcl:"handle"` // User handle (e.g., user.bsky.social)
	PdsHost     *string `hcl:"pds_host"`
}

func ConfigInstance() interface{} {
	return &blueskyConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) (blueskyConfig, error) {
	if connection == nil || connection.Config == nil {
		return blueskyConfig{}, fmt.Errorf("connection config is nil")
	}

	config, ok := connection.Config.(blueskyConfig)
	if !ok {
		return blueskyConfig{}, fmt.Errorf("unable to cast connection config to blueskyConfig")
	}

	// Validate required fields
	if config.Handle == nil {
		return blueskyConfig{}, fmt.Errorf("handle is required")
	}
	if config.AppPassword == nil {
		return blueskyConfig{}, fmt.Errorf("app_password is required")
	}

	// Set default PDS host if not specified
	if config.PdsHost == nil {
		defaultHost := "https://bsky.social"
		config.PdsHost = &defaultHost
	}

	return config, nil
}
