package main

import (
	"crypto/tls"
	"log"
	"net"

	pb "github.com/FergusJJ/realtime-financial-data/internal/proto/pb"
	"github.com/FergusJJ/realtime-financial-data/internal/streamingservice"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

/*
  This is a grpc responsible for sending data to clients over grpc
  it should receive data from the datafeed service
*/

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatalln("failed to create listener: ", err)
	}
	tlsCredentials, err := loadTLS()
	if err != nil {
		log.Fatalln("error loading tls:", err)
	}
	s := grpc.NewServer(
		grpc.Creds(tlsCredentials),
	)
	reflection.Register(s)

	pb.RegisterFinancialDataServer(s, &streamingservice.Server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalln("failed to serve:", err)
	}
}

func loadTLS() (credentials.TransportCredentials, error) {
	serverCert, err := tls.LoadX509KeyPair("cert/server-cert.pem", "cert/server-key.pem")
	if err != nil {
		return nil, err
	}
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}
	return credentials.NewTLS(config), nil
}
