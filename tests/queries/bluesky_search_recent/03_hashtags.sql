-- Test: Search for posts with specific hashtags
select
  uri,
  text,
  author,
  created_at,
  hashtags
from
  bluesky_search_recent
where
  query = '#steampipe'
limit 20; 