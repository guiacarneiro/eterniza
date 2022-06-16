package util

import (
	"fmt"
	"github.com/guiacarneiro/eterniza/config"
	"github.com/guiacarneiro/eterniza/erro"
	"github.com/guiacarneiro/eterniza/logger"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	jwtgo "github.com/dgrijalva/jwt-go"
)

type TokenData struct {
	UserID uint   `json:"userId"`
	Login  string `json:"login"`
}

func CreateToken(userId uint, userName string) (string, error) {
	claims := jwtgo.MapClaims{}
	claims["authorized"] = true
	claims["uid"] = userId
	claims["unm"] = userName
	claims["iat"] = logger.NowSiga().Unix()
	claims["exp"] = logger.NowSiga().Add(time.Hour * 8).Unix() // Token expires after 1 hour
	token := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.GetPropriedade("api_secret")))
}

type AuthHeader struct {
	Authorization string `header:"Authorization"`
}

func ExtractToken(c *gin.Context) string {
	h := AuthHeader{}
	err := c.BindHeader(&h)
	if err != nil {
		return ""
	}
	bearerToken := h.Authorization
	bearerToken = strings.TrimSpace(bearerToken)
	if len(strings.Split(bearerToken, " ")) >= 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func ExtractTokenData(c *gin.Context) (*TokenData, error) {
	tokenString := ExtractToken(c)
	token, err := jwtgo.Parse(tokenString, func(token *jwtgo.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwtgo.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.GetPropriedade("api_secret")), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwtgo.MapClaims)
	if ok && token.Valid {
		if ok && token.Valid {
			uid := claims["uid"]
			unm := claims["unm"]
			return &TokenData{
				UserID: uid.(uint),
				Login:  unm.(string),
			}, nil
		}
	}
	return nil, erro.SemPermissao
}
