package api

import (
	"context"
	"errors"

	"github.com/cdriehuys/stuff/api/internal/models"
	"github.com/go-playground/validator/v10"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func (s *Server) GetModels(ctx context.Context, req GetModelsRequestObject) (GetModelsResponseObject, error) {
	models, err := s.models.ListModels(ctx)
	if err != nil {
		return GetModels500JSONResponse{}, err
	}

	return GetModels200JSONResponse(modelCollection(models)), nil
}

func (s *Server) PostModels(ctx context.Context, req PostModelsRequestObject) (PostModelsResponseObject, error) {
	model := internalNewModel(*req.Body)
	newModel, err := s.models.Create(ctx, model)
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			return PostModels400JSONResponse{InvalidRequestJSONResponse(s.validationError(ctx, ve))}, nil
		}

		if errors.Is(err, models.ErrVendorNotFound) {
			return PostModels400JSONResponse{s.modelVendorNotFound(ctx)}, nil
		}

		if errors.Is(err, models.ErrModelNotUnique) {
			return PostModels400JSONResponse{s.modelNotUnique(ctx, model.Model)}, nil
		}

		return PostModels500JSONResponse{}, err
	}

	return PostModels201JSONResponse(externalModel(newModel)), nil
}

func (s *Server) GetModelsModelID(ctx context.Context, req GetModelsModelIDRequestObject) (GetModelsModelIDResponseObject, error) {
	model, err := s.models.GetByID(ctx, int64(req.ModelID))
	if err != nil {
		return GetModelsModelID500JSONResponse{}, err
	}

	return GetModelsModelID200JSONResponse(externalModel(model)), nil
}

func (s *Server) PutModelsModelID(ctx context.Context, req PutModelsModelIDRequestObject) (PutModelsModelIDResponseObject, error) {
	update := internalNewModel(*req.Body)
	model, err := s.models.UpdateByID(ctx, int64(req.ModelID), update)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return PutModelsModelID404JSONResponse{s.modelNotFoundByID(ctx, req.ModelID)}, nil
		}

		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			return PutModelsModelID400JSONResponse{InvalidRequestJSONResponse(s.validationError(ctx, ve))}, nil
		}

		if errors.Is(err, models.ErrVendorNotFound) {
			return PutModelsModelID400JSONResponse{s.modelVendorNotFound(ctx)}, nil
		}

		if errors.Is(err, models.ErrModelNotUnique) {
			return PutModelsModelID400JSONResponse{s.modelNotUnique(ctx, update.Model)}, nil
		}

		return PutModelsModelID500JSONResponse{}, err
	}

	return PutModelsModelID200JSONResponse(externalModel(model)), nil
}

func (s *Server) DeleteModelsModelID(ctx context.Context, req DeleteModelsModelIDRequestObject) (DeleteModelsModelIDResponseObject, error) {
	err := s.models.DeleteByID(ctx, int64(req.ModelID))
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return DeleteModelsModelID404JSONResponse{s.modelNotFoundByID(ctx, req.ModelID)}, nil
		}

		return DeleteModelsModelID500JSONResponse{}, err
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

func (s *Server) modelNotFoundByID(ctx context.Context, id int) NotFoundJSONResponse {
	message := s.mustLocalize(ctx, &i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "model.error.notFoundByID",
			Other: "Model {{.ID}} does not exist.",
		},
		TemplateData: map[string]int{
			"ID": id,
		},
	})

	return NotFoundJSONResponse(APIError{Message: &message})
}

func (s *Server) modelVendorNotFound(ctx context.Context) InvalidRequestJSONResponse {
	fields := []FieldError{
		{Field: "vendorID", Message: s.mustLocalize(ctx, &i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "model.error.vendorNotFound",
				Other: "This vendor does not exist.",
			},
		})},
	}

	return InvalidRequestJSONResponse{Fields: &fields}
}

func (s *Server) modelNotUnique(ctx context.Context, model string) InvalidRequestJSONResponse {
	fields := []FieldError{
		{Field: "model", Message: s.mustLocalize(ctx, &i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "model.error.modelNotUnique",
				Other: "This vendor already has the model '{{.Model}}'.",
			},
			TemplateData: map[any]string{
				"Model": model,
			},
		})},
	}

	return InvalidRequestJSONResponse{Fields: &fields}
}
