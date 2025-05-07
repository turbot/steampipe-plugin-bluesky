-- Test: Get mentions with hashtags
select
  m.uri,
  m.text,
  m.author,
  m.created_at,
  m.hashtags
from
  bluesky_user_mention m
  join bluesky_user u on m.target_did = u.did
where
  u.handle = 'matty.wtf'
  and m.hashtags is not null
limit 20; 