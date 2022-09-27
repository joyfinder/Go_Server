package main

import "time"

var rdr *renderer.Render
var db *mgo.Database

const (
	hostname       string = "localhost:27001"
	dbName         string = "demo_todo"
	collectionName string = "todo"
	port           string = ":9000"
)

type (
	todoModel struct {
		ID        bson.ObjectId `bson:"_id,omitempty"`
		Title     string        `bson:"title"`
		Completed bool          `bson:"completed"`
		CreatedAt time.Time     `bson:"createAt"`
	}
	todo struct {
		ID        string    `json:"id"`
		Title     string    `json:"title"`
		Completed bool      `bson:"completed"`
		CreatedAt time.Time `json:"created_at`
	}
)

func main() {

}
