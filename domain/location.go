package domain

import (
	"time"
)

type Location struct {
	ID int `json:"id" gorm:"primary_key"`
	Title string `json:"title"`
	Content string `json:"content"`
	Image string `json:"image"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
}

func GetLocation(id int) (location Location, err error) {
	db.First(&location, id)
	return
}

func GetAllLocations() (locations []Location, err error) {
	db.Find(&locations)
	return
}

func (location Location) Create() (err error) {
	db.Create(&location)
	return
}

func (location Location) Update(newLocation Location) (err error) {
	db.Model(&location).Updates(newLocation)
	return
}

func (location Location) Delete() (err error) {
	db.Delete(&location)
	return
}
