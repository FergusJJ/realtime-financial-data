package api

import (
	"context"

	"github.com/FergusJJ/realtime-financial-data/internal/proto/pb"
)

func (c *FinancialDataClient) PingRequest(body *pb.PingRequest) (*pb.PingResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()
	return c.client.Ping(ctx, body)
}
