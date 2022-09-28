package main

import (
	"github.com/gofiber/fiber/v2"
)

type MongoInstance struct {
	Client
	DB
}

type Employee struct {
	ID     string
	Age    string
	Name   string
	Salary float64
}

func Connect() error {

}

var mg MongoInstance

const dbName = "human_resource"
const mongoURI = "mongodb://localhost:27068" + dbName

func main() {
	app := fiber.New()

	app.Get("/employee")
	app.Post()
	app.Put()
	app.Delete()
}
