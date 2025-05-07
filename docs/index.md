---
organization: Turbot
category: ["media"]
icon_url: "/images/plugins/turbot/twitter.svg"
brand_color: "#50b0ff"
display_name: Bluesky
name: bluesky
description: Steampipe plugin to query posts, users and followers from Bluesky.
og_description: Query Bluesky with SQL! Open source CLI. No DB required.
og_image: "/images/plugins/turbot/twitter-social-graphic.png"
engines: ["steampipe", "sqlite", "postgres", "export"]
---


# Bluesky + Steampipe

[Steampipe](https://steampipe.io) is an open-source zero-ETL engine to instantly query cloud APIs using SQL.

[Bluesky](https://bsky.social/about) is a microblogging social media service. Users can share short posts containing text, images, and videos.

```sql
select
  cid,
  author,
  text
from
  bluesky_search_recent
where
  query = '#tailpipe'
```

```
+-------------------------------------------------------------+----------------------------+--------------------------------------------------------------------------------------------------------------------+
| cid                                                         | author                     | text                                                                                                               |
+-------------------------------------------------------------+----------------------------+--------------------------------------------------------------------------------------------------------------------+
| bafyreig62yqzx6ef5ruxe4he53p5bmwii6lxh5qxbtdtqj2vkp5vkzicmi | awscmblogposts.bsky.social | ✍️ New blog post by bob-bot                                                                                        |
|                                                             |                            |                                                                                                                    |
|                                                             |                            | Visualize AWS CloudTrail Logs Locally for Advanced Threat Detection                                                |
|                                                             |                            |                                                                                                                    |
|                                                             |                            | #aws #opensource #tailpipe #powerpipe                                                                              |
| bafyreien7kl7jsjb23ikr6maxgliaumni5orizawr3z6rqm4hullncp6tm | awscmblogposts.bsky.social | ✍️ New blog post by bob-bot                                                                                        |
|                                                             |                            |                                                                                                                    |
|                                                             |                            | Query AWS CloudTrail Logs Locally with SQL                                                                         |
|                                                             |                            |                                                                                                                    |
|                                                             |                            | #aws #security #opensource #tailpipe                                                                               |
+-------------------------------------------------------------+----------------------------+--------------------------------------------------------------------------------------------------------------------+


```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/bluesky/tables)**

## Get started

### Install

Download and install the latest Bluesky plugin:

```bash
steampipe plugin install bluesky
```

### Credentials

| Item | Description |
| - | - |
| Credentials | All API requests require a Bluesky [app password](https://bsky.social/settings/app-passwords). |
| Permissions | Default permissions are sufficient, access to Direct Messages is not required. |
| Radius | Each connection represents a single set of Bluesky credentials. |
| Resolution |  1. `handle`, `app_password` in Steampipe config.<br />2. `BLUESKY_HANDLE`, `BLUESKY_APP_PASSWORD` environment variables.

### Configuration

Installing the latest Bluesky plugin will create a config file (`~/.steampipe/config/bluesky.spc`) with a single connection named `bluesky`:

```hcl
connection "bluesky" {
  plugin = "bluesky"
  
  # Required: Your Bluesky handle (e.g., user.bsky.social)
  handle = "your.handle.bsky.social"
  
  # Required: Your Bluesky app password
  app_password = "your-app-password"
  
  # Optional: Custom PDS host (defaults to https://bsky.social)
  pds_host = "https://bsky.social"
}
```
