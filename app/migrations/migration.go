package main

import (
	"gin-app/infra"
	"gin-app/models"
	"log"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()

	if err := db.AutoMigrate(&models.Item{}, &models.User{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}
