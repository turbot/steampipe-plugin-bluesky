-- Test: Get follower profile media
select
  handle,
  display_name,
  avatar,
  banner
from
  bluesky_user_follower
where
  target_did = 'did:plc:vipregezugaizr3kfcjijzrv'
limit 20; 