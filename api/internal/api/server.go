package api

import (
	"context"
	"log/slog"

	"github.com/cdriehuys/stuff/api/internal/models"
	"github.com/go-playground/validator/v10"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

// ensure we're implementing the generated interface
var _ StrictServerInterface = (*Server)(nil)

type assetModel interface {
	Create(ctx context.Context, asset models.NewAsset) (models.Asset, error)
	DeleteByID(ctx context.Context, id int64) error
	GetByID(ctx context.Context, id int64) (models.Asset, error)
	List(ctx context.Context) ([]models.Asset, error)
	UpdateByID(ctx context.Context, id int64, updated models.NewAsset) (models.Asset, error)
}

type modelModel interface {
	Create(ctx context.Context, model models.NewModel) (models.Model, error)
	DeleteByID(ctx context.Context, id int64) error
	GetByID(ctx context.Context, id int64) (models.Model, error)
	ListByVendorID(ctx context.Context, vendorID int64) ([]models.Model, error)
	ListModels(ctx context.Context) ([]models.Model, error)
	UpdateByID(ctx context.Context, id int64, model models.NewModel) (models.Model, error)
}

type vendorModel interface {
	Create(ctx context.Context, vendor models.NewVendor) (models.Vendor, error)
	DeleteByID(ctx context.Context, id int64) error
	GetByID(ctx context.Context, id int64) (models.Vendor, error)
	ListVendors(ctx context.Context) ([]models.Vendor, error)
}

type Server struct {
	logger *slog.Logger
	bundle *i18n.Bundle

	assets  assetModel
	models  modelModel
	vendors vendorModel

	validate *validator.Validate
}

func NewServer(
	logger *slog.Logger,
	bundle *i18n.Bundle,
	validate *validator.Validate,
	assets assetModel,
	models modelModel,
	vendors vendorModel,
) *Server {
	return &Server{
		logger: logger,
		bundle: bundle,

		assets:  assets,
		models:  models,
		vendors: vendors,

		validate: validate,
	}
}
