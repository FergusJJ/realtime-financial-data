package api

import (
	"time"

	"github.com/FergusJJ/realtime-financial-data/internal/proto/pb"
	"google.golang.org/grpc"
)

type FinancialDataClient struct {
	client  pb.FinancialDataClient
	timeout time.Duration
	conn    *grpc.ClientConn
}

func NewFinancialDataClient(timeout time.Duration) (*FinancialDataClient, error) {
	conn, err := getConn()
	if err != nil {
		return nil, err
	}
	client := newClient(conn)
	return &FinancialDataClient{
		client:  client,
		conn:    conn,
		timeout: timeout,
	}, nil
}
