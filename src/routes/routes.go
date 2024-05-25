package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/steevedevops/simplecruda/src/app/user"
	"github.com/steevedevops/simplecruda/src/shared/middleware"
)

func Routes(r *gin.Engine) {
	r.GET("/users", middleware.StarMigle, user.FetchUserHandler)
	r.POST("/users", user.CreateUserHandler)
}
