package main

import (
	"log"

	"github.com/norapanel/nora-node/internal/certutil"
	appconfig "github.com/norapanel/nora-node/internal/config"
	"github.com/norapanel/nora-node/internal/node"
)

func main() {
	settings := appconfig.Load()

	if err := certutil.EnsureServerCertificate(settings.SSLCertFile, settings.SSLKeyFile); err != nil {
		log.Fatalf("failed to prepare TLS certificate: %v", err)
	}

	server, err := node.New(settings)
	if err != nil {
		log.Fatalf("failed to initialize node service: %v", err)
	}

	log.Printf("Nora Node gRPC service running on %s:%d", settings.ServiceHost, settings.ServicePort)
	if err := server.ListenAndServeGRPC(); err != nil {
		log.Fatal(err)
	}
}
