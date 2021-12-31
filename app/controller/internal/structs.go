package internal

import "github.com/gin-gonic/gin"

type Response interface {
	Build(c *gin.Context)
}

type Request interface{}

var _ Response = &NodeAvailableResponse{}
var _ Request = &NodeAvailableRequest{}
