package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type User struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	Image       string        `bson:"image" json:"image"`
	DateOfBirth time.Time     `bson:"dateOfBirth" json:"dateOfBirth"`
	Active      bool          `bson:"active" json:"active"`
}
