package controllers

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

type OrderItemPack struct {
	Table_id    *string
	Order_items []models.OrderItem
}

var orderItemCollection *mongo.Collection = database.OpenCollection(database.Client, "orderItem")

func GetOrderItemsByOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	}
}

func GetOrderItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		result, err := orderItemCollection.Find(context.TODO(), bson.M{})
	}
}

func CreateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func ItemsByOrders(id string) (OrderItems []primitive.M, err error) {

}

func UpdateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
