package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if err := http.ListenAndServe(":"+port, http.FileServer(http.Dir("./"))); err != nil {
		log.Fatalf("server: %s", err)
	}
}
