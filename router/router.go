package router

import (
	"net/http"
	"github.com/gin-gonic/gin"
  )


  //Initialize é uma função que inicia o servidor
  //funcões deve ser com letra maiuscula para serem exportadas
func Initialize() {
	//r seria router
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
	  c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	  })
	})
	//esta rodadno na porta 9090 (por padrão roda na 8080)
	r.Run(":9090")
}