package dao

import (
	"log"

	. "github.com/gguibittencourt/go-restapi/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UsersDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "users"
)

func (usersDAO *UsersDAO) Connect() {
	session, err := mgo.Dial(usersDAO.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(usersDAO.Database)
}

func (usersDAO *UsersDAO) List() ([]User, error) {
	var users []User
	err := db.C(COLLECTION).Find(bson.M{}).All(&users)
	return users, err
}

func (usersDAO *UsersDAO) GetByID(id string) (User, error) {
	var user User
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

func (usersDAO *UsersDAO) Create(user User) error {
	err := db.C(COLLECTION).Insert(&user)
	return err
}

func (usersDAO *UsersDAO) Delete(id string) error {
	err := db.C(COLLECTION).RemoveId(bson.ObjectIdHex(id))
	return err
}

func (usersDAO *UsersDAO) Update(id string, user User) error {
	err := db.C(COLLECTION).UpdateId(bson.ObjectIdHex(id), &user)
	return err
}
