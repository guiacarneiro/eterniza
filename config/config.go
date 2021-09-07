package config

import (
	"fmt"
	"github.com/guiacarneiro/eterniza/logger"
	"sync"

	"github.com/spf13/viper"
)

var (
	//Versao - versao da aplicacao
	Versao                  = 0
	nomeArquivoConfiguracao string
	once                    sync.Once
)

func init() {
	nomeArquivoConfiguracao = "eterniza_config"
	Versao = 1
}

//GetPropriedade - busca propriedade no arquivo de configuracoes
func GetPropriedade(propriedade string) string {
	once.Do(func() {
		viper.SetConfigName(nomeArquivoConfiguracao)
		viper.AddConfigPath("./")
		viper.AddConfigPath("./config/")

		errConfig := viper.ReadInConfig()
		if errConfig != nil {
			logger.LogError.Println(fmt.Errorf("erro ao buscar arquivo de configurações: %s", errConfig))
			panic(fmt.Errorf("erro ao buscar arquivo de configurações: %s", errConfig))
		}
	})
	return viper.GetString(propriedade)
}

//GetPropriedade - busca propriedade no arquivo de configuracoes
func GetPropriedadeDefault(propriedade, defaultVal string) (val string) {
	once.Do(func() {
		viper.SetConfigName(nomeArquivoConfiguracao)
		viper.AddConfigPath("./")
		viper.AddConfigPath("./config/")

		errConfig := viper.ReadInConfig()
		if errConfig != nil {
			logger.LogError.Println(fmt.Errorf("erro ao buscar arquivo de configurações: %s", errConfig))
			panic(fmt.Errorf("erro ao buscar arquivo de configurações: %s", errConfig))
		}
	})
	val = viper.GetString(propriedade)
	if val == "" {
		val = defaultVal
	}
	return
}

//GetPropriedadeInt - busca propriedade no arquivo de configuracoes
func GetPropriedadeInt(propriedade string) int {
	once.Do(func() {
		viper.SetConfigName(nomeArquivoConfiguracao)
		viper.AddConfigPath("./")
		viper.AddConfigPath("./config/")

		errConfig := viper.ReadInConfig()
		if errConfig != nil {
			logger.LogError.Println(fmt.Errorf("erro ao buscar arquivo de configurações: %s", errConfig))
			panic(fmt.Errorf("erro ao buscar arquivo de configurações: %s", errConfig))
		}
	})
	return viper.GetInt(propriedade)
}

//UnmarshalProperties - busca propriedade no arquivo de configuracoes
func UnmarshalProperties(obj interface{}) error {
	once.Do(func() {
		viper.SetConfigName(nomeArquivoConfiguracao)
		viper.AddConfigPath("./")
		viper.AddConfigPath("./config/")

		errConfig := viper.ReadInConfig()
		if errConfig != nil {
			logger.LogError.Println(fmt.Errorf("erro ao buscar arquivo de configurações: %s", errConfig))
			panic(fmt.Errorf("erro ao buscar arquivo de configurações: %s", errConfig))
		}
	})
	err := viper.Unmarshal(obj)
	return err
}
