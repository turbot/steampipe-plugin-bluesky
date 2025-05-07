-- Test: Get mentions with high engagement
select
  m.uri,
  m.text,
  m.author,
  m.created_at,
  m.like_count,
  m.repost_count
from
  bluesky_user_mention m
  join bluesky_user u on m.target_did = u.did
where
  u.handle = 'matty.wtf'
  and m.like_count > 10
limit 20; 