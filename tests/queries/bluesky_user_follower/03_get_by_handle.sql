-- Test: Get followers by handle
select
  f.did,
  f.handle,
  f.display_name,
  f.description
from
  bluesky_user_follower as f
  join bluesky_user as u on f.target_did = u.did
where
  u.handle = 'matty.wtf'
limit 20; 