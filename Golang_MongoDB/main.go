package main

import (
	"net/http"

	"github.com/Go_Server/Golang_MongoDB/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {

	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

func getSession() *mgo.Session {

	// Helps connect MongoDB
	s, err := mgo.Dial("mongodb://localhost:28088")
	if err != nil {
		panic(err)
	}
	return s
}
