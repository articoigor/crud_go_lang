package controller

import (
	"fmt"

	validator "github.com/articoigor/crud_go_lang/src/configuration/validation"
	"github.com/articoigor/crud_go_lang/src/controller/model/request"
	"github.com/articoigor/crud_go_lang/src/controller/model/response"
	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	var user_request request.UserRequest

	if err := c.ShouldBindJSON((&user_request)); err != nil {
		mapped_err := validator.ValidateUserError(err)

		c.JSON(mapped_err.Code, mapped_err)

		return
	}

	fmt.Println(user_request)

	response := response.UserResponse{
		Id:    "1",
		Name:  "Igor",
		Email: "igor@email.com",
		Age:   25,
	}

	fmt.Println(response)
}
