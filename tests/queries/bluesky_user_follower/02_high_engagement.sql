-- Test: Get followers with high engagement
select
  handle,
  display_name,
  follower_count,
  following_count,
  post_count
from
  bluesky_user_follower
where
  target_did = 'did:plc:vipregezugaizr3kfcjijzrv'
  and follower_count > 1000
limit 20; 