-- Test: Get posts with high engagement
select
  uri,
  text,
  created_at,
  like_count,
  repost_count
from
  bluesky_user_post
where
  handle = 'matty.wtf'
  and like_count > 100
limit 20; 