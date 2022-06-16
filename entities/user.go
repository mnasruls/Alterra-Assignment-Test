package entities

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email       string `gorm:"unique"`
	Name        string
	Password    string
	DOB         time.Time
	Gender      string
	Address     string
	PhoneNumber string
	Role        string
}

type CreateUserRequest struct {
	Email       string `form:"email" validate:"required,email"`
	Password    string `form:"password" validate:"required"`
	Name        string `form:"name" validate:"required"`
	DOB         string `form:"dob" validate:"required"`
	Gender      string `form:"gender" validate:"required"`
	Address     string `form:"address" validate:"required"`
	PhoneNumber string `form:"phone_number" validate:"required"`
	Role        string `form:"role" validate:"required"`
}

type UpdateUserRequest struct {
	Email       string `form:"email"`
	Password    string `form:"password"`
	Name        string `form:"name"`
	DOB         string `form:"dob"`
	Gender      string `form:"gender"`
	Address     string `form:"address"`
	PhoneNumber string `form:"phone_number"`
	Role        string `form:"role"`
}

type UserResponse struct {
	ID          uint      `json:"id"`
	Email       string    `json:"email"`
	Name        string    `json:"name"`
	DOB         time.Time `json:"dob"`
	Gender      string    `json:"gender"`
	Address     string    `json:"address"`
	PhoneNumber string    `json:"phone_number"`
	Role        string    `json:"role"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
