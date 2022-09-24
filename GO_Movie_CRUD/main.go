package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Movie struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Origin       string    `json:"origin"`
	Length_Hours float32   `json:"length_hours"`
	Producer     *Producer `json:"producer"`
}

type Producer struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicaiton/json")
	json.NewEncoder(w).Encode(movies)
}

func main() {
	fmt.Print("Welcome to Movie intro!")
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "01", Name: "Narco", Origin: "SK", Length_Hours: 2.5, Producer: &Producer{FirstName: "Kim", LastName: "Tony"}})
	r.HandleFunc("/movies", getMovies).Method("GET")
	r.HandleFunc("/movies/{id}", getMovie).Method("GET")
	r.HandleFunc("/movies", createMovie).Method("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Method("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Method("DELETE")

}
