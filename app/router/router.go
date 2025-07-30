package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"gin-app/controllers"
	"gin-app/middlewares"
	"gin-app/repositories"
	"gin-app/services"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	// item
	itemRepository := repositories.NewItemRepository(db)
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemService)
	// auth
	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authController := controllers.NewAuthController(authService)

	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, World!!!!!"})
	})

	itemRouter := r.Group("/items")
	itemRputerWithAuth := r.Group("/items", middlewares.AuthMiddleware(authService))
	authRouter := r.Group("/auth")

	itemRouter.GET("", itemController.FindAll)
	itemRputerWithAuth.GET("/:id", itemController.FindById)
	itemRputerWithAuth.POST("", itemController.Create)
	itemRputerWithAuth.PUT("/:id", itemController.Update)
	itemRputerWithAuth.DELETE("/:id", itemController.Delete)

	authRouter.POST("/signup", authController.Signup)
	authRouter.POST("/login", authController.Login)
	return r
}
