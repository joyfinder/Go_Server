package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST reqeust successful")
	name := r.FormValue("name")
	phone_num := r.FormValue("phone_number")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Phone number = %s\n", phone_num)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Checking whether path is correct
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello")
}

func main() {

	fileServer := http.FileServer(http.Dir("./server_web"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Print("Start server port 8080!\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
