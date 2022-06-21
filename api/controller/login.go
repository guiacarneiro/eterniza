package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/guiacarneiro/eterniza/api/request"
	"github.com/guiacarneiro/eterniza/api/response"
	"github.com/guiacarneiro/eterniza/api/util"
	"github.com/guiacarneiro/eterniza/database/model"
	"github.com/guiacarneiro/eterniza/service"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Login(c *gin.Context) {
	req := request.Login{}
	if err := c.BindJSON(&req); err != nil {
		util.ERROR(c, http.StatusBadRequest, err)
		return
	}
	user := model.User{
		Email:    req.Email,
		Password: req.Password,
	}

	user.Prepare()
	err := user.Validate("login")
	if err != nil {
		util.ERROR(c, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := SignIn(user.Email, user.Password)
	if err != nil {
		formattedError := util.FormatError(err.Error())
		util.ERROR(c, http.StatusUnprocessableEntity, formattedError)
		return
	}
	c.JSON(http.StatusOK, response.Login{
		Basic: response.Basic{
			Success: true,
		},
		Token: token,
	})
}

func SignIn(email, password string) (string, error) {
	user, err := service.GetUsuarioLogin(email)
	if err != nil {
		return "", err
	}
	err = model.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return util.CreateToken(user.ID, user.Nickname)
}
