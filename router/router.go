package router

//Initialize é uma função que inicia o servidor
//funcões deve ser com letra maiuscula para serem exportadas

//tudo que esta no mesmo pacote pode
//ser acessado por qualquer arquivo do mesmo pacote

//tudo que esta em um pacote diferente precisa ser
//exportado para ser acessado

import (
	"github.com/gin-gonic/gin"
)


func Initialize() {

	//inicializa as rotas
	router  := gin.Default()

	InitializeRoutes(router)
	
	// //esta rodadno na porta 9090 (por padrão roda na 8080)
	// //run the server
	router.Run(":9090")

}