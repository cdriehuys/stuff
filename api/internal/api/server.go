package api

import (
	"context"
	"log/slog"
	"reflect"
	"strings"

	"github.com/cdriehuys/stuff/api/internal/models"
	"github.com/go-playground/validator/v10"
)

// ensure we're implementing the generated interface
var _ StrictServerInterface = (*Server)(nil)

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
	logger  *slog.Logger
	models  modelModel
	vendors vendorModel

	validate *validator.Validate
}

func NewServer(logger *slog.Logger, models modelModel, vendors vendorModel) *Server {
	validate := validator.New(validator.WithRequiredStructEnabled())
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		// skip if tag key says it should be ignored
		if name == "-" {
			return ""
		}
		return name
	})

	return &Server{
		logger:  logger,
		models:  models,
		vendors: vendors,

		validate: validate,
	}
}
