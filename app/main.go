package main

import (
	"gin-app/infra"

	"gin-app/router"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()
	r := router.SetupRouter(db)
	r.Run(":8080")
}
