package router

import (
	"github.com/gin-gonic/gin"

	"my_diary/diary"
	"my_diary/user"
	"my_diary/middleware"
)

func initRoutes(r *gin.RouterGroup) {
	user.RegisterRoutes(r)
	diary.RegisterRoutes(r)
}

func RouterHelper() *gin.Engine {
	r := gin.Default()
	router := r.Group("/diary")
	router.Use(middleware.AuthMiddleware())
	router.Use(middleware.CorsMiddleware())
	initRoutes(router)
	return r
}