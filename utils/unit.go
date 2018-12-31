package utils

import (
	"github.com/pkg/errors"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	m "ymo.me/sbum-end/models"
	//d "ymo.me/sbum-end/utils"
)

// UnitCollection - Unit Model CRUD operation
var UnitCollection *mgo.Collection

// FindAllUnit - All Avaliable Unit measure in system
func FindAllUnit() ([]m.Unit, error) {
	var units []m.Unit
	err := UnitCollection.Find(bson.M{}).All(&units)
	return units, err
}

// FindByName - Get Unit with Name - Name is Unique
func FindByName(name string) (m.Unit, error) {
	var unit m.Unit
	err := UnitCollection.Find(bson.M{"name": name}).One(&unit)
	return unit, err
}

// InsertUnit - Create new Unit
func InsertUnit(unit m.Unit) error {
	//Check Name is already exsit or not
	expUnit, err := FindByName(unit.Name)
	if expUnit.Name == unit.Name {
		return errors.New("ထိုယူနစ်နာမည်အား အသုံးပြုပြီးသားဖြစ်ပါသည်။")
	}
	err = UnitCollection.Insert(&unit)
	return err
}

// DeleteUnit - Delete Unit
func DeleteUnit(unit m.Unit) error {
	err := UnitCollection.Remove(&unit)
	return err
}

// InsertUnitConverter - Insert Unit Converter
func InsertUnitConverter(unit m.Unit, con m.Converter) error {
	unit.Converters = append(unit.Converters, con)
	err := UnitCollection.UpdateId(unit.ID, &unit)
	return err
}

// DeleteUnitConverter - Delete Unit Converter
func DeleteUnitConverter(unit m.Unit, con m.Converter) error {
	//unit.Converters
	contains(unit.Converters, con)
	err := UnitCollection.UpdateId(unit.ID, &unit)
	return err
}

// Conatains - Check and delete
func contains(s []m.Converter, e m.Converter) bool {
	for i, a := range s {
		if a == e {
			copy(s[i:], s[i+1:])
			//	s[len(s)-1] =
			s = s[:len(s)-1]
			return true
		}
	}
	return false
}
