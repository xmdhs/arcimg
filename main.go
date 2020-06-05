package main

import (
	"arcimg/arcimg"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/http2"
)

func main() {
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		Handler:      mux,
	}
	Middleware := arcimg.NewMiddleware(arcimg.Img)
	Middleware.Add(arcimg.Anticc)
	Middleware.Add(arcimg.Log)
	mux.HandleFunc("/favicon.ico", http.NotFound)
	mux.HandleFunc("/", Middleware.Use)
	http2.ConfigureServer(server, &http2.Server{})
	log.Println(server.ListenAndServe())

}
