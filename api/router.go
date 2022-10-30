package api

import (
	"encoding/json"
	"fmt"
	"github.com/guiacarneiro/eterniza/api/controller"
	"github.com/guiacarneiro/eterniza/api/util"
	fs2 "github.com/guiacarneiro/eterniza/fs"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/contrib/static"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	assetfs "github.com/elazarl/go-bindata-assetfs"

	"github.com/gin-gonic/gin"
)

// @title Siganet API
// @version 1.0
// @description Swagger API.
// @contact.email contato@visual.com.br
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-API-KEY
// @BasePath /api/v1
// @tag.name Emissão
// @tag.description Métodos de emissão de senha
// @tag.name Autenticação
// @tag.description Método de autenticação do atendente para operar o guichê
// @tag.name Guichê
// @tag.description Métodos de operação de guichê
func ProcessaAPI() *gin.Engine {
	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)
	ginrouter := gin.New()
	ginrouter.Use(gin.Recovery(), util.CORSMiddleware(), util.TransactionMiddleware())

	apiV1 := ginrouter.Group("/api/v1")

	apiV1.POST("/login", util.LogMiddleware(), util.NoCacheMiddleware(), controller.Login)

	// Echo Route
	apiV1.POST("/echo", Echo)

	// User route
	apiV1.POST("/user", controller.CreateUser)
	//apiV1.GET("/user", util.UserAuthenticationMiddleware(), controller.GetUsers)
	//apiV1.GET("/user/:id", util.UserAuthenticationMiddleware(), controller.GetUser)
	//apiV1.PUT("/user/:id", util.UserAuthenticationMiddleware(), controller.UpdateUser)
	//apiV1.DELETE("/user/:id", util.UserAuthenticationMiddleware(), controller.DeleteUser)

	// producao route
	producao := apiV1.Group("/producao")
	producao.Use(util.UserAuthenticationMiddleware())

	//Materia prima routes
	producao.POST("/materiaprima", controller.CreateMateriaPrima)
	//producao.GET("/materiaprima", controller.GetMateriaPrimas)
	//producao.GET("/materiaprima/{id}", controller.GetMateriaPrima)
	//producao.PUT("/materiaprima/{id}", controller.UpdateMateriaPrima)
	//producao.DELETE("/materiaprima/{id}", controller.DeleteMateriaPrima)

	//Produto routes
	producao.POST("/ficha", controller.SalvaFicha)
	producao.POST("/produtos", controller.BuscaProduto)
	producao.POST("/fichas", controller.BuscaFicha)
	//producao.GET("/materiaprima/{id}", controller.GetMateriaPrima)
	//producao.PUT("/materiaprima/{id}", controller.UpdateMateriaPrima)
	producao.DELETE("/ficha/:id", controller.ApagaFicha)

	ginrouter.Use(static.Serve("/", BinaryFileSystem("interface")))
	ginrouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return ginrouter
}

func Echo(c *gin.Context) {
	b, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		util.ERROR(c, http.StatusUnprocessableEntity, err)
	}
	defer c.Request.Body.Close()
	fmt.Println(string(b))
	var val map[string]interface{}
	json.Unmarshal(b, &val)
	c.JSON(http.StatusOK, val)
}

type binaryFileSystem struct {
	fs http.FileSystem
}

func (b *binaryFileSystem) Open(name string) (http.File, error) {
	return b.fs.Open(name)
}

func (b *binaryFileSystem) Exists(prefix string, filepath string) bool {
	if p := strings.TrimPrefix(filepath, prefix); len(p) < len(filepath) {
		if _, err := b.fs.Open(p); err != nil {
			return false
		}
		return true
	}
	return false
}

func BinaryFileSystem(root string) *binaryFileSystem {
	fs := &assetfs.AssetFS{Asset: fs2.Asset, AssetDir: fs2.AssetDir, AssetInfo: fs2.AssetInfo, Prefix: root}
	return &binaryFileSystem{
		fs,
	}
}
