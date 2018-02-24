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
	err = db.First(&location, id).Error
	return
}

func GetAllLocations() (locations []Location, err error) {
	err = db.Find(&locations).Error
	return
}

func (location Location) Create() (err error) {
	err = db.Create(&location).Error
	return
}

func (location Location) Update(newLocation Location) (err error) {
	err = db.Model(&location).Updates(newLocation).Error
	return
}

func (location Location) Delete() (err error) {
	err = db.Delete(&location).Error
	return
}
