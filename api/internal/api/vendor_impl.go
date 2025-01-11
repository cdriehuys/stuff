package api

import (
	"context"
	"errors"
	"fmt"

	"github.com/cdriehuys/stuff/api/internal/apierrors"
)

func (s *Server) GetVendors(ctx context.Context, req GetVendorsRequestObject) (GetVendorsResponseObject, error) {
	vendors, err := s.vendors.ListVendors(ctx)
	if err != nil {
		return GetVendors200JSONResponse{}, err
	}

	return GetVendors200JSONResponse{Items: vendors}, nil
}

func (s *Server) PostVendors(ctx context.Context, req PostVendorsRequestObject) (PostVendorsResponseObject, error) {
	newVendor := NewVendor(*req.Body)

	if err := s.validate.Struct(newVendor); err != nil {
		s.logger.DebugContext(ctx, "Invalid new vendor.", "error", err)

		return PostVendors400JSONResponse(validationError(err)), nil
	}

	vendor, err := s.vendors.Create(ctx, newVendor)
	if err != nil {
		return PostVendors201JSONResponse{}, err
	}

	return PostVendors201JSONResponse(vendor), nil
}

func (s *Server) GetVendorsVendorID(ctx context.Context, req GetVendorsVendorIDRequestObject) (GetVendorsVendorIDResponseObject, error) {
	vendor, err := s.vendors.GetByID(ctx, int64(req.VendorID))
	if err != nil {
		if errors.Is(err, apierrors.ErrNotFound) {
			message := fmt.Sprintf("No vendor with ID %d.", req.VendorID)
			return GetVendorsVendorID404JSONResponse{Message: &message}, nil
		}

		return GetVendorsVendorID200JSONResponse{}, err
	}

	return GetVendorsVendorID200JSONResponse(vendor), nil
}

func (s *Server) DeleteVendorsVendorID(ctx context.Context, req DeleteVendorsVendorIDRequestObject) (DeleteVendorsVendorIDResponseObject, error) {
	if err := s.vendors.DeleteByID(ctx, int64(req.VendorID)); err != nil {
		if errors.Is(err, apierrors.ErrNotFound) {
			return DeleteVendorsVendorID404JSONResponse{}, nil
		}

		return DeleteVendorsVendorID204Response{}, fmt.Errorf("failed to delete vendor: %v", err)
	}

	return DeleteVendorsVendorID204Response{}, nil
}

func (s *Server) GetVendorsVendorIDModels(ctx context.Context, req GetVendorsVendorIDModelsRequestObject) (GetVendorsVendorIDModelsResponseObject, error) {
	_, err := s.vendors.GetByID(ctx, int64(req.VendorID))
	if err != nil {
		if errors.Is(err, apierrors.ErrNotFound) {
			message := "Vendor does not exist."
			return GetVendorsVendorIDModels404JSONResponse{Message: &message}, nil
		}

		return GetVendorsVendorIDModels200JSONResponse{}, err
	}

	s.logger.DebugContext(ctx, "Vendor exists.", "vendorID", req.VendorID)

	models, err := s.models.ListByVendorID(ctx, int64(req.VendorID))
	if err != nil {
		return GetVendorsVendorIDModels200JSONResponse{}, err
	}

	return GetVendorsVendorIDModels200JSONResponse{Items: models}, nil
}

func (s *Server) PostVendorsVendorIDModels(ctx context.Context, req PostVendorsVendorIDModelsRequestObject) (PostVendorsVendorIDModelsResponseObject, error) {
	newModel, err := s.models.Create(ctx, int64(req.VendorID), *req.Body)
	if err != nil {
		if errors.Is(err, apierrors.ErrNotFound) {
			message := fmt.Sprintf("No vendor with ID %d.", req.VendorID)
			return PostVendorsVendorIDModels404JSONResponse{Message: &message}, nil
		}

		return PostVendorsVendorIDModels201JSONResponse{}, err
	}

	return PostVendorsVendorIDModels201JSONResponse(newModel), nil
}
