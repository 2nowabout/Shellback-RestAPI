package models

import (
	"gorm.io/gorm"
)

type Notification struct {
	ID       uint    `json:"ID" gorm:"primary_key"`
	Type     Type    `json:"Type" gorm:"not null"`
	IpAdress Company `json:"IpAdress" gorm:"not null"`
	Value    string  `json:"Value" gorm:"not null"`
}

func CreateNotification(db *gorm.DB, Notification *Notification) (err error) {
	err = db.Create(Notification).Error
	if err != nil {
		return err
	}
	return nil
}

func GetNotifications(db *gorm.DB, Notification *[]Notification, ipadress string) (err error) {
	err = db.Where("IpAdress = ?", ipadress).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteNotifications(db *gorm.DB, Notification *[]Notification, ipadress string) (err error) {
	db.Where("IpAdress = ?", ipadress).Delete(Notification)
	return nil
}
