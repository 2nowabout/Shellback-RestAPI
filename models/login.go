package models

import (
	"errors"

	"Shellback.nl/Restapi/token"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	Email    string `json:"email" gorm:"unique"`
}

type UpdateUser struct {
	ID          uint   `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	NewPassword string `json:"newpassword"`
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (u *User) PrepareGive() {
	u.Password = ""
}

func GetUserByID(db *gorm.DB, uid uint) (User, error) {
	var u User

	if err := db.First(&u, uid).Error; err != nil {
		return u, errors.New("User not found!")
	}

	u.PrepareGive()
	return u, nil
}

func ChangePassword(db *gorm.DB, username string, password string, newpassword string) error {
	u := User{}
	err := db.Model(User{}).Where("username = ?", username).Take(&u).Error
	if err != nil {
		return err
	}
	err = CheckPassword(u.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return err
	}
	if err := u.HashPassword(newpassword); err != nil {
		return err
	}
	err = db.Model(User{}).Where("username = ?", username).Update("password", u.Password).Error
	if err != nil {
		return err
	}
	return nil
}

func CheckPassword(password, providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}

func CreateUser(db *gorm.DB, user *User) error {
	if err := user.HashPassword(user.Password); err != nil {
		return err
	}
	err := db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func AuthenticateUser(db *gorm.DB, username string, password string) (string, error) {
	u := User{}
	err := db.Model(User{}).Where("username = ?", username).Take(&u).Error
	if err != nil {
		return "", err
	}
	err = CheckPassword(u.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	token, err := token.GenerateToken(u.Username, u.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}
