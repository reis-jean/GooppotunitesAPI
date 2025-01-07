package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/reis-jean/GooppotunitesAPI.git/schemas"

)

func DeleteOpeningHandler(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		sendError(c, http.StatusBadRequest, "id is required")
		return
	}	

	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		sendError(c, http.StatusNotFound, "opening not found")
		return
	}

	if err := db.Delete(&schemas.Opening{}, id).Error; err != nil {
		sendError(c, http.StatusInternalServerError, "error deleting opening")
		return
	}

	sendSuccess(c, http.StatusOK, "opening deleted")
}