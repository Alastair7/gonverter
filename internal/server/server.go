package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Alastair7/gonverter/api/router"
)

type GonverterServer struct {
	server *http.Server
}

func (g *GonverterServer) RunServer() {
	log.Printf("Server is starting in address %s...", g.server.Addr)
	log.Println("GET /checkhealth to check the server status.")

	serverErr := g.server.ListenAndServe()
	if serverErr != nil {
		log.Fatalf("Unexpected server error: %s", serverErr)
	}
}

func New(addr string) *GonverterServer {
	server := &http.Server{
		Addr:              addr,
		Handler:           router.SetupRouter(),
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      30 * time.Second,
	}

	return &GonverterServer{server: server}
}
