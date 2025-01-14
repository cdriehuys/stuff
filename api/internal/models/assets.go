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

type NewAsset struct {
	ModelID  int64  `json:"modelID"`
	Serial   string `json:"serial"   conform:"trim" validate:"required,min=1,max=150"`
	Comments string `json:"comments" conform:"trim" validate:"max=1000"`
}

type Asset struct {
	ID        int64
	ModelID   int64
	Serial    string
	Comments  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AssetModel struct {
	logger   *slog.Logger
	q        *queries.Queries
	validate *validator.Validate
}

func NewAssetModel(logger *slog.Logger, db queries.DBTX, validate *validator.Validate) *AssetModel {
	return &AssetModel{
		logger:   logger,
		q:        queries.New(db),
		validate: validate,
	}
}

func (m *AssetModel) Create(ctx context.Context, asset NewAsset) (Asset, error) {
	conform.Strings(&asset)

	if err := m.validate.Struct(asset); err != nil {
		return Asset{}, err
	}

	args := queries.CreateAssetParams{
		ModelID:  asset.ModelID,
		Serial:   asset.Serial,
		Comments: asset.Comments,
	}

	row, err := m.q.CreateAsset(ctx, args)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			// foreign_key_violation
			if pgErr.Code == "23503" {
				return Asset{}, fmt.Errorf("failed to add asset: %w", ErrModelNotFound)
			}

			// unique_violation
			if pgErr.Code == "23505" {
				return Asset{}, fmt.Errorf("cannot add asset: %w", ErrAssetSerialNotUnique)
			}
		}

		return Asset{}, fmt.Errorf("failed to create asset: %v", err)
	}

	m.logger.InfoContext(ctx, "Created asset.", "assetID", row.ID)

	return assetFromRow(row), nil
}

func (m *AssetModel) DeleteByID(ctx context.Context, id int64) error {
	rows, err := m.q.DeleteAssetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete asset %d: %v", id, err)
	}

	if rows == 0 {
		return fmt.Errorf("no asset deleted: %w", ErrAssetNotFound)
	}

	m.logger.InfoContext(ctx, "Deleted asset.", "id", id)

	return nil
}

func (m *AssetModel) GetByID(ctx context.Context, id int64) (Asset, error) {
	row, err := m.q.GetAssetByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return Asset{}, fmt.Errorf("no asset with ID %d: %w", id, ErrAssetNotFound)
		}

		return Asset{}, fmt.Errorf("failed to retrieve asset %d: %v", id, err)
	}

	return assetFromRow(row), nil
}

func (m *AssetModel) List(ctx context.Context) ([]Asset, error) {
	rows, err := m.q.ListAssets(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list Assets: %v", err)
	}

	Assets := make([]Asset, len(rows))
	for i, row := range rows {
		Assets[i] = assetFromRow(row)
	}

	return Assets, nil
}

func (m *AssetModel) UpdateByID(ctx context.Context, id int64, asset NewAsset) (Asset, error) {
	conform.Strings(&asset)

	if err := m.validate.Struct(asset); err != nil {
		return Asset{}, err
	}

	args := queries.UpdateAssetByIDParams{
		ModelID:  asset.ModelID,
		Serial:   asset.Serial,
		Comments: asset.Comments,
		AssetID:  id,
	}

	row, err := m.q.UpdateAssetByID(ctx, args)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return Asset{}, fmt.Errorf("no asset with ID %d: %w", id, ErrAssetNotFound)
		}

		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			// foreign_key_violation
			if pgErr.Code == "23503" {
				return Asset{}, fmt.Errorf("failed to update asset for model %d: %w", asset.ModelID, ErrVendorNotFound)
			}

			// unique_violation
			if pgErr.Code == "23505" {
				return Asset{}, fmt.Errorf("cannot add serial %q under model %d: %w", asset.Serial, asset.ModelID, ErrAssetSerialNotUnique)
			}
		}

		return Asset{}, fmt.Errorf("failed to update Asset: %v", err)
	}

	return assetFromRow(row), nil
}

func assetFromRow(row queries.Asset) Asset {
	return Asset{
		Comments:  row.Comments,
		CreatedAt: row.CreatedAt.Time,
		ID:        row.ID,
		ModelID:   row.ModelID,
		Serial:    row.Serial,
		UpdatedAt: row.UpdatedAt.Time,
	}
}
