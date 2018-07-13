package dao

import (
	. "api_sacolao/models"

	"gopkg.in/mgo.v2/bson"
)

const (
	COLLECTION3 = "fornecedores"
)

func (f *FrutasDAO) InsertFornecedor(fornecedor Fornecedor) error {
	err := db.C(COLLECTION3).Insert(&fornecedor)
	return err
}

func (f *FrutasDAO) GetFornecedor() ([]Fornecedor, error) {
	var fornecedores []Fornecedor
	err := db.C(COLLECTION3).Find(bson.M{}).All(&fornecedores)
	return fornecedores, err
}
