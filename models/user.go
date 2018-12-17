package models

import "gopkg.in/mgo.v2/bson"

// User - System User
type User struct {
	ID    bson.ObjectId `bson:"_id,omitempty"`
	Name  string        `json:"name"  `
	Pass  string        `json:"pass"  `
	Level string        `json:"level"  `
}

// Admins - Mock data set for Admin User
var Admins = []User{
	User{"507f1f77bcf86cd799439011", "U Myint", "1234", "1"},
	User{"507f1f77bcf86cd799439012", "U Kyaw", "1234", "2"},
	User{"507f1f77bcf86cd799439013", "Aung Khaing", "1234", "1"},
}
