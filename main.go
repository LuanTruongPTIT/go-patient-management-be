package main

import (
	"net/http"

	"github.com/LuanTruongPTIT/patient/pkg/server"
	"github.com/gin-gonic/gin"
)

var svr *server.Server

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
