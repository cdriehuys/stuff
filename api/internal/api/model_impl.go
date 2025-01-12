package api

import (
	"context"
	"errors"

	"github.com/cdriehuys/stuff/api/internal/models"
	"github.com/go-playground/validator/v10"
)

func (s *Server) GetModels(ctx context.Context, req GetModelsRequestObject) (GetModelsResponseObject, error) {
	models, err := s.models.ListModels(ctx)
	if err != nil {
		s.logger.ErrorContext(ctx, "Failed to list models.", "error", err)
		return GetModels500JSONResponse{}, nil
	}

	return GetModels200JSONResponse(modelCollection(models)), nil
}

func (s *Server) PostModels(ctx context.Context, req PostModelsRequestObject) (PostModelsResponseObject, error) {
	model := internalNewModel(*req.Body)
	newModel, err := s.models.Create(ctx, model)
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			return PostModels400JSONResponse{InvalidRequestJSONResponse(validationError(ve))}, nil
		}

		if errors.Is(err, models.ErrVendorNotFound) {
			fields := []FieldError{
				{Field: "vendorID", Message: "This vendor does not exist."},
			}

			return PostModels400JSONResponse{InvalidRequestJSONResponse{Fields: &fields}}, nil
		}

		if errors.Is(err, models.ErrModelNotUnique) {
			fields := []FieldError{
				{Field: "model", Message: "This model already exists for this vendor."},
			}

			return PostModels400JSONResponse{InvalidRequestJSONResponse{Fields: &fields}}, nil
		}

		s.logger.ErrorContext(ctx, "Failed to create model.", "error", err)
		return PostModels500JSONResponse{}, nil
	}

	return PostModels201JSONResponse(externalModel(newModel)), nil
}

func (s *Server) GetModelsModelID(ctx context.Context, req GetModelsModelIDRequestObject) (GetModelsModelIDResponseObject, error) {
	model, err := s.models.GetByID(ctx, int64(req.ModelID))
	if err != nil {
		s.logger.ErrorContext(ctx, "Failed to get model by ID.", "error", "err", "modelID", req.ModelID)
		return GetModelsModelID500JSONResponse{}, nil
	}

	return GetModelsModelID200JSONResponse(externalModel(model)), nil
}

func (s *Server) PutModelsModelID(ctx context.Context, req PutModelsModelIDRequestObject) (PutModelsModelIDResponseObject, error) {
	update := internalNewModel(*req.Body)
	model, err := s.models.UpdateByID(ctx, int64(req.ModelID), update)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			message := "Model does not exist."
			return PutModelsModelID404JSONResponse{NotFoundJSONResponse{Message: &message}}, nil
		}

		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			return PutModelsModelID400JSONResponse{InvalidRequestJSONResponse(validationError(ve))}, nil
		}

		if errors.Is(err, models.ErrVendorNotFound) {
			fields := []FieldError{
				{Field: "vendorID", Message: "This vendor does not exist."},
			}

			return PutModelsModelID400JSONResponse{InvalidRequestJSONResponse{Fields: &fields}}, nil
		}

		if errors.Is(err, models.ErrModelNotUnique) {
			fields := []FieldError{
				{Field: "model", Message: "This model already exists for this vendor."},
			}

			return PutModelsModelID400JSONResponse{InvalidRequestJSONResponse{Fields: &fields}}, nil
		}

		s.logger.ErrorContext(ctx, "Failed to update model.", "error", err, "modelID", req.ModelID)
		return PutModelsModelID500JSONResponse{}, nil
	}

	return PutModelsModelID200JSONResponse(externalModel(model)), nil
}

func (s *Server) DeleteModelsModelID(ctx context.Context, req DeleteModelsModelIDRequestObject) (DeleteModelsModelIDResponseObject, error) {
	err := s.models.DeleteByID(ctx, int64(req.ModelID))
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			message := "Model does not exist."
			return DeleteModelsModelID404JSONResponse{NotFoundJSONResponse{Message: &message}}, nil
		}

		s.logger.ErrorContext(ctx, "Failed to delete model.", "error", err, "modelID", req.ModelID)
		return DeleteModelsModelID500JSONResponse{}, nil
	}

	return DeleteModelsModelID204Response{}, nil
}

func externalModel(model models.Model) Model {
	return Model{
		CreatedAt: model.CreatedAt,
		Id:        int(model.ID),
		Model:     model.Model,
		Name:      model.Name,
		UpdatedAt: model.UpdatedAt,
		VendorID:  int(model.VendorID),
	}
}

func modelCollection(models []models.Model) ModelCollection {
	reps := make([]Model, len(models))
	for i, model := range models {
		reps[i] = externalModel(model)
	}

	return ModelCollection{Items: reps}
}

func internalNewModel(model NewModel) models.NewModel {
	internal := models.NewModel{
		VendorID: int64(model.VendorID),
		Model:    model.Model,
	}

	if model.Name != nil {
		internal.Name = *model.Name
	}

	return internal
}
