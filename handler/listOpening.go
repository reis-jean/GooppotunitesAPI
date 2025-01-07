package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/reis-jean/GooppotunitesAPI.git/schemas"

)

func ListOpeningHandler(c *gin.Context) {

	var openings []schemas.Opening

    if err := db.Find(&openings).Error; err != nil {
        sendError(c, http.StatusInternalServerError, "error listing openings")
        return
    }

    sendSuccess(c, http.StatusOK, openings)


}