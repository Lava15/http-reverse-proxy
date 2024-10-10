package main

import (
	"github.com/lava15/http-reverse-proxy/config"
	"github.com/lava15/http-reverse-proxy/handlers"
	"github.com/lava15/http-reverse-proxy/proxy"
	"github.com/lava15/http-reverse-proxy/tlsconfig"
	"log"
	"net/http"
)

func main() {
	cfg := config.LoadConfig()
	proxyHandler, _ := proxy.NewProxy(cfg.BackendURL)
	tlsConf, _ := tlsconfig.LoadTLSConfig(cfg.CertFile, cfg.KeyFile)
	httpHandler := handlers.NewHandler(proxyHandler)
	server := &http.Server{
		Addr:      cfg.Address,
		Handler:   httpHandler,
		TLSConfig: tlsConf,
	}
	log.Printf("Starting HTTP/2 reverse proxy on %s", cfg.Address)
	if err := server.ListenAndServeTLS(cfg.CertFile, cfg.KeyFile); err != nil {
		log.Fatalf("Failed to start HTTP/2 reverse proxy on %s: %v", cfg.Address, err)
	}
}
