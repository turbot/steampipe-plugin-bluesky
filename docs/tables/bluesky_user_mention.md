---
title: "Steampipe Table: bluesky_user_mention - Query Bluesky User Mentions using SQL"
description: "Allows users to query posts that mention a specific Bluesky user, providing insights into mentions, engagement metrics, and content details."
---

# Table: bluesky_user_mention - Query Bluesky User Mentions using SQL

Bluesky is a decentralized social network protocol that allows users to create and share content. The `bluesky_user_mention` table provides access to posts that mention a specific Bluesky user, including mention content, engagement metrics, and metadata.

## Table Usage Guide

The `bluesky_user_mention` table provides insights into posts that mention a specific Bluesky user. As a data analyst or social media manager, explore mention-specific details through this table, including content, engagement metrics, and metadata. Utilize it to uncover information about mention patterns, engagement rates, and the impact of specific mentions.

**Important Notes**
- The `target_did` field must be specified in the `where` clause.
- The DID must be in the format `did:plc:...` or `did:web:...`
- To query by handle, use a join with the `bluesky_user` table as shown in the examples below.

## Examples

### Get all mentions by DID
List all posts that mention a specific user using their DID.

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
  bluesky_user_mention
where
  target_did = 'did:plc:vipregezugaizr3kfcjijzrv';
```

### Get all mentions by handle
List all posts that mention a specific user using their handle.

```sql
select
  m.uri,
  m.text,
  m.author,
  m.created_at,
  m.like_count,
  m.repost_count,
  m.hashtags,
  m.mentioned_handles_names,
  m.external_links
from
  bluesky_user_mention m
  join bluesky_user u on m.target_did = u.did
where
  u.handle = 'matty.wtf';
```

### Get mentions with high engagement
Find mentions that have received significant engagement (likes and reposts).

```sql
select
  m.uri,
  m.text,
  m.author,
  m.created_at,
  m.like_count,
  m.repost_count
from
  bluesky_user_mention m
  join bluesky_user u on m.target_did = u.did
where
  u.handle = 'matty.wtf'
  and m.like_count > 10;
```

### Get mentions with external links
Find mentions that include external links.

```sql
select
  m.uri,
  m.text,
  m.author,
  m.created_at,
  m.external_links
from
  bluesky_user_mention m
  join bluesky_user u on m.target_did = u.did
where
  u.handle = 'matty.wtf'
  and m.has_external_links = true;
```

### Get mentions with hashtags
Find mentions that include hashtags.

```sql
select
  m.uri,
  m.text,
  m.author,
  m.created_at,
  m.hashtags
from
  bluesky_user_mention m
  join bluesky_user u on m.target_did = u.did
where
  u.handle = 'matty.wtf'
  and m.hashtags is not null;
``` 