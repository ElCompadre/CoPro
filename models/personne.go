package models

import (
	"time"
)

type Personne struct {
	Base
	PersonneID string `json:"id" gorm"type uuid"`
	Nom string `json:"nom" gorm:"type: varchar(200)"`
	Prenom string `json:"prenom" gorm:"type: varchar(200)"`
	SecondPrenom string `json:"second_prenom" gorm:"type: varchar(200)"`
	DateDeNaissance time.Time `json:"date_de_naissance"`
	Adresse Adresse `json:"adresse" gorm:"foreignkey:ID"`
	CompteBancaire CompteBancaire `json:"compte_bancaire" gorm:"foreignkey:ID"`
	User `json:"user" gorm:"foreignkey:ID"`
}