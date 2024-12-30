package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from the Server!")
	})
	fmt.Println("Server running on port 9090")
	http.ListenAndServe(":9090", nil)

	fmt.Println("Exited")
}

