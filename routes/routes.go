package routes

import (
	"github.com/Fadhelbulloh/DBO-Test-Case/controllers"
	"github.com/Fadhelbulloh/DBO-Test-Case/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// Auth routes
	auth := r.Group("/auth")
	{
		auth.POST("/login", controllers.Login)
		auth.POST("/register", controllers.Register)
	}

	// Customer routes
	customers := r.Group("/customers", middleware.AuthMiddleware)
	{
		customers.GET("/", controllers.GetCustomers)
		customers.GET("/:id", controllers.GetCustomerDetail)
		customers.POST("/", controllers.AddCustomer)
		customers.PUT("/:id", controllers.UpdateCustomer)
		customers.DELETE("/:id", controllers.DeleteCustomer)
	}

	//Order routes
	orders := r.Group("/orders", middleware.AuthMiddleware)
	{
		orders.GET("/", controllers.GetOrders)
		orders.GET("/:id", controllers.GetOrderDetail)
		orders.POST("/", controllers.AddOrder)
		orders.PUT("/:id", controllers.UpdateOrder)
		orders.DELETE("/:id", controllers.DeleteOrder)
	}
}
