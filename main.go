package main

import (
	"GolangProject/config"
	"GolangProject/controllers"
	"GolangProject/controllers/Admin"
	"GolangProject/middleware"
	"GolangProject/repositories"
	"GolangProject/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
	"os"
)

var db *gorm.DB = config.SetupDatabaseConnection()
var userRepository repositories.UserRepository = repositories.NewUserRepository(db)
var productRepositiry repositories.ProductRepository = repositories.NewProductRepository(db)
var orderRepository repositories.OrderRepository = repositories.NewOrderRepository(db)
var jwtService services.JWTService = services.NewJWTService()
var userService services.UserService = services.NewUserService(userRepository)
var authService services.AuthService = services.NewAuthService(userRepository)
var orderService services.OrderService = services.NewOrderService(orderRepository)
var productService services.ProductService = services.NewProductService(productRepositiry)
var authController controllers.AuthController = controllers.NewAuthController(authService, jwtService)
var userController controllers.UserController = controllers.NewUserController(userService, jwtService)
var productController controllers.ProductController = controllers.NewProductController(productService, jwtService)
var notificationController Admin.NotificationController = Admin.NewNotificationController(userService, jwtService)
var orderController controllers.OrderController = controllers.NewOrderController(orderService, jwtService)

func main() {
	defer config.CloseDatabaseConnection(db)
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	//routes.AuthRoutes(router)

	authRoutes := router.Group("api/")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := router.Group("api/user", middleware.Authenticate(jwtService))
	{
		userRoutes.GET("/", userController.Profile)
	}

	productRoutes := router.Group("api/products", middleware.Authenticate(jwtService))
	{
		productRoutes.GET("/", productController.All)
		productRoutes.GET("/:id", productController.Show)
		productRoutes.POST("/", productController.Insert)
		productRoutes.PUT("/:id", productController.Update)
		productRoutes.DELETE("/:id", productController.Delete)
	}

	adminRoutes := router.Group("admin/")
	{
		adminRoutes.POST("/notification", notificationController.Send)
	}

	orderRoutes := router.Group("api/orders", middleware.Authenticate(jwtService))
	{
		orderRoutes.GET("/", orderController.All)
		orderRoutes.POST("/", orderController.Insert)
		orderRoutes.GET("/:id", orderController.Show)
		orderRoutes.PUT("/:id", orderController.Update)
		orderRoutes.DELETE("/:id", orderController.Delete)
	}

	router.Run(":" + port)
}
