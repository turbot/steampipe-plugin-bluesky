-- Test: Get user profile information by handle
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