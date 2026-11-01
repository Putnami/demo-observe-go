package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":8080"
	}
	http.HandleFunc("/quote", quote)
	log.Fatal(http.ListenAndServe(addr, nil))
}
