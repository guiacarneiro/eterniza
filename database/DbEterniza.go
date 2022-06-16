package database

import (
	"fmt"
	"github.com/guiacarneiro/eterniza/config"
	"github.com/guiacarneiro/eterniza/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"sync"
	"time"

	sqliteEncrypt "github.com/jackfr0st13/gorm-sqlite-cipher"
	gormlog "gorm.io/gorm/logger"
)

// Database ...
var Database struct {
	Conn        *gorm.DB
	Transaction *gorm.DB
	sync.RWMutex
}

var (
	contadorSemaforoDebug int
	start                 time.Time
)

func init() {
	err := AbreConexaoBancoProcesso()
	if err != nil {
		logger.LogDebug.Println(err)
	}
}

func getLoggerConfig() gormlog.Interface {
	newLogger := gormlog.New(
		log.New(logger.LogError.Writer(), "", log.LstdFlags), // io writer
		gormlog.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  gormlog.Error, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)
	return newLogger
}

// AbreConexaoBancoProcesso ...
func AbreConexaoBancoProcesso() error {
	var err error
	dbType := config.GetPropriedade("DB_TYPE")
	if dbType == "" {
		dbType = "sqlite"
	}
	dbString := config.GetPropriedade("DB_STRING")
	if dbString == "" {
		dbString = "siganet.db"
	}
	if dbType == "sqlite" {
		key := "0(]BiX@rD20h6eAJJFn&bSN4no:gyI"
		filePath, err := config.GetConfigPath()
		if err != nil {
			filePath = "./"
		}
		dbnameWithDSN := filePath + dbString + fmt.Sprintf("?_pragma_key=%s&_pragma_cipher_page_size=4096", key)
		logger.LogInfo.Println(filePath + dbString)
		Database.Conn, err = gorm.Open(sqliteEncrypt.Open(dbnameWithDSN), &gorm.Config{Logger: getLoggerConfig(), DisableForeignKeyConstraintWhenMigrating: true})
		if err != nil {
			logger.LogError.Println(err)
		}
		// Database.Conn, err = gorm.Open(sqlite.Open(dbString), &gorm.Config{Logger: getLoggerConfig()})
	} else {
		Database.Conn, err = gorm.Open(mysql.Open(dbString), &gorm.Config{Logger: getLoggerConfig()})
		if err != nil {
			logger.LogError.Println(err)
		}
	}
	return err
}

// CriaConexaoBancoRaw ...
func CriaConexaoBancoRaw() error {
	var err error
	Database.Lock()

	contadorSemaforoDebug++
	if contadorSemaforoDebug != 1 {
		contadorSemaforoDebug--
		Database.Unlock()
		return fmt.Errorf("### ERRO SOLICITANDO SEMAFORO %d ###", contadorSemaforoDebug)
	}

	start = time.Now()

	if Database.Conn != nil {
		sqlDB, err := Database.Conn.DB()
		if err != nil {
			return err
		}
		err = sqlDB.Ping()
		if err != nil {
			logger.LogDebug.Printf("Conexao Caiu: %v", err)
			err = AbreConexaoBancoProcesso()
			if err != nil {
				logger.LogDebug.Printf("Erro Conectando: %v", err)
				contadorSemaforoDebug--
				Database.Unlock()
				return err
			}
		}
	} else {
		err = AbreConexaoBancoProcesso()
		if err != nil {
			logger.LogDebug.Printf("Erro Conectando: %v", err)
			contadorSemaforoDebug--
			Database.Unlock()
			return err
		}
	}
	Database.Transaction = Database.Conn.Begin()

	return nil
}

// Rollback ...
func Rollback(err error) {
	logger.LogDebug.Printf("Rollback: " + err.Error())
	Database.Transaction.Rollback()

	Database.Transaction = Database.Conn.Begin()
}

// FechaConexaoBancoRaw ...
func FechaConexaoBancoRaw() {
	contadorSemaforoDebug--
	if contadorSemaforoDebug != 0 {
		logger.LogDebug.Printf("### ERRO SOLICITANDO SEMAFORO %d ###", contadorSemaforoDebug)
	}

	Database.Transaction.Commit()

	Database.Unlock()

	elapsed := time.Since(start)
	if elapsed.Nanoseconds() > int64(150000000) {
		logger.LogDebug.Printf("Tempo conexao: %s", elapsed)
	}
}

func Migrate(entidade interface{}) {
	//logger.LogDebug.Printf("Migrate %T", entidade)
	err := Database.Conn.AutoMigrate(entidade)
	if err != nil {
		logger.LogDebug.Println(err)
	}
}
