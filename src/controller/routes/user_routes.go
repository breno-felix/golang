package routes

import (
	"github.com/breno-felix/golang/src/controller"
	"github.com/gin-gonic/gin"
)

func InitUserRoutes(
	r *gin.RouterGroup,
	userController controller.UserControllerInterface,
) {
	r.GET("/:userId", userController.FindUserById)
	r.GET("/email/:userEmail", userController.FindUserByEmail)
	r.POST("/", userController.CreateUser)
	r.PUT("/:userId", userController.UpdateUser)
	r.DELETE("/:userId", userController.DeleteUser)
}
