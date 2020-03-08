package models

import (
	"github.com/graphql-go/graphql"
)

type Adresse struct {
	Base
	Adresse string `json:"adresse" gorm:"type:varchar(500)"`
	Boite string `json:"boite" gorm="type:varchar(100)"`
	CodePostal *CodePostal `json:"code_postal" gorm:"association_foreignkey:ID"`
}

var adresseType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Address",
		Fields: graphql.Fields{
			"adresse": &graphql.Field{
				Type: graphql.String,
			},
			"boite": &graphql.Field{
				Type: graphql.String,
			},
			"code_postal": &graphql.Field{
				Type: graphql.String,
			},
		},
	}) 