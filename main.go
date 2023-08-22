/* Sample Web Application
   Author : Ansil H
   Email: ansilh@gmail.com */

package main

import (
	"fmt"
	"net/http"
        "log"
)

func demoDefault(w http.ResponseWriter, r *http.Request) {

	data := `<html><head></head><title></title><body><div> <h2>Hello World!</h2>`
	fmt.Fprintf(w, data+`</div></body></html>`)
}

func main() {
	http.HandleFunc("/", demoDefault)
        log.Printf("Starting web server on port 9090")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
