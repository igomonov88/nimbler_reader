package handlers

import (
	"context"
	"net/http"

	"go.opencensus.io/trace"
	"google.golang.org/grpc/status"

	"github.com/igomonov88/nimbler_reader/internal/storage"
	pb "github.com/igomonov88/nimbler_reader/proto"
)

func (s *Server) GetOriginalURL(ctx context.Context, req *pb.GetOriginalUrlRequest) (*pb.GetOriginalUrlResponse, error) {
	ctx, span := trace.StartSpan(ctx, "handlers.GetOriginalURL")
	defer span.End()

	url, err := storage.RetrieveOriginalURL(ctx, s.DB, req.GetUrlHash())
	if err != nil {
		return &pb.GetOriginalUrlResponse{}, status.Error(http.StatusInternalServerError, err.Error())
	}

	return &pb.GetOriginalUrlResponse{OriginalUrl: url}, nil
}
