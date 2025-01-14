package api

import (
	"context"
	"net/http"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/oapi-codegen/runtime/strictmiddleware/nethttp"
)

type apiContextKey string

const ctxKeyLocalizer apiContextKey = "ctxKeyLocalizer"

func (s *Server) LocalizationMiddleware() StrictMiddlewareFunc {
	return func(next nethttp.StrictHTTPHandlerFunc, operationID string) nethttp.StrictHTTPHandlerFunc {
		return func(ctx context.Context, w http.ResponseWriter, r *http.Request, request any) (response any, err error) {
			lang := r.FormValue("lang")
			accept := r.Header.Get("Accept-Language")
			localizer := i18n.NewLocalizer(s.bundle, lang, accept)

			withLocalizer := context.WithValue(r.Context(), ctxKeyLocalizer, localizer)

			return next(withLocalizer, w, r, request)
		}
	}
}

func (s *Server) localize(ctx context.Context, config *i18n.LocalizeConfig) (string, error) {
	localizer, ok := ctx.Value(ctxKeyLocalizer).(*i18n.Localizer)
	if !ok {
		s.logger.WarnContext(ctx, "No localizer in context. Falling back to the default.")
		localizer = i18n.NewLocalizer(s.bundle)
	}

	return localizer.Localize(config)
}

func (s *Server) mustLocalize(ctx context.Context, config *i18n.LocalizeConfig) string {
	value, err := s.localize(ctx, config)
	if err != nil {
		panic(err)
	}

	return value
}
