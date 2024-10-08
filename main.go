package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var (
	backendUrl *url.URL
	err        error
)

func init() {
	backendUrl, err = url.Parse("http://localhost:8080")
	if err != nil {
		log.Fatalf("Failed to parse backend URL: %v", err)
	}
}
func main() {
	server := &http.Server{
		Addr:    ":8443",
		Handler: proxyHandler(),
	}

	cert, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
	if err != nil {
		log.Fatalf("Failed to load TLS certificates: %v", err)
	}

	server.TLSConfig = &tls.Config{
		Certificates: []tls.Certificate{cert},
		NextProtos:   []string{"h2", "http/1.1"},
	}

	log.Println("Starting HTTP/2 reverse proxy on https://localhost:8443")
	err = server.ListenAndServeTLS("", "")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func proxyHandler() http.Handler {
	proxy := httputil.NewSingleHostReverseProxy(backendUrl)
	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
	}
	proxy.ModifyResponse = func(resp *http.Response) error {
		return nil
	}
	return proxy
}
