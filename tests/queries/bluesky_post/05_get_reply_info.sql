-- Test: Get reply information
select
  uri,
  text,
  reply_root,
  reply_parent,
  created_at
from
  bluesky_post
where
  uri = 'at://did:plc:example/app.bsky.feed.post/3k2m6q5dpl42g'; 