package routes

import (
	"github.com/articoigor/crud_go_lang/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/retrieveUserById/:userId", controller.RetrieveUserById)
	r.GET("/retrieveUserByEmail/:userEmail", controller.RetrieveUserByEmail)
	r.POST("/registerUser", controller.RegisterUser)
	r.PUT("/updateUser/:userId", controller.UpdateUser)
	r.DELETE("/deleteUser/:userId", controller.DeleteUser)
}
