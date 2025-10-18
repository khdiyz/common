package httpserver

import (
	"context"
	"net/http"
	"strconv"
	"time"
)

const (
	// DefaultMaxHeaderBytes is the default maximum size of request headers (1 MB)
	DefaultMaxHeaderBytes = 1 << 20
	// DefaultReadTimeout is the default timeout for reading the entire request
	DefaultReadTimeout = 10 * time.Second
	// DefaultWriteTimeout is the default timeout for writing the response
	DefaultWriteTimeout = 10 * time.Second
)

// Config holds the configuration for the HTTP server.
type Config struct {
	Host           string        // Host is the hostname to bind to (empty for all interfaces)
	Port           int           // Port is the port number to listen on
	MaxHeaderBytes int           // MaxHeaderBytes is the maximum size of request headers
	ReadTimeout    time.Duration // ReadTimeout is the timeout for reading the entire request
	WriteTimeout   time.Duration // WriteTimeout is the timeout for writing the response
}

// Server wraps http.Server and provides methods for running and gracefully shutting down an HTTP server.
type Server struct {
	httpServer *http.Server
	config     *Config
}

// New creates a new Server instance with the provided configuration.
// If config is nil, default values will be used.
func New(config *Config) *Server {
	if config == nil {
		config = &Config{
			MaxHeaderBytes: DefaultMaxHeaderBytes,
			ReadTimeout:    DefaultReadTimeout,
			WriteTimeout:   DefaultWriteTimeout,
		}
	}

	// Apply defaults for zero values
	if config.MaxHeaderBytes == 0 {
		config.MaxHeaderBytes = DefaultMaxHeaderBytes
	}
	if config.ReadTimeout == 0 {
		config.ReadTimeout = DefaultReadTimeout
	}
	if config.WriteTimeout == 0 {
		config.WriteTimeout = DefaultWriteTimeout
	}

	return &Server{
		config: config,
	}
}

// Run starts the HTTP server with the given handler.
// It uses the configuration provided during server creation.
//
// Parameters:
//   - handler: the HTTP handler to serve requests
//
// Returns an error if the server fails to start or encounters a fatal error.
func (s *Server) Run(handler http.Handler) error {
	addr := s.buildAddress()

	s.httpServer = &http.Server{
		Addr:           addr,
		Handler:        handler,
		MaxHeaderBytes: s.config.MaxHeaderBytes,
		ReadTimeout:    s.config.ReadTimeout,
		WriteTimeout:   s.config.WriteTimeout,
	}

	return s.httpServer.ListenAndServe()
}

// Shutdown gracefully shuts down the server without interrupting active connections.
// It waits for active connections to finish or for the context to be canceled.
//
// Parameters:
//   - ctx: context to control the shutdown timeout
//
// Returns an error if the shutdown process fails.
func (s *Server) Shutdown(ctx context.Context) error {
	if s.httpServer == nil {
		return nil
	}
	return s.httpServer.Shutdown(ctx)
}

// buildAddress constructs the server address from host and port configuration.
func (s *Server) buildAddress() string {
	if s.config.Host != "" {
		return s.config.Host + ":" + strconv.Itoa(s.config.Port)
	}
	return ":" + strconv.Itoa(s.config.Port)
}
