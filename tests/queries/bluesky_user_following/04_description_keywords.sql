-- Test: Get following with specific keywords in description
select
  did,
  handle,
  display_name,
  description
from
  bluesky_user_following
where
  handle = 'matty.wtf'
  and description ilike '%developer%'
limit 20; 