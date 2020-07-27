package middleware

import (
	"log"
	"net/http"
)

func HttpSet(hf func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		log.Println(req.Method+" the url requesting is ", req.URL)
		rw.Header().Set("Content-Type", "application/json")
		hf(rw, req)
	})
}
