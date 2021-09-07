package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/guiacarneiro/eterniza/api"
	"github.com/guiacarneiro/eterniza/api/producao"
	"io/ioutil"
	"net/http"
	"strconv"
)

func CreateMateriaPrima(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		api.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	materiaPrima := producao.MateriaPrima{}
	err = json.Unmarshal(body, &materiaPrima)
	if err != nil {
		api.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	if err != nil {
		api.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	materiaPrima.Save()

	if err != nil {

		formattedError := api.FormatError(err.Error())

		api.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, materiaPrima.ID))
	api.JSON(w, http.StatusCreated, materiaPrima)
}

func GetMateriaPrimas(w http.ResponseWriter, r *http.Request) {

	listaMateriaPrima, err := producao.ListaMateriaPrima(0, 100)
	if err != nil {
		api.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	api.JSON(w, http.StatusOK, listaMateriaPrima)
}

func GetMateriaPrima(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		api.ERROR(w, http.StatusBadRequest, err)
		return
	}
	materiaPrima, err := producao.FindMateriaPrimaById(uint(uid))
	if err != nil {
		api.ERROR(w, http.StatusBadRequest, err)
		return
	}
	api.JSON(w, http.StatusOK, materiaPrima)
}

func UpdateMateriaPrima(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		api.ERROR(w, http.StatusBadRequest, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		api.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	materiaPrima := producao.MateriaPrima{}
	materiaPrima.ID = uint(uid)
	err = json.Unmarshal(body, &materiaPrima)
	if err != nil {
		api.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	tokenID, err := api.ExtractTokenID(r)
	if err != nil {
		api.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if tokenID != uint32(uid) {
		api.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	//materiaPrima.Prepare()
	//err = materiaPrima.Validate("update")
	//if err != nil {
	//	api.ERROR(w, http.StatusUnprocessableEntity, err)
	//	return
	//}

	materiaPrimaSalva, err := materiaPrima.Save()
	if err != nil {
		formattedError := api.FormatError(err.Error())
		api.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	api.JSON(w, http.StatusOK, materiaPrimaSalva)
}

func DeleteMateriaPrima(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		api.ERROR(w, http.StatusBadRequest, err)
		return
	}
	tokenID, err := api.ExtractTokenID(r)
	if err != nil {
		api.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if tokenID != 0 && tokenID != uint32(uid) {
		api.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	_, err = producao.DeleteMateriaPrimaById(uint(uid))
	if err != nil {
		api.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", uid))
	api.JSON(w, http.StatusNoContent, "")
}
