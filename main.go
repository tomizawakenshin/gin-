package main

import (
	"gin-fleamarket/controller"
	"gin-fleamarket/infra"
	"gin-fleamarket/middlewares"

	// "gin-fleamarket/models"
	"gin-fleamarket/reposotories"
	"gin-fleamarket/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupRouter(db *gorm.DB) *gin.Engine {
	itemRepository := reposotories.NewItemMemoryRepository(db)
	itemService := services.NewItemService(itemRepository)
	itemController := controller.NewItemController(itemService)

	authRepository := reposotories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authController := controller.NewAuthController(authService)

	hanabiRepository := reposotories.NewHanabiRepository(db)
	hanabiService := services.NewHanabiService(hanabiRepository)
	hanabiController := controller.NewHanabiController(hanabiService)

	r := gin.Default()
	r.Use(cors.Default())

	//itemのエンドポイント
	itemRouter := r.Group("/items")
	itemRouterWithAuth := r.Group("/items", middlewares.AuthMiddleware(authService))

	itemRouter.GET("", itemController.FindAll)
	itemRouterWithAuth.GET("/:id", itemController.FindById)
	itemRouterWithAuth.POST("", itemController.Create)
	itemRouter.POST("/restore/:id", itemController.Restore)
	itemRouterWithAuth.PUT("/:id", itemController.Update)
	itemRouterWithAuth.DELETE("/:id", itemController.Delete)

	//hanabiのエンドポイント
	//hanabiRouter := r.Group("/hanabi")
	hanabiRouterWithAuth := r.Group("/hanabi", middlewares.AuthMiddleware(authService))
	hanabiRouterWithAuth.POST("/create", hanabiController.Create)
	hanabiRouterWithAuth.GET("/getAll", hanabiController.FindAll)
	hanabiRouterWithAuth.GET("/getByID/:id", hanabiController.FindByID)

	//user認証関連のエンドポイント
	authRouter := r.Group("/auth")
	authRouter.POST("/signup", authController.SignUp)
	authRouter.POST("/login", authController.Login)

	return r
}
func main() {
	infra.Initialize()
	db := infra.SetupDB()

	r := setupRouter(db)
	r.Run("localhost:8080") // 0.0.0.0:8080 でサーバーを立てます。
}
