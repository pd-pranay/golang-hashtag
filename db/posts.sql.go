// Code generated by sqlc. DO NOT EDIT.
// source: posts.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

const createPost = `-- name: CreatePost :one
INSERT INTO posts (post_name, post_data, tags) VALUES ($1, $2, $3) RETURNING id, post_name, post_data, tags
`

type CreatePostParams struct {
	PostName string   `json:"post_name"`
	PostData string   `json:"post_data"`
	Tags     []string `json:"tags"`
}

type CreatePostRow struct {
	ID       uuid.UUID `json:"id"`
	PostName string    `json:"post_name"`
	PostData string    `json:"post_data"`
	Tags     []string  `json:"tags"`
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (CreatePostRow, error) {
	row := q.queryRow(ctx, q.createPostStmt, createPost, arg.PostName, arg.PostData, pq.Array(arg.Tags))
	var i CreatePostRow
	err := row.Scan(
		&i.ID,
		&i.PostName,
		&i.PostData,
		pq.Array(&i.Tags),
	)
	return i, err
}

const getAllPosts = `-- name: GetAllPosts :many
SELECT id, post_name, post_data, tags FROM posts
`

type GetAllPostsRow struct {
	ID       uuid.UUID `json:"id"`
	PostName string    `json:"post_name"`
	PostData string    `json:"post_data"`
	Tags     []string  `json:"tags"`
}

func (q *Queries) GetAllPosts(ctx context.Context) ([]GetAllPostsRow, error) {
	rows, err := q.query(ctx, q.getAllPostsStmt, getAllPosts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllPostsRow
	for rows.Next() {
		var i GetAllPostsRow
		if err := rows.Scan(
			&i.ID,
			&i.PostName,
			&i.PostData,
			pq.Array(&i.Tags),
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

const getPostByID = `-- name: GetPostByID :one
SELECT id, post_name, post_data, tags FROM posts WHERE id = ($1)
`

type GetPostByIDRow struct {
	ID       uuid.UUID `json:"id"`
	PostName string    `json:"post_name"`
	PostData string    `json:"post_data"`
	Tags     []string  `json:"tags"`
}

func (q *Queries) GetPostByID(ctx context.Context, id uuid.UUID) (GetPostByIDRow, error) {
	row := q.queryRow(ctx, q.getPostByIDStmt, getPostByID, id)
	var i GetPostByIDRow
	err := row.Scan(
		&i.ID,
		&i.PostName,
		&i.PostData,
		pq.Array(&i.Tags),
	)
	return i, err
}

const getPostsByTag = `-- name: GetPostsByTag :many
SELECT id, post_name, post_data, tags FROM posts WHERE tags @> ($1)
`

type GetPostsByTagRow struct {
	ID       uuid.UUID `json:"id"`
	PostName string    `json:"post_name"`
	PostData string    `json:"post_data"`
	Tags     []string  `json:"tags"`
}

func (q *Queries) GetPostsByTag(ctx context.Context, tags []string) ([]GetPostsByTagRow, error) {
	rows, err := q.query(ctx, q.getPostsByTagStmt, getPostsByTag, pq.Array(tags))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPostsByTagRow
	for rows.Next() {
		var i GetPostsByTagRow
		if err := rows.Scan(
			&i.ID,
			&i.PostName,
			&i.PostData,
			pq.Array(&i.Tags),
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
