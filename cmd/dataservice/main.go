package main

import (
	"context"
	"log"
	"time"

	"github.com/FergusJJ/realtime-financial-data/internal/dataservice"
	"github.com/FergusJJ/realtime-financial-data/internal/dataservice/finnhub"
	"github.com/FergusJJ/realtime-financial-data/internal/proto/pb"
	"github.com/FergusJJ/realtime-financial-data/pkg/config"
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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	pingRes, err := client.Ping(ctx, &pb.PingRequest{Ping: 42})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(pingRes)
	ds := setup()
	if ds == nil {
		log.Fatalln("error in setup function")
	}
	go func() {
		s := finnhub.NewFinnhubSession(ds.API_Keys["finnhub"])
		s.MonitorFinnhub()

	}()
}

// want to initialise connection to server?
// need to setup connections to apis
// setup
func setup() *dataservice.DataService {
	ds := &dataservice.DataService{}
	finnhubAPIKey, err := config.Load("FINNHUB_API_KEY")
	if err != nil {
		return nil
	}
	ds.API_Keys["finnhub"] = finnhubAPIKey
	return ds
}
