package api

import (
	"context"
)

func (s *Server) GetVendors(ctx context.Context, req GetVendorsRequestObject) (GetVendorsResponseObject, error) {
	vendors, err := s.vendors.ListVendors(ctx)
	if err != nil {
		return GetVendors200JSONResponse{}, err
	}

	return GetVendors200JSONResponse{Items: vendors}, nil
}
