package main

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-drive/mongo"
	// "go.mongodb.org/mongo-drive/mongo/options"
	// "go.mongodb.org/mongo-drive/bson"
	// "go.mongodb.org/mongo-drive/bson/primitive"
)

// Data type: includes Client & db
type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

type Employee struct {
	ID     string  `json: "id,omitempty" bson:"_id, omitempty"`
	Age    string  `json: "name"`
	Name   string  `json: "salary"`
	Salary float64 `json: "age"`
}

func Connect() error {

	// Connect to client host
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancer()

	err = client.Connect(ctx)
	db := client.Database(dbName)

	if err != nil {
		return err
	}

	mg = MongoInstance{
		Client: client,
		Db:     db,
	}
	return nil
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
