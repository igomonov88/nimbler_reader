package handlers

import (
	"context"
	"net/http"

	"go.opencensus.io/trace"
	"google.golang.org/grpc/status"

	"github.com/igomonov88/nimbler_reader/internal/storage"
	pb "github.com/igomonov88/nimbler_reader/proto"
)

func (s *Server) GetUserInfo(ctx context.Context, req *pb.GetUserInfoRequest) (*pb.GetUserInfoResponse, error) {
	ctx, span := trace.StartSpan(ctx, "handlers.GetUserInfo")
	defer span.End()

	u, err := storage.GetUserInfo(ctx, s.DB, req.GetUserID())
	if err != nil {
		return &pb.GetUserInfoResponse{}, status.Error(http.StatusInternalServerError, err.Error())
	}

	return &pb.GetUserInfoResponse{Name: u.Name, Email: u.Email}, nil
}
