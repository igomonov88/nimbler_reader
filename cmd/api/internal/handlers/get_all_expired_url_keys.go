package handlers

import (
	"context"
	"net/http"

	"go.opencensus.io/trace"
	"google.golang.org/grpc/status"

	"github.com/igomonov88/nimbler_reader/internal/storage"
	pb "github.com/igomonov88/nimbler_reader/proto"
)

func (s *Server) GetAllExpiredUrlKeys(ctx context.Context, req *pb.GetAllExpiredUrlKeysRequest) (*pb.GetAllExpiredUrlKeysResponse, error) {
	ctx, span := trace.StartSpan(ctx, "handlers.GetAllExpiredUrlKeys")
	defer span.End()

	urls, err := storage.RetrieveAllExpiredURLKeysFromDate(ctx, s.DB, req.ExpirationDate.AsTime(), req.GetLimit())
	if err != nil {
		return &pb.GetAllExpiredUrlKeysResponse{}, status.Error(http.StatusInternalServerError, err.Error())
	}

	return &pb.GetAllExpiredUrlKeysResponse{UrlKeys: urls}, nil
}
