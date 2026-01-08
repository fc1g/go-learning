-- name: GetPostsForUser :many
SELECT posts.* FROM posts
INNER JOIN feeds ON posts.feed_id = feeds.id
INNER JOIN feed_follow ON feeds.id = feed_follow.feed_id
WHERE feed_follow.user_id = $1
ORDER BY posts.published_at DESC
LIMIT $2;