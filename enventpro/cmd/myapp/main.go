package main

import (
	_ "enventpro/cmd/myapp/router"
	"log"
	"net/http"
)

func main() {
	log.Println(http.ListenAndServe(":8080", nil))
}