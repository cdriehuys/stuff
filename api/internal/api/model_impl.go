package api

import (
	"context"
	"fmt"
)

func (s *Server) GetModels(ctx context.Context, req GetModelsRequestObject) (GetModelsResponseObject, error) {
	models, err := s.models.ListModels(ctx)
	if err != nil {
		return GetModels200JSONResponse{}, fmt.Errorf("failed to list models: %v", err)
	}

	return GetModels200JSONResponse{Items: models}, nil
}
