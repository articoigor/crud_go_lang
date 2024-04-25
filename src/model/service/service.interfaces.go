package service

import (
	error_map "github.com/articoigor/crud_go_lang/src/configuration/error_mapping"
	"github.com/articoigor/crud_go_lang/src/model"
)

func NewUserService(model.UserDomainInterface) UserServiceInterface {
	return &userServiceInterface{}
}

type userServiceInterface struct {
}

type UserServiceInterface interface {
	RegisterUser(model.UserDomainInterface) *error_map.CurrError
	RetrieveUser(string) (*model.UserDomainInterface, *error_map.CurrError)
	DeleteUser(string) *error_map.CurrError
	UpdateUser(string, model.UserDomainInterface) *error_map.CurrError
}
