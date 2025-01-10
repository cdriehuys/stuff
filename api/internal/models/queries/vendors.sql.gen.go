// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: vendors.sql

package queries

import (
	"context"
)

const createVendor = `-- name: CreateVendor :one
INSERT INTO vendors(name)
VALUES ($1)
RETURNING id, name, created_at, updated_at
`

func (q *Queries) CreateVendor(ctx context.Context, name string) (Vendor, error) {
	row := q.db.QueryRow(ctx, createVendor, name)
	var i Vendor
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteVendorByID = `-- name: DeleteVendorByID :execrows
DELETE FROM vendors
WHERE id = $1
`

func (q *Queries) DeleteVendorByID(ctx context.Context, id int64) (int64, error) {
	result, err := q.db.Exec(ctx, deleteVendorByID, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

const getVendorByID = `-- name: GetVendorByID :one
SELECT id, name, created_at, updated_at
FROM vendors
WHERE id = $1
`

func (q *Queries) GetVendorByID(ctx context.Context, id int64) (Vendor, error) {
	row := q.db.QueryRow(ctx, getVendorByID, id)
	var i Vendor
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listVendors = `-- name: ListVendors :many
SELECT id, name, created_at, updated_at FROM vendors ORDER BY id LIMIT 50
`

func (q *Queries) ListVendors(ctx context.Context) ([]Vendor, error) {
	rows, err := q.db.Query(ctx, listVendors)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Vendor
	for rows.Next() {
		var i Vendor
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}