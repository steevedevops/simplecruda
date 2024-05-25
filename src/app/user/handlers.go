package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUserHandler(ctx *gin.Context) {
	// var user User
	// if ctx.ShouldBind(&user) != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{
	// 		"message": "parametro vinculado errado",
	// 		"status":  http.StatusBadRequest,
	// 	})
	// }

	// userCreated, err := user.CreatUser(user)
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{
	// 		"message": userCreated,
	// 		"status":  http.StatusBadRequest,
	// 	})
	// }
	ctx.JSON(http.StatusCreated, gin.H{
		"usuario": "OK",
		// "usuario": userCreated,
	})
}

func FetchUserHandler(ctx *gin.Context) {
	var user *User
	names := ctx.Query("q")
	users, err := user.FetchUser(names)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err,
			"status":  http.StatusBadRequest,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"usuarios": users,
	})
}
