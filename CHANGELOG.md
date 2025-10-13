## v1.1.0 [2025-10-13]

_Dependencies_

- Recompiled plugin with Go version `1.24`. ([#14](https://github.com/turbot/steampipe-plugin-bluesky/pull/14))
- Recompiled plugin with [steampipe-plugin-sdk v5.13.1](https://github.com/turbot/steampipe-plugin-sdk/blob/develop/CHANGELOG.md#v5131-2025-09-25) that addresses critical and high vulnerabilities in dependent packages. ([#18](https://github.com/turbot/steampipe-plugin-bluesky/pull/18))

## v1.0.1 [2025-05-15]

_Bug fixes_

- Fixed the import path and the module path in main.go and go.mod files respectively to use the full GitHub URL.

## v1.0.0 [2025-05-15]

_What's new?_

- New tables added
  - [bluesky_search_recent](https://hub.steampipe.io/plugins/turbot/bluesky/tables/bluesky_search_recent)
  - [bluesky_post](https://hub.steampipe.io/plugins/turbot/bluesky/tables/bluesky_post)
  - [bluesky_user](https://hub.steampipe.io/plugins/turbot/bluesky/tables/bluesky_user)
  - [bluesky_user_follower](https://hub.steampipe.io/plugins/turbot/bluesky/tables/bluesky_user_follower)
  - [bluesky_user_following](https://hub.steampipe.io/plugins/turbot/bluesky/tables/bluesky_user_following)
  - [bluesky_user_mention](https://hub.steampipe.io/plugins/turbot/bluesky/tables/bluesky_user_mention)
  - [bluesky_user_post](https://hub.steampipe.io/plugins/turbot/bluesky/tables/bluesky_user_post)