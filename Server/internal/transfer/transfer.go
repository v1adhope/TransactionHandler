package transfer

import (
	"context"
	"time"

	"github.com/v1adhope/TransactionHandler/Server/internal/couch"
	"github.com/v1adhope/TransactionHandler/Server/pb/v2"
)

type SvrTransactions struct {
	pb.UnimplementedSendServer
}

func (g *SvrTransactions) Send(ctx context.Context, req *pb.SendRequest) (*pb.SendResponse, error) {
	currentTime := time.Now().Format(time.RFC1123Z)
	data := couch.NewData(currentTime, req.From, req.To, req.Amount)
	couch.WriteDB(&data)
	return &pb.SendResponse{Result: "Translation done!"}, nil
}
