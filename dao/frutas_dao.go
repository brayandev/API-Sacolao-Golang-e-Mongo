package dao

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"

	. "api_sacolao/models"

	"gopkg.in/mgo.v2/bson"
)

type FrutasDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "frutas"
)

func (f *FrutasDAO) Connect() {
	session, err := mgo.Dial(f.Server)

	if err != nil {
		log.Fatal(err)
	}

	db = session.DB(f.Database)
}

func (f *FrutasDAO) FindAll() ([]Fruta, error) {
	var fruta []Fruta
	err := db.C(COLLECTION).Find(bson.M{}).All(&fruta)
	if err != nil {
		fmt.Printf(err.Error())
	}
	return fruta, err
}

func (f *FrutasDAO) FindById(id string) (Fruta, error) {
	var fruta Fruta
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&fruta)
	return fruta, err
}

func (f *FrutasDAO) Insert(fruta Fruta) error {
	err := db.C(COLLECTION).Insert(&fruta)
	return err
}

func (f *FrutasDAO) Delete(fruta Fruta) error {
	err := db.C(COLLECTION).Remove(bson.M{"nome": fruta.Nome})
	return err
}

func (f *FrutasDAO) Update(fruta Fruta) error {
	err := db.C(COLLECTION).UpdateId(fruta.ID, &fruta)
	return err
}
