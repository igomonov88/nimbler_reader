package handlers

import (
	"context"
	"net/http"

	"go.opencensus.io/trace"
	"google.golang.org/grpc/status"

	"github.com/igomonov88/nimbler_reader/internal/storage"
	pb "github.com/igomonov88/nimbler_reader/proto"
)

func (s *Server) DoesCustomAliasExist(ctx context.Context, req *pb.DoesCustomAliasExistRequest) (*pb.DoesCustomAliasExistResponse, error) {
	ctx, span := trace.StartSpan(ctx, "handlers.DoesCustomAliasExist")
	defer span.End()

	exist, err := storage.DoesURLAliasExist(ctx, s.DB, req.GetCustomAlias())
	if err != nil {
		return &pb.DoesCustomAliasExistResponse{}, status.Error(http.StatusInternalServerError, err.Error())
	}

	return &pb.DoesCustomAliasExistResponse{Exist: exist}, nil
}
