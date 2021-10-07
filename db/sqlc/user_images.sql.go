// Code generated by sqlc. DO NOT EDIT.
// source: user_images.sql

package db

import (
	"context"
)

const createUserImage = `-- name: CreateUserImage :one

INSERT INTO user_images (
    user_id,
    url,
    type
)
VALUES (
   $1,
   $2,
   $3
    )
    RETURNING
    id, user_id, url, type
`

type CreateUserImageParams struct {
	UserID int32  `json:"user_id"`
	Url    string `json:"url"`
	Type   int32  `json:"type"`
}

// query.sql
func (q *Queries) CreateUserImage(ctx context.Context, arg CreateUserImageParams) (UserImage, error) {
	row := q.db.QueryRowContext(ctx, createUserImage, arg.UserID, arg.Url, arg.Type)
	var i UserImage
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Url,
		&i.Type,
	)
	return i, err
}

const listImage = `-- name: ListImage :many
SELECT id, user_id, url, type FROM user_images
WHERE user_id = $1
`

func (q *Queries) ListImage(ctx context.Context, userID int32) ([]UserImage, error) {
	rows, err := q.db.QueryContext(ctx, listImage, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []UserImage{}
	for rows.Next() {
		var i UserImage
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Url,
			&i.Type,
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
