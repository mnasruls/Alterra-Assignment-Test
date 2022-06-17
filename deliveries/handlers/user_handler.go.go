package handlers

import (
	"alterra/configs"
	"alterra/deliveries/helpers"
	"alterra/entities/web"
	userService "alterra/services/user"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService *userService.UserService
}

func NewUserHandler(service *userService.UserService) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

func (handler UserHandler) GetSingleUser(c echo.Context) error {
	// Get param
	id, tx := strconv.Atoi(c.Param("id"))
	links := map[string]string{"self": configs.Get().App.BaseURL + "/api/users/" + c.Param("id")}
	if tx != nil {
		return c.JSON(http.StatusBadRequest, web.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "ERROR",
			Error:  "invalid parameter",
			Links:  links,
		})
	}

	// Get user data
	user, err := handler.userService.Find(id)
	if err != nil {
		return helpers.WebErrorResponse(c, err, links)
	} else if user.ID == 0 {
		return c.JSON(http.StatusBadRequest, web.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "ERROR",
			Error:  "bad request",
			Links:  links,
		})
	}
	// response
	return c.JSON(200, web.SuccessResponse{
		Status: "OK",
		Code:   200,
		Error:  nil,
		Links:  links,
		Data:   user,
	})
}

func (handler UserHandler) GetAllUser(c echo.Context) error {

	// Translate query param to map of filters
	filters := []map[string]string{}
	name := c.QueryParam("name")
	if name != "" {
		filters = append(filters, map[string]string{
			"field":    "name",
			"operator": "LIKE",
			"value":    "%" + name + "%",
		})
	}
	gender := c.QueryParam("gender")
	if gender != "" {
		filters = append(filters, map[string]string{
			"field":    "gender",
			"operator": "=",
			"value":    gender,
		})
	}

	// Sort parameter
	sorts := []map[string]interface{}{}
	sortName := c.QueryParam("sort_name")
	sorts = append(sorts, map[string]interface{}{
		"field": "name",
		"desc":  map[string]bool{"1": true, "0": false}[sortName],
	})

	links := map[string]string{"self": configs.Get().App.BaseURL + "/api/users?limit=" + c.QueryParam("limit") + "&page=" + c.QueryParam("page")}

	// pagination param
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return c.JSON(400, helpers.MakeErrorResponse("ERROR", 400, "Limit Parameter format is invalid", links))
	}
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		links := map[string]string{"self": configs.Get().App.BaseURL}
		return c.JSON(400, helpers.MakeErrorResponse("ERROR", 400, "page Parameter format is invalid", links))
	}
	links["self"] = configs.Get().App.BaseURL + "/api/drivers?limit=" + c.QueryParam("limit") + "&page=" + c.QueryParam("page")

	// Get all drivers
	driversRes, err := handler.userService.FindAll(limit, page, filters, sorts)
	if err != nil {
		return helpers.WebErrorResponse(c, err, links)
	}

	// Get pagination data
	pagination, err := handler.userService.GetPaginationUser(limit, page, filters)
	if err != nil {
		return helpers.WebErrorResponse(c, err, links)
	}

	links["first"] = configs.Get().App.BaseURL + "/api/users?limit=" + c.QueryParam("limit") + "&page=1"
	links["last"] = configs.Get().App.BaseURL + "/api/users?limit=" + c.QueryParam("limit") + "&page=" + strconv.Itoa(pagination.TotalPages)
	if pagination.Page > 1 {
		links["prev"] = configs.Get().App.BaseURL + "/api/users?limit=" + c.QueryParam("limit") + "&page=" + strconv.Itoa(pagination.Page-1)
	}
	if pagination.Page < pagination.TotalPages {
		links["next"] = configs.Get().App.BaseURL + "/api/users?limit=" + c.QueryParam("limit") + "&page=" + strconv.Itoa(pagination.Page+1)
	}

	// success response
	return c.JSON(200, web.SuccessListResponse{
		Status:     "OK",
		Code:       200,
		Error:      nil,
		Links:      links,
		Data:       driversRes,
		Pagination: pagination,
	})
}
