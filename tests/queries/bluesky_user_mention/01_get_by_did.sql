-- Test: Get all mentions by DID
select
  uri,
  text,
  author,
  created_at,
  like_count,
  repost_count,
  hashtags,
  mentioned_handles_names,
  external_links
from
  bluesky_user_mention
where
  target_did = 'did:plc:vipregezugaizr3kfcjijzrv'
limit 20; 