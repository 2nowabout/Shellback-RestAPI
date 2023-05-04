package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Company struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	CompanyName string    `json:"companyname"`
	IpAdress    string    `json:"ipadress" gorm:"index"`
	StartDate   time.Time `json:"startdate"`
	LastActive  time.Time `json:"lastactive"`
}

func CreateCompany(db *gorm.DB, Company *Company) (err error) {
	err = db.Create(&Company).Error
	if err != nil {
		return err
	}
	return nil
}

func GetCompanys(db *gorm.DB, Company *[]Company) (err error) {
	err = db.Find(&Company).Error
	if err != nil {
		return err
	}
	return nil
}

func GetCompany(db *gorm.DB, Company *Company, id int) (err error) {
	err = db.Where("id = ?", id).First(&Company).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteCompany(db *gorm.DB, Company *Company, id int) (err error) {
	db.Where("id = ?", id).Delete(&Company)
	return nil
}

func UpdateActive(db *gorm.DB, ip string) (err error) {
	currentTime := time.Now()
	fmt.Println(currentTime.Format("2006-01-02 15:04:05"))
	fmt.Println(ip)
	db.Model(Company{}).Where("ip_adress = ?", ip).Update("last_active", currentTime.Format("2006-01-02 15:04:05"))
	return nil
}
