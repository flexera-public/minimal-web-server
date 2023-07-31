package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
    fmt.Println("Minimal web server is listening on 0.0.0.0:8080!")
    fmt.Printf("/ is the only route available.\n\n")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Printf("I received a %s request from %s\n", r.Method, r.Host)
		fmt.Fprintf(w, "Hello, World!")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
