package models


type CoPro struct {
	Base
	Nom string `json:"nom" gorm:"type: varchar(300)"`
	BCE string `json:"bce" gorm:"type: varchar(200)"`
	CompteBancaireID string `json:"compte_bancaire" sql:"type:uuid REFERENCES CompteBancaire(ID)"`
}
