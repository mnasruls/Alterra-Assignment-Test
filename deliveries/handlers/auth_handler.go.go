package handlers

import (
	"alterra/configs"
	"alterra/deliveries/middlewares"
	"alterra/entities"
	"alterra/entities/web"
	authService "alterra/services/auth"
	"net/http"
	"reflect"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authService *authService.AuthService
}

func NewAuthHandler(service *authService.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: service,
	}
}

/*
 * Auth Handler - Login
 * -------------------------------
 * Login user berdasarkan email dan password
 * dan mengembalikan response berupa token
 */
func (handler AuthHandler) Login(c echo.Context) error {
	// Populate request input
	authReq := entities.AuthRequest{
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}

	// define link hateoas
	links := map[string]string{"self": configs.Get().App.BaseURL + "/api/auth"}

	// call auth service login
	authRes, err := handler.authService.Login(authReq)
	if err != nil {

		// return error response khusus jika err termasuk webError
		if reflect.TypeOf(err).String() == "web.WebError" {
			webErr := err.(web.WebError)
			return c.JSON(webErr.Code, web.ErrorResponse{
				Status: "ERROR",
				Code:   webErr.Code,
				Error:  webErr.Error(),
				Links:  links,
			})
		}

		// return error 500 jika bukan webError
		return c.JSON(http.StatusInternalServerError, web.ErrorResponse{
			Status: "ERROR",
			Code:   http.StatusInternalServerError,
			Error:  err.Error(),
			Links:  links,
		})
	}

	// send response
	return c.JSON(200, web.SuccessResponse{
		Status: "OK",
		Code:   200,
		Error:  nil,
		Links:  links,
		Data:   authRes,
	})
}

/*
 * Auth Handler - Me
 * -------------------------------
 * Mengambil data user yang sedang login
 */
func (handler AuthHandler) Me(c echo.Context) error {

	// Token
	token := c.Get("user")
	ID, _, err := middlewares.ReadToken(token)

	// Define link
	links := map[string]string{"self": configs.Get().App.BaseURL + "/api/auth/me"}
	if err != nil {
		return c.JSON(http.StatusUnauthorized, web.ErrorResponse{
			Code:   http.StatusUnauthorized,
			Status: "ERROR",
			Error:  "unauthorized",
			Links:  links,
		})
	}

	// Memanggil service auth me
	authRes, err := handler.authService.Me(ID, token)
	if err != nil {

		// return error response khusus jika err termasuk webError
		if reflect.TypeOf(err).String() == "web.WebError" {
			webErr := err.(web.WebError)
			return c.JSON(webErr.Code, web.ErrorResponse{
				Status: "ERROR",
				Code:   webErr.Code,
				Error:  webErr.Error(),
				Links:  links,
			})
		}

		// return error 500 jika bukan webError
		return c.JSON(http.StatusInternalServerError, web.ErrorResponse{
			Status: "ERROR",
			Code:   http.StatusInternalServerError,
			Error:  err.Error(),
			Links:  links,
		})
	}

	// Response
	return c.JSON(200, web.SuccessResponse{
		Status: "OK",
		Code:   200,
		Error:  nil,
		Links:  links,
		Data:   authRes,
	})
}
