package handler

import (
	"github.com/flmn28/mapic_api/domain"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func GetLocation(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	location, err := domain.GetLocation(id)
	if err != nil {
		return
	}
	return c.JSON(http.StatusOK, &location)
}

func GetAllLocations(c echo.Context) (err error) {
	location, _ := domain.GetAllLocations()
	return c.JSON(http.StatusOK, &location)
}

func PostLocation(c echo.Context) (err error) {
	location := new(domain.Location)
	err = c.Bind(location)
	if err != nil {
		return
	}
	err = location.Create()
	if err != nil {
		return
	}
	return c.JSON(http.StatusOK, &location)
}

func PutLocation(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	location, err := domain.GetLocation(id)
	if err != nil {
		return
	}
	newLocation := new(domain.Location)
	err = c.Bind(newLocation)
	if err != nil {
		return
	}
	err = location.Update(*newLocation)
	if err != nil {
		return
	}
	return c.String(http.StatusOK, "put!")
}

func DeleteLocation(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	location, err := domain.GetLocation(id)
	if err != nil {
		return
	}
	err = location.Delete()
	if err != nil {
		return
	}
	return c.String(http.StatusOK, "deleted!")
}
