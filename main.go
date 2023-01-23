package main

import (
	"time"

	"github.com/Hulhay/goldfish/config"
	"github.com/Hulhay/goldfish/controller"
	"github.com/Hulhay/goldfish/middleware"
	"github.com/Hulhay/goldfish/model"
	"github.com/Hulhay/goldfish/repository"
	"github.com/Hulhay/goldfish/usecase"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()

	userRepo        repository.UserRepository        = repository.NewUserRepository(db)
	memberRepo      repository.MemberRepository      = repository.NewMemberRepository(db)
	familyRepo      repository.FamilyRepository      = repository.NewFamilyRepository(db)
	categoryRepo    repository.CategoryRepository    = repository.NewCategoryRepository(db)
	transactionRepo repository.TransactionRepository = repository.NewTransactionRepository(db)
	walletRepo      repository.WalletRepository      = repository.NewWalletRepository(db)
	timeRepo        repository.TimeRepository        = repository.NewTimeRepository()

	tokenUC       usecase.Token       = usecase.NewTokenUc()
	authUC        usecase.Auth        = usecase.NewAuthUC(userRepo)
	familyUC      usecase.Family      = usecase.NewFamilyUC(familyRepo)
	memberUC      usecase.Member      = usecase.NewMemberUC(memberRepo, familyUC)
	controllerUC  usecase.Category    = usecase.NewCategoryUC(categoryRepo)
	transactionUC usecase.Transaction = usecase.NewTransactionUC(transactionRepo, familyRepo, walletRepo, categoryRepo, timeRepo)
	profileUC     usecase.Profile     = usecase.NewProfileUC(userRepo)

	ac controller.AuthController        = controller.NewAuthController(authUC, tokenUC)
	hc controller.HealthController      = controller.NewHealthController()
	mc controller.MemberContoller       = controller.NewMemberController(memberUC)
	cc controller.CategoryController    = controller.NewCategoryController(controllerUC)
	tc controller.TransactionController = controller.NewTransactionRepository(transactionUC)
	pc controller.ProfileController     = controller.NewProfileController(profileUC)
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
		authRoutes.POST("/logout", middleware.AuthorizeJWT(tokenUC, []string{}), ac.Logout)
		authRoutes.POST("/change-password", middleware.AuthorizeJWT(tokenUC, []string{}), ac.ChangePassword)
	}

	memberRoutes := r.Group("api/member")
	{
		memberRoutes.POST("", middleware.AuthorizeJWT(tokenUC, []string{}), mc.InsertMember)
		memberRoutes.GET("/list", middleware.AuthorizeJWT(tokenUC, []string{}), mc.GetMember)
		memberRoutes.GET("/detail", middleware.AuthorizeJWT(tokenUC, []string{}), mc.GetDetailMember)
		memberRoutes.PATCH("/:member-id", middleware.AuthorizeJWT(tokenUC, []string{}), mc.EditMember)
	}

	categoryRoutes := r.Group("api/category")
	{
		categoryRoutes.POST("", middleware.AuthorizeJWT(tokenUC, []string{model.SUPER_USER}), cc.InsertCategory)
		categoryRoutes.GET("/list", middleware.AuthorizeJWT(tokenUC, []string{}), cc.GetListCategory)
	}

	transactionRoutes := r.Group("api/transaction")
	{
		transactionRoutes.POST("/new", middleware.AuthorizeJWT(tokenUC, []string{}), tc.CreateTransaction)
		transactionRoutes.GET("/history", middleware.AuthorizeJWT(tokenUC, []string{}), tc.GetHistoryTransaction)
	}

	profileRoutes := r.Group("api/profile")
	{
		profileRoutes.GET("/me", middleware.AuthorizeJWT(tokenUC, []string{}), pc.GetProfile)
	}

	r.Run()
}
