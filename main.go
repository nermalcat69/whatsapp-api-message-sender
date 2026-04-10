package main

import (
	"fmt"
	"net/http"
	"os"

	handler "whatsapp-sender/api"
)

func main() {
	http.HandleFunc("/", handler.Index)
	http.HandleFunc("/parse-vcf", handler.Handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "6979"
	}
	fmt.Printf("Server running at http://localhost:%s\n", port)
	http.ListenAndServe(":"+port, nil)
}
