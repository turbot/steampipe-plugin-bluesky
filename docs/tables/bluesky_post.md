---
title: "Steampipe Table: bluesky_post - Query Bluesky Posts using SQL"
description: "Allows users to query individual Bluesky posts, providing insights into post content, engagement metrics, and metadata."
folder: "Post"
---

# Table: bluesky_post - Query Bluesky Posts using SQL

Bluesky is a decentralized social network protocol that allows users to create and share content. The `bluesky_post` table provides access to individual posts on Bluesky, including post content, engagement metrics, and metadata.

## Table Usage Guide

The `bluesky_post` table in Steampipe provides a comprehensive view of individual Bluesky posts, allowing you to query and analyze post content, engagement metrics, and metadata.

**Important Notes**

- You must specify either the `uri` or `http_url` in the `where` clause
- The `uri` should be in the format `at://did:plc:example/app.bsky.feed.post/postid`
- The `http_url` can be in either of these formats:
  - `https://bsky.app/profile/example.bsky.social/post/postid`
  - `https://bsky.app/profile/did:plc:example/post/postid`
  - URLs can optionally start with `@`
- The table provides comprehensive post information including content, engagement metrics, and metadata

## Examples

### Look up a post by URI

```sql+postgres
select
  uri,
  http_url,
  text,
  author,
  created_at,
  like_count,
  repost_count
from
  bluesky_post
where
  uri = 'at://did:plc:example/app.bsky.feed.post/3k2m6q5dpl42g';
```

```sql+sqlite
select
  uri,
  http_url,
  text,
  author,
  created_at,
  like_count,
  repost_count
from
  bluesky_post
where
  uri = 'at://did:plc:example/app.bsky.feed.post/3k2m6q5dpl42g';
```

### Look up a post by HTTP URL (using handle)

```sql+postgres
select
  uri,
  http_url,
  text,
  author,
  created_at,
  like_count,
  repost_count
from
  bluesky_post
where
  http_url = 'https://bsky.app/profile/example.bsky.social/post/3k2m6q5dpl42g';
```

```sql+sqlite
select
  uri,
  http_url,
  text,
  author,
  created_at,
  like_count,
  repost_count
from
  bluesky_post
where
  http_url = 'https://bsky.app/profile/example.bsky.social/post/3k2m6q5dpl42g';
```

### Look up a post by HTTP URL (using DID)

```sql+postgres
select
  uri,
  http_url,
  text,
  author,
  created_at,
  like_count,
  repost_count
from
  bluesky_post
where
  http_url = 'https://bsky.app/profile/did:plc:example/post/3k2m6q5dpl42g';
```

```sql+sqlite
select
  uri,
  http_url,
  text,
  author,
  created_at,
  like_count,
  repost_count
from
  bluesky_post
where
  http_url = 'https://bsky.app/profile/did:plc:example/post/3k2m6q5dpl42g';
```

### Get post metadata

```sql+postgres
select
  uri,
  text,
  hashtags,
  mentioned_handles,
  mentioned_handles_names,
  external_links,
  has_external_links,
  image_count
from
  bluesky_post
where
  uri = 'at://did:plc:example/app.bsky.feed.post/3k2m6q5dpl42g';
```

```sql+sqlite
select
  uri,
  text,
  hashtags,
  mentioned_handles,
  mentioned_handles_names,
  external_links,
  has_external_links,
  image_count
from
  bluesky_post
where
  uri = 'at://did:plc:example/app.bsky.feed.post/3k2m6q5dpl42g';
``` 