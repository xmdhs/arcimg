package main

import (
	"arcimg/arcimg"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/http2"
)

func main() {
	/*	args := os.Args
		if args[1] == "true" {
			arcimg.Logoutfiles = true
			go arcimg.Logw()
		} else if args[1] == "false" {
			arcimg.Logoutfiles = false
		} else {
			os.Exit(0)
		}*/
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
	log.Println(server.ListenAndServe())

}
