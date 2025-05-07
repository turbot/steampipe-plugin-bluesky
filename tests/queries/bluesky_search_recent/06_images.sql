-- Test: Search for posts with images
select
  uri,
  http_url,
  text,
  author,
  created_at,
  image_count
from
  bluesky_search_recent
where
  query = '#steampipe'
  and image_count > 0
limit 20; 