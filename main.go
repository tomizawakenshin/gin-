package main

import (
	"gin-fleamarket/controller"
	"gin-fleamarket/infra"
	"gin-fleamarket/middlewares"
	"time"

	"gin-fleamarket/reposotories"
	"gin-fleamarket/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupRouter(db *gorm.DB) *gin.Engine {

	authRepository := reposotories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authController := controller.NewAuthController(authService)

	hanabiRepository := reposotories.NewHanabiRepository(db)
	hanabiService := services.NewHanabiService(hanabiRepository)
	hanabiController := controller.NewHanabiController(hanabiService)

	commentRepository := reposotories.NewCommentMemoryRepository(db)
	commentService := services.NewCommentService(commentRepository, hanabiRepository)
	commentController := controller.NewCommentController(commentService)

	likeRepository := reposotories.NewLikeRepository(db)
	likeService := services.NewLikeService(likeRepository)
	likeController := controller.NewLikeController(likeService)

	r := gin.Default()
	// r.Use(cors.Default())
	// CORS 設定
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://your-frontend.vercel.app"}, // フロントエンドのドメインを許可
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},                   // 許可するHTTPメソッド
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},                   // 許可するリクエストヘッダー
		ExposeHeaders:    []string{"Content-Length"},                                            // クライアントに公開するレスポンスヘッダー
		AllowCredentials: true,                                                                  // 認証情報（クッキーなど）の送信を許可
		MaxAge:           12 * time.Hour,                                                        // プリフライトリクエストのキャッシュ時間
	}))

	//hanabiのエンドポイント
	//hanabiRouter := r.Group("/hanabi")
	hanabiRouterWithAuth := r.Group("/hanabi", middlewares.AuthMiddleware(authService))
	hanabiRouterWithAuth.POST("/create", hanabiController.Create)
	hanabiRouterWithAuth.GET("/getAll", hanabiController.FindAll)
	hanabiRouterWithAuth.GET("/getByID/:id", hanabiController.FindByID)

	//commentのエンドポイント
	commentRouterWithAuth := r.Group("/comment", middlewares.AuthMiddleware(authService))
	commentRouterWithAuth.POST("/create/:hanabiId", commentController.Create)

	//likeのエンドポイント
	likeRouterWithAuth := r.Group("/like", middlewares.AuthMiddleware(authService))
	likeRouterWithAuth.POST("/like/:commentId", likeController.Like)
	likeRouterWithAuth.DELETE("/unlike/:commentId", likeController.Unlike)

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
