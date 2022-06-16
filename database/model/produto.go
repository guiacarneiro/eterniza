package model

import (
	"errors"
	"strings"
	"time"

	"github.com/guiacarneiro/eterniza/database"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Produto struct {
	gorm.Model
	Referencia  string
	Descricao   string
	Fotos       []Foto
	Componentes []Componente
	Precos      []Preco
}

func init() {
	database.Migrate(&Produto{})
}

func (p *Produto) BeforeSave(tx *gorm.DB) error {
	return nil
}

func (p *Produto) Prepare() {
	p.ID = 0
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *Produto) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if p.Referencia == "" {
			return errors.New("Referência obrigatório")
		}
		return nil
	default:
		if p.Referencia == "" {
			return errors.New("Referência obrigatório")
		}
		return nil
	}
}

func (p *Produto) Save(db *gorm.DB) (*Produto, error) {
	err := db.Debug().Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&p).Error
	if err != nil {
		return &Produto{}, err
	}
	return p, nil
}

func (p *Produto) FindAllProdutos(db *gorm.DB) (*[]Produto, error) {
	var err error
	users := []Produto{}
	err = db.Debug().Model(&Produto{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]Produto{}, err
	}
	return &users, err
}

func (p *Produto) FindProdutoByID(db *gorm.DB, uid uint32) (*Produto, error) {
	err := db.Debug().Model(Produto{}).Where("id = ?", uid).Take(&p).Error
	if err != nil {
		return &Produto{}, err
	}
	if err == gorm.ErrRecordNotFound {
		return &Produto{}, errors.New("Usuário não encontrado")
	}
	return p, err
}

func (p *Produto) DeleteProduto(db *gorm.DB, uid uint32) (int64, error) {
	db = db.Debug().Model(&Produto{}).Where("id = ?", uid).Take(&Produto{}).Delete(&Produto{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
