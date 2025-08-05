package gapi

import (
	"context"
	"net/http"
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

type ResponseRecorder struct {
	http.ResponseWriter
	StatusCode int
	Body       []byte
}

func (recorder *ResponseRecorder) WriteHeader(statusCode int) {
	recorder.StatusCode = statusCode
	recorder.ResponseWriter.WriteHeader(statusCode)
}

func (recorder *ResponseRecorder) Write(data []byte) (int, error) {
	recorder.Body = data
	return recorder.ResponseWriter.Write(data)
}

func HttpLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		startTime := time.Now()
		recorder := &ResponseRecorder{
			ResponseWriter: res,
			StatusCode:     http.StatusOK,
		}
		handler.ServeHTTP(recorder, req)
		duration := time.Since(startTime)

		logger := log.Info()
		if recorder.StatusCode != http.StatusOK {
			logger = log.Error().Bytes("body", recorder.Body)
		}

		logger.Str("protocol", "http").
			Str("method", req.Method).
			Str("path", req.RequestURI).
			Int("status_code", recorder.StatusCode).
			Str("status_text", http.StatusText(recorder.StatusCode)).
			Dur("duration", duration).
			Msgf("received a HTTP request")
	})
}