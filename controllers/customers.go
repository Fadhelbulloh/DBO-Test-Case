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

// GetCustomers handles paginated list of customers
func GetCustomers(c *gin.Context) {
	var (
		param      params.Customer
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
	param.Email = c.DefaultQuery("email", "")
	param.Name = c.DefaultQuery("name", "")
	param.Phone = c.DefaultQuery("phone", "")

	customers, err := services.GetCustomers(pagination, param.Email, param.Name, param.Phone)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ReturnErrorTemplate(err, "internal server error"))
		return
	}

	c.JSON(http.StatusOK, utils.ReturnSuccessTemplate(customers, pagination))
}

// GetCustomerDetail handles logic of getting a single customer
func GetCustomerDetail(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, utils.ReturnErrorTemplate(nil, "id query param is required"))
		return
	}

	customer, err := services.GetCustomerDetail(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ReturnErrorTemplate(err, "internal server error"))
		return
	}

	c.JSON(http.StatusOK, utils.ReturnSuccessTemplate(customer))
}

// AddCustomer handles adding a new customer
func AddCustomer(c *gin.Context) {
	var newCustomer models.Customer

	if err := c.ShouldBindJSON(&newCustomer); err != nil {
		c.JSON(http.StatusBadRequest, utils.ReturnErrorTemplate(err, "invalid payload"))
		return
	}

	err := services.AddCustomer(newCustomer)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ReturnErrorTemplate(err, "internal server error"))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Customer added successfully"})
}

// UpdateCustomer handles updating an existing customer
func UpdateCustomer(c *gin.Context) {
	var dataCustomer models.Customer

	if err := c.ShouldBindJSON(&dataCustomer); err != nil {
		c.JSON(http.StatusBadRequest, utils.ReturnErrorTemplate(err, "invalid payload"))
		return
	}

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, utils.ReturnErrorTemplate(nil, "id query param is required"))
		return
	}

	err := services.UpdateCustomer(dataCustomer, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ReturnErrorTemplate(err, "internal server error"))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Update customer success"})
}

// DeleteCustomer handles deleting an existing customer
func DeleteCustomer(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, utils.ReturnErrorTemplate(nil, "id query param is required"))
		return
	}

	err := services.DeleteCustomer(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ReturnErrorTemplate(err, "internal server error"))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Delete customer success"})
}
