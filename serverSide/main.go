package main

import (
	"fmt"
	"net/http"
)

func main() {

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// CORS 
		w.Header().Add("Access-Control-Allow-Origin", "http://127.0.0.1:5500") 
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Add("Access-Control-Allow-Methods", "POST")
		w.Header().Add("Access-Control-Allow-Credentials", "true")	

		fmt.Fprintf(w, "{message: Hello World}")
	})

		http.ListenAndServe(":4000", nil)
		fmt.Println("Server is running on port 4000")
}

