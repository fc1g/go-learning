-- name: GetFeedFollowForUser :many
SELECT
    feed_follow.*,
    users.name AS user_name,
    feeds.name AS feed_name
FROM feed_follow
INNER JOIN users ON users.id = feed_follow.user_id
INNER JOIN feeds ON feeds.id = feed_follow.feed_id
WHERE feed_follow.user_id = $1;