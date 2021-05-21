package controller

import (
	"eterniza/api"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	api.JSON(w, http.StatusOK, "Welcome To This Awesome API")
}
