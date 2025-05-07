-- Test: Get all following by handle
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
  handle = 'matty.wtf'
limit 20; 