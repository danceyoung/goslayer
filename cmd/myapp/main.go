package main

import (
	"log"
	"net/http"

	_ "github.com/danceyoung/goslayer/cmd/myapp/router"
)

func main() {
	log.Println(http.ListenAndServe(":8080", nil))
}
