package service

import (
	"github.com/guiacarneiro/eterniza/database"
	"github.com/guiacarneiro/eterniza/database/model"
	"github.com/guiacarneiro/eterniza/erro"
	"gorm.io/gorm/clause"
)

func GetUsuarioLogin(login string) (*model.User, error) {
	var user model.User
	err := database.Database.Transaction.Preload(clause.Associations).
		Where("email = ?", login).
		First(&user).Error
	if err != nil {
		return nil, erro.UsuarioNaoEncontrado
	}
	return &user, nil
}
