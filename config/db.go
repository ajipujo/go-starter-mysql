package config

import (
	"fmt"
	"gostart/helper"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB_HOST = helper.GoDotEnvVariable("DB_HOST")
var DB_NAME = helper.GoDotEnvVariable("DB_NAME")
var DB_PORT = helper.GoDotEnvVariable("DB_PORT")
var DB_USERNAME = helper.GoDotEnvVariable("DB_USERNAME")
var DB_PASSWORD = helper.GoDotEnvVariable("DB_PASSWORD")

func DB() *gorm.DB {
	dsn := fmt.Sprintln("host=", DB_HOST, " user=", DB_USERNAME, " password=", DB_PASSWORD, " dbname=", DB_NAME, " port=", DB_PORT, "sslmode=disable TimeZone=Asia/Jakarta")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil
	}

	return db
}
