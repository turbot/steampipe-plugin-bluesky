-- Test: Get mentions with external links
select
  m.uri,
  m.text,
  m.author,
  m.created_at,
  m.external_links
from
  bluesky_user_mention m
  join bluesky_user u on m.target_did = u.did
where
  u.handle = 'matty.wtf'
  and m.has_external_links = true
limit 20; 