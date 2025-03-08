package routes

import (
	"github.com/breno-felix/golang/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(
	r *gin.Engine,
	userController controller.UserControllerInterface,
) {
	InitUserRoutes(r.Group("/users"), userController)
}
