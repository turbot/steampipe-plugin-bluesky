-- Test: Look up a post by HTTP URL (using handle)
select
  uri,
  http_url,
  text,
  author,
  created_at,
  like_count,
  repost_count
from
  bluesky_post
where
  http_url = 'https://bsky.app/profile/example.bsky.social/post/3k2m6q5dpl42g'; 