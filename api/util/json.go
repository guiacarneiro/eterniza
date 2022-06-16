package util

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func ERROR(c *gin.Context, statusCode int, err error) {
	if err != nil {
		c.JSON(statusCode, struct {
			Mensagem string `json:"message"`
			Success  bool   `json:"success"`
		}{
			Mensagem: err.Error(),
			Success:  false,
		})
		return
	}
	c.JSON(http.StatusBadRequest, nil)
}
