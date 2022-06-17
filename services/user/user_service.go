package user

import (
	"alterra/deliveries/helpers"
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

func (service UserService) CreateUser(userRequest entities.CreateUserRequest) (entities.UserResponse, error) {

	// Validation
	err := validation.ValidateCreateUserRequest(service.validate, userRequest)
	if err != nil {
		return entities.UserResponse{}, err
	}

	// Konversi user request menjadi domain untuk diteruskan ke repository
	user := entities.User{}
	copier.Copy(&user, &userRequest)

	// Konversi datetime untuk field datetime (dob)
	dob, err := time.Parse("2006-01-02", userRequest.DOB)
	if err != nil {
		return entities.UserResponse{}, web.WebError{Code: 400, Message: "date of birth format is invalid"}
	}
	user.DOB = dob

	// Password hashing menggunakan bcrypt
	hashedPassword, _ := helpers.HashPassword(user.Password)
	user.Password = hashedPassword

	// Insert ke sistem melewati repository
	user, err = service.userRepo.Store(user)
	if err != nil {
		return entities.UserResponse{}, err
	}

	// Konversi hasil repository menjadi user response
	userRes := entities.UserResponse{}
	copier.Copy(&userRes, &user)

	return userRes, nil
}

/*
 * User Service - FindAll
 * -------------------------------
 * Mencari semua user
 */
func (service UserService) FindAll(limit, page int, filters []map[string]string, sorts []map[string]interface{}) ([]entities.UserResponse, error) {

	offset := (page - 1) * limit

	usersRes := []entities.UserResponse{}

	// Mengambil data user dari repository
	users, err := service.userRepo.FindAllUser(limit, offset, filters, sorts)
	if err != nil {
		return []entities.UserResponse{}, err
	}

	// proses menjadi user response
	copier.Copy(&usersRes, &users)

	return usersRes, err
}

/*
 * User Service - Find
 * -------------------------------
 * Mencari user berdasarkan ID
 */
func (service UserService) Find(id int) (entities.UserResponse, error) {

	// Mengambil data user dari repository
	user, err := service.userRepo.Find(id)
	if err != nil {
		return entities.UserResponse{}, err
	}

	// proses menjadi user response
	userRes := entities.UserResponse{}
	copier.Copy(&userRes, &user)

	return userRes, err
}

/*
 * User Service - FindBy
 * -------------------------------
 * Mencari user berdasarkan column dan value
 */
func (service UserService) FindBy(field string, value string) (entities.UserResponse, error) {

	// Mengambil data user dari repository
	user, err := service.userRepo.FindBy(field, value)
	if err != nil {
		return entities.UserResponse{}, err
	}

	// proses menjadi user response
	userRes := entities.UserResponse{}
	copier.Copy(&userRes, &user)

	return userRes, err
}

func (service UserService) UpdateUser(userRequest entities.UpdateUserRequest, id int) (entities.UserResponse, error) {

	// Get user by ID via repository
	user, err := service.userRepo.Find(id)
	if err != nil {
		return entities.UserResponse{}, web.WebError{Code: 400, Message: "The requested ID doesn't match with any record"}
	}

	// Konversi dari request ke domain entities user - mengabaikan nilai kosong pada request
	copier.CopyWithOption(&user, &userRequest, copier.Option{IgnoreEmpty: true, DeepCopy: true})

	// Hanya hash password jika password juga diganti (tidak kosong)
	if userRequest.Password != "" {
		hashedPassword, _ := helpers.HashPassword(user.Password)
		user.Password = hashedPassword
	}

	// Update via repository
	user, err = service.userRepo.Update(user)
	// Konversi user domain menjadi user response
	userRes := entities.UserResponse{}
	copier.Copy(&userRes, &user)

	return userRes, err
}

func (service UserService) DeleteUser(id int) error {

	// Cari user berdasarkan ID via repo
	_, err := service.userRepo.Find(id)
	if err != nil {
		return web.WebError{Code: 400, Message: "The request ID has been deleted or not exist"}
	}

	// Delete via repository
	err = service.userRepo.Delete(id)
	return err
}
