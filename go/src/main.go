package main

import (
	"fmt"
	"log"
	"net/http"
)

func greeting(message string) string {
	return fmt.Sprintf("<b>%v</b>", message)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", greeting("Code.education Rocks!!!"))
}

func main() {

	http.HandleFunc("/", handler)

	port := "8000"
	log.Printf("Listening on port %s...", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
