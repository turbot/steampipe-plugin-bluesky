-- Test: Get all following by DID
select
  did,
  handle,
  display_name,
  description,
  follower_count,
  following_count,
  post_count
from
  bluesky_user_following
where
  target_did = 'did:plc:vipregezugaizr3kfcjijzrv'
limit 20; 