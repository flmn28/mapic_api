package domain

import (
	"time"
)

type User struct {
	ID int `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
	Image string `json:"image"`
	Email string `json:"email"`
	Password string `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
}

func GetUser(id int) (user User, err error) {
	db.First(&user, id)
	return
}

func (user User) Create() (err error) {
	db.Create(&user)
	return
}

func (user User) Update(newUser User) (err error) {
	db.Model(&user).Updates(newUser)
	return
}

func (user User) Delete() (err error) {
	db.Delete(&user)
	return
}
