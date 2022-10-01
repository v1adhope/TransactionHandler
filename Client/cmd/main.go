package main

import (
	"context"
	"flag"
	"log"
	"strconv"

	"github.com/v1adhope/TransactionHandler/Client/pb/v2"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func connect() *grpc.ClientConn {
	conn, err := grpc.Dial(":9000", grpc.WithTransportCredentials(insecure.NewCredentials())) //TODO no sec
	if err != nil {
		log.Fatalf("connect: %v", err)
	}
	return conn
}

func main() {
	flag.Parse()

	switch flag.Arg(0) {
	case "send":
		{
			if flag.NArg() != 4 {
				log.Fatal("not enought arg: addressFrom addressTo amount")
			}
			amount, err := strconv.ParseFloat(flag.Arg(3), 64)
			if err != nil {
				log.Fatalf("number entered incorrectly: %v\n", err)
			}

			conn := connect()
			defer conn.Close()

			client := pb.NewSendClient(conn)
			res, err := client.Send(context.Background(), &pb.SendRequest{From: flag.Arg(1), To: flag.Arg(2), Amount: float64(amount)})
			if err != nil {
				log.Fatal(err)
			}
			log.Println(res.GetResult())
		}
	case "getlast":
		{
			if flag.NArg() != 2 {
				log.Fatal("not enought arg: countLastTransaction")
			}
			count, err := strconv.Atoi(flag.Arg(1))
			if err != nil || count <= 0 {
				log.Fatalf("the number of requested transactions must be greater than 0: %v", err)
			}

			conn := connect()
			defer conn.Close()

			client := pb.NewGetLastClient(conn)
			res, err := client.GetLast(context.Background(), &pb.GetLastRequest{Count: int32(count)})
			if err != nil {
				log.Fatal(err)
			}
			log.Println(res.GetDocs())
		}
	default:
		{
			log.Println("Action does not exist: send or getlast")
		}
	}
}
