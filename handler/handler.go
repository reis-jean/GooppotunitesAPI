package handler

import (
	"github.com/reis-jean/GooppotunitesAPI.git/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandler(){
	//quando iniciar o handler é inicial o logger e o cd globalmente
	logger = config.GetLogger("handler")
	db = config.GetDB()
}
