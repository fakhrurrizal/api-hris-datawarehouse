package controllers

import (
	"hris-datawarehouse/app/models"
	repository "hris-datawarehouse/app/repositories"
	"hris-datawarehouse/app/reqres"
	"hris-datawarehouse/app/utils"
	"net/http"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/labstack/echo/v4"
)

// SignIn godoc
// @Summary SignIn
// @Description SignIn
// @Tags Auth
// @Accept json
// @Param x-csrf-token header string false "csrf token"
// @Produce json
// @Param signin body reqres.SignInRequest true "SignIn user"
// @Success 200
// @Router /v1/auth/signin [post]
// @Security ApiKeyAuth
func SignIn(c echo.Context) error {

	var req reqres.SignInRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewUnprocessableEntityError(err.Error()))
	}

	if err := req.Validate(); err != nil {
		errVal := err.(validation.Errors)
		return c.JSON(http.StatusBadRequest, utils.NewInvalidInputError(errVal))
	}

	user, accessToken, err := repository.SignIn(req.Email, req.Password)
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"status": 400,
			"error":  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": map[string]interface{}{
			"user":         user,
			"access_token": accessToken,
			"expiration":   time.Now().Add(time.Hour * 72).Format("2006-01-02 15:04:05"),
		},
	})
}

// GetSignInUser godoc
// @Summary Get Sign In User
// @Description Get Sign In User
// @Tags Auth
// @Produce json
// @Success 200
// @Router /v1/auth/user [get]
// @Security ApiKeyAuth
// @Security JwtToken
func GetSignInUser(c echo.Context) error {

	id := c.Get("user_id").(int)
	user, err := repository.GetUserByIDPlain(id)
	if err != nil {
		return c.JSON(500, utils.Respond(500, err.Error(), "Failed to get user data"))
	}

	data, err := GetSignInUserProcess(user)
	if err != nil {
		return c.JSON(500, utils.Respond(500, err.Error(), "failed to get user"))
	}

	return c.JSON(200, map[string]interface{}{
		"status":  200,
		"data":    data,
		"message": "Success to get user data",
	})
}

func GetSignInUserProcess(user models.GlobalUser) (data reqres.GlobalUserAuthResponse, err error) {
	data.ID = int(user.ID)

	if data.EncodedID == "" {
		data.EncodedID = utils.EndcodeID(int(user.ID))
	}
	data.Fullname = user.Fullname
	data.Avatar = user.Avatar
	data.Email = user.Email

	if user.EmailVerifiedAt.Time.IsZero() {
		data.EmailVerified = false
	} else {
		data.EmailVerified = true
	}

	return
}
