package main

import (
	_ "goslayer/cmd/myapp/router"
	"log"
	"net/http"
)

func main() {
	log.Println(http.ListenAndServe(":8080", nil))
}