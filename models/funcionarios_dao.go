package models

import "gopkg.in/mgo.v2/bson"

type Funcionario struct {
	ID    bson.ObjectId `bson:"_id" json:"id"`
	Nome  string        `bson:"nome" json:"nome"`
	Idade int           `bson:"idade" json:"idade"`
	Cargo string        `bson:"cargo" json:"cargo"`
}
