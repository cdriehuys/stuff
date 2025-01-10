package models

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/cdriehuys/stuff/api/internal/api"
	"github.com/cdriehuys/stuff/api/internal/models/queries"
)

type ModelModel struct {
	logger *slog.Logger
	q      *queries.Queries
}

func NewModelModel(logger *slog.Logger, db queries.DBTX) *ModelModel {
	return &ModelModel{logger: logger, q: queries.New(db)}
}

func (m *ModelModel) ListModels(ctx context.Context) ([]api.Model, error) {
	rows, err := m.q.ListModels(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list models: %v", err)
	}

	models := make([]api.Model, len(rows))
	for i, row := range rows {
		models[i] = api.Model{
			CreatedAt: row.CreatedAt.Time,
			Id:        int(row.ID),
			Model:     row.Model,
			Name:      row.Name,
			UpdatedAt: row.UpdatedAt.Time,
			VendorID:  int(row.VendorID),
		}
	}

	return models, nil
}
