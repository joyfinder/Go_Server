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

func main() {
	fmt.Print("Hello")
}
