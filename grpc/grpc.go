package grpc

import (
	"github.com/charmbracelet/log"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/wcodesoft/mosha-service-common/logger"
	"google.golang.org/grpc"
	"os"
)

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
