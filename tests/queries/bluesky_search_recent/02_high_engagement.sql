-- Test: Search for posts with high engagement
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
  and like_count > 10
limit 20; 