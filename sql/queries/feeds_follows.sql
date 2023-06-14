-- name: FollowFeed :one
INSERT INTO feeds_follows(id, created_at, updated_at, user_id, feed_id) 
VALUES ($1,$2,$3,$4,$5)
RETURNING *;


-- name: GetFollowFeed :many

SELECT *  
FROM feeds_follows
WHERE feeds_follows.user_id = $1;