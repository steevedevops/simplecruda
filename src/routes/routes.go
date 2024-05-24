package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/steevedevops/simplecruda/src/app/user"
)

func Routes(r *gin.Engine) {
	r.GET("/users", user.FetchUserHandler)
	r.POST("/users", user.CreateUserHandler)
}
