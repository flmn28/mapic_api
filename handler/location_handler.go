package handler

import (
	"net/http"
	"strconv"

	"github.com/johskw/mapic_api/domain"
	"github.com/labstack/echo"
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
	locations, _ := domain.GetAllLocations()
	return c.JSON(http.StatusOK, &locations)
}

func GetMyLocations(c echo.Context) (err error) {
	userID, _ := strconv.Atoi(c.Param("user_id"))
	locations, _ := domain.GetLocationsByUserId(userID)
	return c.JSON(http.StatusOK, &locations)
}

func PostLocation(c echo.Context) (err error) {
	location := new(domain.Location)
	err = c.Bind(location)
	if err != nil {
		return
	}
	createdLocation, err := location.Create()
	if err != nil {
		return
	}
	err = location.SaveImage(createdLocation.ID)
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
	return c.JSON(http.StatusOK, &newLocation)
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
	return c.JSON(http.StatusOK, &location)
}
