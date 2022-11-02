package controllers

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		foodId := c.Param("food_id")
		var food models.Food

		foodCollection.FindOne(ctx, bson, M{"food_id": foodId}).Decode(&food)
		defer cancel()
	}
}

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func round(num float64) int {

}

func toFixed(num float64) float64 {

}

func UpdateFood() gin.HandlerFunc {

}
