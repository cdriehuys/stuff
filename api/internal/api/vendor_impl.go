package api

import (
	"context"
	"errors"

	"github.com/cdriehuys/stuff/api/internal/models"
	"github.com/go-playground/validator/v10"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func (s *Server) GetVendors(ctx context.Context, req GetVendorsRequestObject) (GetVendorsResponseObject, error) {
	vendors, err := s.vendors.ListVendors(ctx)
	if err != nil {
		return GetVendors500JSONResponse{}, err
	}

	return GetVendors200JSONResponse(vendorCollection(vendors)), nil
}

func (s *Server) PostVendors(ctx context.Context, req PostVendorsRequestObject) (PostVendorsResponseObject, error) {
	newVendor := internalNewVendor(*req.Body)

	vendor, err := s.vendors.Create(ctx, newVendor)
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			return PostVendors400JSONResponse{InvalidRequestJSONResponse(s.validationError(ctx, ve))}, nil
		}

		return PostVendors500JSONResponse{}, err
	}

	return PostVendors201JSONResponse(externalVendor(vendor)), nil
}

func (s *Server) GetVendorsVendorID(ctx context.Context, req GetVendorsVendorIDRequestObject) (GetVendorsVendorIDResponseObject, error) {
	vendor, err := s.vendors.GetByID(ctx, int64(req.VendorID))
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return GetVendorsVendorID404JSONResponse{s.vendorNotFoundByID(ctx, req.VendorID)}, nil
		}

		return GetVendorsVendorID500JSONResponse{}, err
	}

	return GetVendorsVendorID200JSONResponse(externalVendor(vendor)), nil
}

func (s *Server) DeleteVendorsVendorID(ctx context.Context, req DeleteVendorsVendorIDRequestObject) (DeleteVendorsVendorIDResponseObject, error) {
	if err := s.vendors.DeleteByID(ctx, int64(req.VendorID)); err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return DeleteVendorsVendorID404JSONResponse{s.vendorNotFoundByID(ctx, req.VendorID)}, nil
		}

		if errors.Is(err, models.ErrVendorHasModels) {
			message := s.mustLocalize(ctx, &i18n.LocalizeConfig{
				DefaultMessage: &i18n.Message{
					ID:    "vendor.error.deleteWithModels",
					Other: "A vendor with models cannot be deleted.",
				},
			})

			return DeleteVendorsVendorID400JSONResponse{InvalidRequestJSONResponse{Message: &message}}, nil
		}

		s.logger.ErrorContext(ctx, "Failed to delete vendor.", "error", err, "vendorID", req.VendorID)
		return DeleteVendorsVendorID500JSONResponse{}, nil
	}

	return DeleteVendorsVendorID204Response{}, nil
}

func (s *Server) GetVendorsVendorIDModels(ctx context.Context, req GetVendorsVendorIDModelsRequestObject) (GetVendorsVendorIDModelsResponseObject, error) {
	_, err := s.vendors.GetByID(ctx, int64(req.VendorID))
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return GetVendorsVendorIDModels404JSONResponse{s.vendorNotFoundByID(ctx, req.VendorID)}, nil
		}

		return GetVendorsVendorIDModels500JSONResponse{}, err
	}

	s.logger.DebugContext(ctx, "Vendor exists.", "vendorID", req.VendorID)

	models, err := s.models.ListByVendorID(ctx, int64(req.VendorID))
	if err != nil {
		return GetVendorsVendorIDModels500JSONResponse{}, err
	}

	return GetVendorsVendorIDModels200JSONResponse(modelCollection(models)), nil
}

func externalVendor(vendor models.Vendor) Vendor {
	return Vendor{
		CreatedAt: vendor.CreatedAt,
		Id:        int(vendor.ID),
		Name:      vendor.Name,
		UpdatedAt: vendor.UpdatedAt,
	}
}

func vendorCollection(vendors []models.Vendor) VendorCollection {
	reps := make([]Vendor, len(vendors))
	for i, vendor := range vendors {
		reps[i] = externalVendor(vendor)
	}

	return VendorCollection{Items: reps}
}

func internalNewVendor(vendor NewVendor) models.NewVendor {
	return models.NewVendor{
		Name: vendor.Name,
	}
}

func (s *Server) vendorNotFoundByID(ctx context.Context, id int) NotFoundJSONResponse {
	message := s.mustLocalize(ctx, &i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "vendor.error.notFoundByID",
			Other: "Vendor {{.ID}} does not exist.",
		},
		TemplateData: map[string]int{
			"ID": id,
		},
	})

	return NotFoundJSONResponse(APIError{Message: &message})
}
