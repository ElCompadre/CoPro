package models


type Lot struct {
	Base
	NumLot string `json:"num_lot" gorm:"type: varchar(50)"`
	NumCave string `json:"num_cave" gorm:"type: varchar(50)"`
	NumParking string `json:"num_parking" gorm:"type: varchar(50)"`
	NumBoite string `json:"num_boite" gorm:"type: varchar(10)"`
	Quotite float64 `json:"quotite"`
	Etage int `json:"etage"`
	SurfaceSol float64 `json:"surface_sol"`
	SurfaceJardin float64 `json:"surface_jardin"`
	SurfaceTerrasse float64 `json:"surface_terrasse"`
	Copro CoPro `json:"copro" gorm:"foreignkey:ID"`
	TypeLot TypeLot `json:"typelot" gorm:"foreignkey:ID"`
} 
