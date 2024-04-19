package routes

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/retrieveUserById/:userId", controller.retrieveUserById)
	r.GET("/retrieveUserByEmail/:userEmail", controller.retrieveUserByEmail)
	r.POST("/registerUser", controller.registerUser)
	r.PUT("/updateUser/:userId", controller.updateUser)
	r.DELETE("/deleteUser/:userId", controller.deleteUser)
}
