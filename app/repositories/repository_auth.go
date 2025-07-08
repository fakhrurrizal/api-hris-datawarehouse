package repository

import (
	"errors"
	"hris-datawarehouse/app/middlewares"
	"hris-datawarehouse/app/models"
	"hris-datawarehouse/config"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func SignIn(email, password string) (user models.GlobalUser, token string, err error) {
	err = config.DB.
		Where("email = '" + strings.ToLower(email) + "'").First(&user).Error
	if err != nil {
		return
	}
	err = middlewares.VerifyPassword(password, user.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		err = errors.New("incorrect password")
		return
	}
	if user.EmailVerifiedAt.IsZero() {
		err = errors.New("please verify your email")
		return
	}
	token, err = middlewares.AuthMakeToken(user)
	if err != nil {
		return
	}
	return
}

func GetUserByIDPlain(id int) (response models.GlobalUser, err error) {
	err = config.DB.First(&response, id).Error
	return
}
