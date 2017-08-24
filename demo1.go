package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/foo", fooHandler)
	fmt.Println("Listening on 3000....")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Golang")
}
