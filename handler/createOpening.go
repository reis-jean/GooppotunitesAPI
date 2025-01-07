package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/reis-jean/GooppotunitesAPI.git/schemas"

)

func CreateOpeningHandler(c *gin.Context) {

	var request CreateOpeningRequest

	c.BindJSON(&request) 

	if err := request.validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(c, http.StatusBadRequest, err.Error())
		return
		// se o handler terminar antes  do fim da função automaticamente é retorndado o erro status 400
	}

	opening := schemas.Opening{
		Role: request.Role,
		Company: request.Company,
		Remote: *request.Remote,
		Salary: request.Salary,
	}


 
	if err  := db.Create(&opening).Error; err != nil{
		logger.Errorf("error creating opening: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "error creating opening")
		return
	}

	sendSuccess(c, http.StatusCreated, opening)

}