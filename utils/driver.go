package utils

import mgo "gopkg.in/mgo.v2"

var session *mgo.Session

// TodosCollection -  Todo Model CRUD operation
var TodosCollection *mgo.Collection

func InitDriver() {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)

	// get a Collection of todo
	TodosCollection = session.DB("test-todo").C("todo")
}
