package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {

	param1 := r.URL.Query().Get("param1")

	if len(param1) < 1 {
		log.Println("Param1 are missing")
	}

	version := os.Getenv("VERSION")
	if version == "" {
		version = "v1"
	}
	fmt.Fprintf(w, "Hello World %s %s\n", version, param1)
}

func main() {
	log.Print("Hello world sample started.")
	http.HandleFunc("/api/hello", handler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
