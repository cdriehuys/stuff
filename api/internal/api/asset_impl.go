package api

import (
	"context"
	"errors"

	"github.com/cdriehuys/stuff/api/internal/models"
	"github.com/go-playground/validator/v10"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func (s *Server) GetAssets(ctx context.Context, req GetAssetsRequestObject) (GetAssetsResponseObject, error) {
	assets, err := s.assets.List(ctx)
	if err != nil {
		return GetAssets500JSONResponse{}, err
	}

	return GetAssets200JSONResponse(assetCollection(assets)), nil
}

func (s *Server) PostAssets(ctx context.Context, req PostAssetsRequestObject) (PostAssetsResponseObject, error) {
	newAsset := internalNewAsset(*req.Body)
	asset, err := s.assets.Create(ctx, newAsset)
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			return PostAssets400JSONResponse{s.validationError(ctx, ve)}, nil
		}

		if errors.Is(err, models.ErrModelNotFound) {
			return PostAssets400JSONResponse{s.assetModelNotFound(ctx)}, nil
		}

		if errors.Is(err, models.ErrAssetSerialNotUnique) {
			return PostAssets400JSONResponse{s.assetSerialNotUnique(ctx, newAsset.Serial)}, nil
		}

		return PostAssets500JSONResponse{}, err
	}

	return PostAssets201JSONResponse(externalAsset(asset)), nil
}

func (s *Server) GetAssetsAssetID(ctx context.Context, req GetAssetsAssetIDRequestObject) (GetAssetsAssetIDResponseObject, error) {
	asset, err := s.assets.GetByID(ctx, int64(req.AssetID))
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return GetAssetsAssetID404JSONResponse{s.assetNotFoundByID(ctx, req.AssetID)}, nil
		}

		return GetAssetsAssetID500JSONResponse{}, err
	}

	return GetAssetsAssetID200JSONResponse(externalAsset(asset)), nil
}

func (s *Server) PutAssetsAssetID(ctx context.Context, req PutAssetsAssetIDRequestObject) (PutAssetsAssetIDResponseObject, error) {
	update := internalNewAsset(*req.Body)
	asset, err := s.assets.UpdateByID(ctx, int64(req.AssetID), update)
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			return PutAssetsAssetID400JSONResponse{s.validationError(ctx, ve)}, nil
		}

		if errors.Is(err, models.ErrModelNotFound) {
			return PutAssetsAssetID400JSONResponse{s.assetModelNotFound(ctx)}, nil
		}

		if errors.Is(err, models.ErrAssetSerialNotUnique) {
			return PutAssetsAssetID400JSONResponse{s.assetSerialNotUnique(ctx, update.Serial)}, nil
		}

		return PutAssetsAssetID500JSONResponse{}, err
	}

	return PutAssetsAssetID200JSONResponse(externalAsset(asset)), nil
}

func (s *Server) DeleteAssetsAssetID(ctx context.Context, req DeleteAssetsAssetIDRequestObject) (DeleteAssetsAssetIDResponseObject, error) {
	err := s.assets.DeleteByID(ctx, int64(req.AssetID))
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return DeleteAssetsAssetID404JSONResponse{s.assetNotFoundByID(ctx, req.AssetID)}, nil
		}

		return DeleteAssetsAssetID500JSONResponse{}, err
	}

	return DeleteAssetsAssetID204Response{}, nil
}

func (s *Server) GetModelsModelIDAssets(ctx context.Context, req GetModelsModelIDAssetsRequestObject) (GetModelsModelIDAssetsResponseObject, error) {
	_, err := s.models.GetByID(ctx, int64(req.ModelID))
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return GetModelsModelIDAssets404JSONResponse{s.modelNotFoundByID(ctx, req.ModelID)}, nil
		}

		return GetModelsModelIDAssets500JSONResponse{}, err
	}

	s.logger.DebugContext(ctx, "Model exists.", "modelID", req.ModelID)

	assets, err := s.assets.ListByModel(ctx, int64(req.ModelID))
	if err != nil {
		return GetModelsModelIDAssets500JSONResponse{}, err
	}

	return GetModelsModelIDAssets200JSONResponse(assetCollection(assets)), nil
}

func externalAsset(asset models.Asset) Asset {
	return Asset{
		Comments:  &asset.Comments,
		CreatedAt: asset.CreatedAt,
		Id:        int(asset.ID),
		ModelID:   int(asset.ModelID),
		VendorID:  int(asset.VendorID),
		Serial:    &asset.Serial,
		UpdatedAt: asset.UpdatedAt,
	}
}

func assetCollection(assets []models.Asset) AssetCollection {
	reps := make([]Asset, len(assets))
	for i, asset := range assets {
		reps[i] = externalAsset(asset)
	}

	return AssetCollection{Items: reps}
}

func internalNewAsset(rep NewAsset) models.NewAsset {
	asset := models.NewAsset{
		ModelID: int64(rep.ModelID),
		Serial:  rep.Serial,
	}

	if rep.Comments != nil {
		asset.Comments = *rep.Comments
	}

	return asset
}

func (s *Server) assetNotFoundByID(ctx context.Context, id int) NotFoundJSONResponse {
	message := s.mustLocalize(ctx, &i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "asset.error.notFoundByID",
			Other: "Asset {{.ID}} does not exist.",
		},
		TemplateData: map[string]int{
			"ID": id,
		},
	})

	return NotFoundJSONResponse(APIError{Message: &message})
}

func (s *Server) assetModelNotFound(ctx context.Context) InvalidRequestJSONResponse {
	fields := []FieldError{
		{Field: "modelID", Message: s.mustLocalize(ctx, &i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "asset.error.modelNotFound",
				Other: "This model does not exist.",
			},
		})},
	}

	return InvalidRequestJSONResponse{Fields: &fields}
}

func (s *Server) assetSerialNotUnique(ctx context.Context, serial string) InvalidRequestJSONResponse {
	fields := []FieldError{
		{Field: "serial", Message: s.mustLocalize(ctx, &i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "asset.error.serialNotUnique",
				Other: "This model already has an asset with the serial '{{.Serial}}'.",
			},
			TemplateData: map[any]string{
				"Serial": serial,
			},
		})},
	}

	return InvalidRequestJSONResponse{Fields: &fields}
}
