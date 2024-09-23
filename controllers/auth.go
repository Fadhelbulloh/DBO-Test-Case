package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login handles login request
func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Login"})
}

// register handles register request
func Register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Register"})
}
