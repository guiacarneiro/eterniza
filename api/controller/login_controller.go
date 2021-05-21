package controller

import (
	"encoding/json"
	"eterniza/api"
	"eterniza/database"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
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

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		api.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := SignIn(user.Email, user.Password)
	if err != nil {
		formattedError := api.FormatError(err.Error())
		api.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	api.JSON(w, http.StatusOK, token)
}

func SignIn(email, password string) (string, error) {

	var err error

	user := api.User{}

	err = database.DB.Debug().Model(api.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", err
	}
	err = api.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return api.CreateToken(user.ID)
}
