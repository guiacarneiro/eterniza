package main

import (
	"github.com/guiacarneiro/eterniza/api"
	"github.com/guiacarneiro/eterniza/config"
	"github.com/guiacarneiro/eterniza/logger"
	"log"
	"net/http"
	"os"
)

func main() {

	router := api.ProcessaAPI()

	server := config.GetPropriedade("SERVER")
	logger.LogInfo.Println("Listening to " + server)
	tls := config.GetPropriedadeBool("TLS")
	if tls {
		configPath, err := config.GetConfigPath()
		if err != nil {
			configPath = "./"
		}
		cert := configPath + config.GetPropriedade("TLS_CERT")
		key := configPath + config.GetPropriedade("TLS_KEY")
		if _, err := os.Stat(cert); err == nil {
			if _, err := os.Stat(key); err == nil {
				logger.LogInfo.Println("HTTPS habilitado em " + server)
				log.Fatal(http.ListenAndServeTLS(server, cert, key, router))
			}
		}
		logger.LogInfo.Println("HTTPS nao habilitado em " + server)
		logger.LogInfo.Println("HTTPS cert " + cert)
		logger.LogInfo.Println("HTTPS key " + key)
	}
	// se o tls falhar ou estiver desabilitado, roda no http
	log.Fatal(http.ListenAndServe(server, router))
}
