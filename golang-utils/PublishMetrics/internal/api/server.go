package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// Server abstraction of a server
type Server interface {
	ListenAndServe() error
	Shutdown(ctx context.Context) error
}

type server struct {
	httpPort   string
	httpServer *http.Server
}

// NewServer creates a new server
func NewServer() (Server, error) {
	result := &server{}

	// find HTTP port to listen to
	//port := os.Getenv("PORT")
	port := "8085"
	if port == "" {
		return nil, fmt.Errorf("PORT environment variable not defined")
	}
	result.httpPort = port

	// setup HTTP routes
	r := mux.NewRouter()

	// IP Pool methods
	ipPool := r.PathPrefix("/metrics").Subrouter()
	ipPool.HandleFunc("/publish", result.PostInstanceMessageHandler).Methods("POST")

	srv := &http.Server{
		Addr: fmt.Sprintf("0.0.0.0:%v", port),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	result.httpServer = srv
	return result, nil
}

// ListenAndServe listen and server the http server and close connections of db
func (s *server) ListenAndServe() error {
	logrus.Infof("Staring HTTP server on port %s", s.httpPort)
	return s.httpServer.ListenAndServe()
}

// Shutdown shutdown the http server
func (s *server) Shutdown(ctx context.Context) error {
	logrus.Infof("Shutting down API server")
	// shutdown server
	err := s.httpServer.Shutdown(ctx)

	return err
}
