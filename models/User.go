package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name      string    `json:"name"`
	BirthDate time.Time `json:"birthDate"`
}

func CreateUser(db *gorm.DB, User *User) (err error) {
	err = db.Create(User).Error
	if err != nil {
		return err
	}
	return nil
}

func ListUsers(db *gorm.DB, Users *[]User) (err error) {
	err = db.Find(Users).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUser(db *gorm.DB, User *User, id int) (err error) {
	err = db.Where("id = ?", id).First(User).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateUser(db *gorm.DB, User *User) (err error) {
	db.Save(User)
	return nil
}

func DeleteUser(db *gorm.DB, User *User, id string) (err error) {
	db.Where("id = ?", id).Delete(User)
	return nil
}
