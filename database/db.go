package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBInfo struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_NAME     string
	DB_HOST     string
	DB_PORT     string
}

var Db *gorm.DB

func InitDb() *gorm.DB {
	Db = connectDB(getValues())
	return Db
}

func connectDB(info DBInfo) *gorm.DB {
	var err error
	dsn := info.DB_USERNAME + ":" + info.DB_PASSWORD + "@tcp" + "(" + info.DB_HOST + ":" + info.DB_PORT + ")/" + info.DB_NAME + "?" + "parseTime=true&loc=Local"
	fmt.Println("dsn : ", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to database : error=%v", err)
		return nil
	}

	return db
}

func getValues() DBInfo {
	dbInfo := new(DBInfo)
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("Error loading environment variables file")
	}
	dbInfo.DB_USERNAME = os.Getenv("DB_USERNAME")
	dbInfo.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	dbInfo.DB_NAME = os.Getenv("DB_NAME")
	dbInfo.DB_HOST = os.Getenv("DB_HOST")
	dbInfo.DB_PORT = os.Getenv("DB_PORT")
	fmt.Printf("DB INFO HERE:")
	fmt.Printf(os.Getenv("DB_USERNAME"))
	return *dbInfo
}
