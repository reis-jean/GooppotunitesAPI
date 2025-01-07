package handler

import (
	"github.com/gin-gonic/gin"
)

func sendError(c *gin.Context, code int, msg string){
	c.Header("Content-Type", "application/json")
	c.JSON(code, gin.H{
		"error": msg, 
		"errorCode": code, 
	})
}


func sendSuccess(c *gin.Context, code int, data interface{}){
	c.Header("Content-Type", "application/json")
	c.JSON(code, gin.H{
		"data": data,
	})
}