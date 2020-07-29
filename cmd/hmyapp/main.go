package main

import (
	"log"
	"net/http"

	_ "github.com/danceyoung/goslayer/cmd/hmyapp/router"
)

func main() {
	log.Println("dd")
	log.Println(http.ListenAndServe(":8080", nil))
}
