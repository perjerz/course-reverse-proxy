package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(handler))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// config upstream server
	r.URL.Scheme = "https"
	r.URL.Host = "93.184.216.34"

	// forward request to upstream
	resp, err := http.DefaultTransport.RoundTrip(r)
	if err != nil {
		http.Error(w, "Bad Gateway", http.StatusBadGateway)
		log.Println(err)
		return
	}

	// copy response to client
	for k, v := range resp.Header {
		w.Header()[k] = v
	}
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
