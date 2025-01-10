// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: models.sql

package queries

import (
	"context"
)

const listModels = `-- name: ListModels :many
SELECT id, model, vendor_id, name, created_at, updated_at FROM models ORDER BY id LIMIT 50
`

func (q *Queries) ListModels(ctx context.Context) ([]Model, error) {
	rows, err := q.db.Query(ctx, listModels)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Model
	for rows.Next() {
		var i Model
		if err := rows.Scan(
			&i.ID,
			&i.Model,
			&i.VendorID,
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