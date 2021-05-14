// +build main

package main

import (
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/xmdhs/arcimg/arcimg"
)

func main() {
	mux := httprouter.New()
	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      mux,
	}
	mux.HandlerFunc("GET", "/user/:uid/img", arcimg.Chain(arcimg.Img, arcimg.Anticc, arcimg.Log))
	log.Println(server.ListenAndServe())

}
