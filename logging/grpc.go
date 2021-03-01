package logging

import (
	"context"
	"path"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

const logMessage = "gRPC"

func logIncomingCall(ctx context.Context, event *zerolog.Event, method string, start time.Time, req interface{}) {
	event.Time("time", start).
		Str("service", path.Dir(method)[1:]).
		Str("method", path.Base(method)).
		Dur("duration", time.Since(start))
	logIP(ctx, event)
}

func logIP(ctx context.Context, logger *zerolog.Event) {
	if p, ok := peer.FromContext(ctx); ok {
		*logger = *logger.Str("ip", p.Addr.String())
	}
}

func logStatusError(logger *zerolog.Event, err error) {
	statusErr := status.Convert(err)
	*logger = *logger.Err(err).Str("code", statusErr.Code().String()).Str("message", statusErr.Message()).Interface("details", statusErr.Details())
}

// UnaryInterceptor is a gRPC Server Option that uses NewUnaryServerInterceptor() to log gRPC Requests.
// It uses global zerolog instance for logging.
func UnaryInterceptor() grpc.ServerOption {
	return grpc.UnaryInterceptor(NewUnaryLoggingInterceptor())
}

func NewUnaryLoggingInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		startTime := time.Now()
		resp, err := handler(ctx, req)
		if err != nil {
			event := log.Error()
			logIncomingCall(ctx, event, info.FullMethod, startTime, req)
			logStatusError(event, err)
			event.Msg(logMessage)
		} else {
			event := log.Info()
			logIncomingCall(ctx, event, info.FullMethod, startTime, req)
			event.Msg(logMessage)
		}

		return resp, err
	}
}
