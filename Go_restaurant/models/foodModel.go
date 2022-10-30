package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Food struct {
	ID    primitive.ObjectID `bson:"_id"`
	Name  *string
	Price *float64
}
