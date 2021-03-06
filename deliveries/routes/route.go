package routes

import (
	"alterra/deliveries/handlers"
	"alterra/deliveries/middlewares"

	"github.com/labstack/echo/v4"
)

func RegisterAdminRoute(e *echo.Echo, adminHandler *handlers.AdminHandler) {
	group := e.Group("/api/admins", middlewares.JWTMiddleware())
	group.POST("", adminHandler.Create)       // Registration
	group.PUT("/:id", adminHandler.Update)    //Edit
	group.DELETE("/:id", adminHandler.Delete) //Delete
}

func RegisterAuthRoute(e *echo.Echo, authHandler *handlers.AuthHandler) {
	e.POST("api/auth", authHandler.Login)                             // Login
	e.GET("api/auth/me", authHandler.Me, middlewares.JWTMiddleware()) //Mendapatkan data profile
}

func RegisterUserRoute(e *echo.Echo, userHandler *handlers.UserHandler) {
	group := e.Group("/api/users", middlewares.JWTMiddleware())
	group.GET("", userHandler.GetAllUser)
	group.GET("/:id", userHandler.GetSingleUser)
}
