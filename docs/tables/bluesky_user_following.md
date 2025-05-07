---
title: "Steampipe Table: bluesky_user_following - Query Bluesky User Following using SQL"
description: "Allows users to query the list of users that a specific Bluesky user follows, providing insights into following relationships and user profiles."
---

# Table: bluesky_user_following - Query Bluesky User Following using SQL

Bluesky is a decentralized social network protocol that allows users to create and share content. The `bluesky_user_following` table provides access to the list of users that a specific Bluesky user follows, including profile information and relationship details.

## Table Usage Guide

The `bluesky_user_following` table provides insights into the users that a specific Bluesky user follows. As a data analyst or social media manager, explore following-specific details through this table, including profile information, engagement metrics, and relationship data. Utilize it to uncover information about following patterns, user connections, and network analysis.

**Important Notes**
- Either `target_did` or `handle` must be specified in the `where` clause.
- If using `target_did`, it must be in the format `did:plc:...` or `did:web:...`
- If using `handle`, it can be provided with or without the `@` prefix

## Examples

### Get all following by DID
List all users that a specific user follows using their DID.

```sql
select
  did,
  handle,
  display_name,
  description,
  follower_count,
  following_count,
  post_count
from
  bluesky_user_following
where
  target_did = 'did:plc:vipregezugaizr3kfcjijzrv';
```

### Get all following by handle
List all users that a specific user follows using their handle.

```sql
select
  did,
  handle,
  display_name,
  description,
  follower_count,
  following_count,
  post_count
from
  bluesky_user_following
where
  handle = 'matty.wtf';
```

### Get following with high follower counts
Find followed users who have a significant number of followers.

```sql
select
  did,
  handle,
  display_name,
  follower_count,
  following_count,
  post_count
from
  bluesky_user_following
where
  handle = 'matty.wtf'
  and follower_count > 1000;
```

### Get following with specific keywords in description
Find followed users whose descriptions contain specific keywords.

```sql
select
  did,
  handle,
  display_name,
  description
from
  bluesky_user_following
where
  handle = 'matty.wtf'
  and description ilike '%developer%';
```

### Get following with profile images
Find followed users who have both an avatar and banner image.

```sql
select
  did,
  handle,
  display_name,
  avatar,
  banner
from
  bluesky_user_following
where
  handle = 'matty.wtf'
  and avatar is not null
  and banner is not null;
``` 