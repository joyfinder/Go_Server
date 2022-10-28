package database

import (
	"fmt"
	"log"
)

func DBinstance() *mongo.Client {
	MongoDB := "mongodb://localhost:27017"
	fmt.Print(MongoDB)

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDB))
	if err != nil {
		log.Fatal(err)
	}
}
