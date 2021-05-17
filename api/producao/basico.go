package producao

import (
	"eterniza/database"
	"gorm.io/gorm"
)

// DB buscado
var db *gorm.DB

func init() {
	db = database.DB
}
