package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/oapi-codegen/runtime/strictmiddleware/nethttp"
)

func (s *Server) validationError(ctx context.Context, err error) InvalidRequestJSONResponse {
	ve := err.(validator.ValidationErrors)

	fieldErrors := make([]FieldError, len(ve))
	for i, fe := range ve {
		fieldErrors[i] = FieldError{Field: fe.Field(), Message: s.fieldErrorMessage(ctx, fe)}
	}

	return InvalidRequestJSONResponse(APIError{Fields: &fieldErrors})
}

func (s *Server) fieldErrorMessage(ctx context.Context, fe validator.FieldError) string {
	switch fe.Tag() {
	case "max":
		return s.mustLocalize(ctx, &i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "validation.maxLength",
				Other: "This field has a maximum length of {{.Length}}",
			},
			TemplateData: map[string]any{
				"Length": fe.Param(),
			},
		})

	case "required":
		return s.mustLocalize(ctx, &i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "validation.required",
				Other: "This field is required.",
			},
		})
	}

	return s.mustLocalize(ctx, &i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "validation.unknownTag",
			Other: "Failed validation tag '{{.Tag}}'",
		},
		TemplateData: map[string]any{
			"Tag": fe.Tag(),
		},
	})
}

func (s *Server) ErrorMiddleware() StrictMiddlewareFunc {
	return func(next nethttp.StrictHTTPHandlerFunc, operationID string) nethttp.StrictHTTPHandlerFunc {
		return func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
			response, err := next(ctx, w, r, request)
			if err != nil {
				s.logger.ErrorContext(ctx, "Unhandled error from handler.", "error", err, "operationID", operationID)

				message := s.mustLocalize(ctx, &i18n.LocalizeConfig{
					DefaultMessage: &i18n.Message{
						ID:    "error.unknown",
						Other: "An unknown error occurred.",
					},
				})
				errResponse := APIError{Message: &message}

				w.WriteHeader(http.StatusInternalServerError)
				if err := json.NewEncoder(w).Encode(errResponse); err != nil {
					s.logger.ErrorContext(ctx, "Failed to write error message.", "error", err)
				}

				return nil, nil
			}

			return response, nil
		}
	}
}

func (s *Server) PanicRecoveryMiddleware() StrictMiddlewareFunc {
	return func(next nethttp.StrictHTTPHandlerFunc, operationID string) nethttp.StrictHTTPHandlerFunc {
		return func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (response interface{}, err error) {
			defer func() {
				if panicErr := recover(); panicErr != nil {
					s.logger.ErrorContext(ctx, "Recovered a panic.", "operationID", operationID)

					w.Header().Set("Connection", "close")

					response = nil
					err = fmt.Errorf("unhandled panic: %v", panicErr)
				}
			}()

			response, err = next(ctx, w, r, request)

			return
		}
	}
}
