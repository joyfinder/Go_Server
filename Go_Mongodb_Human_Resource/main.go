package main

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

var mg MongoInstance

const dbName = "human_resource"
const mongoURI = "mongodb://localhost:27068" + dbName
