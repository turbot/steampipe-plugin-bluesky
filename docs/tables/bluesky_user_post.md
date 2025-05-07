---
title: "Steampipe Table: bluesky_user_post - Query Bluesky User Posts using SQL"
description: "Allows users to query posts published by a specific Bluesky user, providing insights into post content, engagement metrics, and metadata."
---

# Table: bluesky_user_post - Query Bluesky User Posts using SQL

Bluesky is a decentralized social network protocol that allows users to create and share content. The `bluesky_user_post` table provides access to posts published by a specific Bluesky user, including post content, engagement metrics, and metadata.

## Table Usage Guide

The `bluesky_user_post` table provides insights into posts published by a specific Bluesky user. As a data analyst or social media manager, explore post-specific details through this table, including content, engagement metrics, and metadata. Utilize it to uncover information about posting patterns, engagement rates, and the impact of specific posts.

**Important Notes**
- Either `target_did` or `handle` must be specified in the `where` clause.
- If using `target_did`, it must be in the format `did:plc:...` or `did:web:...`
- If using `handle`, it can be provided with or without the `@` prefix

## Examples

### Get all posts by DID
List all posts published by a specific user using their DID.

```sql
select
  uri,
  text,
  author,
  created_at,
  like_count,
  repost_count,
  hashtags,
  mentioned_handles_names,
  external_links
from
  bluesky_user_post
where
  target_did = 'did:plc:vipregezugaizr3kfcjijzrv';
```

### Get all posts by handle
List all posts published by a specific user using their handle.

```sql
select
  uri,
  text,
  author,
  created_at,
  like_count,
  repost_count,
  hashtags,
  mentioned_handles_names,
  external_links
from
  bluesky_user_post
where
  handle = 'matty.wtf';
```

### Get posts with high engagement
Find posts that have received significant engagement (likes and reposts).

```sql
select
  uri,
  text,
  created_at,
  like_count,
  repost_count
from
  bluesky_user_post
where
  handle = 'matty.wtf'
  and like_count > 100;
```

### Get posts with hashtags
Find posts that contain specific hashtags.

```sql
select
  uri,
  text,
  created_at,
  hashtags
from
  bluesky_user_post
where
  handle = 'matty.wtf'
  and hashtags is not null;
``` 