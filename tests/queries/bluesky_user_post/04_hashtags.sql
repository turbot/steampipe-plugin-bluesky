-- Test: Get posts with hashtags
select
  uri,
  text,
  created_at,
  hashtags
from
  bluesky_user_post
where
  handle = 'matty.wtf'
  and hashtags is not null
limit 20; 