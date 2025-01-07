package router

import (
	"github.com/gin-gonic/gin"
	"github.com/reis-jean/GooppotunitesAPI.git/handler"
  )

//   funções deve ser com letra maiuscula para serem exportadas
// se forem minusculas só podem ser acessadas dentro do mesmo pacote
func initializeRoutes(router *gin.Engine) {

	//iniciar  o handler

	handler.InitializeHandler()
	
	v1 := router.Group("/api/v1")
    {	
		//show opening
        v1.GET("/opening", handler.ShowOpeningHandler)
		//create opening
        v1.POST("/opening", handler.CreateOpeningHandler)
		//delete opening
        v1.DELETE("/opening", handler.DeleteOpeningHandler)
		//update opening
        v1.PUT("/opening", handler.UpdateOpeningHandler)
		//list openings
        v1.GET("/openings", handler.ListOpeningHandler)
    }


}
