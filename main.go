package main

import (
	_ "github.com/GoAdminGroup/go-admin/adapter/gin"
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/examples/datamodel"
	"github.com/GoAdminGroup/go-admin/modules/config"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/GoAdminGroup/themes/adminlte"
	"github.com/gin-gonic/gin"
	"time"
)

//var db *gorm.DB = config.SetupDatabaseConnection()
//var userRepository repositories.UserRepository = repositories.NewUserRepository(db)
//var productRepositiry repositories.ProductRepository = repositories.NewProductRepository(db)
//var orderRepository repositories.OrderRepository = repositories.NewOrderRepository(db)
//var jwtService services.JWTService = services.NewJWTService()
//var userService services.UserService = services.NewUserService(userRepository)
//var authService services.AuthService = services.NewAuthService(userRepository)
//var orderService services.OrderService = services.NewOrderService(orderRepository)
//var productService services.ProductService = services.NewProductService(productRepositiry)
//var authController controllers.AuthController = controllers.NewAuthController(authService, jwtService)
//var userController controllers.UserController = controllers.NewUserController(userService, jwtService)
//var productController controllers.ProductController = controllers.NewProductController(productService, jwtService)
//var notificationController Admin.NotificationController = Admin.NewNotificationController(userService, jwtService)
//var orderController controllers.OrderController = controllers.NewOrderController(orderService, jwtService)

//func main() {
//	defer config.CloseDatabaseConnection(db)
//	err := godotenv.Load(".env")
//
//	if err != nil {
//		log.Fatal("Error loading .env file")
//	}
//	port := os.Getenv("PORT")
//
//	if port == "" {
//		port = "8000"
//	}
//
//	router := gin.New()
//	router.Use(gin.Logger())
//
//	//routes.AuthRoutes(router)
//
//	authRoutes := router.Group("api/")
//	{
//		authRoutes.POST("/login", authController.Login)
//		authRoutes.POST("/register", authController.Register)
//		authRoutes.POST("/support", authController.Support)
//	}
//
//	userRoutes := router.Group("api/user", middleware.Authenticate(jwtService))
//	{
//		userRoutes.GET("/", userController.Profile)
//	}
//
//	productRoutes := router.Group("api/products", middleware.Authenticate(jwtService))
//	{
//		productRoutes.GET("/", productController.All)
//		productRoutes.GET("/:id", productController.Show)
//		productRoutes.POST("/", productController.Insert)
//		productRoutes.PUT("/:id", productController.Update)
//		productRoutes.DELETE("/:id", productController.Delete)
//	}
//
//	adminRoutes := router.Group("admin/")
//	{
//		adminRoutes.POST("/notification", notificationController.Send)
//	}
//
//	orderRoutes := router.Group("api/orders", middleware.Authenticate(jwtService))
//	{
//		orderRoutes.GET("/", orderController.All)
//		orderRoutes.POST("/", orderController.Insert)
//		orderRoutes.GET("/:id", orderController.Show)
//		orderRoutes.PUT("/:id", orderController.Update)
//		orderRoutes.DELETE("/:id", orderController.Delete)
//	}
//
//	router.Run(":" + port)
//}

func main() {
	r := gin.Default()

	eng := engine.Default()

	// global config
	cfg := config.Config{
		Databases: config.DatabaseList{
			"default": {
				Host:            "127.0.0.1",
				Port:            "3306",
				User:            "root",
				Pwd:             "",
				Name:            "golang",
				MaxIdleConns:    50,
				MaxOpenConns:    150,
				ConnMaxLifetime: time.Hour,
				Driver:          "mysql",
			},
		},
		UrlPrefix: "admin",
		// STORE is important. And the directory should has permission to write.
		Store: config.Store{
			Path:   "./uploads",
			Prefix: "uploads",
		},
		Language: language.EN,
		// debug mode
		Debug: true,
		// log file absolute path
		InfoLogPath:   "/var/logs/info.log",
		AccessLogPath: "/var/logs/access.log",
		ErrorLogPath:  "/var/logs/error.log",
		ColorScheme:   adminlte.ColorschemeSkinBlack,
	}

	// add component chartjs
	template.AddComp(chartjs.NewChart())

	_ = eng.AddConfig(&cfg).
		AddGenerators(datamodel.Generators).
		// add generator, first parameter is the url prefix of table when visit.
		// example:
		//
		// "user" => http://localhost:9033/admin/info/user
		//
		AddGenerator("user", datamodel.GetUserTable).
		Use(r)

	// customize your pages
	eng.HTML("GET", "/admin", datamodel.GetContent)

	_ = r.Run(":9033")
}
