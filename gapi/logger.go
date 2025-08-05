package gapi

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GrpcLogger(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	startTime := time.Now()
	result, err := handler(ctx, req)

	statusCode := status.Code(err)
	if statusCode != codes.OK {
		log.Error().Str("protocol", "grpc").Str("method", info.FullMethod).Int("status_code", int(statusCode)).Msgf("error while calling %s: %v", info.FullMethod, err)
	}

	logger := log.Info()
	if err != nil {
		logger = log.Error()
	}

	duration := time.Since(startTime)
	logger.Str("protocol", "grpc").
		Str("method", info.FullMethod).
		Int("status_code", int(statusCode)).
		Dur("duration", duration).
		Msgf("request - %v - %v", info.FullMethod, req)
	return result, err
}
