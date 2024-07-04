package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Define routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "Welcome to the Network API!")
	})

	http.HandleFunc("/check-net", func(w http.ResponseWriter, r *http.Request) {

		queryParam := r.URL.Query()


		fmt.Fprintf(w, "%s", queryParam)

	})

	// Start server
	port := ":8080"
	fmt.Printf("Server listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
