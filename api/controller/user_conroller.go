package controller

import (
	"encoding/json"
	"errors"
	"eterniza/api"
	"eterniza/database"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		api.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	user := api.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		api.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user.Prepare()
	err = user.Validate("")
	if err != nil {
		api.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	userCreated, err := user.SaveUser(database.DB)

	if err != nil {

		formattedError := api.FormatError(err.Error())

		api.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, userCreated.ID))
	api.JSON(w, http.StatusCreated, userCreated)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {

	user := api.User{}

	users, err := user.FindAllUsers(database.DB)
	if err != nil {
		api.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	api.JSON(w, http.StatusOK, users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		api.ERROR(w, http.StatusBadRequest, err)
		return
	}
	user := api.User{}
	userGotten, err := user.FindUserByID(database.DB, uint32(uid))
	if err != nil {
		api.ERROR(w, http.StatusBadRequest, err)
		return
	}
	api.JSON(w, http.StatusOK, userGotten)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

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
	user := api.User{}
	err = json.Unmarshal(body, &user)
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
	user.Prepare()
	err = user.Validate("update")
	if err != nil {
		api.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	updatedUser, err := user.UpdateAUser(database.DB, uint32(uid))
	if err != nil {
		formattedError := api.FormatError(err.Error())
		api.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	api.JSON(w, http.StatusOK, updatedUser)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	user := api.User{}

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
	_, err = user.DeleteAUser(database.DB, uint32(uid))
	if err != nil {
		api.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", uid))
	api.JSON(w, http.StatusNoContent, "")
}
