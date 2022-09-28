package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-drive/mongo"
	// "go.mongodb.org/mongo-drive/mongo/options"
	// "go.mongodb.org/mongo-drive/bson"
	// "go.mongodb.org/mongo-drive/bson/primitive"
)

type MongoInstance struct {
	Client
	Db
}

type Employee struct {
	ID     string
	Age    string
	Name   string
	Salary float64
}

func Connect() error {

	// Connect to client host
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
}

var mg MongoInstance

const dbName = "human_resource"
const mongoURI = "mongodb://localhost:27068" + dbName

func main() {

	if err := Connect(); err != nil {
		log.Fatal(err)
	}
	app := fiber.New()

	app.Get("/employee", func(c *fiber.Ctx) error {

	})
	app.Post("/employee")
	app.Put("/employee/:id")
	app.Delete("/employee/:id")
}
