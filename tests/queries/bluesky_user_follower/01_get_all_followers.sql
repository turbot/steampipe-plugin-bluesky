-- Test: Get all followers of a user
select
  did,
  handle,
  display_name,
  description,
  follower_count,
  following_count,
  post_count
from
  bluesky_user_follower
where
  target_did = 'did:plc:vipregezugaizr3kfcjijzrv'
limit 20; 