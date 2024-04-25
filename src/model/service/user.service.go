package service

import (
	"fmt"

	error_map "github.com/articoigor/crud_go_lang/src/configuration/error_mapping"
	logger "github.com/articoigor/crud_go_lang/src/configuration/logging"
	"github.com/articoigor/crud_go_lang/src/model"
	"go.uber.org/zap"
)

func (u *userServiceInterface) RegisterUser(
	userDomain model.UserDomainInterface,
) *error_map.CurrError {
	logger.Info("User MODEL STARTED", zap.String("method", "RegisterUser"))

	userDomain.EncryptUserPassword()

	fmt.Println(userDomain.GetUserProps().Password)

	return nil
}

func (u *userServiceInterface) UpdateUser(userId string, userDomain model.UserDomainInterface) *error_map.CurrError {
	logger.Info("User model iniated.", zap.String("method", "UpdateUser"))

	return nil
}

func (u *userServiceInterface) RetrieveUser(string) (*model.UserDomainInterface, *error_map.CurrError) {
	logger.Info("User model iniated.", zap.String("method", "RetrieveUser"))

	return nil, nil
}

func (u *userServiceInterface) DeleteUser(string) *error_map.CurrError {
	logger.Info("User model iniated.", zap.String("method", "DeleteUser"))

	return nil
}
