package main

import (
	"log"
	"net/http"

	_ "github.com/danceyoung/goslayer/cmd/hmyapp/router"
)

func main() {
	log.Println("[http Handler] Listening and serving HTTP on :8080")
	log.Println(http.ListenAndServe(":8080", nil))
}
