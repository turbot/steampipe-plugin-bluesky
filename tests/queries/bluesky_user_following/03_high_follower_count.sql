-- Test: Get following with high follower counts
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
  and follower_count > 1000
limit 20; 