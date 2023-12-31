// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: feeds_follows.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const deleteFollowFeed = `-- name: DeleteFollowFeed :one
DELETE  
FROM feeds_follows
WHERE feeds_follows.user_id = $1
AND feeds_follows.feed_id = $2
RETURNING id, created_at, updated_at, user_id, feed_id
`

type DeleteFollowFeedParams struct {
	UserID uuid.UUID
	FeedID uuid.UUID
}

func (q *Queries) DeleteFollowFeed(ctx context.Context, arg DeleteFollowFeedParams) (FeedsFollow, error) {
	row := q.db.QueryRowContext(ctx, deleteFollowFeed, arg.UserID, arg.FeedID)
	var i FeedsFollow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.FeedID,
	)
	return i, err
}

const deleteFollowFeedById = `-- name: DeleteFollowFeedById :exec
DELETE  
FROM feeds_follows
WHERE feeds_follows.id = $1
AND feeds_follows.user_id = $2
`

type DeleteFollowFeedByIdParams struct {
	ID     uuid.UUID
	UserID uuid.UUID
}

func (q *Queries) DeleteFollowFeedById(ctx context.Context, arg DeleteFollowFeedByIdParams) error {
	_, err := q.db.ExecContext(ctx, deleteFollowFeedById, arg.ID, arg.UserID)
	return err
}

const followFeed = `-- name: FollowFeed :one
INSERT INTO feeds_follows(id, created_at, updated_at, user_id, feed_id) 
VALUES ($1,$2,$3,$4,$5)
RETURNING id, created_at, updated_at, user_id, feed_id
`

type FollowFeedParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uuid.UUID
	FeedID    uuid.UUID
}

func (q *Queries) FollowFeed(ctx context.Context, arg FollowFeedParams) (FeedsFollow, error) {
	row := q.db.QueryRowContext(ctx, followFeed,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.UserID,
		arg.FeedID,
	)
	var i FeedsFollow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.FeedID,
	)
	return i, err
}

const getFollowFeed = `-- name: GetFollowFeed :many
SELECT id, created_at, updated_at, user_id, feed_id  
FROM feeds_follows
WHERE feeds_follows.user_id = $1
`

func (q *Queries) GetFollowFeed(ctx context.Context, userID uuid.UUID) ([]FeedsFollow, error) {
	rows, err := q.db.QueryContext(ctx, getFollowFeed, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FeedsFollow
	for rows.Next() {
		var i FeedsFollow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UserID,
			&i.FeedID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
