package models

type RolePersonne struct {
	Base
	Role string `json:"role" gorm:"type:varchar(100)"`
	CoProID string `json:"copro_id" sql:"type:uuid REFERENCES CoPro(ID)"`
	PersonneID string `json:"personne_id" sql:"type:uuid REFERENCES Personne(ID)"`
}