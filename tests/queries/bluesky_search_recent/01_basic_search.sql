-- Test: Basic search for posts containing a specific term
select
  uri,
  text,
  author,
  created_at,
  like_count,
  repost_count
from
  bluesky_search_recent
where
  query = 'steampipe'
limit 20; 