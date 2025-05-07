-- Test: Search for posts with external links
select
  uri,
  http_url,
  text,
  author,
  created_at,
  external_links
from
  bluesky_search_recent
where
  query = '#steampipe'
  and has_external_links = true
limit 20; 