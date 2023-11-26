package main

import (
	"context"
	"log"
	"time"

	"github.com/FergusJJ/realtime-financial-data/internal/proto/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

/*
  this service is responsible for connecting to external apis
  it should process data so it is suitable to send to grpc servic
*/

func main() {
  conn, err := grpc.Dial("localhost:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewFinancialDataClient(conn)
 	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
  pingRes, err := client.Ping(ctx, &pb.PingRequest{Ping: 42})
  if err != nil {
    log.Fatal(err)
  }
  log.Println(pingRes)
}
