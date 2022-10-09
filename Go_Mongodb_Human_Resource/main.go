package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/pelletier/go-toml/query"
	"go.mongodb.org/mongo-drive/mongo"
	"gopkg.in/mgo.v2/bson"
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
		// Slice: functioning similarly with Array, but containing objects
		// In other words, storing number of employees' id
		query := bson.D{{}}

		cursor, err := mg.dB.Collection("employee").Find(c.Context(), query)

		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		var employees []Employee = make([Employee, 0])

		if err := cursor.All(c.Context(), &employees), err != nil {
			return c.Status(500).SendString(err.Error())
		}

		return c.JSON(employees)
	})
	app.Post("/employee", func (c *fiber.Ctx) error {
		collection := mg.Db.Collection("employees")

		employee := new(Employee)

		if err := c.BodyParser(employee); err != nil {
			return c.Status(400).SendString(err.Error())	
		}

		employee.ID = ""

		insertionResult, err := collection.InsertOne(c.Context(), employee)
		if err := nil {
			return c.Status(400).SendString(err.Error())
		}
	})
	app.Put("/employee/:id")
	app.Delete("/employee/:id")
}
