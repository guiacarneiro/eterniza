/*
 * Created on: 15/01/2019
 *     Author: guilhermehenrique
 */

package logger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/guiacarneiro/eterniza/config"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	Debug   = 4
	Info    = 3
	Warning = 2
	Error   = 1
	Fatal   = -1
)

var (
	onceLogger        sync.Once
	mutexLogger       = &sync.Mutex{}
	caminhoArquivoLog string
	nomeArquivoLog    string
	logLevel          int
	versao            string

	// LogFatal - log do tipo Fatal
	LogFatal = log.New(logWriter{Fatal}, "FATAL: ", 0)

	// LogError - log do tipo Error
	LogError = log.New(logWriter{Error}, "ERROR: ", 0)

	// LogWarning - log do tipo Warning
	LogWarning = log.New(logWriter{Warning}, "WARN: ", 0)

	// LogInfo - log do tipo Info
	LogInfo = log.New(logWriter{Info}, "INFO: ", 0)

	// LogDebug - log do tipo Debug
	LogDebug = log.New(logWriter{Debug}, "DEBUG: ", 0)
)

func GetPathLog() string {
	var err error
	logPath := config.GetPropriedade("CAMINHO_ARQUIVO_LOG")

	if logPath == "" {
		logPath, err = config.GetConfigPath()
		if err != nil {
			logPath = "./"
		}
		logPath = logPath + "log/"
	}
	return logPath
}

func init() {
	Inicializa(config.GetPropriedadeInt("LOG_LEVEL"),
		GetPathLog(), "siganet3go", config.Versao)
	logFatal := GetPathLog() + "/fatal.log"
	f, err := os.OpenFile(logFatal, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0640)
	if err != nil {
		LogError.Printf("Failed to open log file: %v", err)
	}
	_, err = f.Seek(0, io.SeekEnd)
	if err != nil {
		LogError.Printf("Failed to seek log file to end: %v", err)
	}
	redirectStderr(f)

	gin.DefaultWriter = logWriter{Info}
	gin.DefaultErrorWriter = logWriter{Error}
}

func Inicializa(level int, logPath string, logFileName string, projectVersion string) {
	onceLogger.Do(func() {
		logLevel = level
		caminhoArquivoLog = logPath
		nomeArquivoLog = logFileName
		versao = projectVersion
		_, err := os.Stat(caminhoArquivoLog)
		if os.IsNotExist(err) {
			err = os.MkdirAll(caminhoArquivoLog, os.ModePerm)
			if err != nil {
				panic(fmt.Errorf("erro ao criar diretorio de log: %s", err))
			}
		}
	})
}

type logWriter struct {
	logLevel int
}

func (f logWriter) Write(p []byte) (n int, err error) {
	pc, file, line, ok := runtime.Caller(4)
	if !ok {
		file = "?"
		line = 0
	}

	fn := runtime.FuncForPC(pc)
	var fnName string
	if fn == nil {
		fnName = "?()"
	} else {
		dotName := filepath.Ext(fn.Name())
		fnName = strings.TrimLeft(dotName, ".") + "()"
	}

	gravaLog(Error, fmt.Sprintf("{%s} %s:%d %s: %s", versao, filepath.Base(file), line, fnName, p))
	return len(p), nil
}

func gravaLog(tipoLog int, mensagem string) {
	if logLevel == 0 {
		logLevel = 4
	}
	if caminhoArquivoLog == "" {
		panic("caminho dos arquivos logs nao informado")
	}
	if versao == "" {
		panic("versao do projeto nao informada")
	}

	mutexLogger.Lock()
	defer mutexLogger.Unlock()

	if tipoLog > logLevel {
		return
	}

	current := NowSiga()

	log := strings.Join([]string{
		caminhoArquivoLog, "/", nomeArquivoLog, "_", strconv.Itoa(current.Year()),
		fmt.Sprintf("%02d", current.Month()),
		fmt.Sprintf("%02d", current.Day()), ".log",
	}, "")

	file, err := os.OpenFile(log, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0o644)
	if err != nil {
		panic(fmt.Errorf("erro ao abrir arquivo de log: %s", err))
	}
	defer file.Close()

	_, err = file.WriteString(NowSiga().Format("[02/01/2006 - 15:04:05]") + mensagem)
	if err != nil {
		panic(fmt.Errorf("erro ao escrever no arquivo de log: %s", err))
	}

	err = file.Sync()
	if err != nil {
		panic(fmt.Errorf("erro ao salvar arquivo de log: %s", err))
	}

	if tipoLog == Fatal {
		panic(mensagem)
	}
}

func LimpaLogAntigo() {
	LogDebug.Println("Iniciando limpeza de logs antigos")
	fileInfo, err := ioutil.ReadDir(caminhoArquivoLog)
	if err != nil {
		LogWarning.Println(err.Error())
		return
	}
	now := NowSiga()
	for _, info := range fileInfo {
		if info.IsDir() {
			continue
		}
		if diff := now.Sub(info.ModTime()); diff > 30*24*time.Hour {
			strLog := fmt.Sprintf("%s.*\\.log", nomeArquivoLog)
			match, err := regexp.Match(strLog, []byte(info.Name()))
			if err != nil {
				LogWarning.Println(err.Error())
				return
			}
			if match {
				LogDebug.Printf("Apagando logs antigos: %s (%s)\n", info.Name(), diff)
				err = os.Remove(caminhoArquivoLog + "/" + info.Name())
				if err != nil {
					LogWarning.Println(err)
				}
			}
		}
	}
}
