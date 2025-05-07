-- Test: Get all mentions by handle
select
  m.uri,
  m.text,
  m.author,
  m.created_at,
  m.like_count,
  m.repost_count,
  m.hashtags,
  m.mentioned_handles_names,
  m.external_links
from
  bluesky_user_mention m
  join bluesky_user u on m.target_did = u.did
where
  u.handle = 'matty.wtf'
limit 20; 