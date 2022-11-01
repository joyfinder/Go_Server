package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderItem struct {
	ID            primitive.ObjectID
	Quantity      *string
	Unit_price    *float64
	Order_Date    time.Time `json:"created_at"`
	Created_at    time.Time `json:"created_at"`
	Updated_at    time.Time `json:"updated_at"`
	food_id       *string   `json:"food_id" validate:"required"`
	Order_item_id string    `json:"order_item_id"`
	Order_id      string    `json:"order_id" validate:"required"`
}
