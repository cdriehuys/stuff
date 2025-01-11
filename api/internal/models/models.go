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
	"github.com/jackc/pgx/v5/pgconn"
)

type ModelModel struct {
	logger *slog.Logger
	q      *queries.Queries
}

func NewModelModel(logger *slog.Logger, db queries.DBTX) *ModelModel {
	return &ModelModel{logger: logger, q: queries.New(db)}
}

func (m *ModelModel) Create(ctx context.Context, vendorID int64, model api.NewModel) (api.Model, error) {
	args := queries.CreateModelParams{
		Model:    model.Model,
		VendorID: vendorID,
	}
	if model.Name != nil {
		args.Name = *model.Name
	}

	row, err := m.q.CreateModel(ctx, args)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			// foreign_key_violation
			if pgErr.Code == "23503" {
				return api.Model{}, fmt.Errorf("failed to add model for vendor %d: %w", vendorID, apierrors.ErrNotFound)
			}
		}

		return api.Model{}, fmt.Errorf("failed to create model: %v", err)
	}

	m.logger.InfoContext(ctx, "Created model for vendor.", "vendorID", vendorID, "modelID", row.ID)

	return modelFromRow(row), nil
}

func (m *ModelModel) DeleteByID(ctx context.Context, id int64) error {
	rows, err := m.q.DeleteModelByID(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete model %d: %v", id, err)
	}

	if rows == 0 {
		return fmt.Errorf("no model deleted: %w", apierrors.ErrNotFound)
	}

	m.logger.InfoContext(ctx, "Deleted model.", "id", id)

	return nil
}

func (m *ModelModel) GetByID(ctx context.Context, id int64) (api.Model, error) {
	row, err := m.q.GetModelByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return api.Model{}, fmt.Errorf("no model with ID %d: %w", id, apierrors.ErrNotFound)
		}

		return api.Model{}, fmt.Errorf("failed to retrieve model: %v", err)
	}

	return modelFromRow(row), nil
}

func (m *ModelModel) ListByVendorID(ctx context.Context, vendorID int64) ([]api.Model, error) {
	rows, err := m.q.ListModelsByVendorID(ctx, vendorID)
	if err != nil {
		return nil, fmt.Errorf("failed to list models: %v", err)
	}

	models := make([]api.Model, len(rows))
	for i, row := range rows {
		models[i] = modelFromRow(row)
	}

	return models, nil
}

func (m *ModelModel) ListModels(ctx context.Context) ([]api.Model, error) {
	rows, err := m.q.ListModels(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list models: %v", err)
	}

	models := make([]api.Model, len(rows))
	for i, row := range rows {
		models[i] = modelFromRow(row)
	}

	return models, nil
}

func (m *ModelModel) UpdateByID(ctx context.Context, id int64, model api.NewModel) (api.Model, error) {
	args := queries.UpdateModelByIDParams{ID: id, Model: model.Model}
	if model.Name != nil {
		args.Name = *model.Name
	}

	row, err := m.q.UpdateModelByID(ctx, args)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return api.Model{}, fmt.Errorf("no model with ID %d: %w", id, apierrors.ErrNotFound)
		}

		return api.Model{}, fmt.Errorf("failed to update model: %v", err)
	}

	return modelFromRow(row), nil
}

func modelFromRow(row queries.Model) api.Model {
	return api.Model{
		CreatedAt: row.CreatedAt.Time,
		Id:        int(row.ID),
		Model:     row.Model,
		Name:      row.Name,
		UpdatedAt: row.UpdatedAt.Time,
		VendorID:  int(row.VendorID),
	}
}
