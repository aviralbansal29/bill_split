package server

import (
	"log"
	"net/http"

	"github.com/aviralbansal29/bill_split/config"
	httpRoutes "github.com/aviralbansal29/bill_split/transport/http"
	"github.com/labstack/echo/v4"
)

// RestServer contains global objects
type RestServer struct {
	mux    *echo.Echo
	server *http.Server
}

// Setup creates router instance
func (s *RestServer) Setup() {
	s.mux = echo.New()
	s.server = &http.Server{
		Addr:    config.GetEnv().GetString("server_address"),
		Handler: s.mux,
	}
	httpRoutes.SetupRouter(s.mux)
}

// Start starts the server
func (s *RestServer) Start() {
	if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

// Stop stops the server
func (RestServer) Stop() {

}
