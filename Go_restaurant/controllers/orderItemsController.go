package controllers

import "github.com/gin-gonic/gin"

type OrderItemPack struct {
	Table_id    *string
	Order_items []models.OrderItem
}

var orderItemCollection *mongo.Collection = database.OpenCollection(database.Client, "orderItem")

func GetOrderItemsByOrder() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func GetOrderItems() gin.HandlerFunc {
	return func(c *gin.Context) {

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
