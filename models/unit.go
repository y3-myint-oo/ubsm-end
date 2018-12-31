package models

import "gopkg.in/mgo.v2/bson"

// Unit - အတိုင်းအတာယူနစ်များ
type Unit struct {
	ID         bson.ObjectId `bson:"_id,omitempty"`
	Name       string        `json:"name"`
	Count      string        `json:"count"`
	Converters []Converter   `json:"converters"`
}

// Converter - နှန်းထားများ
type Converter struct {
	ID    bson.ObjectId `bson:"_id,omitempty"`
	Name  string        `json:"name"`
	Unit  bson.ObjectId `json:"mid"` // mapping unit id ( ref: Unit )
	Cross string        `json:"cross"`
}

// MockUnits - အတိုင်းအတာယူနစ် example များ
var MockUnits = []Unit{
	Unit{"507f1f77bcf86cd799439011", "ပိသာ", "၁", []Converter{
		{"507f1f77bcf86cd799410012", "တင်းရေ", "507f1f77bcf86cd799439012", "၄"},
		{"507f1f77bcf86cd799410012", "ပြည်", "507f1f77bcf86cd799439013", "၁၉"},
	}},
	Unit{"507f1f77bcf86cd799439012", "တင်းရေ", "၁", []Converter{
		{"507f1f77bcf86cd799410012", "ပိသာ", "507f1f77bcf86cd799439012", "၁"},
		{"507f1f77bcf86cd799410012", "ပြည်", "507f1f77bcf86cd799439013", "၁၉"},
	}},
	Unit{"507f1f77bcf86cd799439013", "ပြည်", "၁", []Converter{
		{"507f1f77bcf86cd799410012", "တင်းရေ", "507f1f77bcf86cd799439012", "၁"},
		{"507f1f77bcf86cd799410012", "ပြည်", "507f1f77bcf86cd799439013", "၁၉"}}},
}
