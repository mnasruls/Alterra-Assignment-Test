package routes

import (
	"alterra/deliveries/handlers"
	"alterra/deliveries/middlewares"

	"github.com/labstack/echo/v4"
)

func RegisterAdminRoute(e *echo.Echo, userHandler *handlers.AdminHandler) {
	group := e.Group("/api/admins", middlewares.JWTMiddleware())
	group.POST("", userHandler.Create)       // Registration
	group.PUT("/:id", userHandler.Update)    //Edit
	group.DELETE("/:id", userHandler.Delete) //Delete
}

func RegisterAuthRoute(e *echo.Echo, authHandler *handlers.AuthHandler) {
	e.POST("api/auth", authHandler.Login)                             // Login
	e.GET("api/auth/me", authHandler.Me, middlewares.JWTMiddleware()) //Mendapatkan data profile
}
