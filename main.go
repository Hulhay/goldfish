package main

import (
	"time"

	"github.com/Hulhay/goldfish/config"
	"github.com/Hulhay/goldfish/controller"
	"github.com/Hulhay/goldfish/middleware"
	"github.com/Hulhay/goldfish/repository"
	"github.com/Hulhay/goldfish/usecase"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db      *gorm.DB                  = config.SetupDatabaseConnection()
	repo    repository.UserRepository = repository.NewUserRepository(db)
	tokenUC usecase.Token             = usecase.NewTokenUc()
	authUC  usecase.Auth              = usecase.NewAuthUC(repo)
	ac      controller.AuthController = controller.NewAuthController(authUC, tokenUC)
)

func main() {

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/register", ac.Register)
		authRoutes.POST("/login", ac.Login)
		authRoutes.POST("/logout", middleware.AuthorizeJWT(tokenUC), ac.Logout)
	}

	r.Run()
}
