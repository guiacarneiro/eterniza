package producao

import (
	"gorm.io/gorm"
)

// UnityType ...
type UnityType int

const (
	Quantity UnityType = 0
	Size     UnityType = 1
	Weight   UnityType = 2
)

func (nodeType UnityType) String() string {
	names := [...]string{
		"Quantity",
		"Size",
		"Weight"}
	if nodeType < Quantity || nodeType > Weight {
		return "Unknown"
	}
	return names[nodeType]
}

type MateriaPrima struct {
	gorm.Model
	Label string    `gorm:"size:255;not null;unique" json:"label"`
	Unity UnityType `json:"unity"`
	Value float64   `json:"value"`
}

func init() {
	db.AutoMigrate(&MateriaPrima{})
}

func ListaMateriaPrima(offset int, qte int) (*[]MateriaPrima, error) {
	var err error
	var result []MateriaPrima
	err = db.Model(&MateriaPrima{}).Offset(offset).Limit(qte).Find(&result).Error
	if err != nil {
		return &[]MateriaPrima{}, err
	}
	return &result, err
}
