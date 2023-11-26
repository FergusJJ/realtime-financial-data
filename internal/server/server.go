package server
import (
	"context"

	"github.com/FergusJJ/realtime-financial-data/internal/proto/pb"
)

type server struct {
	pb.UnimplementedFinancialDataServer
}

func NewServer() *server {
  s := &server{}
  return s
}

func (s *server) Ping(
	ctx context.Context, in *pb.PingRequest,
) (*pb.PingResponse, error) {
	return &pb.PingResponse{
		Pong: in.GetPing(),
	}, nil
}
