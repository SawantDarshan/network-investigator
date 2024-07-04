package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func main() {
	// Define routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the Network API!")
	})

	http.HandleFunc("/check-host", func(w http.ResponseWriter, r *http.Request) {
		queryParam := r.URL.Query().Get("domain")
		ips, err := net.LookupHost(queryParam)
		if err != nil {
			http.Error(w,"Error: ", http.StatusInternalServerError)
			return
		}
		for _, ip := range ips {
			fmt.Fprintf(w, "IP address: %s", ip)
		}
	})

	// Start server
	port := ":8080"
	fmt.Printf("Server listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
