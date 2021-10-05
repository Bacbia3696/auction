// Code generated by sqlc. DO NOT EDIT.
// source: user_images.sql

package db

import (
	"context"
)

const createUserImage = `-- name: CreateUserImage :one

INSERT INTO user_images (
    UserId,
    Url
)
VALUES (
   $1,
   $2
    )
    RETURNING
    id, userid, url
`

type CreateUserImageParams struct {
	Userid int32  `json:"userid"`
	Url    string `json:"url"`
}

// query.sql
func (q *Queries) CreateUserImage(ctx context.Context, arg CreateUserImageParams) (UserImage, error) {
	row := q.db.QueryRowContext(ctx, createUserImage, arg.Userid, arg.Url)
	var i UserImage
	err := row.Scan(&i.ID, &i.Userid, &i.Url)
	return i, err
}

const listImage = `-- name: ListImage :many
SELECT id, userid, url FROM user_images
WHERE UserId = $1
`

func (q *Queries) ListImage(ctx context.Context, userid int32) ([]UserImage, error) {
	rows, err := q.db.QueryContext(ctx, listImage, userid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []UserImage{}
	for rows.Next() {
		var i UserImage
		if err := rows.Scan(&i.ID, &i.Userid, &i.Url); err != nil {
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
