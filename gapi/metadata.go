package gapi

import (
	"context"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
)

type Metadata struct {
	UserAgent string
	ClientIP  string
}

func (server *Server) extractMetadata(ctx context.Context) (*Metadata, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Internal, "failed to get metadata")
	}

	// Extract User-Agent and Client IP from metadata
	userAgents := md.Get("user-agent")
	clientIPs := md.Get("x-forwarded-for")
	
	if len(userAgents) == 0 {
		userAgents = md.Get("grpcgateway-user-agent")
	}
	
	if len(clientIPs) == 0 {
		clientIPs = md.Get("x-real-ip")
	}

	metadata := &Metadata{
		UserAgent: "",
		ClientIP:  "",
	}

	peer, ok := peer.FromContext(ctx)
	if ok {
		metadata.ClientIP = peer.Addr.String()
	}

	if len(userAgents) > 0 {
		metadata.UserAgent = userAgents[0]
	}
	
	if len(clientIPs) > 0 {
		metadata.ClientIP = clientIPs[0]
	}

	return metadata, nil
}