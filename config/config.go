package config

import (
	"fmt"
	"github.com/guiacarneiro/eterniza/config/viper"
	"os"
	"path/filepath"
)

var (
	// Versao - versao da aplicacao
	Versao                  = ""
	nomeArquivoConfiguracao string
)

func GetConfigPath() (string, error) {
	fullexecpath, err := os.Executable()
	if err != nil {
		return "", err
	}

	dir, _ := filepath.Split(fullexecpath)
	// ext := filepath.Ext(execname)
	// name := execname[:len(execname)-len(ext)]

	return dir, nil
}

func init() {
	Inicializa("eterniza", "0.0.1")
}

// Inicializa - inicializa configuracoes
func Inicializa(nomeArquivoConfig string, versao string) {
	nomeArquivoConfiguracao = nomeArquivoConfig
	Versao = versao
	viper.SetConfigName(nomeArquivoConfiguracao)
	path, err := GetConfigPath()
	if err == nil {
		viper.AddConfigPath(path)
	}
	viper.AddConfigPath("./")
	viper.AddConfigPath("./config/")

	errConfig := viper.ReadInConfig()
	if errConfig != nil {
		fmt.Printf("erro ao buscar arquivo de configurações: %s\n", errConfig)
	}
}

// GetPropriedade - busca propriedade no arquivo de configuracoes
func GetPropriedade(propriedade string) string {
	return viper.GetString(propriedade)
}

// GetPropriedadeInt - busca propriedade no arquivo de configuracoes
func GetPropriedadeInt(propriedade string) int {
	return viper.GetInt(propriedade)
}

// UnmarshalProperties - busca propriedade no arquivo de configuracoes
func UnmarshalProperties(obj interface{}) error {
	err := viper.Unmarshal(obj)
	return err
}

// GetPropriedadeBool - busca propriedade no arquivo de configuracoes
func GetPropriedadeBool(propriedade string) bool {
	return viper.GetBool(propriedade)
}
