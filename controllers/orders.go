package controllers

import (
	"net/http"

	"github.com/Fadhelbulloh/DBO-Test-Case/models"
	"github.com/gin-gonic/gin"
)

// GetOrders handles paginated list of orders
func GetOrders(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get orders"})
}

// AddOrder handles adding a new order
func AddOrder(c *gin.Context) {
	var newOrder models.Order

	if err := c.ShouldBindJSON(&newOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order added successfully"})
}

// UpdateOrder handles updating an existing order
func UpdateOrder(c *gin.Context) {
	var dataOrder models.Order
	id := c.Param("id")
	_ = id
	if err := c.ShouldBindJSON(&dataOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Update order"})
}

// DeleteOrder handles deleting an existing order
func DeleteOrder(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Delete order"})
}
