package model

import (
	"github.com/guiacarneiro/eterniza/database"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
		"Weight",
	}
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
	database.Migrate(&MateriaPrima{})
}

func ListaMateriaPrima(offset int, qte int) (*[]MateriaPrima, error) {
	var err error
	var result []MateriaPrima
	err = database.Database.Transaction.Model(&MateriaPrima{}).Offset(offset).Limit(qte).Find(&result).Error
	if err != nil {
		return &[]MateriaPrima{}, err
	}
	return &result, err
}

func (a *MateriaPrima) Save() error {
	err := database.Database.Transaction.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&a).Error
	if err != nil {
		return err
	}
	return nil
}

func FindMateriaPrimaById(id uint) (*MateriaPrima, error) {
	var err error
	var result MateriaPrima
	err = database.Database.Transaction.Model(&MateriaPrima{}).First(&result, id).Error
	if err != nil {
		return nil, err
	}
	return &result, err
}

func DeleteMateriaPrimaById(id uint) (*MateriaPrima, error) {
	var err error
	var result MateriaPrima
	err = database.Database.Transaction.Model(&MateriaPrima{}).Delete(&result, id).Error
	if err != nil {
		return nil, err
	}
	return &result, err
}
