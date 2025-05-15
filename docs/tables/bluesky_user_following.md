---
title: "Steampipe Table: bluesky_user_following - Query Bluesky User Following using SQL"
description: "Allows users to query accounts that a Bluesky user is following, providing insights into following profiles, engagement metrics, and relationship details."
folder: "Following"
---

# Table: bluesky_user_following - Query Bluesky User Following using SQL

Bluesky is a decentralized social network protocol that allows users to create and share content. The `bluesky_user_following` table provides access to the accounts that a specific Bluesky user is following, including profile information, engagement metrics, and relationship details.

## Table Usage Guide

The `bluesky_user_following` table provides insights into the accounts that a specific Bluesky user is following. As a data analyst or social media manager, explore following-specific details through this table, including profile information, engagement metrics, and relationship details. Utilize it to uncover information about following demographics, engagement patterns, and network growth.

**Important Notes**

- The `did` field must be set in the `where` clause
- The DID must be in the format `did:plc:...` or `did:web:...`
- To query by handle, use a join with the `bluesky_user` table
- The table provides comprehensive following profile information including engagement metrics and media URLs

## Examples

### Get all accounts a user is following
List all accounts that a specific user is following, including their profile information.

```sql+postgres
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
  did = 'did:plc:vipregezugaizr3kfcjijzrv';
```

```sql+sqlite
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
  did = 'did:plc:vipregezugaizr3kfcjijzrv';
```

### Get following with high engagement
Find accounts the user is following who have a significant number of followers, indicating potential influence.

```sql+postgres
select
  handle,
  display_name,
  follower_count,
  following_count,
  post_count
from
  bluesky_user_following
where
  did = 'did:plc:vipregezugaizr3kfcjijzrv'
  and follower_count > 1000;
```

```sql+sqlite
select
  handle,
  display_name,
  follower_count,
  following_count,
  post_count
from
  bluesky_user_following
where
  did = 'did:plc:vipregezugaizr3kfcjijzrv'
  and follower_count > 1000;
```

### Get following by handle
Look up accounts a user is following by their handle instead of DID.

```sql+postgres
select
  f.did,
  f.handle,
  f.display_name,
  f.description,
  f.follower_count,
  f.following_count,
  f.post_count
from
  bluesky_user_following f
  join bluesky_user u on f.did = u.did
where
  u.handle = 'matty.wtf';
```

```sql+sqlite
select
  f.did,
  f.handle,
  f.display_name,
  f.description,
  f.follower_count,
  f.following_count,
  f.post_count
from
  bluesky_user_following f
  join bluesky_user u on f.did = u.did
where
  u.handle = 'matty.wtf';
``` 