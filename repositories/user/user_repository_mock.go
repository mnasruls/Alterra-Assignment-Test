package user

import (
	"alterra/entities"
	"time"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type UserRepositoryMock struct {
	Mock *mock.Mock
}

func NewUserRepositoryMock(mock *mock.Mock) *UserRepositoryMock {
	return &UserRepositoryMock{
		Mock: mock,
	}
}

var UserCollection = []entities.User{
	{
		Model:       gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		Name:        "test1",
		Email:       "test1@mail.com",
		Password:    "test",
		PhoneNumber: "082111222333",
		Gender:      "male",
		DOB:         time.Now(),
		Address:     "jl. reformasi",
		Role:        "user",
	},
	{
		Model:       gorm.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		Name:        "test2",
		Email:       "test2@mail.com",
		Password:    "test",
		PhoneNumber: "082111222333",
		Gender:      "male",
		DOB:         time.Now(),
		Address:     "jl. reformasi",
		Role:        "admin",
	},
}

func (repo UserRepositoryMock) FindAllUser(limit int, offset int, filters []map[string]string, sorts []map[string]interface{}) ([]entities.User, error) {
	args := repo.Mock.Called()
	return args.Get(0).([]entities.User), args.Error(1)
}
func (repo UserRepositoryMock) Find(id int) (entities.User, error) {
	args := repo.Mock.Called()
	return args.Get(0).(entities.User), args.Error(1)
}
func (repo UserRepositoryMock) FindBy(field string, value string) (entities.User, error) {
	args := repo.Mock.Called()
	return args.Get(0).(entities.User), args.Error(1)
}
func (repo UserRepositoryMock) Store(user entities.User) (entities.User, error) {
	args := repo.Mock.Called()
	return args.Get(0).(entities.User), args.Error(1)
}
func (repo UserRepositoryMock) Update(user entities.User) (entities.User, error) {
	args := repo.Mock.Called()
	return args.Get(0).(entities.User), args.Error(1)
}
func (repo UserRepositoryMock) Delete(id int) error {
	args := repo.Mock.Called()
	return args.Error(0)
}
func (repo UserRepositoryMock) CountAllUser(filters []map[string]string) (int64, error) {
	args := repo.Mock.Called()
	return int64(args.Int(0)), args.Error(1)
}
