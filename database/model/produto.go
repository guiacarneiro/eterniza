package model

import (
	"errors"
	"github.com/guiacarneiro/eterniza/database"
	"github.com/guiacarneiro/eterniza/integracao/tiny"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strings"
)

type ItemFicha struct {
	Texto          string `json:"texto,omitempty"`
	FichaTecnicaID uint   `json:"fichaTecnicaId,omitempty"`
}

type FichaTecnica struct {
	gorm.Model
	Variacao      string        `json:"variacao,omitempty"`
	Fotos         []Foto        `json:"fotos,omitempty"`
	Componentes   []ItemFicha   `json:"componentes,omitempty"`
	ProdutoTinyID string        `json:"produtoTinyID,omitempty"`
	ProdutoTiny   *tiny.Produto `gorm:"-" json:"produtoTiny,omitempty"`
}

func (FichaTecnica) TableName() string {
	return "ficha_tecnica"
}

func (ItemFicha) TableName() string {
	return "item_ficha"
}

func init() {
	database.Migrate(&ItemFicha{})
	database.Migrate(&FichaTecnica{})
}

func (p *FichaTecnica) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if p.ProdutoTinyID == "" {
			return errors.New("Referência obrigatório")
		}
		return nil
	default:
		if p.ProdutoTinyID == "" {
			return errors.New("Referência obrigatório")
		}
		return nil
	}
}

func (p *FichaTecnica) Save() (*FichaTecnica, error) {
	var err error
	if p.ID != 0 {
		err = database.Database.Transaction.Where("ficha_tecnica_id = ?", p.ID).Delete(&ItemFicha{}).Error
	}
	if err != nil {
		return &FichaTecnica{}, err
	}
	err = database.Database.Transaction.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&p).Error
	if err != nil {
		return &FichaTecnica{}, err
	}
	return p, nil
}

func (p *FichaTecnica) FindAllFichasTecnicas(db *gorm.DB) (*[]FichaTecnica, error) {
	var err error
	users := []FichaTecnica{}
	err = db.Debug().Model(&FichaTecnica{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]FichaTecnica{}, err
	}
	return &users, err
}

func (p *FichaTecnica) FindFichaTecnicaByID(db *gorm.DB, uid uint32) (*FichaTecnica, error) {
	err := db.Debug().Model(FichaTecnica{}).Where("id = ?", uid).Take(&p).Error
	if err != nil {
		return &FichaTecnica{}, err
	}
	if err == gorm.ErrRecordNotFound {
		return &FichaTecnica{}, errors.New("Ficha não encontrado")
	}
	return p, err
}

func FindFichasTecnicasByIDTiny(uid string) ([]FichaTecnica, error) {
	var p []FichaTecnica
	err := database.Database.Transaction.Preload(clause.Associations).Model(FichaTecnica{}).Where("produto_tiny_id = ?", uid).Find(&p).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return make([]FichaTecnica, 0), nil
		}
		return nil, err
	}
	return p, err
}

func (p *FichaTecnica) DeleteFichaTecnica(db *gorm.DB, uid uint32) (int64, error) {
	db = db.Debug().Model(&FichaTecnica{}).Where("id = ?", uid).Take(&FichaTecnica{}).Delete(&FichaTecnica{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
