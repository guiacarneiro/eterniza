package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/guiacarneiro/eterniza/api/request"
	"github.com/guiacarneiro/eterniza/api/response"
	"github.com/guiacarneiro/eterniza/api/util"
	"github.com/guiacarneiro/eterniza/database/model"
	"github.com/guiacarneiro/eterniza/erro"
	"github.com/guiacarneiro/eterniza/service"
	"net/http"
)

func CreateUser(c *gin.Context) {
	req := request.CreateUser{}
	if err := c.BindJSON(&req); err != nil {
		util.ERROR(c, http.StatusBadRequest, err)
		return
	}
	_, err := service.GetUsuarioLogin(req.Email)
	if err == nil {
		util.ERROR(c, http.StatusUnprocessableEntity, erro.ParametrosIncorretos)
		return
	}
	user := model.User{
		Nickname: req.FisrtName + " " + req.LastName,
		Email:    req.Email,
		Password: req.Password,
	}

	user.Prepare()
	err = user.Validate("update")
	if err != nil {
		util.ERROR(c, http.StatusUnprocessableEntity, err)
		return
	}
	err = user.Save()
	if err != nil {
		formattedError := util.FormatError(err.Error())
		util.ERROR(c, http.StatusUnprocessableEntity, formattedError)
		return
	}
	c.JSON(http.StatusOK, response.Basic{
		Success: true,
	})
}

//
//func GetUsers(w http.ResponseWriter, r *http.Request) {
//
//	user := model.User{}
//
//	users, err := user.FindAllUsers(database.DB)
//	if err != nil {
//		util.ERROR(w, http.StatusInternalServerError, err)
//		return
//	}
//	util.JSON(w, http.StatusOK, users)
//}
//
//func GetUser(w http.ResponseWriter, r *http.Request) {
//
//	vars := mux.Vars(r)
//	uid, err := strconv.ParseUint(vars["id"], 10, 32)
//	if err != nil {
//		util.ERROR(w, http.StatusBadRequest, err)
//		return
//	}
//	user := model.User{}
//	userGotten, err := user.FindUserByID(database.DB, uint32(uid))
//	if err != nil {
//		util.ERROR(w, http.StatusBadRequest, err)
//		return
//	}
//	util.JSON(w, http.StatusOK, userGotten)
//}
//
//func UpdateUser(w http.ResponseWriter, r *http.Request) {
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
//	user := model.User{}
//	err = json.Unmarshal(body, &user)
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
//	user.Prepare()
//	err = user.Validate("update")
//	if err != nil {
//		util.ERROR(w, http.StatusUnprocessableEntity, err)
//		return
//	}
//	updatedUser, err := user.UpdateAUser(database.DB, uint32(uid))
//	if err != nil {
//		formattedError := util.FormatError(err.Error())
//		util.ERROR(w, http.StatusInternalServerError, formattedError)
//		return
//	}
//	util.JSON(w, http.StatusOK, updatedUser)
//}
//
//func DeleteUser(w http.ResponseWriter, r *http.Request) {
//
//	vars := mux.Vars(r)
//
//	user := model.User{}
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
//	_, err = user.DeleteAUser(database.DB, uint32(uid))
//	if err != nil {
//		util.ERROR(w, http.StatusInternalServerError, err)
//		return
//	}
//	w.Header().Set("Entity", fmt.Sprintf("%d", uid))
//	util.JSON(w, http.StatusNoContent, "")
//}
