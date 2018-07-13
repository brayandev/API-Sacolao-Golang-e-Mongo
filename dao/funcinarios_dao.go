package dao //Funcionarios
import (
	. "api_sacolao/models"

	"gopkg.in/mgo.v2/bson"
)

const (
	COLLECTION2 = "funcionarios"
)

func (f *FrutasDAO) FindAllF() ([]Funcionario, error) {
	var funcionarios []Funcionario
	err := db.C(COLLECTION2).Find(bson.M{}).All(&funcionarios)
	return funcionarios, err
}

func (f *FrutasDAO) InsertF(funcionarios Funcionario) error {
	err := db.C(COLLECTION2).Insert(&funcionarios)
	return err
}
