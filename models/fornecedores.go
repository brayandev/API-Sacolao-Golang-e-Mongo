package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Fornecedor struct {
	ID              bson.ObjectId `bson:"_id" json:"id"`
	Nome_fornecedor string        `bson:"nome"json:"nome"`
}
