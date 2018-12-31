package utils

import (
	"fmt"

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
	//Make Sure at least one unit have
	units, err := FindAllUnit()
	if err != nil {
		return err
	}
	if len(units) > 1 {
		err = UnitCollection.Remove(&unit)
	} else {
		return errors.New("အတိုင်းအတာယူနစ် တစ်ခု တည်ရှိရပါမည်")
	}
	return err
}

// InsertUnitConverter - Insert Unit Converter
func InsertUnitConverter(con m.Converter) error {
	con.ID = bson.NewObjectId()
	pUnit, err := FindByName(con.Map)
	if err != nil {
		return err
	}
	fmt.Print(" pUnit ID ", pUnit.ID)
	fmt.Print(" con ID ", con.ID)
	pUnit.Converters = append(pUnit.Converters, con)
	err = UnitCollection.UpdateId(pUnit.ID, &pUnit)
	return err
}

// DeleteUnitConverter - Delete Unit Converter
func DeleteUnitConverter(con m.Converter) error {
	pUnit, err := FindByName(con.Map)
	if err != nil {
		return err
	}
	fmt.Println(" before remove - ", len(pUnit.Converters))
	pUnit.Converters = removeConverter(pUnit.Converters, &con)
	fmt.Println(" After remove - ", len(pUnit.Converters))
	err = UnitCollection.UpdateId(pUnit.ID, &pUnit)
	return err
}

func removeConverter(clients []m.Converter, connClient *m.Converter) []m.Converter {
	for i := len(clients) - 1; i >= 0; i-- {
		if clients[i].Name == connClient.Name {
			copy(clients[i:], clients[i+1:])
			clients[len(clients)-1] = m.Converter{}
			clients = clients[:len(clients)-1]
		}
	}
	return clients
}
