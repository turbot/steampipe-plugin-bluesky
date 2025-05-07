-- Test: Get following with profile images
select
  did,
  handle,
  display_name,
  avatar,
  banner
from
  bluesky_user_following
where
  handle = 'matty.wtf'
  and avatar is not null
  and banner is not null
limit 20; 