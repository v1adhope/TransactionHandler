package getter

import (
	"context"

	"github.com/v1adhope/TransactionHandler/Server/internal/couch"
	"github.com/v1adhope/TransactionHandler/Server/pb/v2"
)

type SvrGetter struct {
	pb.UnimplementedGetLastServer
}

func (s *SvrGetter) GetLast(ctx context.Context, req *pb.GetLastRequest) (*pb.GetLastResponse, error) {
	count := req.Count
	str := couch.GetDB(count)

	return &pb.GetLastResponse{Docs: str}, nil
}
