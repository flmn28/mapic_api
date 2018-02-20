package user

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var db *gorm.DB

type User struct {
	ID int `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
	Image string `json:"image"`
	Email string `json:"email"`
	Password string `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
}

func init() {
	var err error
	db, err = gorm.Open("mysql", "root@/mapic_api_development?charset=utf8&parseTime=True")
	if err != nil {
		panic(err)
	}
}

func Get(id int) (usr User, err error) {
	db.First(&usr, id)
	return
}

func (usr User) Create() (err error) {
	db.Create(&usr)
	return
}

func (usr User) Update(newUsr User) (err error) {
	db.Model(&usr).Updates(newUsr)
	return
}

func (usr User) Delete() (err error) {
	db.Delete(&usr)
	return
}
