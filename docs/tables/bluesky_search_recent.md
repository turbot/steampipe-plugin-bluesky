---
title: "Steampipe Table: bluesky_search_recent - Query Bluesky Posts using SQL"
description: "Allows users to search for recent posts on Bluesky, with configurable result limits."
---

# Table: bluesky_search_recent

Search for recent posts on Bluesky. Results are limited to 100 by default, but can be changed using the limit parameter.

## Usage

This table is useful for:
- Searching for posts containing specific text, hashtags, or mentions
- Getting recent posts with engagement metrics
- Analyzing post content and metadata
- Finding posts by specific users

## Important Notes

- The `query` parameter is required and must be specified in the `where` clause
- Results are limited to 100 by default
- You can change the limit using the `limit` parameter
- The search is case-insensitive
- Results are returned in reverse chronological order (newest first)

## Examples

### Basic search with default limit

```sql
select
  uri,
  text,
  author,
  created_at,
  like_count,
  repost_count
from
  bluesky_search_recent
where
  query = 'steampipe';
```

### Search with custom limit

```sql
select
  uri,
  text,
  author,
  created_at,
  like_count,
  repost_count
from
  bluesky_search_recent
where
  query = 'steampipe'
  and limit = 50;
```

### Search for posts with hashtags

```sql
select
  uri,
  text,
  author,
  created_at,
  hashtags
from
  bluesky_search_recent
where
  query = '#steampipe'
  and limit = 200;
```

### Search for posts mentioning a user

```sql
select
  uri,
  text,
  author,
  created_at,
  mentioned_handles
from
  bluesky_search_recent
where
  query = '@mattstratton'
  and limit = 100;
```

### Get posts with high engagement

```sql
select
  uri,
  text,
  author,
  created_at,
  like_count,
  repost_count
from
  bluesky_search_recent
where
  query = 'steampipe'
  and limit = 100
order by
  like_count desc;
```

## Schema

| Column | Type | Description |
|--------|------|-------------|
| uri | string | The URI of the post |
| http_url | string | The HTTP URL for the post on bsky.app |
| cid | string | The CID of the post |
| author | string | The handle of the post author |
| text | string | The text content of the post |
| reply_root | string | The URI of the root post if this is a reply |
| reply_parent | string | The URI of the parent post if this is a reply |
| created_at | string | When the post was created |
| indexed_at | string | When the post was indexed |
| like_count | int | Number of likes on the post |
| repost_count | int | Number of reposts of the post |
| has_external_links | bool | Whether the post contains external links |
| image_count | int | Number of images in the post |
| hashtags | json | List of hashtags in the post |
| mentioned_handles | json | List of handles mentioned in the post |
| mentioned_handles_names | json | List of handle names (not DIDs) mentioned in the post |
| external_links | json | List of external links in the post |
| query | string | The search query used to find this post |

## Key Columns

| Column | Type | Description |
|--------|------|-------------|
| query | string | The search query to find posts (required) |

## Notes

- The search query is required and must be specified in the `where` clause.
- The search API returns posts from the last 7 days.
- Results are paginated and will automatically fetch additional pages as needed.
- The search is case-insensitive.
- You can use hashtags (e.g., `#steampipe`) and mentions (e.g., `@mattstratton`) in your search query.
- The table includes metadata about the post such as hashtags, mentions, and external links.

## Table Usage Guide

The `bluesky_search_recent` table provides insights into recent posts on Bluesky. As a data analyst or social media manager, explore post-specific details through this table, including content, engagement metrics, and metadata. Utilize it to uncover information about trending topics, engagement patterns, and user interactions.

**Important Notes**
- The `query` field must be specified in the `where` clause.
- The search query can include hashtags, mentions, or plain text.
- Results are returned in reverse chronological order (newest first).

### Search for posts containing a specific hashtag
Find recent posts that contain a specific hashtag.

```sql
select
  uri,
  http_url,
  text,
  author,
  created_at,
  like_count,
  repost_count,
  hashtags
from
  bluesky_search_recent
where
  query = '#steampipe';
```

### Search for posts mentioning a user
Find recent posts that mention a specific user.

```sql
select
  uri,
  http_url,
  text,
  author,
  created_at,
  like_count,
  repost_count,
  mentioned_handles_names
from
  bluesky_search_recent
where
  query = '@matty.wtf';
```

### Search for posts with external links
Find recent posts that contain external links.

```sql
select
  uri,
  http_url,
  text,
  author,
  created_at,
  external_links
from
  bluesky_search_recent
where
  query = '#steampipe'
  and has_external_links = true;
```

### Search for posts with images
Find recent posts that contain images.

```sql
select
  uri,
  http_url,
  text,
  author,
  created_at,
  image_count
from
  bluesky_search_recent
where
  query = '#steampipe'
  and image_count > 0;
``` 