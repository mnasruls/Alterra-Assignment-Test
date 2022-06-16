package main

import (
	"alterra/configs"
	"alterra/deliveries/handlers"
	"alterra/deliveries/routes"
	userRepository "alterra/repositories/user"
	userService "alterra/services/user"
	"alterra/utils"

	"github.com/labstack/echo/v4"
)

func main() {
	config := configs.Get()
	db := utils.NewMysqlGorm(config)
	utils.Migrate(db)

	e := echo.New()

	// User
	userRepository := userRepository.NewUserRepository(db)
	userService := userService.NewUserService(userRepository)
	adminHandler := handlers.NewAdminHandler(userService)
	routes.RegisterAdminRoute(e, adminHandler)
	e.Logger.Fatal(e.Start(":" + config.App.Port))
}
