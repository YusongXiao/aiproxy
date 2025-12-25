package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	// The target URL for Google Gemini API
	target := "https://generativelanguage.googleapis.com"
	targetURL, err := url.Parse(target)
	if err != nil {
		log.Fatalf("Failed to parse target URL: %v", err)
	}

	// Create a reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	// Configure the proxy to flush immediately for streaming responses
	proxy.FlushInterval = -1

	// Custom error handler
	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("Proxy error: %v", err)
		w.WriteHeader(http.StatusBadGateway)
	}

	// Modify the Director to ensure the Host header is set correctly
	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		// Important: Set the Host header to the target's host
		// Google's servers require the Host header to match the target domain
		req.Host = targetURL.Host
	}

	// Create a handler that uses the proxy
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Serve robots.txt
		if r.URL.Path == "/robots.txt" && r.Method == http.MethodGet {
			http.ServeFile(w, r, "robots.txt")
			return
		}

		// Serve the index page for the root path
		if r.URL.Path == "/" && r.Method == http.MethodGet {
			http.ServeFile(w, r, "index.html")
			return
		}

		log.Printf("Proxying request: %s %s", r.Method, r.URL.String())
		proxy.ServeHTTP(w, r)
	})

	// Start the server
	port := ":8080"
	log.Printf("Starting proxy server on %s", port)
	log.Printf("Target URL: %s", target)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
