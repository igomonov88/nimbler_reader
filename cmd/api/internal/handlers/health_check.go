package handlers

import (
	"context"
	"net/http"

	"go.opencensus.io/trace"
	"google.golang.org/grpc/status"

	"github.com/igomonov88/nimbler_reader/internal/platform/database"
	pb "github.com/igomonov88/nimbler_reader/proto"

)

func (s *Server) HealthCheck(ctx context.Context, req *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	ctx, span := trace.StartSpan(ctx, "handlers.Check.Health")
	defer span.End()

	if err := database.StatusCheck(ctx, s.DB); err != nil {
		return &pb.HealthCheckResponse{}, status.Error(http.StatusInternalServerError, "database is not ready")
	}

	return &pb.HealthCheckResponse{Version: "develop"}, nil
}

