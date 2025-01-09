package models

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/cdriehuys/stuff/api/internal/api"
	"github.com/cdriehuys/stuff/api/internal/models/queries"
	"github.com/jackc/pgx/v5/pgxpool"
)

type VendorModel struct {
	logger  *slog.Logger
	queries *queries.Queries
}

func NewVendorModel(logger *slog.Logger, db *pgxpool.Pool) *VendorModel {
	return &VendorModel{logger, queries.New(db)}
}

func (m *VendorModel) ListVendors(ctx context.Context) ([]api.Vendor, error) {
	rows, err := m.queries.ListVendors(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list vendors: %v", err)
	}

	vendors := make([]api.Vendor, len(rows))
	for i, row := range rows {
		vendors[i] = api.Vendor{
			Id:        int(row.ID),
			Name:      row.Name,
			CreatedAt: row.CreatedAt.Time,
			UpdatedAt: row.UpdatedAt.Time,
		}
	}

	return vendors, nil
}
