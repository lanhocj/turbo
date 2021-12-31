package internal

import "github.com/gin-gonic/gin"

type NodeAvailableRequest struct {
	Addr string `json:"addr"`
	Port string `json:"port"`
	Tag  string `json:"tag"`
}

type NodeAvailableResponse struct {
	State int `json:"state"`
}

func (r *NodeAvailableResponse) Build(c *gin.Context) {
	c.SecureJSON(200, r)
	return
}

type NodeResponse struct {
	Addr   string `json:"addr"`
	Region string `json:"region"`
}

type NodeListResponse []*NodeResponse

type NodeAddRequest struct{}

type NodeAddResponse struct {
	Addr string `json:"addr"`
	Port string `json:"port"`
	Tag  string `json:"tag"`
}

func (r *NodeAddResponse) Build(c *gin.Context) {}
