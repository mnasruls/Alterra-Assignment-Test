package user_test

import (
	"testing"

	"alterra/entities"
	"alterra/entities/web"
	userRepository "alterra/repositories/user"
	userService "alterra/services/user"

	"github.com/jinzhu/copier"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFind(t *testing.T) {
	t.Run("success", func(t *testing.T) {

		userSample := userRepository.UserCollection[0]
		userRepositoryMock := userRepository.NewUserRepositoryMock(&mock.Mock{})
		userRepositoryMock.Mock.On("Find").Return(userSample, nil)

		Service := userService.NewUserService(
			userRepositoryMock,
		)
		_, err := Service.Find(int(userSample.ID))

		assert.Nil(t, err)
	})

	t.Run("repo-fail", func(t *testing.T) {
		sampleCustomer := userRepository.UserCollection[0]
		userRepositoryMock := userRepository.NewUserRepositoryMock(&mock.Mock{})
		userRepositoryMock.Mock.On("Find").Return(entities.User{}, web.WebError{})

		Service := userService.NewUserService(
			userRepositoryMock,
		)

		actual, err := Service.Find(int(sampleCustomer.ID))
		assert.Error(t, err)
		assert.Equal(t, entities.UserResponse{}, actual)
	})
}

func TestCreate(t *testing.T) {
	sampleCentral := userRepository.UserCollection[0]
	sampleRequestCentral := entities.CreateUserRequest{}
	copier.Copy(&sampleRequestCentral, &sampleCentral)
	sampleRequestCentral.DOB = "1999-12-12"

	t.Run("success", func(t *testing.T) {
		sampleUser := sampleCentral
		sampleRequest := sampleRequestCentral

		userRepositoryMock := userRepository.NewUserRepositoryMock(&mock.Mock{})
		userRepositoryMock.Mock.On("Store").Return(sampleUser, nil)

		Service := userService.NewUserService(
			userRepositoryMock,
		)
		actual, err := Service.CreateUser(sampleRequest)

		expected := entities.UserResponse{}
		copier.Copy(&expected, &sampleUser)

		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})
	t.Run("validation-fail", func(t *testing.T) {
		sampleUser := sampleCentral
		sampleRequest := sampleRequestCentral
		sampleRequest.Name = ""
		userRepositoryMock := userRepository.NewUserRepositoryMock(&mock.Mock{})
		userRepositoryMock.Mock.On("Store").Return(sampleUser, nil)

		Service := userService.NewUserService(
			userRepositoryMock,
		)
		actual, err := Service.CreateUser(sampleRequest)

		expected := entities.UserResponse{}
		copier.Copy(&expected, &sampleUser)

		assert.Error(t, err)
		assert.Equal(t, entities.UserResponse{}, actual)
	})
	t.Run("invalid-dob", func(t *testing.T) {
		sampleUser := sampleCentral
		sampleRequest := sampleRequestCentral

		sampleRequest.DOB = "2022222222"
		userRepositoryMock := userRepository.NewUserRepositoryMock(&mock.Mock{})
		userRepositoryMock.Mock.On("Store").Return(sampleUser, nil)

		Service := userService.NewUserService(
			userRepositoryMock,
		)
		actual, err := Service.CreateUser(sampleRequest)

		expected := entities.UserResponse{}
		copier.Copy(&expected, &sampleUser)

		assert.Error(t, err)
		assert.Equal(t, entities.UserResponse{}, actual)
	})
	t.Run("store-fail", func(t *testing.T) {
		sampleRequest := sampleRequestCentral

		userRepositoryMock := userRepository.NewUserRepositoryMock(&mock.Mock{})
		userRepositoryMock.Mock.On("Store").Return(entities.User{}, web.WebError{})

		Service := userService.NewUserService(
			userRepositoryMock,
		)
		actual, err := Service.CreateUser(sampleRequest)

		expected := entities.UserResponse{}
		assert.Error(t, err)
		assert.Equal(t, expected, actual)
	})
}

func TestUpdate(t *testing.T) {
	sampleUserCentral := userRepository.UserCollection[0]
	sampleRequestCentral := entities.UpdateUserRequest{}
	copier.Copy(&sampleRequestCentral, &sampleUserCentral)
	sampleRequestCentral.DOB = "1999-12-12"
	t.Run("success", func(t *testing.T) {
		sampleRequest := sampleRequestCentral
		sampleUser := sampleUserCentral

		userRepositoryMock := userRepository.NewUserRepositoryMock(&mock.Mock{})
		userRepositoryMock.Mock.On("Find").Return(sampleUser, nil)

		userOutput := sampleUser
		copier.CopyWithOption(&userOutput, &sampleRequest, copier.Option{IgnoreEmpty: true, DeepCopy: true})
		userRepositoryMock.Mock.On("Update").Return(userOutput, nil)

		Service := userService.NewUserService(
			userRepositoryMock,
		)
		actual, err := Service.UpdateUser(sampleRequest, int(sampleUser.ID))
		expected := entities.UserResponse{}
		copier.Copy(&expected, &userOutput)

		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})
	t.Run("find-fail", func(t *testing.T) {
		sampleRequest := sampleRequestCentral
		sampleUser := sampleUserCentral

		userRepositoryMock := userRepository.NewUserRepositoryMock(&mock.Mock{})
		userRepositoryMock.Mock.On("Find").Return(entities.User{}, web.WebError{})

		userOutput := sampleUser
		copier.CopyWithOption(&userOutput, &sampleRequest, copier.Option{IgnoreEmpty: true, DeepCopy: true})
		userRepositoryMock.Mock.On("Update").Return(userOutput, nil)

		Service := userService.NewUserService(
			userRepositoryMock,
		)
		actual, err := Service.UpdateUser(sampleRequest, int(sampleUser.ID))
		assert.Error(t, err)
		assert.Equal(t, entities.UserResponse{}, actual)
	})
}

func TestDelete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		sampleCustomer := userRepository.UserCollection[0]
		userRepositoryMock := userRepository.NewUserRepositoryMock(&mock.Mock{})
		userRepositoryMock.Mock.On("Find").Return(sampleCustomer, nil)

		userRepositoryMock.Mock.On("Delete").Return(nil)

		Service := userService.NewUserService(
			userRepositoryMock,
		)
		err := Service.DeleteUser(int(sampleCustomer.ID))
		assert.Nil(t, err)
	})
	t.Run("repo-fail", func(t *testing.T) {
		sampleCustomer := userRepository.UserCollection[0]
		userRepositoryMock := userRepository.NewUserRepositoryMock(&mock.Mock{})
		userRepositoryMock.Mock.On("Find").Return(entities.User{}, web.WebError{})

		userRepositoryMock.Mock.On("Delete").Return(nil)

		Service := userService.NewUserService(
			userRepositoryMock,
		)
		err := Service.DeleteUser(int(sampleCustomer.ID))
		assert.Error(t, err)
	})
	t.Run("fail", func(t *testing.T) {
		sampleCustomer := userRepository.UserCollection[0]
		userRepositoryMock := userRepository.NewUserRepositoryMock(&mock.Mock{})
		userRepositoryMock.Mock.On("Find").Return(sampleCustomer, nil)

		userRepositoryMock.Mock.On("Delete").Return(web.WebError{})

		Service := userService.NewUserService(
			userRepositoryMock,
		)
		err := Service.DeleteUser(int(sampleCustomer.ID))
		assert.Error(t, err)
	})
}

func TestGetPaginationUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		userRepositoryMock := userRepository.NewUserRepositoryMock(&mock.Mock{})
		userRepositoryMock.Mock.On("CountAllUser").Return(20, nil)

		Service := userService.NewUserService(
			userRepositoryMock,
		)
		actual, err := Service.GetPaginationUser(5, 1, []map[string]string{})

		expected := web.Pagination{
			Page:       1,
			Limit:      5,
			TotalPages: int(4),
		}
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})
	t.Run("repo-fail", func(t *testing.T) {
		userRepositoryMock := userRepository.NewUserRepositoryMock(&mock.Mock{})
		userRepositoryMock.Mock.On("CountAllUser").Return(0, web.WebError{})

		Service := userService.NewUserService(
			userRepositoryMock,
		)
		actual, err := Service.GetPaginationUser(5, 1, []map[string]string{})
		assert.Error(t, err)
		assert.Equal(t, web.Pagination{}, actual)
	})
	t.Run("limit-zero", func(t *testing.T) {
		userRepositoryMock := userRepository.NewUserRepositoryMock(&mock.Mock{})
		userRepositoryMock.Mock.On("CountAllUser").Return(20, nil)

		Service := userService.NewUserService(
			userRepositoryMock,
		)
		actual, err := Service.GetPaginationUser(0, 1, []map[string]string{})

		expected := web.Pagination{
			Page:       1,
			Limit:      1,
			TotalPages: int(20),
		}
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})
	t.Run("added-page-on-active-module", func(t *testing.T) {
		userRepositoryMock := userRepository.NewUserRepositoryMock(&mock.Mock{})
		userRepositoryMock.Mock.On("CountAllUser").Return(22, nil)

		Service := userService.NewUserService(
			userRepositoryMock,
		)
		actual, err := Service.GetPaginationUser(5, 1, []map[string]string{})

		expected := web.Pagination{
			Page:       1,
			Limit:      5,
			TotalPages: int(5),
		}
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})
}

func TestFindAllUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		userSamples := userRepository.UserCollection
		userRepositoryMock := userRepository.NewUserRepositoryMock(&mock.Mock{})
		userRepositoryMock.Mock.On("FindAllUser").Return(userSamples, nil)

		Service := userService.NewUserService(
			userRepositoryMock,
		)

		actual, err := Service.FindAll(0, 0, []map[string]string{}, []map[string]interface{}{})
		expected := []entities.UserResponse{}
		copier.Copy(&expected, &actual)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})
	t.Run("fail", func(t *testing.T) {
		userRepositoryMock := userRepository.NewUserRepositoryMock(&mock.Mock{})
		userRepositoryMock.Mock.On("FindAllUser").Return([]entities.User{}, web.WebError{})

		Service := userService.NewUserService(
			userRepositoryMock,
		)

		actual, err := Service.FindAll(0, 0, []map[string]string{}, []map[string]interface{}{})
		assert.Error(t, err)
		assert.Equal(t, []entities.UserResponse{}, actual)
	})
}

func TestFindBy(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		userSample := userRepository.UserCollection[0]
		userRepositoryMock := userRepository.NewUserRepositoryMock(&mock.Mock{})
		userRepositoryMock.Mock.On("FindBy").Return(userSample, nil)

		Service := userService.NewUserService(
			userRepositoryMock,
		)
		actual, err := Service.FindBy("id", "1")
		expected := entities.UserResponse{}
		copier.Copy(&expected, &actual)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})
	t.Run("user-repo-fail", func(t *testing.T) {
		userRepositoryMock := userRepository.NewUserRepositoryMock(&mock.Mock{})
		userRepositoryMock.Mock.On("FindBy").Return(entities.User{}, web.WebError{})

		Service := userService.NewUserService(
			userRepositoryMock,
		)
		actual, err := Service.FindBy("id", "1")
		expected := entities.UserResponse{}
		copier.Copy(&expected, &actual)
		assert.Error(t, err)
		assert.Equal(t, expected, actual)
	})
}
