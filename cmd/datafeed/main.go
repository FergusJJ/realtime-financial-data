package main

import (
	"fmt"
	"log"
	"time"

	"github.com/FergusJJ/realtime-financial-data/internal/datafeed/api"
	"github.com/FergusJJ/realtime-financial-data/internal/proto/pb"
)

/*
  this service is responsible for connecting to external apis
  it should process data so it is suitable to send to grpc servic
*/

func main() {
	client, err := api.NewFinancialDataClient(5 * time.Second)
	if err != nil {
		log.Fatal(err)
	}
	res, err := client.PingRequest(&pb.PingRequest{
		Ping: 42,
	})
	if err != nil {
		log.Println("error sending request:", err)
	}
	fmt.Println("response:", res)

}
