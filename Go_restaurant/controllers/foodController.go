package controllers

import (
	"context"
	"net/http"
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
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while fetching the food"})
		}
		c.JSON(http.StatusOK, food)
	}
}

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var food models.Food
		var menu models.Menu

		if err := c.BindJSON(&food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
}

func round(num float64) int {

}

func toFixed(num float64) float64 {

}

func UpdateFood() gin.HandlerFunc {

}
