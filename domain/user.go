package domain

import (
	"time"
)

type User struct {
	ID        int       `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	Image     string    `json:"image"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetUser(id int) (user User, err error) {
	err = db.First(&user, id).Error
	return
}

func GetUserByEmail(email string) (user User, err error) {
	err = db.Where("email = ?", email).First(&user).Error
	return
}

func (user User) Create() (createdUser User, err error) {
	err = db.Create(&user).Error
	createdUser = user
	return
}

func (user User) Update(newUser User) (err error) {
	err = db.Model(&user).Updates(newUser).Error
	return
}

func (user User) Delete() (err error) {
	err = db.Delete(&user).Error
	return
}
