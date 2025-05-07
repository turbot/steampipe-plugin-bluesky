---
title: "Steampipe Table: bluesky_post - Query Bluesky Posts using SQL"
description: "Allows users to query individual Bluesky posts, providing insights into post content, engagement metrics, and metadata."
---

# Table: bluesky_post

Query a specific Bluesky post by its URI or HTTP URL.

## Usage

This table is useful for:
- Looking up specific posts by their URI or HTTP URL
- Getting detailed information about a post, including its content, engagement metrics, and metadata
- Analyzing post content, mentions, and hashtags

## Important Notes

- You must specify either the `uri` or `http_url` in the `where` clause
- The `uri` should be in the format `at://did:plc:example/app.bsky.feed.post/postid`
- The `http_url` can be in either of these formats:
  - `https://bsky.app/profile/example.bsky.social/post/postid`
  - `https://bsky.app/profile/did:plc:example/post/postid`
  - URLs can optionally start with `@`

## Examples

### Look up a post by URI

```sql
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

```sql
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

```sql
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

```sql
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

### Get reply information

```sql
select
  uri,
  text,
  reply_root,
  reply_parent,
  created_at
from
  bluesky_post
where
  uri = 'at://did:plc:example/app.bsky.feed.post/3k2m6q5dpl42g';
``` 