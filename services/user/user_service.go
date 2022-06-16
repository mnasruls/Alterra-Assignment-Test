package user

import (
	"alterra/deliveries/helpers"
	"alterra/deliveries/middlewares"
	"alterra/deliveries/validation"
	"alterra/entities"
	"alterra/entities/web"
	userRepository "alterra/repositories/user"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
)

type UserService struct {
	userRepo userRepository.UserRepositoryInterface
	validate *validator.Validate
}

func NewUserService(repository userRepository.UserRepositoryInterface) *UserService {
	return &UserService{
		userRepo: repository,
		validate: validator.New(),
	}
}

func (service UserService) CreateUser(userRequest entities.CreateUserRequest) (entities.UserAuthResponse, error) {

	// Validation
	err := validation.ValidateCreateUserRequest(service.validate, userRequest)
	if err != nil {
		return entities.UserAuthResponse{}, err
	}

	// Konversi user request menjadi domain untuk diteruskan ke repository
	user := entities.User{}
	copier.Copy(&user, &userRequest)

	// Konversi datetime untuk field datetime (dob)
	dob, err := time.Parse("2006-01-02", userRequest.DOB)
	if err != nil {
		return entities.UserAuthResponse{}, web.WebError{Code: 400, Message: "date of birth format is invalid"}
	}
	user.DOB = dob

	// Password hashing menggunakan bcrypt
	hashedPassword, _ := helpers.HashPassword(user.Password)
	user.Password = hashedPassword

	// Insert ke sistem melewati repository
	user, err = service.userRepo.Store(user)
	if err != nil {
		return entities.UserAuthResponse{}, err
	}

	// Konversi hasil repository menjadi user response
	userRes := entities.UserResponse{}
	copier.Copy(&userRes, &user)

	// generate token
	token, err := middlewares.CreateToken(int(user.ID), user.Name, user.Role)
	if err != nil {
		return entities.UserAuthResponse{}, err
	}

	// Buat auth response untuk dimasukkan token dan user
	authRes := entities.UserAuthResponse{
		Token: token,
		User:  userRes,
	}
	return authRes, nil
}
