package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Notification struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	TypeID    uint      `json:"type" gorm:"not null"`
	Type      Type      `gorm:"foreignKey:TypeID;references:ID"`
	IpAdress  string    `json:"ipadress" gorm:"not null;index"`
	Company   Company   `gorm:"foreignKey:IpAdress;references:IpAdress"`
	Value     string    `json:"value" gorm:"not null"`
	TimeStamp time.Time `json:"timeStamp" gorm:"not null"`
}

func CreateNotification(db *gorm.DB, Notification *Notification) (err error) {
	err = db.Create(Notification).Error
	if err != nil {
		return err
	}
	return nil
}

func GetNotifications(db *gorm.DB, Notification *[]Notification, ipadress string) (err error) {
	fmt.Println("TRYING TO FIND IT")
	err = db.Find(&Notification, "ip_adress = ?", ipadress).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteNotifications(db *gorm.DB, Notification *[]Notification, ipadress string) (err error) {
	db.Where("ip_adress = ?", ipadress).Delete(Notification)
	return nil
}
