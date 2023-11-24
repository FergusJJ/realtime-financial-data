package streamingservice

import (
	"context"
	"log"

	"github.com/FergusJJ/realtime-financial-data/internal/proto/pb"
)

type Server struct {
	pb.UnimplementedFinancialDataServer
}

func (s *Server) Ping(
	ctx context.Context, in *pb.PingRequest,
) (*pb.PingResponse, error) {
	log.Println("got ping:", in.GetPing())
	return &pb.PingResponse{
		Pong: in.GetPing(),
	}, nil
}
