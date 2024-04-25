package controller

import (
	"net/http"

	logger "github.com/articoigor/crud_go_lang/src/configuration/logging"
	validator "github.com/articoigor/crud_go_lang/src/configuration/validation"
	"github.com/articoigor/crud_go_lang/src/controller/model/request"
	"github.com/articoigor/crud_go_lang/src/model"
	"github.com/articoigor/crud_go_lang/src/model/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RegisterUser(c *gin.Context) {
	logger.Info("registerUser CONTROLLER STARTED",
		zap.String("method", "RegisterUser"),
	)

	var user_request request.UserRequest

	if err := c.ShouldBindJSON((&user_request)); err != nil {
		logger.Error("Error trying to validate user input.", err,
			zap.String("method", "RegisterUser"),
		)

		mapped_err := validator.ValidateUserError(err)

		c.JSON(mapped_err.Code, mapped_err)

		return
	}

	domain := model.NewUserDomain(
		user_request.Name,
		user_request.Email,
		user_request.Password,
		user_request.Age,
	)

	service := service.NewUserService(domain)

	if err := service.RegisterUser(domain); err != nil {
		c.JSON(err.Code, err)

		return
	}

	logger.Info("User registered succesfully.",
		zap.String("method", "RegisterUser"),
	)

	c.String(http.StatusOK, "")
}

func UpdateUser(c *gin.Context) {

}

func RetrieveUserById(c *gin.Context) {

}

func RetrieveUserByEmail(c *gin.Context) {

}

func DeleteUser(c *gin.Context) {

}
