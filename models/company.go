package models

import (
	"time"

	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	ID          uint      `json:id`
	CompanyName string    `json:companyname`
	IpAdress    string    `json:ipadress`
	StartDate   time.Time `json:startdate`
	LastActive  time.Time `json:lastactive`
}

func CreateCompany(db *gorm.DB, Company *Company) (err error) {
	err = db.Create(Company).Error
	if err != nil {
		return err
	}
	return nil
}

func GetCompanys(db *gorm.DB, Company *[]Company) (err error) {
	err = db.Find(Company).Error
	if err != nil {
		return err
	}
	return nil
}

func GetCompany(db *gorm.DB, Company *Company, id int) (err error) {
	err = db.Where("id = ?", id).First(Company).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteCompany(db *gorm.DB, Company *Company, id int) (err error) {
	db.Where("id = ?", id).Delete(Company)
	return nil
}

func UpdateActive(db *gorm.DB, Company *Company, id int) (err error) {
	currentTime := time.Now()
	db.Where("id = ?", id).Update("LastActive", currentTime.Format("2006-01-02 3:4:5"))
	return nil
}
