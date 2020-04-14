package main

import (
	"go-blog/bootstrap"
	"log"
	"net/http"
)

func main() {
	bootstrap.Bootstrap();

	//http.Handle("/", service.GetMyUrlHandler())

	log.Fatal(http.ListenAndServe(":8080", nil))
}
