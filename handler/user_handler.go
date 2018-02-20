package handler

import (
	"github.com/flmn28/mapic_api/models/user"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func GetUser(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	usr, err := user.Get(id)
	if err != nil {
		return
	}
	return c.JSON(http.StatusOK, &usr)
}

func PostUser(c echo.Context) (err error) {
	usr := new(user.User)
	err = c.Bind(usr)
	if err != nil {
		return
	}
	err = usr.Create()
	if err != nil {
		return
	}
	return c.JSON(http.StatusOK, &usr)
}

func PutUser(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	usr, err := user.Get(id)
	if err != nil {
		return
	}
	newUsr := new(user.User)
	err = c.Bind(newUsr)
	if err != nil {
		return
	}
	err = usr.Update(*newUsr)
	if err != nil {
		return
	}
	return c.String(http.StatusOK, "put!")
}

func DeleteUser(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	usr, err := user.Get(id)
	if err != nil {
		return
	}
	err = usr.Delete()
	if err != nil {
		return
	}
	return c.String(http.StatusOK, "deleted!")
}
