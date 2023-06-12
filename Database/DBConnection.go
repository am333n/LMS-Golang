package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "host=localhost user=ameen password=ameen#123 dbname=LMS port=5432 sslmode=disable TimeZone=Asia/Shanghai"



func GetDB()(*gorm.DB, error) {
	DB, err = gorm.Open(postgres.Open(DNS), &gorm.Config{})
	if err != nil {
		panic("couldn't open database")
	}
	return DB,err

}

