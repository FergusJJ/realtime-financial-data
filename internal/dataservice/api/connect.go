package api

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"time"

	"github.com/FergusJJ/realtime-financial-data/internal/proto/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	PORT = "8000"
	HOST = "127.0.0.1"
)

func newClient(conn *grpc.ClientConn) pb.FinancialDataClient {
	return pb.NewFinancialDataClient(conn)
}

func getConn() (*grpc.ClientConn, error) {
	serverAddr := flag.String("server", fmt.Sprintf("%s:%s", HOST, PORT), "The server address in the format of host:port")
	flag.Parse()

	creds := credentials.NewTLS(&tls.Config{InsecureSkipVerify: false})

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	conn, err := grpc.DialContext(ctx, *serverAddr, opts...)
	if err != nil {
		return nil, fmt.Errorf("unable to dial: %w", err)
	}
	return conn, nil
}
