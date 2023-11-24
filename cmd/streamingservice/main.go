package main

import (
	"log"
	"net"

	pb "github.com/FergusJJ/realtime-financial-data/internal/proto/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

/*
  This is a grpc responsible for sending data to clients over grpc
  it should receive data from the datafeed service
*/

type server struct {
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("failed to create listener: ", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)

	pb.RegisterFinancialDataServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalln("failed to serve:", err)
	}
}
