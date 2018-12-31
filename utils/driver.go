package utils

import (
	"fmt"

	mgo "gopkg.in/mgo.v2"
)

var session *mgo.Session

// TodosCollection -  Todo Model CRUD operation
var TodosCollection *mgo.Collection

const (
	DTATBASE = "test001"
)

var DB *mgo.Database

// InitDriver - Init the MongoDb Connection
func InitDriver() error {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	DB = session.DB(DTATBASE)
	TodosCollection = DB.C("todo")
	UnitCollection = DB.C("unit")
	fmt.Print(" Init Dirver Collection ", UnitCollection.FullName)
	return nil
}
