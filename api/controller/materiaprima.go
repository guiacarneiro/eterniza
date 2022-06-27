package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/guiacarneiro/eterniza/api/request"
	"github.com/guiacarneiro/eterniza/api/response"
	"github.com/guiacarneiro/eterniza/api/util"
	"github.com/guiacarneiro/eterniza/database/model"
	"net/http"
	"strconv"
)

func CreateMateriaPrima(c *gin.Context) {
	req := request.CreateMateriaPrima{}
	if err := c.BindJSON(&req); err != nil {
		util.ERROR(c, http.StatusBadRequest, err)
		return
	}
	valor, err := strconv.ParseFloat(req.Value, 64)
	if err != nil {
		util.ERROR(c, http.StatusBadRequest, errors.New("Valor inválido"))
		return
	}
	mp := model.MateriaPrima{
		Unity: model.UnityType(req.Unity),
		Label: req.Label,
		Value: valor,
	}

	err = mp.Save()
	if err != nil {
		formattedError := util.FormatError(err.Error())
		util.ERROR(c, http.StatusUnprocessableEntity, formattedError)
		return
	}
	c.JSON(http.StatusOK, response.Basic{
		Success: true,
	})
}

//func CreateMateriaPrima(w http.ResponseWriter, r *http.Request) {
//
//	body, err := ioutil.ReadAll(r.Body)
//	if err != nil {
//		util.ERROR(w, http.StatusUnprocessableEntity, err)
//	}
//	materiaPrima := model.MateriaPrima{}
//	err = json.Unmarshal(body, &materiaPrima)
//	if err != nil {
//		util.ERROR(w, http.StatusUnprocessableEntity, err)
//		return
//	}
//	if err != nil {
//		util.ERROR(w, http.StatusUnprocessableEntity, err)
//		return
//	}
//	materiaPrima.Save()
//
//	if err != nil {
//
//		formattedError := util.FormatError(err.Error())
//
//		util.ERROR(w, http.StatusInternalServerError, formattedError)
//		return
//	}
//	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, materiaPrima.ID))
//	util.JSON(w, http.StatusCreated, materiaPrima)
//}
//
//func GetMateriaPrimas(w http.ResponseWriter, r *http.Request) {
//
//	listaMateriaPrima, err := model.ListaMateriaPrima(0, 100)
//	if err != nil {
//		util.ERROR(w, http.StatusInternalServerError, err)
//		return
//	}
//	util.JSON(w, http.StatusOK, listaMateriaPrima)
//}
//
//func GetMateriaPrima(w http.ResponseWriter, r *http.Request) {
//
//	vars := mux.Vars(r)
//	uid, err := strconv.ParseUint(vars["id"], 10, 32)
//	if err != nil {
//		util.ERROR(w, http.StatusBadRequest, err)
//		return
//	}
//	materiaPrima, err := model.FindMateriaPrimaById(uint(uid))
//	if err != nil {
//		util.ERROR(w, http.StatusBadRequest, err)
//		return
//	}
//	util.JSON(w, http.StatusOK, materiaPrima)
//}
//
//func UpdateMateriaPrima(w http.ResponseWriter, r *http.Request) {
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
//	materiaPrima := model.MateriaPrima{}
//	materiaPrima.ID = uint(uid)
//	err = json.Unmarshal(body, &materiaPrima)
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
//	//materiaPrima.Prepare()
//	//err = materiaPrima.Validate("update")
//	//if err != nil {
//	//	api.ERROR(w, http.StatusUnprocessableEntity, err)
//	//	return
//	//}
//
//	materiaPrimaSalva, err := materiaPrima.Save()
//	if err != nil {
//		formattedError := util.FormatError(err.Error())
//		util.ERROR(w, http.StatusInternalServerError, formattedError)
//		return
//	}
//	util.JSON(w, http.StatusOK, materiaPrimaSalva)
//}
//
//func DeleteMateriaPrima(w http.ResponseWriter, r *http.Request) {
//
//	vars := mux.Vars(r)
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
//	_, err = model.DeleteMateriaPrimaById(uint(uid))
//	if err != nil {
//		util.ERROR(w, http.StatusInternalServerError, err)
//		return
//	}
//	w.Header().Set("Entity", fmt.Sprintf("%d", uid))
//	util.JSON(w, http.StatusNoContent, "")
//}
