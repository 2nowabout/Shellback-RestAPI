package models

import (
	"gorm.io/gorm"
)

type Type struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	NotiType string `json:"notitype"`
}

// create a Type
func CreateType(db *gorm.DB, Type *Type) (err error) {
	err = db.Create(Type).Error
	if err != nil {
		return err
	}
	return nil
}

// get Types
func GetTypes(db *gorm.DB, Type *[]Type) (err error) {
	err = db.Find(&Type).Error
	if err != nil {
		return err
	}
	return nil
}

// get Type by id
func GetType(db *gorm.DB, Type *Type, id int) (err error) {
	err = db.Where("id = ?", id).First(&Type).Error
	if err != nil {
		return err
	}
	return nil
}

// update Type
func UpdateType(db *gorm.DB, Type *Type) (err error) {
	db.Save(Type)
	return nil
}

// delete Type
func DeleteType(db *gorm.DB, Type *Type, id int) (err error) {
	db.Where("id = ?", id).Delete(Type)
	return nil
}
