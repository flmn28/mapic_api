package handler

import (
	"github.com/johskw/mapic_api/domain"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func GetUser(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := domain.GetUser(id)
	if err != nil {
		return
	}
	return c.JSON(http.StatusOK, &user)
}

func PostUser(c echo.Context) (err error) {
	user := new(domain.User)
	err = c.Bind(user)
	if err != nil {
		return
	}
	err = user.Create()
	if err != nil {
		return
	}
	return c.JSON(http.StatusOK, &user)
}

func PutUser(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := domain.GetUser(id)
	if err != nil {
		return
	}
	newUser := new(domain.User)
	err = c.Bind(newUser)
	if err != nil {
		return
	}
	err = user.Update(*newUser)
	if err != nil {
		return
	}
	return c.Redirect(http.StatusMovedPermanently, "/users/" + c.Param("id"))
}

func DeleteUser(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := domain.GetUser(id)
	if err != nil {
		return
	}
	err = user.Delete()
	if err != nil {
		return
	}
	return c.Redirect(http.StatusMovedPermanently, "/locations")
}
