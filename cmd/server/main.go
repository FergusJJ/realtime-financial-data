package main

import (
	"log"
	"log/slog"
	"net"
	"os"

	pb "github.com/FergusJJ/realtime-financial-data/internal/proto/pb"
	"github.com/FergusJJ/realtime-financial-data/internal/server"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

/*
  This is a grpc responsible for sending data to clients over grpc
  it should receive data from the datafeed service
*/

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatalln("failed to create listener: ", err)
	}
  s := initServer()
	if err := s.Serve(listener); err != nil {
		log.Fatalln("failed to serve:", err)
	}
}

func initServer() *grpc.Server {
  logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))
  slopts := []logging.Option{
    logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
  }

  var opts []grpc.ServerOption
//  opts = append(opts, grpc.ChainStreamInterceptor(logging.StreamServerInterceptor(server.InterceptorLogger(logger), slopts...)))
  opts = append(opts, grpc.ChainUnaryInterceptor(logging.UnaryServerInterceptor(server.InterceptorLogger(logger), slopts...)))
  s := grpc.NewServer(
    opts...,
  )
  reflection.Register(s)
  fds := server.NewServer()
  pb.RegisterFinancialDataServer(s, fds)
  return s
}
