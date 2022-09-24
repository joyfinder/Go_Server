package main

import "fmt"

type Movie struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Origin   string    `json:"origin"`
	Length   int       `json:"length"`
	Producer *Producer `json:"producer"`
}

type Producer struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

func main() {
	fmt.Print("Welcome to Movie intro!")
	r := mux.NewRouter()

	r.HandleFunc("/movies", getMovies).Method("GET")
	r.HandleFunc("/movies/{id}", getMovie).Method("GET")
	r.HandleFunc("/movies", createMovie).Method("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Method("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Method("DELETE")

}
