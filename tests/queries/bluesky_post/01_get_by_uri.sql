-- Test: Look up a post by URI
select
  uri,
  http_url,
  text,
  author,
  created_at,
  like_count,
  repost_count
from
  bluesky_post
where
  uri = 'at://did:plc:example/app.bsky.feed.post/3k2m6q5dpl42g'; 