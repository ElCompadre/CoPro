package models




type CodePostal struct {
	Base
	CodePostal string `json:"code_postal" gorm:"type: varchar(50)"`
	Commune string `json:"commune" gorm:"type: varchar(200)"`
	Localite string `json:"localite" gorm:"type: varchar(200)"`
	Province string `json:"province" gorm:"type: varchar(300)"`
	Pays string `json:"pays" gorm:"type: varchar(200)"`
}
