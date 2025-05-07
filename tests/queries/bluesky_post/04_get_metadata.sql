-- Test: Get post metadata
select
  uri,
  text,
  hashtags,
  mentioned_handles,
  mentioned_handles_names,
  external_links,
  has_external_links,
  image_count
from
  bluesky_post
where
  uri = 'at://did:plc:example/app.bsky.feed.post/3k2m6q5dpl42g'; 