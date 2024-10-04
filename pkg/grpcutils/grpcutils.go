// Package grpcutils provides utilities for gRPC services.
package grpcutils

import (
	"context"
	"time"

	"github.com/JorgeO3/flowcast/pkg/logger"
	"github.com/google/uuid"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// contextKey is a custom type for context keys in this package.
type contextKey string

const (
	requestIDKey contextKey = "RequestID"
	realIPKey    contextKey = "RealIP"
)

// ServerOption represents a functional option for configuring the gRPC server.
type ServerOption func(*serverOptions)

type serverOptions struct {
	logger             logger.Interface
	unaryInterceptors  []grpc.UnaryServerInterceptor
	streamInterceptors []grpc.StreamServerInterceptor
}

// WithLogger sets the logger for the server.
func WithLogger(l logger.Interface) ServerOption {
	return func(o *serverOptions) {
		o.logger = l
	}
}

// WithUnaryInterceptor adds a unary interceptor to the server.
func WithUnaryInterceptor(i grpc.UnaryServerInterceptor) ServerOption {
	return func(o *serverOptions) {
		o.unaryInterceptors = append(o.unaryInterceptors, i)
	}
}

// WithStreamInterceptor adds a stream interceptor to the server.
func WithStreamInterceptor(i grpc.StreamServerInterceptor) ServerOption {
	return func(o *serverOptions) {
		o.streamInterceptors = append(o.streamInterceptors, i)
	}
}

// NewServer creates a new gRPC server with the given options.
func NewServer(opts ...ServerOption) *grpc.Server {
	options := &serverOptions{}
	for _, opt := range opts {
		opt(options)
	}

	// Create server options
	var serverOpts []grpc.ServerOption

	// Collect all unary interceptors
	var unaryInterceptors []grpc.UnaryServerInterceptor

	// Add logger interceptor if logger is set
	if options.logger != nil {
		unaryInterceptors = append(unaryInterceptors, UnaryLoggingInterceptor(options.logger))
		options.streamInterceptors = append(options.streamInterceptors, StreamLoggingInterceptor(options.logger))
	}

	// Add custom unary interceptors
	if len(options.unaryInterceptors) > 0 {
		unaryInterceptors = append(unaryInterceptors, options.unaryInterceptors...)
	}

	// Chain unary interceptors
	if len(unaryInterceptors) > 0 {
		serverOpts = append(serverOpts, grpc.ChainUnaryInterceptor(unaryInterceptors...))
	}

	// Chain stream interceptors
	if len(options.streamInterceptors) > 0 {
		serverOpts = append(serverOpts, grpc.ChainStreamInterceptor(options.streamInterceptors...))
	}

	// Create and return the server
	return grpc.NewServer(serverOpts...)
}

// Middleware Implementations

// RequestIDInterceptor injects a unique RequestID into the context.
func RequestIDInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		id := uuid.New().String()
		ctx = context.WithValue(ctx, requestIDKey, id)
		return handler(ctx, req)
	}
}

// RealIPInterceptor retrieves the client's real IP address.
func RealIPInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		p, ok := peer.FromContext(ctx)
		if ok {
			addr := p.Addr.String()
			ctx = context.WithValue(ctx, realIPKey, addr)
		}
		return handler(ctx, req)
	}
}

// UnaryLoggingInterceptor logs gRPC unary requests using the provided logger.
func UnaryLoggingInterceptor(log logger.Interface) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		start := time.Now()

		// Get client peer information
		p, _ := peer.FromContext(ctx)
		clientIP := ""
		if p != nil {
			clientIP = p.Addr.String()
		}

		// Process the request
		resp, err := handler(ctx, req)
		duration := time.Since(start)

		// Log the request
		logFields := []interface{}{
			"grpc_method", info.FullMethod,
			"client_ip", clientIP,
			"duration_ms", duration.Milliseconds(),
		}

		if err != nil {
			log.Error("Handled gRPC request with error",
				append(logFields, "error", err.Error())...,
			)
		} else {
			log.Info("Handled gRPC request",
				logFields...,
			)
		}

		return resp, err
	}
}

// StreamLoggingInterceptor creates a new logger interceptor for streaming calls.
func StreamLoggingInterceptor(l logger.Interface) grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		start := time.Now()

		// Obtener información del cliente
		ctx := stream.Context()
		p, _ := peer.FromContext(ctx)
		clientIP := ""
		if p != nil {
			clientIP = p.Addr.String()
		}

		// Obtener el RequestID si está disponible
		requestID, _ := ctx.Value(requestIDKey).(string)

		// Procesar la solicitud
		err := handler(srv, stream)
		duration := time.Since(start)

		// Registrar la respuesta
		logFields := []interface{}{
			"grpc_method", info.FullMethod,
			"client_ip", clientIP,
			"duration_ms", duration.Milliseconds(),
			"is_client_stream", info.IsClientStream,
			"is_server_stream", info.IsServerStream,
		}

		if requestID != "" {
			logFields = append(logFields, "request_id", requestID)
		}

		if err != nil {
			l.Error("Handled gRPC stream request with error",
				append(logFields, "error", err.Error())...,
			)
		} else {
			l.Info("Handled gRPC stream request",
				logFields...,
			)
		}

		return err
	}
}

// RecoveryInterceptor handles panics and prevents server crashes.
func RecoveryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = status.Errorf(codes.Internal, "panic triggered: %v", r)
			}
		}()
		return handler(ctx, req)
	}
}

// RequestSizeLimitInterceptor limits the size of incoming requests.
func RequestSizeLimitInterceptor(maxSize int) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if reqProto, ok := req.(proto.Message); ok {
			if reqSize := proto.Size(reqProto); reqSize > maxSize {
				return nil, status.Errorf(codes.InvalidArgument, "request size exceeded: %d > %d", reqSize, maxSize)
			}
		}
		return handler(ctx, req)
	}
}

// TimeoutInterceptor adds a timeout to each request.
func TimeoutInterceptor(duration time.Duration) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		ctx, cancel := context.WithTimeout(ctx, duration)
		defer cancel()
		return handler(ctx, req)
	}
}

// RateLimitInterceptor limits the number of requests processed per unit of time.
func RateLimitInterceptor(r *rate.Limiter) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if err := r.Wait(ctx); err != nil {
			return nil, status.Errorf(codes.ResourceExhausted, "rate limit exceeded")
		}
		return handler(ctx, req)
	}
}

// JWTAuthInterceptor validates JWT tokens in the request metadata.
func JWTAuthInterceptor(validateToken func(string) error) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(codes.Unauthenticated, "missing metadata")
		}

		authHeader, ok := md["authorization"]
		if !ok || len(authHeader) == 0 {
			return nil, status.Errorf(codes.Unauthenticated, "missing authorization token")
		}

		token := authHeader[0]
		if err := validateToken(token); err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "invalid token: %v", err)
		}

		return handler(ctx, req)
	}
}

// LoggerErrorHandlingInterceptor handles errors and logs them appropriately.
func LoggerErrorHandlingInterceptor(l logger.Interface) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		resp, err := handler(ctx, req)
		if err != nil {
			l.Error("Server error occurred",
				"grpc_method", info.FullMethod,
				"error", err.Error(),
			)
		}
		return resp, err
	}
}
