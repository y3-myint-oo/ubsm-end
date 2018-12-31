package models

import "gopkg.in/mgo.v2/bson"

// Supply -
type Supply struct {
	ID       bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name     string        `json:"name"  `
	Address  string        `json:"address"`
	Phone    string        `json:"phone"`
	Division string        `json:"division"`
	Note     string        `json:"note"`
}
