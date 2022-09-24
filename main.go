package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request){
	// Checking whether path is correct
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET"{
		http.Error(w, "method is not supported", http.StatusNotFound)
	}
}

func main() {

	fileServer := http.FileServer(http.Dir("./server_web"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Print("Start server port 8080!\n")
	if err := http.ListenAndServe(":8080", nil); err != nil (
		log.Fatal(err)
	)
		
}
