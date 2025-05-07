-- Test: Search for posts with hashtag and engagement metrics
select
  uri,
  http_url,
  text,
  author,
  created_at,
  like_count,
  repost_count,
  hashtags
from
  bluesky_search_recent
where
  query = '#steampipe'
limit 20; 