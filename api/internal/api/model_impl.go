package api

import (
	"context"
	"errors"
	"fmt"

	"github.com/cdriehuys/stuff/api/internal/apierrors"
)

func (s *Server) GetModels(ctx context.Context, req GetModelsRequestObject) (GetModelsResponseObject, error) {
	models, err := s.models.ListModels(ctx)
	if err != nil {
		return GetModels200JSONResponse{}, fmt.Errorf("failed to list models: %v", err)
	}

	return GetModels200JSONResponse{Items: models}, nil
}

func (s *Server) GetModelsModelID(ctx context.Context, req GetModelsModelIDRequestObject) (GetModelsModelIDResponseObject, error) {
	model, err := s.models.GetByID(ctx, int64(req.ModelID))
	if err != nil {
		return GetModelsModelID200JSONResponse{}, err
	}

	return GetModelsModelID200JSONResponse(model), nil
}

func (s *Server) PutModelsModelID(ctx context.Context, req PutModelsModelIDRequestObject) (PutModelsModelIDResponseObject, error) {
	model, err := s.models.UpdateByID(ctx, int64(req.ModelID), *req.Body)
	if err != nil {
		if errors.Is(err, apierrors.ErrNotFound) {
			message := "Model does not exist."
			return PutModelsModelID404JSONResponse{Message: &message}, nil
		}

		return PutModelsModelID200JSONResponse{}, err
	}

	return PutModelsModelID200JSONResponse(model), nil
}

func (s *Server) DeleteModelsModelID(ctx context.Context, req DeleteModelsModelIDRequestObject) (DeleteModelsModelIDResponseObject, error) {
	err := s.models.DeleteByID(ctx, int64(req.ModelID))
	if err != nil {
		if errors.Is(err, apierrors.ErrNotFound) {
			message := "Model does not exist."
			return DeleteModelsModelID404JSONResponse{Message: &message}, nil
		}

		return DeleteModelsModelID204Response{}, err
	}

	return DeleteModelsModelID204Response{}, nil
}
