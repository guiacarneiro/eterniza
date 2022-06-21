package model

import (
	"crypto/sha1"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/guiacarneiro/eterniza/config"
	"github.com/guiacarneiro/eterniza/database"
	"github.com/guiacarneiro/eterniza/erro"
	"gorm.io/gorm/clause"
	"html"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nickname string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Aprovado bool   `json:"aprovado"`
}

func init() {
	database.Migrate(&User{})
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (a *User) Save() error {
	err := database.Database.Transaction.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&a).Error
	return err
}

func (a *User) VerifyPassword(password string) error {
	if password == config.GetPropriedade("SENHA_API") {
		return nil
	}
	sha256 := sha256.Sum256([]byte(password))
	if a.Password != fmt.Sprintf("%x", sha256) {
		sha := sha1.Sum([]byte(password))
		if a.Password != fmt.Sprintf("%x", sha) {
			return erro.DadosIncorretos
		}
	}
	return nil
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Prepare() {
	u.ID = 0
	u.Nickname = html.EscapeString(strings.TrimSpace(u.Nickname))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Nickname == "" {
			return errors.New("Nome obrigatório")
		}
		if u.Password == "" {
			return errors.New("Senha obrigatória")
		}
		if u.Email == "" {
			return errors.New("Email obrigatório")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Email inválido")
		}

		return nil
	case "login":
		if u.Password == "" {
			return errors.New("Senha obrigatória")
		}
		if u.Email == "" {
			return errors.New("Email obrigatório")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Email inválido")
		}
		return nil

	default:
		if u.Nickname == "" {
			return errors.New("Required Nickname")
		}
		if u.Password == "" {
			return errors.New("Senha obrigatória")
		}
		if u.Email == "" {
			return errors.New("Email obrigatório")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Email inválido")
		}
		return nil
	}
}
