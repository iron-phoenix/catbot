package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", helloWorldHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.ListenAndServe(":"+port, nil)
}

func helloWorldHandler(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("Hello world")
}
