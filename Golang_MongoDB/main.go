package main

import (
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {

	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("")
	r.POST("")
	r.DELETE("")
}

func getSession() *mgo.Session {

	// Helps connect MongoDB
	s, err := mgo.Dial("mongodb://localhost:28088")
	if err != nil {
		panic(err)
	}
	return s
}
