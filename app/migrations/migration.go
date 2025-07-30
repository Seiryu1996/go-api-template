package main

import (
	"gin-app/infra"
	"gin-app/models"
	"log"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()
	if err := db.AutoMigrate(&models.User{}, &models.Item{}); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	log.Println("Migration completed successfully")
}
