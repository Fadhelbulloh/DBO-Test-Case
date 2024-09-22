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
		auth.POST("/login", nil)
		auth.POST("/register", nil)
	}

	// Customer routes
	customers := r.Group("/customers", middleware.AuthMiddleware)
	{
		customers.GET("/", controllers.GetCustomers)
		customers.POST("/", controllers.AddCustomer)
		customers.PUT("/:id", nil)
		customers.DELETE("/:id", nil)
	}

	//Order routes
	orders := r.Group("/orders", middleware.AuthMiddleware)
	{
		orders.GET("/", nil)
		orders.POST("/", nil)
		orders.PUT("/:id", nil)
		orders.DELETE("/:id", nil)
	}
}
