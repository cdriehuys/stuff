package models

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/cdriehuys/stuff/api/internal/api"
	"github.com/cdriehuys/stuff/api/internal/apierrors"
	"github.com/cdriehuys/stuff/api/internal/models/queries"
	"github.com/jackc/pgx/v5"
)

type VendorModel struct {
	logger  *slog.Logger
	queries *queries.Queries
}

func NewVendorModel(logger *slog.Logger, db queries.DBTX) *VendorModel {
	return &VendorModel{logger, queries.New(db)}
}

func (m *VendorModel) Create(ctx context.Context, vendor api.NewVendor) (api.Vendor, error) {
	row, err := m.queries.CreateVendor(ctx, vendor.Name)
	if err != nil {
		return api.Vendor{}, fmt.Errorf("failed to create vendor: %v", err)
	}

	m.logger.InfoContext(ctx, "Created new vendor.", "id", row.ID, "name", row.Name)

	return vendorFromRow(row), nil
}

func (m *VendorModel) DeleteByID(ctx context.Context, id int64) error {
	rows, err := m.queries.DeleteVendorByID(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete vendor %d: %v", id, err)
	}

	if rows == 0 {
		return fmt.Errorf("no vendor deleted: %w", apierrors.ErrNotFound)
	}

	m.logger.InfoContext(ctx, "Deleted vendor.", "id", id)

	return nil
}

func (m *VendorModel) GetByID(ctx context.Context, id int64) (api.Vendor, error) {
	row, err := m.queries.GetVendorByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return api.Vendor{}, fmt.Errorf("no vendor with ID %d: %w", id, apierrors.ErrNotFound)
		}

		return api.Vendor{}, fmt.Errorf("failed to retrive vendor: %v", err)
	}

	return vendorFromRow(row), nil
}

func (m *VendorModel) ListVendors(ctx context.Context) ([]api.Vendor, error) {
	rows, err := m.queries.ListVendors(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list vendors: %v", err)
	}

	vendors := make([]api.Vendor, len(rows))
	for i, row := range rows {
		vendors[i] = vendorFromRow(row)
	}

	return vendors, nil
}

func vendorFromRow(row queries.Vendor) api.Vendor {
	return api.Vendor{
		Id:        int(row.ID),
		Name:      row.Name,
		CreatedAt: row.CreatedAt.Time,
		UpdatedAt: row.UpdatedAt.Time,
	}
}
