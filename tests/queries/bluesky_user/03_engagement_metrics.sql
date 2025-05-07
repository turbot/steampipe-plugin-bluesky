-- Test: Get user engagement metrics
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