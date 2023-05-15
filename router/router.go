package router

import (
	"github.com/gin-gonic/gin"
	"github.com/labadiejoaco/gin-gorm-api/controllers"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/api/posts", controllers.GetPosts)
	r.GET("/api/posts/:id", controllers.GetPostById)
	r.POST("/api/posts", controllers.CreatePost)
	r.PATCH("/api/posts/:id", controllers.UpdatePost)
	r.DELETE("/api/posts", controllers.DeletePosts)
	r.DELETE("/api/posts/:id", controllers.DeletePostById)

	return r
}
