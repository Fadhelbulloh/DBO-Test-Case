package controllers

import (
	"net/http"
	"strconv"

	"github.com/Fadhelbulloh/DBO-Test-Case/models"
	"github.com/Fadhelbulloh/DBO-Test-Case/params"
	"github.com/Fadhelbulloh/DBO-Test-Case/services"
	"github.com/Fadhelbulloh/DBO-Test-Case/utils"
	"github.com/gin-gonic/gin"
)

// GetOrders handles paginated list of orders
func GetOrders(c *gin.Context) {
	var (
		param      params.Order
		pagination params.Pagination
		err        error
	)
	pagination.Page, err = strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ReturnErrorTemplate(err, "page query param must be an integer"))
		return
	}

	pagination.Limit, err = strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ReturnErrorTemplate(err, "limit query param must be an integer"))
		return
	}

	pagination.Offset = (pagination.Page - 1) * pagination.Limit
	param.CustomerID = c.DefaultQuery("customer_id", "")
	param.Status = c.DefaultQuery("status", "")

	orders, err := services.GetOrders(pagination, param.CustomerID, param.Status)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ReturnErrorTemplate(err, "internal server error"))
		return
	}
	c.JSON(http.StatusOK, utils.ReturnSuccessTemplate(orders))
}

// GetOrderDetail handles logic of getting a single order
func GetOrderDetail(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, utils.ReturnErrorTemplate(nil, "id query param is required"))
		return
	}

	order, err := services.GetOrderDetail(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ReturnErrorTemplate(err, "internal server error"))
		return
	}

	c.JSON(http.StatusOK, utils.ReturnSuccessTemplate(order))
}

// AddOrder handles adding a new order
func AddOrder(c *gin.Context) {
	var newOrder models.Order

	if err := c.ShouldBindJSON(&newOrder); err != nil {
		c.JSON(http.StatusBadRequest, utils.ReturnErrorTemplate(err, err.Error()))
		return
	}

	err := services.AddOrder(newOrder)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ReturnErrorTemplate(err, "internal server error"))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order added successfully"})
}

// UpdateOrder handles updating an existing order
func UpdateOrder(c *gin.Context) {
	var dataOrder models.Order
	if err := c.ShouldBindJSON(&dataOrder); err != nil {
		c.JSON(http.StatusBadRequest, utils.ReturnErrorTemplate(err, err.Error()))
		return
	}

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, utils.ReturnErrorTemplate(nil, "id query param is required"))
		return
	}

	err := services.UpdateOrder(dataOrder, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ReturnErrorTemplate(err, "internal server error"))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "order updated successfully"})
}

// DeleteOrder handles deleting an existing order
func DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, utils.ReturnErrorTemplate(nil, "id query param is required"))
		return
	}

	err := services.DeleteOrder(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ReturnErrorTemplate(err, "internal server error"))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "order deleted successfully"})
}
