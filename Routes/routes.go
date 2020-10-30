package routes

import (
	"github.com/gin-gonic/gin"

	"github/Hiinnn/practice-go/controllers"
)

// SetupRouter -> Config router
func SetupRouter() *gin.Engine {
	r := gin.Default()
	group1 := r.Group("/user-api")
	{
		group1.GET("me", controllers.GetUserData)
		group1.POST("signin", controllers.Signin)
		group1.POST("register", controllers.Register)
	}
	return r
}
