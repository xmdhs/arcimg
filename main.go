package main

import (
	"log"
	"net/http"
	"time"

	"github.com/xmdhs/arcimg/arcimg"
)

func main() {
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      mux,
	}
	mux.HandleFunc("/favicon.ico", http.NotFound)
	mux.HandleFunc("/", arcimg.Chain(arcimg.Img, arcimg.Anticc, arcimg.Log))
	log.Println(server.ListenAndServe())

}
