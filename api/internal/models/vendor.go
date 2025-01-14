package models

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/cdriehuys/stuff/api/internal/models/queries"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/leebenson/conform"
)

type NewVendor struct {
	Name string `json:"name" conform:"trim" validate:"required,min=1,max=150"`
}

type Vendor struct {
	ID        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type VendorModel struct {
	logger   *slog.Logger
	queries  *queries.Queries
	validate *validator.Validate
}

func NewVendorModel(logger *slog.Logger, db queries.DBTX, validate *validator.Validate) *VendorModel {
	return &VendorModel{logger, queries.New(db), validate}
}

func (m *VendorModel) Create(ctx context.Context, vendor NewVendor) (Vendor, error) {
	conform.Strings(&vendor)

	if err := m.validate.Struct(vendor); err != nil {
		return Vendor{}, err
	}

	row, err := m.queries.CreateVendor(ctx, vendor.Name)
	if err != nil {
		return Vendor{}, fmt.Errorf("failed to create vendor: %v", err)
	}

	m.logger.InfoContext(ctx, "Created new vendor.", "id", row.ID, "name", row.Name)

	return vendorFromRow(row), nil
}

func (m *VendorModel) DeleteByID(ctx context.Context, id int64) error {
	rows, err := m.queries.DeleteVendorByID(ctx, id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			// If a vendor has models `foreign_key_violation` will be raised if it's deleted.
			if pgErr.Code == "23503" {
				return fmt.Errorf("cannot delete vendor %d: %w", id, ErrVendorHasModels)
			}
		}

		return fmt.Errorf("failed to delete vendor %d: %v", id, err)
	}

	if rows == 0 {
		return fmt.Errorf("no vendor deleted: %w", ErrNotFound)
	}

	m.logger.InfoContext(ctx, "Deleted vendor.", "id", id)

	return nil
}

func (m *VendorModel) GetByID(ctx context.Context, id int64) (Vendor, error) {
	row, err := m.queries.GetVendorByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return Vendor{}, fmt.Errorf("no vendor with ID %d: %w", id, ErrNotFound)
		}

		return Vendor{}, fmt.Errorf("failed to retrive vendor: %v", err)
	}

	return vendorFromRow(row), nil
}

func (m *VendorModel) ListVendors(ctx context.Context) ([]Vendor, error) {
	rows, err := m.queries.ListVendors(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list vendors: %v", err)
	}

	vendors := make([]Vendor, len(rows))
	for i, row := range rows {
		vendors[i] = vendorFromRow(row)
	}

	return vendors, nil
}

func vendorFromRow(row queries.Vendor) Vendor {
	return Vendor{
		ID:        row.ID,
		Name:      row.Name,
		CreatedAt: row.CreatedAt.Time,
		UpdatedAt: row.UpdatedAt.Time,
	}
}
