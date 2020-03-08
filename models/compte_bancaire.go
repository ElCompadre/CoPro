package models

type CompteBancaire struct {
	Base
	Compte string `json:"compte" gorm:"type: varchar(500)"`
	BIC string `json:"bic" gorm:"type: varchar(200)"`
	IBAN string `json:"iban" gorm:"type: varchar(200)"`
}