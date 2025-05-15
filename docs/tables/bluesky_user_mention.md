---
title: "Steampipe Table: bluesky_user_mention - Query Bluesky User Mentions using SQL"
description: "Allows users to query mentions of a Bluesky user, providing insights into mention content, engagement metrics, and relationship details."
folder: "Mentions"
---

# Table: bluesky_user_mention - Query Bluesky User Mentions using SQL

Bluesky is a decentralized social network protocol that allows users to create and share content. The `bluesky_user_mention` table provides access to mentions of a specific Bluesky user, including mention content, engagement metrics, and relationship details.

## Table Usage Guide

The `bluesky_user_mention` table provides insights into mentions of a specific Bluesky user. As a data analyst or social media manager, explore mention-specific details through this table, including content information, engagement metrics, and relationship details. Utilize it to uncover information about mention patterns, engagement trends, and network interactions.

**Important Notes**

- The `did` field must be set in the `where` clause
- The DID must be in the format `did:plc:...` or `did:web:...`
- To query by handle, use a join with the `bluesky_user` table
- The table provides comprehensive mention information including content, engagement metrics, and media URLs

## Examples

### Get all mentions of a user
List all mentions of a specific user, including the mention content and engagement metrics.

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
  bluesky_user_mention
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
  bluesky_user_mention
where
  did = 'did:plc:vipregezugaizr3kfcjijzrv';
```

### Get mentions with high engagement
Find mentions that have received significant engagement through likes and reposts.

```sql+postgres
select
  uri,
  text,
  created_at,
  like_count,
  repost_count,
  reply_count
from
  bluesky_user_mention
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
  bluesky_user_mention
where
  did = 'did:plc:vipregezugaizr3kfcjijzrv'
  and (like_count > 100 or repost_count > 50);
```

### Get mentions by handle
Look up mentions of a user by their handle instead of DID.

```sql+postgres
select
  m.uri,
  m.text,
  m.created_at,
  m.like_count,
  m.repost_count,
  m.reply_count
from
  bluesky_user_mention m
  join bluesky_user u on m.did = u.did
where
  u.handle = 'matty.wtf';
```

```sql+sqlite
select
  m.uri,
  m.text,
  m.created_at,
  m.like_count,
  m.repost_count,
  m.reply_count
from
  bluesky_user_mention m
  join bluesky_user u on m.did = u.did
where
  u.handle = 'matty.wtf';
``` 