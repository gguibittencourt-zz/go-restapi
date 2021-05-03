package models

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Finished    bool   `json:"finished"`
}

func CreateTask(db *gorm.DB, Task *Task) (err error) {
	err = db.Create(Task).Error
	if err != nil {
		return err
	}
	return nil
}

func ListTasks(db *gorm.DB, Tasks *[]Task) (err error) {
	err = db.Find(Tasks).Error
	if err != nil {
		return err
	}
	return nil
}

func GetTask(db *gorm.DB, Task *Task, id int) (err error) {
	err = db.Where("id = ?", id).First(Task).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateTask(db *gorm.DB, Task *Task) (err error) {
	db.Save(Task)
	return nil
}

func DeleteTask(db *gorm.DB, Task *Task, id string) (err error) {
	db.Where("id = ?", id).Delete(Task)
	return nil
}
