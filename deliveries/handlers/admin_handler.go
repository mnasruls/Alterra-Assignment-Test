package handlers

import (
	"alterra/configs"
	"alterra/deliveries/helpers"
	"alterra/deliveries/middlewares"
	"alterra/entities"
	"alterra/entities/web"
	userService "alterra/services/user"
	"net/http"
	"reflect"
	"strconv"

	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	userService *userService.UserService
}

func NewAdminHandler(service *userService.UserService) *AdminHandler {
	return &AdminHandler{
		userService: service,
	}
}

/*
 * User Handler - Create
 * -------------------------------
 * Registrasi User kedalam sistem dan
 * mengembalikan token
 */
func (handler AdminHandler) Create(c echo.Context) error {

	// Bind request ke user request
	userReq := entities.CreateUserRequest{}
	c.Bind(&userReq)

	// Define links (hateoas)
	links := map[string]string{"self": configs.Get().App.BaseURL + "/api/admins"}

	// registrasi user via call user service
	userRes, err := handler.userService.CreateUser(userReq)
	if err != nil {

		// return error response khusus jika err termasuk webError / ValidationError
		if reflect.TypeOf(err).String() == "web.WebError" {
			webErr := err.(web.WebError)
			return c.JSON(webErr.Code, web.ErrorResponse{
				Status: "ERROR",
				Code:   webErr.Code,
				Error:  webErr.Error(),
				Links:  links,
			})
		} else if reflect.TypeOf(err).String() == "web.ValidationError" {
			valErr := err.(web.ValidationError)
			return c.JSON(valErr.Code, web.ValidationErrorResponse{
				Status: "ERROR",
				Code:   valErr.Code,
				Error:  valErr.Error(),
				Errors: valErr.Errors,
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

	// response
	return c.JSON(http.StatusCreated, web.SuccessResponse{
		Status: "OK",
		Code:   http.StatusCreated,
		Error:  nil,
		Links:  links,
		Data:   userRes,
	})
}

func (handler AdminHandler) Update(c echo.Context) error {

	// Bind request to user request
	userReq := entities.UpdateUserRequest{}
	c.Bind(&userReq)

	// Get token
	id, tx := strconv.Atoi(c.Param("id"))
	token := c.Get("user")
	_, role, err := middlewares.ReadToken(token)
	links := map[string]string{"self": configs.Get().App.BaseURL + "/api/admins" + c.Param("id")}
	if tx != nil {
		return c.JSON(http.StatusBadRequest, web.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "ERROR",
			Error:  "invalid parameter",
			Links:  links,
		})
	}
	if err != nil {
		return c.JSON(http.StatusUnauthorized, web.ErrorResponse{
			Code:   http.StatusUnauthorized,
			Status: "ERROR",
			Error:  "unauthorized",
			Links:  links,
		})
	}
	if role == "admin" {
		return c.JSON(http.StatusUnauthorized, web.ErrorResponse{
			Code:   http.StatusUnauthorized,
			Status: "ERROR",
			Error:  "unauthorized",
			Links:  links,
		})
	}

	if userReq.Name == "" &&
		userReq.Address == "" && userReq.DOB == "" &&
		userReq.Email == "" && userReq.Gender == "" &&
		userReq.Password == "" && userReq.PhoneNumber == "" &&
		userReq.Role == "" {
		return c.JSON(http.StatusBadRequest, web.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "ERROR",
			Error:  "no such data filled",
			Links:  links,
		})
	}
	// Update via user service call
	userRes, err := handler.userService.UpdateUser(userReq, id)
	if err != nil {
		return helpers.WebErrorResponse(c, err, links)
	}

	// response
	return c.JSON(200, web.SuccessResponse{
		Status: "OK",
		Code:   200,
		Error:  nil,
		Links:  links,
		Data:   userRes,
	})
}

func (handler AdminHandler) Delete(c echo.Context) error {

	id, tx := strconv.Atoi(c.Param("id"))
	token := c.Get("user")
	_, role, err := middlewares.ReadToken(token)
	links := map[string]string{"self": configs.Get().App.BaseURL + "/api/admins" + c.Param("id")}
	if tx != nil {
		return c.JSON(http.StatusBadRequest, web.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "ERROR",
			Error:  "invalid parameter",
			Links:  links,
		})
	}
	if role != "admin" {
		return c.JSON(http.StatusUnauthorized, web.ErrorResponse{
			Code:   http.StatusUnauthorized,
			Status: "ERROR",
			Error:  "unauthorized",
			Links:  links,
		})
	}
	if err != nil {
		return c.JSON(http.StatusUnauthorized, web.ErrorResponse{
			Code:   http.StatusUnauthorized,
			Status: "ERROR",
			Error:  "unauthorized",
			Links:  links,
		})
	}

	// call delete service
	err = handler.userService.DeleteUser(id)
	if err != nil {
		return helpers.WebErrorResponse(c, err, links)
	}

	// response
	return c.JSON(200, web.SuccessResponse{
		Status: "OK",
		Code:   200,
		Error:  nil,
		Links:  links,
		Data: map[string]interface{}{
			"id": id,
		},
	})
}
