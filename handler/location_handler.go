package handler

import (
	"github.com/flmn28/mapic_api/models/location"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func GetLocation(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	loc, err := location.Get(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, &loc)
}

func GetAllLocations(c echo.Context) error {
	loc, _ := location.All()
	return c.JSON(http.StatusOK, &loc)
}

func PostLocation(c echo.Context) error {
	loc := new(location.Location)
	err := c.Bind(loc)
	if err != nil {
		return err
	}
	err = loc.Create()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, &loc)
}

func PutLocation(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	loc, _ := location.Get(id)
	newLoc := new(location.Location)
	err := c.Bind(newLoc)
	if err != nil {
		return err
	}
	err = loc.Update(*newLoc)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "put!")
}

func DeleteLocation(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	loc, _ := location.Get(id)
	err := loc.Delete()
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "deleted!")
}
