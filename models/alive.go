package models

import (
	"time"

	"gorm.io/gorm"
)

type Alive struct {
	ID             uint      `json:"id" gorm:"primary_key"`
	IpAdress       string    `json:"ipadress"`
	LastOnlinePing time.Time `json:"lastonline"`
}

func CreateAlive(db *gorm.DB, Alive *Alive) (err error) {
	var EditAlive = Alive
	currentTime := time.Now()
	EditAlive.LastOnlinePing, err = time.Parse("2006-01-02 15:04:05", currentTime.GoString())
	if err != nil {
		return err
	}
	err = db.Create(&EditAlive).Error
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
	db.Model(Alive{}).Where("ip_adress = ?", ip).Update("last_active", currentTime.Format("2006-01-02 15:04:05"))
	return nil
}

func DeleteAlive(db *gorm.DB, Alive *Alive, id int) (err error) {
	db.Where("id = ?", id).Delete(&Alive)
	return nil
}
