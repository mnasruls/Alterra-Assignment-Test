package user

import (
	"alterra/entities"
	"alterra/entities/web"
)

type UserServiceInterface interface {
	FindAllUser(limit, page int, filters []map[string]string, sorts []map[string]interface{}) ([]entities.UserResponse, error)
	CreateUser(UserRequest entities.CreateUserRequest) (entities.UserResponse, error)
	UpdateUser(UserRequest entities.UpdateUserRequest, id int) (entities.UserResponse, error)
	GetPaginationUser(limit, page int, filters []map[string]string) (web.Pagination, error)
	FindByUser(field string, value string) (entities.UserResponse, error)
	FindUser(id int) (entities.UserResponse, error)
	DeleteUser(id int) error
}
