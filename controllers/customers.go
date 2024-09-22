package controllers

import (
	"net/http"

	"github.com/Fadhelbulloh/DBO-Test-Case/models"
	"github.com/gin-gonic/gin"
)

// GetCustomers handles paginated list of customers
func GetCustomers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get customers"})
}

// AddCustomer handles adding a new customer
func AddCustomer(c *gin.Context) {
	var newCustomer models.Customer

	if err := c.ShouldBindJSON(&newCustomer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Customer added successfully"})
}

// UpdateCustomer handles updating an existing customer
func UpdateCustomer(c *gin.Context) {
	var dataCustomer models.Customer
	id := c.Param("id")
	_ = id
	if err := c.ShouldBindJSON(&dataCustomer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Update customer"})
}

// DeleteCustomer handles deleting an existing customer
func DeleteCustomer(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Delete customer"})
}
