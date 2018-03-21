package domain

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Location struct {
	ID        int       `json:"id" gorm:"primary_key"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Image     string    `json:"image"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetLocation(id int) (location Location, err error) {
	err = db.First(&location, id).Error
	return
}

func GetAllLocations() (locations []Location, err error) {
	err = db.Find(&locations).Error
	return
}

func (location Location) Create() (url string, err error) {
	err = db.Create(&location).Error
	if err != nil {
		return
	}
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("ap-northeast-1"),
		Credentials: credentials.NewStaticCredentials(os.Getenv("S3_ACCESS_KEY_ID"), os.Getenv("S3_SECRET_ACCESS_KEY"), ""),
	})
	if err != nil {
		return
	}
	svc := s3.New(sess)
	req, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String("mapic-api-development"),
		Key:    aws.String("locations/" + strconv.Itoa(location.ID)),
		Body:   strings.NewReader("test"),
	})
	url, err = req.Presign(15 * time.Minute)
	if err != nil {
		return
	}
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
