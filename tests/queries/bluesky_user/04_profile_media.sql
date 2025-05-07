-- Test: Get user profile media
select
  handle,
  display_name,
  avatar,
  banner
from
  bluesky_user
where
  handle = 'matty.wtf'; 