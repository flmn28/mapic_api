package domain

import (
	"encoding/base64"
	"os"
	"strconv"
	"time"
)

type Location struct {
	ID        int       `json:"id" gorm:"primary_key"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Image     string    `json:"image"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	User      User      `json:"user"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetLocation(id int) (location Location, err error) {
	err = db.First(&location, id).Error
	if err != nil {
		return
	}
	user, err := GetUser(location.UserID)
	if err != nil {
		return
	}
	location.User = user
	return
}

func GetAllLocations() (locations []Location, err error) {
	err = db.Find(&locations).Error
	for i, location := range locations {
		user, err := GetUser(location.UserID)
		if err != nil {
			break
		}
		locations[i].User = user
		strID := strconv.Itoa(location.ID)
		file, _ := os.Open("images/locations/" + strID + "/" + strID + ".jpg")
		defer file.Close()
		fi, _ := file.Stat()
		size := fi.Size()
		data := make([]byte, size)
		file.Read(data)
		locations[i].Image = base64.StdEncoding.EncodeToString(data)
	}
	return
}

func GetLocationsByUserId(userID int) (locations []Location, err error) {
	err = db.Where("user_id = ?", strconv.Itoa(userID)).Find(&locations).Error
	for i, location := range locations {
		user, err := GetUser(location.UserID)
		if err != nil {
			break
		}
		locations[i].User = user
		strID := strconv.Itoa(location.ID)
		file, _ := os.Open("images/locations/" + strID + "/" + strID + ".jpg")
		defer file.Close()
		fi, _ := file.Stat()
		size := fi.Size()
		data := make([]byte, size)
		file.Read(data)
		locations[i].Image = base64.StdEncoding.EncodeToString(data)
	}
	return
}

func (location Location) Create() (createdlocation Location, err error) {
	location.Image = ""
	err = db.Create(&location).Error
	createdlocation = location
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

func (location Location) SaveImage(id int) (err error) {
	data, err := base64.StdEncoding.DecodeString(location.Image)
	if err != nil {
		return
	}
	strID := strconv.Itoa(id)
	err = os.MkdirAll("images/locations/"+strID, 0777)
	if err != nil {
		panic(err)
	}
	file, _ := os.Create("images/locations/" + strID + "/" + strID + ".jpg")
	if err != nil {
		return
	}
	defer file.Close()
	file.Write(data)
	return
}
