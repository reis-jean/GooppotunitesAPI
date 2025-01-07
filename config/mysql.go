package config

import (
	"os"

	"github.com/reis-jean/GooppotunitesAPI.git/schemas"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeMySQL()(*gorm.DB, error){

	logger := GetLogger("mysql")
	dbPath := "./db/openings.db"

	//check id the database file exist
	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		logger.Info("Database file does not exist, creating new file")

		//create the database file
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			logger.Errorf("Error creating database file: %v", err)
			return nil, err
		}

		file, err := os.Create(dbPath)
		if err != nil {
			logger.Errorf("Error creating database file: %v", err)
			return nil, err
		}
		file.Close()

	} 

	dsn := "root:root@tcp(127.0.0.1:3306)/openings?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        logger.Errorf("Error initializing mysql: %v", err)
        return nil, err
    }

    // Migrate the schema
    err = db.AutoMigrate(&schemas.Opening{})
    if err != nil {
        logger.Errorf("Error migrating schema: %v", err)
        return nil, err
    }

    return db, nil
}