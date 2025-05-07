-- Test: Search for posts mentioning a user with handle names
select
  uri,
  http_url,
  text,
  author,
  created_at,
  like_count,
  repost_count,
  mentioned_handles_names
from
  bluesky_search_recent
where
  query = '@matty.wtf'
limit 20; 