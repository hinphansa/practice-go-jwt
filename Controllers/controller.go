package controllers

import (
	"github/Hiinnn/practice-go/models"
	"github/Hiinnn/practice-go/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Register -> Create new user
func Register(c *gin.Context) {
	var user models.User
	c.Bind(&user)
	err := services.Register(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// Signin -> Check username & password, Create Token
func Signin(c *gin.Context) {
	var user models.User
	c.Bind(&user)
	token, err := services.Signin(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}

// GetUserData -> Get your own data by parsing Token
func GetUserData(c *gin.Context) {
	var (
		data interface{}
		err  error
	)

	token := strings.Split(c.Request.Header.Get("Authorization"), " ")
	if len(token) > 0 && token[0] == "bearer" {
		data, err = services.GetUSerData(token[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
		} else {
			c.JSON(http.StatusOK, data)
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
	}
}
