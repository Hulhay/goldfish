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
	db *gorm.DB = config.SetupDatabaseConnection()

	userRepo   repository.UserRepository   = repository.NewUserRepository(db)
	memberRepo repository.MemberRepository = repository.NewMemberRepository(db)
	familyRepo repository.FamilyRepository = repository.NewFamilyRepository(db)

	tokenUC  usecase.Token  = usecase.NewTokenUc()
	authUC   usecase.Auth   = usecase.NewAuthUC(userRepo)
	familyUC usecase.Family = usecase.NewFamilyUC(familyRepo)
	memberUC usecase.Member = usecase.NewMemberUC(memberRepo, familyUC)

	ac controller.AuthController   = controller.NewAuthController(authUC, tokenUC)
	hc controller.HealthController = controller.NewHealthController()
	mc controller.MemberContoller  = controller.NewMemberController(memberUC)
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

	healthRoutes := r.Group("api/health")
	{
		healthRoutes.GET("", hc.Health)
	}

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/register", ac.Register)
		authRoutes.POST("/login", ac.Login)
		authRoutes.POST("/logout", middleware.AuthorizeJWT(tokenUC), ac.Logout)
	}

	memberRoutes := r.Group("api/member")
	{
		memberRoutes.POST("", middleware.AuthorizeJWT(tokenUC), mc.InsertMember)
		memberRoutes.GET("/list", middleware.AuthorizeJWT(tokenUC), mc.GetMember)
		memberRoutes.GET("/detail", middleware.AuthorizeJWT(tokenUC), mc.GetDetailMember)
		memberRoutes.PATCH("/:member-id", middleware.AuthorizeJWT(tokenUC), mc.EditMember)
	}

	r.Run()
}
