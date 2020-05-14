package main

import (
	"arcimg/arcimg"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/http2"
)

func main() {
	go arcimg.Remove()
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		Handler:      mux,
	}
	mux.HandleFunc("/img.png", arcimg.Img)
	http2.ConfigureServer(server, &http2.Server{})
	log.Fatal(server.ListenAndServe())
}
