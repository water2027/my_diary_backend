package router

import (
	"github.com/gin-gonic/gin"

	"my_diary/diary"
	"my_diary/user"
	"my_diary/middleware"
)

func initRoutes(r *gin.RouterGroup) {
	// r.OPTIONS("/*path", middleware.CorsMiddleware())
	user.RegisterRoutes(r)
	diary.RegisterRoutes(r)
}

func RouterHelper() *gin.Engine {
	r := gin.Default()
	router := r.Group("/diary")
	router.Use(middleware.CorsMiddleware())
	router.Use(middleware.AuthMiddleware())
	initRoutes(router)
	return r
}