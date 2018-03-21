package handler

import (
	"github.com/johskw/mapic_api/domain"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

type Auth struct {
	ID int `json:"id"`
	Password string `json:"password"`
}

func Login(c echo.Context) (err error) {
	auth := new(Auth)
	err = c.Bind(auth)

	user, err := domain.GetUser(auth.ID)
	if err != nil {
		return
	}

	if auth.Password != user.Password {
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = auth.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

func Restricted(c echo.Context) (err error) {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	id := int(claims["id"].(float64))
	user, err := domain.GetUser(id)
	return c.String(http.StatusOK, "Welcome "+ user.Name +"!")
}
