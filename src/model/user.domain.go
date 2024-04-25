package model

import (
	"crypto/md5"
	"encoding/hex"
)

type userDomain struct {
	name     string
	email    string
	password string
	age      int32
}

type UserProps struct {
	Name     string
	Email    string
	Password string
	Age      int32
}

type UserDomainInterface interface {
	GetUserProps() *UserProps
	EncryptUserPassword()
}

func NewUserDomain(
	name, email, password string,
	age int32,
) UserDomainInterface {
	return &userDomain{
		name, email, password, age,
	}
}

func (d *userDomain) GetUserProps() *UserProps {
	return &UserProps{
		Name:     d.name,
		Email:    d.email,
		Password: d.password,
		Age:      d.age,
	}
}

func (u *userDomain) EncryptUserPassword() {
	hash := md5.New()

	defer hash.Reset()

	hash.Write([]byte(u.password))

	u.password = hex.EncodeToString(hash.Sum(nil))
}
