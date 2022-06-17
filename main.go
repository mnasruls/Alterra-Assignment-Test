package main

import (
	"alterra/configs"
	"alterra/deliveries/handlers"
	"alterra/deliveries/routes"
	_userRepository "alterra/repositories/user"
	_authService "alterra/services/auth"
	_userService "alterra/services/user"
	"alterra/utils"

	"github.com/labstack/echo/v4"
)

func main() {
	config := configs.Get()
	db := utils.NewMysqlGorm(config)
	utils.Migrate(db)

	e := echo.New()

	// User
	userRepository := _userRepository.NewUserRepository(db)
	userService := _userService.NewUserService(userRepository)

	//admin
	adminHandler := handlers.NewAdminHandler(userService)
	routes.RegisterAdminRoute(e, adminHandler)

	//Authentication
	authService := _authService.NewAuthService(userRepository)
	authHancdler := handlers.NewAuthHandler(authService)
	routes.RegisterAuthRoute(e, authHancdler)

	e.Logger.Fatal(e.Start(":" + config.App.Port))
}
