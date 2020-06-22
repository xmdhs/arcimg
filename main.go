package main

import (
	"arcimg/arcimg"
	"log"
	"net/http"
	"time"
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
	log.Println(server.ListenAndServe())

}
