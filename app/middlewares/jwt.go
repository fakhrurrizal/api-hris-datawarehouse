package middlewares

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"hris-datawarehouse/app/models"
	"hris-datawarehouse/app/utils"
	"hris-datawarehouse/config"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

func Auth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			authorizationHeader := c.Request().Header.Get("Authorization")
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) != 2 {
				return c.JSON(http.StatusUnauthorized, utils.NewUnauthorizedError("Incorrect Authorization Token"))
			}

			tokenStr := bearerToken[1]

			UserID, err := ValidateToken(tokenStr)
			if err != nil {
				fmt.Println("Token Validation,", err)
				return c.JSON(
					http.StatusUnauthorized,
					utils.NewUnauthorizedError(err.Error()),
				)
			}

			c.Set("user_id", UserID)
			return next(c)
		}
	}
}

func ValidateToken(tokenString string) (userID int, err error) {
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		location = time.Local
		err = nil
	}
	tokenStringbyt, err := hex.DecodeString(tokenString)
	if err != nil {
		err = errors.New("incorrect token format")
		return
	}
	str := string(tokenStringbyt)
	newtStr := strings.Replace(string(str), config.LoadConfig().AppKey, "", -1)
	decoded, err := base64.StdEncoding.DecodeString(newtStr)
	if err != nil {
		err = errors.New("incorrect token format")
		return
	}
	newStr := strings.Replace(string(decoded), config.LoadConfig().AppKey, "", -1)
	newdecoded, err := base64.StdEncoding.DecodeString(newStr)
	if err != nil {
		err = errors.New("incorrect token format")
		return
	}
	parts := strings.Split(string(newdecoded), "&")
	expiredAt, _ := strconv.Atoi(parts[1])
	if expiredAt < int(time.Now().In(location).Unix()) {
		err = errors.New("incorrect token format")
		return
	}
	userID, _ = strconv.Atoi(parts[0])

	return
}

func AuthMakeToken(user models.GlobalUser) (string, error) {
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		location = time.Local
		err = nil
	}

	ExpiresAt := time.Now().In(location).Add(7 * 24 * time.Hour).Unix()
	str := fmt.Sprintf("%v&%v", user.ID, ExpiresAt)
	encoded := base64.StdEncoding.EncodeToString([]byte(str)) + config.LoadConfig().AppKey
	token := base64.StdEncoding.EncodeToString([]byte(encoded)) + config.LoadConfig().AppKey
	token = hex.EncodeToString([]byte(token))
	return token, nil
}

func CheckAPIKey() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			if config.LoadConfig().EnableAPIKey {
				hashedApiKey := c.Request().Header.Get("X-API-KEY")
				fmt.Println("hashedApiKey:", hashedApiKey)
				err := VerifyPassword(config.LoadConfig().APIKey, hashedApiKey)
				if err != nil {
					return c.JSON(http.StatusForbidden, map[string]interface{}{
						"status":  403,
						"message": "Wrong API Key",
						"err":     err.Error(),
					})
				}
			}

			return next(c)
		}
	}
}
