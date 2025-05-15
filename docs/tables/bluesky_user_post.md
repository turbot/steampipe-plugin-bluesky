---
title: "Steampipe Table: bluesky_user_post - Query Bluesky User Posts using SQL"
description: "Allows users to query posts by a Bluesky user, providing insights into post content, engagement metrics, and relationship details."
folder: "User Post"
---

# Table: bluesky_user_post - Query Bluesky User Posts using SQL

Bluesky is a decentralized social network protocol that allows users to create and share content. The `bluesky_user_post` table provides access to posts by a specific Bluesky user, including post content, engagement metrics, and relationship details.

## Table Usage Guide

The `bluesky_user_post` table provides insights into posts by a specific Bluesky user. As a data analyst or social media manager, explore post-specific details through this table, including content information, engagement metrics, and relationship details. Utilize it to uncover information about post patterns, engagement trends, and network interactions.

**Important Notes**
- The `did` field must be set in the `where` clause
- The DID must be in the format `did:plc:...` or `did:web:...`
- To query by handle, use a join with the `bluesky_user` table
- The table provides comprehensive post information including content, engagement metrics, and media URLs

## Examples

### Get all posts by a user
List all posts by a specific user, including the post content and engagement metrics.

```sql+postgres
select
  uri,
  cid,
  text,
  created_at,
  like_count,
  repost_count,
  reply_count
from
  bluesky_user_post
where
  did = 'did:plc:vipregezugaizr3kfcjijzrv';
```

```sql+sqlite
select
  uri,
  cid,
  text,
  created_at,
  like_count,
  repost_count,
  reply_count
from
  bluesky_user_post
where
  did = 'did:plc:vipregezugaizr3kfcjijzrv';
```

### Get posts with high engagement
Find posts that have received significant engagement through likes and reposts.

```sql+postgres
select
  uri,
  text,
  created_at,
  like_count,
  repost_count,
  reply_count
from
  bluesky_user_post
where
  did = 'did:plc:vipregezugaizr3kfcjijzrv'
  and (like_count > 100 or repost_count > 50);
```

```sql+sqlite
select
  uri,
  text,
  created_at,
  like_count,
  repost_count,
  reply_count
from
  bluesky_user_post
where
  did = 'did:plc:vipregezugaizr3kfcjijzrv'
  and (like_count > 100 or repost_count > 50);
```

### Get posts by handle
Look up posts by a user by their handle instead of DID.

```sql+postgres
select
  p.uri,
  p.text,
  p.created_at,
  p.like_count,
  p.repost_count,
  p.reply_count
from
  bluesky_user_post p
  join bluesky_user u on p.did = u.did
where
  u.handle = 'matty.wtf';
```

```sql+sqlite
select
  p.uri,
  p.text,
  p.created_at,
  p.like_count,
  p.repost_count,
  p.reply_count
from
  bluesky_user_post p
  join bluesky_user u on p.did = u.did
where
  u.handle = 'matty.wtf';
``` 