package routers

import (
	"go-user-auth/config"
	"go-user-auth/controllers"
	"go-user-auth/db/repository"
	_ "go-user-auth/docs"
	"go-user-auth/services"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func SetupRoutes(cfg config.AccountsConfig) *fiber.App {
	// create a new fiber app with the specified buffer size
	app := fiber.New(fiber.Config{
		ReadBufferSize: 16384,
		ErrorHandler:   fiber.DefaultErrorHandler,
		// ErrorHandler:   stderrors.Handler(), //TODO: (Setup our own error handler) stderror.Handler(),
	})

	// recover middleware recovers from panics anywhere in the stack without terminating the app.
	app.Use(recover.New())

	// logger middleware logs HTTP request details.
	app.Use(logger.New())

	// compress middleware compresses HTTP response using gzip, deflate or brotli.
	app.Use(compress.New())

	// Serve login and signup HTML pages
	app.Static("/", "./static")

	// setup health controllers
	healthSvc := services.NewHealthService(cfg)
	healthCtrl := controllers.NewHealthController(cfg, healthSvc)

	// setup order controllers
	orderSvc := services.NewOrderService(cfg)
	orderCtrl := controllers.NewOrderController(cfg, orderSvc)

	// setup property controllers
	propertySvc := services.NewPropertyService(cfg)
	propertyCtrl := controllers.NewPropertyController(cfg, propertySvc)

	// setup repository
	mongoRepo := repository.NewMongoRepository(cfg.MongoDBName)

	// setup user controllers
	userRepo := repository.NewUserRepository(cfg, mongoRepo)
	userLibrary := services.NewUserService(cfg, userRepo)
	userController := controllers.NewUserHandler(cfg, userLibrary)

	// setup authorization controllers
	authRepo := repository.NewAuthorizationRepository(cfg)
	authLib := services.NewAuthorizationService(cfg, authRepo)
	authCtrl := controllers.NewAuthenticationHandler(cfg, authLib)

	// setup application controllers
	appRepo := repository.NewApplicationRepository(cfg, mongoRepo)
	appLib := services.NewApplicationService(cfg, appRepo)
	appCtrl := controllers.NewApplicationsHandler(cfg, appLib)

	// home routes
	home := app.Group("/")
	home.Get("/", healthCtrl.Home)
	home.Get("/apphealth", healthCtrl.Apphealth) // Application health endpoint
	home.Get("/docs/*", swagger.HandlerDefault)  // Default Swagger UI endpoint

	// create an api group with middleware
	api := app.Group("/userauth")
	v1 := api.Group("/v1") // API version 1 group

	// TODO: add authentication middleware

	// Property routes
	v1.Post("/properties", propertyCtrl.CreateProperty)
	v1.Get("/properties", propertyCtrl.GetProperties)

	// Order routes
	v1.Post("/orders", orderCtrl.CreateOrder)
	v1.Get("/orders/:id", orderCtrl.GetOrder)

	// User routes
	// v1.Get("/users", userController.Get)
	v1.Get("/users", userController.GetAll)
	v1.Post("/users", userController.Add)
	v1.Put("/users/:userid", userController.Update)
	v1.Delete("/users/:userid", userController.Delete)
	app.Post("/user/login", userController.Login)
	app.Post("/user/reset", userController.ResetPassword)

	// Authorization routes
	v1.Post("/register-client", authCtrl.RegisterClient)
	v1.Get("/authorize", authCtrl.AuthorizeClient)
	v1.Get("/callback", authCtrl.Callback)
	v1.Post("/token", authCtrl.ExchangeToken)

	// v1.Post("/login")
	v1.Post("/signup", userController.SignUp)
	v1.Post("/login", userController.Login)
	v1.Get("/verify-token", userController.VerifyToken)
	v1.Get("/invite/:uuid", userController.Invite)

	// In v1 I need to group some api and apply JWTProtected middleware
	applications := v1.Group("/applications")
	// applications.Use(middlewares.JWTProtected())

	applications.Post("/create", appCtrl.RegisterApplication)
	applications.Get("/get/:tag", appCtrl.GetApplicationByTag)
	applications.Get("/get/:client_id", appCtrl.GetApplicationByClientID)
	applications.Get("/get/all", appCtrl.GetAllApplications)
	applications.Delete("/delete", appCtrl.DeleteApplication)

	return app
}

// Todo:
