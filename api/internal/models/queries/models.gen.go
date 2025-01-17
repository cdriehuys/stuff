// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package queries

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Asset struct {
	ID        int64
	ModelID   int64
	Serial    string
	Comments  string
	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
}

type Model struct {
	ID        int64
	Model     string
	VendorID  int64
	Name      string
	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
}

type Vendor struct {
	ID        int64
	Name      string
	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
}
