package web

import (
	"github.com/gin-gonic/gin"
	"net"
)

func Run(host, port string) {
	srv := gin.Default()

	addr := net.JoinHostPort(host, port)

	if err := srv.Run(addr); err != nil {
		panic(err)
	}
}
