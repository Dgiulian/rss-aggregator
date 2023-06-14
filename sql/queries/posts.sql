-- name: CreatePost :one

INSERT INTO posts (id, created_at, updated_at, title, description, published_at, url, feed_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;


-- name: GetFeedPosts :many

SELECT * FROM posts WHERE feed_id = $1;

-- name: GetPostsForUser :many

SELECT posts.* 
FROM posts 
JOIN  feeds_follows on posts.feed_id = feeds_follows.feed_id 
WHERE user_id = $1;