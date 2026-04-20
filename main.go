package main

import (
	"ascii-art-web/handler"
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handler.HomeHandler)
	http.HandleFunc("/asciiart", handler.AsciiHomeHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Server running on port:", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}