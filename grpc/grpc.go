package grpc

import (
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/wcodesoft/mosha-service-common/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
)

// CreateNewGRPCServer creates a new gRPC server with logger interceptor.
func CreateNewGRPCServer() *grpc.Server {
	l := log.New(os.Stderr)
	loggerOpts := []logging.Option{
		logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
	}
	return grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			// Add logger interceptor to grpc server.
			logging.UnaryServerInterceptor(logger.InterceptorLogger(l), loggerOpts...),
		),
	)
}

// ClientInfo represents a gRPC client.
type ClientInfo struct {
	Name    string
	Address string
}

// NewClientConnection creates a new gRPC client connection.
func (ci *ClientInfo) NewClientConnection() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(ci.Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("could not connect to %s at: %s", ci.Name, ci.Address)
	}
	log.Infof("Connected to %s at: %s", ci.Name, ci.Address)
	return conn, nil
}
