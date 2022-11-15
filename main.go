package main

import (
	"net/http"
	"yt-users-service/app"
	"yt-users-service/controller"
	"yt-users-service/exception"
	"yt-users-service/helper"
	"yt-users-service/middleware"
	"yt-users-service/repository"
	"yt-users-service/service"

	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func main() {
	// load environtment
	envErr := godotenv.Load(".env")
	helper.PanicError(envErr)

	// connection to database
	db := app.Database()

	// validation
	validate := validator.New()

	// repository
	userRepository := repository.NewUserRepositoryImpl()

	// service
	userService := service.NewUserServiceImpl(
		userRepository,
		db,
		validate,
	)

	// controller
	userController := controller.NewUserControllerImpl(userService)

	// initialize
	router := httprouter.New()

	// router
	// [USER]
	router.POST("/api/v1/user", userController.Create)
	router.POST("/api/v1/auth", userController.Auth)
	router.POST("/api/v1/refresh-token", userController.CreateWithRefreshToken)
	router.PUT("/api/v1/user/:user_id", userController.Update)
	router.DELETE("/api/v1/user/:user_id", userController.Delete)
	router.GET("/api/v1/user/:user_id", userController.FindById)
	router.GET("/api/v1/user", userController.FindAll)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicError(err)
}
