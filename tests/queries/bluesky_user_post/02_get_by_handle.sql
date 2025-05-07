-- Test: Get all posts by handle
select
  uri,
  text,
  author,
  created_at,
  like_count,
  repost_count,
  hashtags,
  mentioned_handles_names,
  external_links
from
  bluesky_user_post
where
  handle = 'matty.wtf'
limit 20; 