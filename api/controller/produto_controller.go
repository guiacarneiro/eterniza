package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/guiacarneiro/eterniza/api"
	"github.com/guiacarneiro/eterniza/api/producao"
	"github.com/guiacarneiro/eterniza/database"
	"io/ioutil"
	"net/http"
	"strconv"
)

func CreateProduto(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		api.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	user := producao.Produto{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		api.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	if err != nil {
		api.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	userCreated, err := user.Save(database.DB)

	if err != nil {

		formattedError := api.FormatError(err.Error())

		api.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, userCreated.ID))
	api.JSON(w, http.StatusCreated, userCreated)
}

func GetProdutos(w http.ResponseWriter, r *http.Request) {

	user := producao.Produto{}

	users, err := user.FindAllProdutos(database.DB)
	if err != nil {
		api.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	api.JSON(w, http.StatusOK, users)
}

func GetProduto(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		api.ERROR(w, http.StatusBadRequest, err)
		return
	}
	user := producao.Produto{}
	userGotten, err := user.FindProdutoByID(database.DB, uint32(uid))
	if err != nil {
		api.ERROR(w, http.StatusBadRequest, err)
		return
	}
	api.JSON(w, http.StatusOK, userGotten)
}

func UpdateProduto(w http.ResponseWriter, r *http.Request) {

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
	produto := producao.Produto{}
	produto.ID = uint(uid)
	err = json.Unmarshal(body, &produto)
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
	produto.Prepare()
	err = produto.Validate("update")
	if err != nil {
		api.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	updatedProduto, err := produto.Save(database.DB)
	if err != nil {
		formattedError := api.FormatError(err.Error())
		api.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	api.JSON(w, http.StatusOK, updatedProduto)
}

func DeleteProduto(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	user := producao.Produto{}

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
	_, err = user.DeleteProduto(database.DB, uint32(uid))
	if err != nil {
		api.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", uid))
	api.JSON(w, http.StatusNoContent, "")
}
