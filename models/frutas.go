package models

import "gopkg.in/mgo.v2/bson"

type Fruta struct {
	ID         bson.ObjectId `bson:"_id" json:"id"`
	Nome       string        `bson:"nome" json:"nome"`
	Quantidade int           `bson:"quantidade" json:"quantidade"`
}
