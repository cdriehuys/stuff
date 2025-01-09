package api

import "context"

// ensure we're implementing the generated interface
var _ StrictServerInterface = (StrictServerInterface)(nil)

type Server struct{}

func (s *Server) GetModels(ctx context.Context, req GetModelsRequestObject) (GetModelsResponseObject, error) {
	return GetModels200JSONResponse{
		Ping: "pong",
	}, nil
}
