package main

import "fmt"

type Movie struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Origin   string `json:"origin"`
	Length   int    `json:"length"`
	Producer string `json:"producer"`
}

func main() {
	fmt.Print("Hello")
}
