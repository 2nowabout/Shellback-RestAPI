package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Alive struct {
	ID             uint      `json:"id" gorm:"primary_key"`
	IpAdress       string    `json:"ipadress"`
	LastOnlinePing time.Time `json:"lastonline"`
}

func CreateAlive(db *gorm.DB, Alive *Alive) (err error) {
	err = db.Create(&Alive).Error
	if err != nil {
		return err
	}
	return nil
}

func GetAlives(db *gorm.DB, Alive *[]Alive) (err error) {
	err = db.Find(&Alive).Error
	if err != nil {
		return err
	}
	return nil
}

func GetAlive(db *gorm.DB, Alive *Alive, id int) (err error) {
	err = db.Where("id = ?", id).First(&Alive).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateLastOnline(db *gorm.DB, ip string) (err error) {
	currentTime := time.Now()
	fmt.Println(currentTime.Format("2006-01-02 15:04:05"))
	fmt.Println(ip)
	db.Model(Alive{}).Where("ip_adress = ?", ip).Update("last_active", currentTime.Format("2006-01-02 15:04:05"))
	return nil
}
