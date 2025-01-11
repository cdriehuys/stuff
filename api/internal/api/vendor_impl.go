package api

import (
	"context"
	"errors"
	"fmt"

	"github.com/cdriehuys/stuff/api/internal/models"
	"github.com/go-playground/validator/v10"
)

func (s *Server) GetVendors(ctx context.Context, req GetVendorsRequestObject) (GetVendorsResponseObject, error) {
	vendors, err := s.vendors.ListVendors(ctx)
	if err != nil {
		s.logger.ErrorContext(ctx, "Failed to list vendors.", "error", err)
		return GetVendors500JSONResponse{}, nil
	}

	return GetVendors200JSONResponse(vendorCollection(vendors)), nil
}

func (s *Server) PostVendors(ctx context.Context, req PostVendorsRequestObject) (PostVendorsResponseObject, error) {
	newVendor := internalNewVendor(*req.Body)

	vendor, err := s.vendors.Create(ctx, newVendor)
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			return PostVendors400JSONResponse{InvalidRequestJSONResponse(validationError(ve))}, nil
		}

		s.logger.ErrorContext(ctx, "Failed to create vendor.", "error", err)
		return PostVendors500JSONResponse{}, nil
	}

	return PostVendors201JSONResponse(externalVendor(vendor)), nil
}

func (s *Server) GetVendorsVendorID(ctx context.Context, req GetVendorsVendorIDRequestObject) (GetVendorsVendorIDResponseObject, error) {
	vendor, err := s.vendors.GetByID(ctx, int64(req.VendorID))
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			message := fmt.Sprintf("No vendor with ID %d.", req.VendorID)
			return GetVendorsVendorID404JSONResponse{NotFoundJSONResponse{Message: &message}}, nil
		}

		s.logger.ErrorContext(ctx, "Failed to get vendor by ID.", "error", err, "vendorID", req.VendorID)
		return GetVendorsVendorID500JSONResponse{}, nil
	}

	return GetVendorsVendorID200JSONResponse(externalVendor(vendor)), nil
}

func (s *Server) DeleteVendorsVendorID(ctx context.Context, req DeleteVendorsVendorIDRequestObject) (DeleteVendorsVendorIDResponseObject, error) {
	if err := s.vendors.DeleteByID(ctx, int64(req.VendorID)); err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return DeleteVendorsVendorID404JSONResponse{}, nil
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
			message := "Vendor does not exist."
			return GetVendorsVendorIDModels404JSONResponse{NotFoundJSONResponse{Message: &message}}, nil
		}

		s.logger.ErrorContext(ctx, "Failed to get vendor.", "error", err, "vendorID", req.VendorID)
		return GetVendorsVendorIDModels500JSONResponse{}, nil
	}

	s.logger.DebugContext(ctx, "Vendor exists.", "vendorID", req.VendorID)

	models, err := s.models.ListByVendorID(ctx, int64(req.VendorID))
	if err != nil {
		s.logger.ErrorContext(ctx, "Failed to get models for vendor.", "error", err, "vendorID", req.VendorID)
		return GetVendorsVendorIDModels500JSONResponse{}, nil
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
