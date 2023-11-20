package serverity

import (
	"google.golang.org/grpc"
	"net"
	"os"
	"task/pkg/grpc_handler/api"
)

type GRPCServer struct {
	grpcServer *grpc.Server
}

func (s *GRPCServer) Run(handler *api.GRPCHandlerTask) error {
	s.grpcServer = grpc.NewServer()

	l, err := net.Listen("tcp", os.Getenv("GRPC_APP_IP")+":"+os.Getenv("GRPC_APP_PORT"))
	if err != nil {
		return err
	}
	api.RegisterTaskServer(s.grpcServer, handler)

	return s.grpcServer.Serve(l)
}

func (s *GRPCServer) ShutDown() {
	s.grpcServer.Stop()
}
