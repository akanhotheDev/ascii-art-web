package main

import (
	"ascii-art-web/handler"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.HomeHandler)
	http.HandleFunc("/asciiart", handler.AsciiHomeHandler)

	fmt.Println("Server running at http://localhost:8080")
	//http.ListenAndServe(":8080", nil)
	err := http.ListenAndServe(":8080", nil)
if err != nil {
	fmt.Println("Server error:", err)

	
}
}
