---
title: "Steampipe Table: bluesky_user_follower - Query Bluesky User Followers using SQL"
description: "Allows users to query followers of a Bluesky user, providing insights into follower profiles, engagement metrics, and relationship details."
folder: "Followers"
---

# Table: bluesky_user_follower - Query Bluesky User Followers using SQL

Bluesky is a decentralized social network protocol that allows users to create and share content. The `bluesky_user_follower` table provides access to the followers of a specific Bluesky user, including follower profiles, engagement metrics, and relationship details.

## Table Usage Guide

The `bluesky_user_follower` table provides insights into the followers of a specific Bluesky user. As a data analyst or social media manager, explore follower-specific details through this table, including profile information, engagement metrics, and relationship details. Utilize it to uncover information about follower demographics, engagement patterns, and network growth.


**Important Notes**

- The `target_did` field must be set in the `where` clause
- The DID must be in the format `did:plc:...` or `did:web:...`
- To query by handle, use a join with the `bluesky_user` table
- The table provides comprehensive follower profile information including engagement metrics and media URLs

## Examples

### Get all followers of a user
List all followers of a specific user, including their profile information.

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
  bluesky_user_follower
where
  target_did = 'did:plc:vipregezugaizr3kfcjijzrv';
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
  bluesky_user_follower
where
  target_did = 'did:plc:vipregezugaizr3kfcjijzrv';
```

### Get followers with high engagement
Find followers who have a significant number of followers themselves, indicating potential influence.

```sql+postgres
select
  handle,
  display_name,
  follower_count,
  following_count,
  post_count
from
  bluesky_user_follower
where
  target_did = 'did:plc:vipregezugaizr3kfcjijzrv'
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
  bluesky_user_follower
where
  target_did = 'did:plc:vipregezugaizr3kfcjijzrv'
  and follower_count > 1000;
```

### Get followers by handle
Look up followers of a user by their handle instead of DID.

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
  bluesky_user_follower f
  join bluesky_user u on f.target_did = u.did
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
  bluesky_user_follower f
  join bluesky_user u on f.target_did = u.did
where
  u.handle = 'matty.wtf';
```

### Get follower profile media
Retrieve profile media (avatar and banner) for all followers of a user.

```sql+postgres
select
  handle,
  display_name,
  avatar,
  banner
from
  bluesky_user_follower
where
  target_did = 'did:plc:vipregezugaizr3kfcjijzrv';
```

```sql+sqlite
select
  handle,
  display_name,
  avatar,
  banner
from
  bluesky_user_follower
where
  target_did = 'did:plc:vipregezugaizr3kfcjijzrv';
``` 