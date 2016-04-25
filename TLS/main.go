package main

import (
	"io"
	"log"
	"net/http"
)

func main() {

	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "I know YOU are awesome ")
	})
	log.Println("Listening to 80 ...")
	go http.ListenAndServe(":80", http.RedirectHandler("https://localhost:443/", http.StatusMovedPermanently))

	log.Println("Listening to 443 ...")
	err := http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil)
	logError(err)
}
func logError(err error) {
	if err != nil {
		log.Println(err)
	}
}
