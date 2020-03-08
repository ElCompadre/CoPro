package models


type TypeLot struct {
	Base
	Type string `json:"type" gorm:"type: varchar(200)"`
}