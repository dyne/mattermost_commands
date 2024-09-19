package main

import (
	"net/http"
	"time"
)

func main() {
	handler := http.HandlerFunc(MiniServer)

	server := &http.Server{
		Addr:              ":4444",
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
