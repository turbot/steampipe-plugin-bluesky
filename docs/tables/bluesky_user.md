---
title: "Steampipe Table: bluesky_user - Query Bluesky Users using SQL"
description: "Allows users to query Bluesky user profiles, providing insights into user information, engagement metrics, and profile details."
folder: "User"
---

# Table: bluesky_user - Query Bluesky Users using SQL

Bluesky is a decentralized social network protocol that allows users to create and share content. The `bluesky_user` table provides access to user profiles on Bluesky, including profile information, engagement metrics, and account details.

## Table Usage Guide

The `bluesky_user` table provides insights into Bluesky user profiles. As a data analyst or social media manager, explore user-specific details through this table, including profile information, engagement metrics, and account details. Utilize it to uncover information about user activity, influence, and engagement patterns.

**Important Notes**

- Either `did` or `handle` must be specified in the `where` clause
- If using `did`, it must be in the format `did:plc:...` or `did:web:...`
- If using `handle`, it can be provided with or without the `@` prefix
- The table provides comprehensive user profile information including engagement metrics and media URLs

## Examples

### Get user profile information by DID
Look up a specific user's profile information using their DID.

```sql
select
  did,
  handle,
  display_name,
  description,
  follower_count,
  following_count,
  post_count,
  avatar,
  banner
from
  bluesky_user
where
  did = 'did:plc:vipregezugaizr3kfcjijzrv';
```

### Get user profile information by handle
Look up a specific user's profile information using their handle.

```sql
select
  did,
  handle,
  display_name,
  description,
  follower_count,
  following_count,
  post_count,
  avatar,
  banner
from
  bluesky_user
where
  handle = 'matty.wtf';
```

### Get user engagement metrics
Analyze a user's engagement metrics, including follower count, following count, and post count.

```sql
select
  handle,
  display_name,
  follower_count,
  following_count,
  post_count
from
  bluesky_user
where
  handle = 'matty.wtf';
```

### Get user profile media
Retrieve a user's profile media, including avatar and banner images.

```sql
select
  handle,
  display_name,
  avatar,
  banner
from
  bluesky_user
where
  handle = 'matty.wtf';
``` 