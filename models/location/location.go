package location

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var db *gorm.DB

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

func init() {
	var err error
	db, err = gorm.Open("mysql", "root@/mapic_api_development?charset=utf8&parseTime=True")
	if err != nil {
		panic(err)
	}
}

func Get(id int) (loc Location, err error) {
	db.First(&loc, id)
	return
}

func All() (locs []Location, err error) {
	db.Find(&locs)
	return
}

func (loc Location) Create() (err error) {
	db.Create(&loc)
	return
}

func (loc Location) Update(newLoc Location) (err error) {
	db.Model(&loc).Updates(newLoc)
	return
}

func (loc Location) Delete() (err error) {
	db.Delete(&loc)
	return
}
