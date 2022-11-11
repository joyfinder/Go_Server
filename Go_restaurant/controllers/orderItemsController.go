package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type OrderItemPack struct {
	Table_id    *string
	Order_items []models.OrderItem
}

var orderItemCollection *mongo.Collection = database.OpenCollection(database.Client, "orderItem")

func GetOrderItemsByOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		orderId := c.Param("order_id")

		allOrderItems, err := ItemsByOrders(orderId)

		if err != nil {
			c.JSON{http.StatusInternalServerError, gin.H{"error": "error occurred while listing order items by order"}}
			return
		}
		c.JSON(http.StatusOk, allOrderItems)
	}
}

func GetOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		orderItemId := c.Param("order_item_id")
	}
}

func CreateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func ItemsByOrder(id string) (OrderItems []primitive.M, err error) {

}

func UpdateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
