package controllers

import (
	"github/Hiinnn/practice-go/models"
	"github/Hiinnn/practice-go/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Signup -> Create new user
func Signup(c *gin.Context) {
	var user models.User
	c.Bind(&user)
	err := services.Signup(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
	} else {
		c.JSON(http.StatusOK, user)
	}
}
