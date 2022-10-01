package main

import (
	"log"
	"net"

	"github.com/v1adhope/TransactionHandler/Server/internal/getter"
	"github.com/v1adhope/TransactionHandler/Server/internal/transfer"
	_ "github.com/v1adhope/TransactionHandler/Server/internal/verifyDataBase"
	"github.com/v1adhope/TransactionHandler/Server/pb/v2"

	"google.golang.org/grpc"
)

func main() {

	s := grpc.NewServer() //TODO

	srvSend := &transfer.SvrTransactions{
		UnimplementedSendServer: pb.UnimplementedSendServer{},
	}
	pb.RegisterSendServer(s, srvSend)

	svrGetLast := &getter.SvrGetter{
		UnimplementedGetLastServer: pb.UnimplementedGetLastServer{},
	}
	pb.RegisterGetLastServer(s, svrGetLast)

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	if err := s.Serve(lis); err != nil {
		log.Fatalf("could not serve: %v\n", err)
	}

	defer s.GracefulStop()
}
