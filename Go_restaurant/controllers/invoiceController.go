package controllers

import "github.com/gin-gonic/gin"

type InvoiceViewFormat struct {
	Invoice_id
	Payment_method
	Order_id
	Payment_status
	Payment_due
	Table_number
	Payment_due_date
	Order_details
}

func GetInvoices() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func GetInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func CreateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func UpdateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
