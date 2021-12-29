package structs

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ErrResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"msg"`
}

func (r *ErrResponse) Build(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusOK, r)
	return
}

func Error(c *gin.Context, code int, data interface{}) {
	err := &ErrResponse{
		Code: code,
		Data: data,
	}

	err.Build(c)
}
