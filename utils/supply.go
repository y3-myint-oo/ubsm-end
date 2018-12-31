package utils

import (
	"fmt"

	"github.com/pkg/errors"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	m "ymo.me/sbum-end/models"
	//d "ymo.me/sbum-end/utils"
)

// SupplyCollection - Supply Model CRUD operation
var SupplyCollection *mgo.Collection

// FindAllSupply - All Avaliable Unit measure in system
func FindAllSupply() ([]m.Supply, error) {
	var supply []m.Supply
	err := SupplyCollection.Find(bson.M{}).All(&supply)
	return supply, err
}

// FindBySupplyID - Get Unit with Id - Id is Unique
func FindBySupplyID(id string) (m.Supply, error) {
	var unit m.Supply
	err := SupplyCollection.Find(bson.M{"id": id}).One(&unit)
	return unit, err
}

// InsertSupply - Create new Supply
func InsertSupply(unit m.Supply) error {
	err := SupplyCollection.Insert(&unit)
	return err
}

// DeleteSupply - Delete Supply
func DeleteSupply(supply m.Supply) error {
	supplies, err := FindAllSupply()
	if err != nil {
		return err
	}
	fmt.Println("---- supply length ", len(supplies))
	if len(supplies) > 1 {
		err = SupplyCollection.Remove(&supply)
	} else {
		return errors.New("တောင်သူ တစ်ခု တည်ရှိရပါမည်")
	}
	return err
}

//FindSupplyStock -  Supply with relatived Stocks
func FindSupplyStock(supply m.Supply) (m.Stock, error) {
	return m.Stock{}, nil
}
