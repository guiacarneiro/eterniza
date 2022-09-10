package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/guiacarneiro/eterniza/api/request"
	"github.com/guiacarneiro/eterniza/api/util"
	"github.com/guiacarneiro/eterniza/database/model"
	"github.com/guiacarneiro/eterniza/integracao/tiny"
	"gorm.io/gorm"
	"net/http"
)

func BuscaProduto(c *gin.Context) {
	req := tiny.RequestPesquisaProdutos{}
	if err := c.BindJSON(&req); err != nil {
		util.ERROR(c, http.StatusBadRequest, err)
		return
	}

	if req.Situacao == nil {
		ativo := "A"
		req.Situacao = &ativo
	}
	resposta := tiny.ResponsePesquisaProdutos{}

	err := tiny.ChamaMetodoTiny(req, &resposta, tiny.PesquisaProdutos)
	if err != nil {
		formattedError := util.FormatError(err.Error())
		util.ERROR(c, http.StatusUnprocessableEntity, formattedError)
		return
	}
	c.JSON(http.StatusOK, resposta)
}

func BuscaFicha(c *gin.Context) {
	req := request.BuscaFicha{}
	if err := c.BindJSON(&req); err != nil {
		util.ERROR(c, http.StatusBadRequest, err)
		return
	}

	resposta, err := model.FindFichasTecnicasByIDTiny(req.TinyID)
	if err != nil {
		formattedError := util.FormatError(err.Error())
		util.ERROR(c, http.StatusUnprocessableEntity, formattedError)
		return
	}

	reqTiny := tiny.RequestObterProduto{ID: req.TinyID}
	respostaTiny := tiny.ResponseObterProduto{}

	err = tiny.ChamaMetodoTiny(reqTiny, &respostaTiny, tiny.ObterProduto)
	if err != nil {
		formattedError := util.FormatError(err.Error())
		util.ERROR(c, http.StatusUnprocessableEntity, formattedError)
		return
	}

	for i := range resposta {
		resposta[i].ProdutoTiny = &respostaTiny.Retorno.Produto
	}

	c.JSON(http.StatusOK, resposta)
}

func SalvaFicha(c *gin.Context) {
	req := request.SalvaFicha{}
	if err := c.BindJSON(&req); err != nil {
		util.ERROR(c, http.StatusBadRequest, err)
		return
	}
	if len(req.Componentes) == 0 {
		util.ERROR(c, http.StatusBadRequest, errors.New("A ficha precisa ter componentes"))
		return
	}
	var componentes []model.ItemFicha
	for _, componete := range req.Componentes {
		item := model.ItemFicha{
			Texto: componete,
		}
		componentes = append(componentes, item)
	}
	produto := model.FichaTecnica{
		Model:         gorm.Model{ID: req.ID},
		Variacao:      req.Variacao,
		Fotos:         nil,
		Componentes:   componentes,
		ProdutoTinyID: req.TinyID,
		ProdutoTiny:   nil,
	}
	resposta, err := produto.Save()
	if err != nil {
		formattedError := util.FormatError(err.Error())
		util.ERROR(c, http.StatusUnprocessableEntity, formattedError)
		return
	}

	c.JSON(http.StatusOK, resposta)
}

//
//func CreateProduto(w http.ResponseWriter, r *http.Request) {
//
//	body, err := ioutil.ReadAll(r.Body)
//	if err != nil {
//		util.ERROR(w, http.StatusUnprocessableEntity, err)
//	}
//	user := model.Componentes{}
//	err = json.Unmarshal(body, &user)
//	if err != nil {
//		util.ERROR(w, http.StatusUnprocessableEntity, err)
//		return
//	}
//	if err != nil {
//		util.ERROR(w, http.StatusUnprocessableEntity, err)
//		return
//	}
//	userCreated, err := user.Save(database.DB)
//
//	if err != nil {
//
//		formattedError := util.FormatError(err.Error())
//
//		util.ERROR(w, http.StatusInternalServerError, formattedError)
//		return
//	}
//	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, userCreated.ID))
//	util.JSON(w, http.StatusCreated, userCreated)
//}
//
//func GetProdutos(w http.ResponseWriter, r *http.Request) {
//
//	user := model.Componentes{}
//
//	users, err := user.FindAllFichasTecnicas(database.DB)
//	if err != nil {
//		util.ERROR(w, http.StatusInternalServerError, err)
//		return
//	}
//	util.JSON(w, http.StatusOK, users)
//}
//
//func GetProduto(w http.ResponseWriter, r *http.Request) {
//
//	vars := mux.Vars(r)
//	uid, err := strconv.ParseUint(vars["id"], 10, 32)
//	if err != nil {
//		util.ERROR(w, http.StatusBadRequest, err)
//		return
//	}
//	user := model.Componentes{}
//	userGotten, err := user.FindFichaTecnicaByID(database.DB, uint32(uid))
//	if err != nil {
//		util.ERROR(w, http.StatusBadRequest, err)
//		return
//	}
//	util.JSON(w, http.StatusOK, userGotten)
//}
//
//func UpdateProduto(w http.ResponseWriter, r *http.Request) {
//
//	vars := mux.Vars(r)
//	uid, err := strconv.ParseUint(vars["id"], 10, 32)
//	if err != nil {
//		util.ERROR(w, http.StatusBadRequest, err)
//		return
//	}
//	body, err := ioutil.ReadAll(r.Body)
//	if err != nil {
//		util.ERROR(w, http.StatusUnprocessableEntity, err)
//		return
//	}
//	produto := model.Componentes{}
//	produto.ID = uint(uid)
//	err = json.Unmarshal(body, &produto)
//	if err != nil {
//		util.ERROR(w, http.StatusUnprocessableEntity, err)
//		return
//	}
//	tokenID, err := util.ExtractTokenID(r)
//	if err != nil {
//		util.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
//		return
//	}
//	if tokenID != uint32(uid) {
//		util.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
//		return
//	}
//	produto.Prepare()
//	err = produto.Validate("update")
//	if err != nil {
//		util.ERROR(w, http.StatusUnprocessableEntity, err)
//		return
//	}
//	updatedProduto, err := produto.Save(database.DB)
//	if err != nil {
//		formattedError := util.FormatError(err.Error())
//		util.ERROR(w, http.StatusInternalServerError, formattedError)
//		return
//	}
//	util.JSON(w, http.StatusOK, updatedProduto)
//}
//
//func DeleteFichaTecnica(w http.ResponseWriter, r *http.Request) {
//
//	vars := mux.Vars(r)
//
//	user := model.Componentes{}
//
//	uid, err := strconv.ParseUint(vars["id"], 10, 32)
//	if err != nil {
//		util.ERROR(w, http.StatusBadRequest, err)
//		return
//	}
//	tokenID, err := util.ExtractTokenID(r)
//	if err != nil {
//		util.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
//		return
//	}
//	if tokenID != 0 && tokenID != uint32(uid) {
//		util.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
//		return
//	}
//	_, err = user.DeleteFichaTecnica(database.DB, uint32(uid))
//	if err != nil {
//		util.ERROR(w, http.StatusInternalServerError, err)
//		return
//	}
//	w.Header().Set("Entity", fmt.Sprintf("%d", uid))
//	util.JSON(w, http.StatusNoContent, "")
//}
