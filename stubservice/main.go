package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		dump, err := httputil.DumpRequest(r, false)
		if err != nil {
			fmt.Printf("error dumping request: %v", err)
			return
		}
		log.Printf("received a proxied request: %s\n", dump)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
