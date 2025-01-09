package api

import (
	"context"
	"log/slog"
)

// ensure we're implementing the generated interface
var _ StrictServerInterface = (*Server)(nil)

type vendorModel interface {
	ListVendors(ctx context.Context) ([]Vendor, error)
}

type Server struct {
	logger  *slog.Logger
	vendors vendorModel
}

func NewServer(logger *slog.Logger, vendors vendorModel) *Server {
	return &Server{logger: logger, vendors: vendors}
}

func (s *Server) GetModels(ctx context.Context, req GetModelsRequestObject) (GetModelsResponseObject, error) {
	return GetModels200JSONResponse{
		Ping: "pong",
	}, nil
}
