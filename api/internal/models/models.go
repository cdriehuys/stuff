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

type NewModel struct {
	VendorID int64  `json:"vendorID"`
	Model    string `json:"model"    conform:"trim" validate:"required,min=1,max=150"`
	Name     string `json:"name"     conform:"trim" validate:"max=150"`
}

type Model struct {
	ID        int64
	VendorID  int64
	Model     string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ModelModel struct {
	logger   *slog.Logger
	q        *queries.Queries
	validate *validator.Validate
}

func NewModelModel(logger *slog.Logger, db queries.DBTX, validate *validator.Validate) *ModelModel {
	return &ModelModel{logger: logger, q: queries.New(db), validate: validate}
}

func (m *ModelModel) Create(ctx context.Context, model NewModel) (Model, error) {
	conform.Strings(&model)

	if err := m.validate.Struct(model); err != nil {
		return Model{}, err
	}

	args := queries.CreateModelParams{
		Model:    model.Model,
		VendorID: model.VendorID,
		Name:     model.Name,
	}

	row, err := m.q.CreateModel(ctx, args)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			// foreign_key_violation
			if pgErr.Code == "23503" {
				return Model{}, fmt.Errorf("failed to add model for vendor %d: %w", model.VendorID, ErrVendorNotFound)
			}

			// unique_violation
			if pgErr.Code == "23505" {
				return Model{}, fmt.Errorf("cannot add model %q to vendor %d: %w", model.Model, model.VendorID, ErrModelNotUnique)
			}
		}

		return Model{}, fmt.Errorf("failed to create model: %v", err)
	}

	m.logger.InfoContext(ctx, "Created model.", "modelID", row.ID)

	return modelFromRow(row), nil
}

func (m *ModelModel) DeleteByID(ctx context.Context, id int64) error {
	rows, err := m.q.DeleteModelByID(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete model %d: %v", id, err)
	}

	if rows == 0 {
		return fmt.Errorf("no model deleted: %w", ErrNotFound)
	}

	m.logger.InfoContext(ctx, "Deleted model.", "id", id)

	return nil
}

func (m *ModelModel) GetByID(ctx context.Context, id int64) (Model, error) {
	row, err := m.q.GetModelByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return Model{}, fmt.Errorf("no model with ID %d: %w", id, ErrNotFound)
		}

		return Model{}, fmt.Errorf("failed to retrieve model: %v", err)
	}

	return modelFromRow(row), nil
}

func (m *ModelModel) ListByVendorID(ctx context.Context, vendorID int64) ([]Model, error) {
	rows, err := m.q.ListModelsByVendorID(ctx, vendorID)
	if err != nil {
		return nil, fmt.Errorf("failed to list models: %v", err)
	}

	models := make([]Model, len(rows))
	for i, row := range rows {
		models[i] = modelFromRow(row)
	}

	return models, nil
}

func (m *ModelModel) ListModels(ctx context.Context) ([]Model, error) {
	rows, err := m.q.ListModels(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list models: %v", err)
	}

	models := make([]Model, len(rows))
	for i, row := range rows {
		models[i] = modelFromRow(row)
	}

	return models, nil
}

func (m *ModelModel) UpdateByID(ctx context.Context, id int64, model NewModel) (Model, error) {
	conform.Strings(&model)

	if err := m.validate.Struct(model); err != nil {
		return Model{}, err
	}

	args := queries.UpdateModelByIDParams{
		ID:       id,
		Model:    model.Model,
		Name:     model.Name,
		VendorID: model.VendorID,
	}

	row, err := m.q.UpdateModelByID(ctx, args)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return Model{}, fmt.Errorf("no model with ID %d: %w", id, ErrNotFound)
		}

		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			// foreign_key_violation
			if pgErr.Code == "23503" {
				return Model{}, fmt.Errorf("failed to add model for vendor %d: %w", model.VendorID, ErrVendorNotFound)
			}

			// unique_violation
			if pgErr.Code == "23505" {
				return Model{}, fmt.Errorf("cannot add model %q to vendor %d: %w", model.Model, model.VendorID, ErrModelNotUnique)
			}
		}

		return Model{}, fmt.Errorf("failed to update model: %v", err)
	}

	return modelFromRow(row), nil
}

func modelFromRow(row queries.Model) Model {
	return Model{
		CreatedAt: row.CreatedAt.Time,
		ID:        row.ID,
		Model:     row.Model,
		Name:      row.Name,
		UpdatedAt: row.UpdatedAt.Time,
		VendorID:  row.VendorID,
	}
}
