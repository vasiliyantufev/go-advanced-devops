package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/http2"
)

func main() {
	client := &http.Client{}

	// Create a pool with the server certificate since it is not signed
	// by a known CA
	caCert, err := os.ReadFile("server.crt")
	if err != nil {
		log.Fatalf("Reading server certificate: %s", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Create TLS configuration with the certificate of the server
	tlsConfig := &tls.Config{
		RootCAs: caCertPool,
	}

	// Use the proper transport in the client
	client.Transport = &http2.Transport{
		TLSClientConfig: tlsConfig,
	}

	// Perform the request
	resp, err := client.Post("https://localhost:9191/hello/sayHello", "text/plain", bytes.NewBufferString("Hello Go!"))
	if err != nil {
		log.Fatalf("Failed get: %s", err)
	}
	defer resp.Body.Close()
}
