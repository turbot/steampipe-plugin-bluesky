-- Test: Search for posts mentioning a specific user
select
  uri,
  text,
  author,
  created_at,
  mentioned_handles
from
  bluesky_search_recent
where
  query = '@mattstratton'
limit 20; 