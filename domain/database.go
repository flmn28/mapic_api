package domain

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", "root@/mapic_api_development?charset=utf8&parseTime=True")
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Location{}, &User{})
}
