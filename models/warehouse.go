package models

import "gopkg.in/mgo.v2/bson"

// Warehouse - Sample
type Warehouse struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Name        string        `json:"name"`
	Address     string        `json:"address"`
	Transaction string        `json:"transaction"`
	Stocks      []Stock       `json:"stocks"`
}

//Stock - Stock Item (ကုန်ပစ္စည်းအမျိုးအစား)
type Stock struct {
	ID         bson.ObjectId `bson:"_id,omitempty"`
	Name       string        `json:"name"`
	Code       string        `json:"code"`
	Unit       string        `bson:"unit"`
	RentCharge int           `bson:"rentcharge"`
	RentRate   string        `bson:"rentRate"` //per day or per month or per year
	Quality    int           `bson:"quality"`
	Ownership  Supply        `bson:"customer"`
}
