package database

import (
	"eterniza/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB a ser utilizado pelas entidades
var DB *gorm.DB

func init() {
	dbUser := config.GetPropriedade("db_user")
	dbPassword := config.GetPropriedade("db_password")
	dbHost := config.GetPropriedade("db_host")
	dbPort := config.GetPropriedade("db_port")
	dbName := config.GetPropriedade("db_name")
	sqlConn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	var err error
	DB, err = gorm.Open(mysql.Open(sqlConn), &gorm.Config{})
	if err != nil {
		panic("Sem conexao com banco")
	}
}
