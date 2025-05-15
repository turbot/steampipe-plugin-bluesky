---
title: "Steampipe Table: bluesky_search_recent - Query Bluesky Posts using SQL"
description: "Allows users to search for recent posts on Bluesky, with configurable result limits."
folder: "Search Recent"
---

# Table: bluesky_search_recent - Query Bluesky Posts using SQL

Bluesky is a decentralized social network protocol that allows users to create and share content. The `bluesky_search_recent` table provides access to recent posts on Bluesky, including search results with configurable limits and comprehensive post details.

## Table Usage Guide

The `bluesky_search_recent` table provides insights into recent posts on Bluesky. As a data analyst or social media manager, explore post-specific details through this table, including content, engagement metrics, and metadata. Utilize it to uncover information about trending topics, engagement patterns, and user interactions.

### Key Columns

| Column | Type | Description |
|--------|------|-------------|
| query | string | The search query to find posts (required) |
| limit | int | The maximum number of results to return (optional, defaults to 100) |

**Important Notes**

- The search query is required and must be specified in the `where` clause
- The search API returns posts from the last 7 days
- Results are paginated and will automatically fetch additional pages as needed
- The search is case-insensitive
- You can use hashtags (e.g., `#steampipe`) and mentions (e.g., `@matty.wtf`) in your search query
- The table includes metadata about the post such as hashtags, mentions, and external links

## Examples

### Basic search with default limit

```sql+postgres
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

```sql+sqlite
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

```sql+postgres
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

```sql+sqlite
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

```sql+postgres
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

```sql+sqlite
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

```sql+postgres
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

```sql+sqlite
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

```sql+postgres
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

```sql+sqlite
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

### Search for posts containing a specific hashtag
Find recent posts that contain a specific hashtag.

```sql+postgres
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

```sql+sqlite
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

```sql+postgres
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

```sql+sqlite
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

```sql+postgres
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

```sql+sqlite
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

```sql+postgres
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

```sql+sqlite
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