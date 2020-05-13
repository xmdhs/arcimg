package main

import (
	"arcimg/arcimg"
	"log"
	"net/http"
)

func main() {
        go arcimg.Remove()
	http.HandleFunc("/img.png", arcimg.Img)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
