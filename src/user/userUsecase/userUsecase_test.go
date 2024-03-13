package userUsecase

import (
	"avengers-clinic/model/dto/userDto"
	"avengers-clinic/pkg/utils"
	"avengers-clinic/src/user"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type mockUserRepository struct {
	mock.Mock
}

func (mock *mockUserRepository) GetAllTrash() ([]userDto.User, error) {
	args := mock.Called()
	return args.Get(0).([]userDto.User), args.Error(1)
}

func (mock *mockUserRepository) GetAll() ([]userDto.User, error) {
	args := mock.Called()
	return args.Get(0).([]userDto.User), args.Error(1)
}

func (mock *mockUserRepository) GetByID(userID string) (userDto.User, error) {
	args := mock.Called(userID)
	return args.Get(0).(userDto.User), args.Error(1)
}

func (mock *mockUserRepository) GetUserByID(userID string) (userDto.User, error) {
	args := mock.Called(userID)
	return args.Get(0).(userDto.User), args.Error(1)
}

func (mock *mockUserRepository) GetTrashByID(userID string) (userDto.User, error) {
	args := mock.Called(userID)
	return args.Get(0).(userDto.User), args.Error(1)
}

func (mock *mockUserRepository) GetByUsername(username string) (userDto.User, error) {
	args := mock.Called(username)
	return args.Get(0).(userDto.User), args.Error(1)
}

func (mock *mockUserRepository) Insert(user userDto.User) (string, error) {
	args := mock.Called(user)
	return args.String(0), args.Error(1)
}

func (mock *mockUserRepository) Update(user userDto.User) error {
	args := mock.Called(user)
	return args.Error(0)
}

func (mock *mockUserRepository) UpdatePassword(userId, hashPassword string) error {
	args := mock.Called(userId, hashPassword)
	return args.Error(0)
}

func (mock *mockUserRepository) Delete(userID string) error {
	args := mock.Called(userID)
	return args.Error(0)
}

func (mock *mockUserRepository) SoftDelete(userID string) error {
	args := mock.Called(userID)
	return args.Error(0)
}

func (mock *mockUserRepository) Restore(userID string) error {
	args := mock.Called(userID)
	return args.Error(0)
}

func (mock *mockUserRepository) IsUsernameExists(username string) bool {
	args := mock.Called(username)
	return args.Bool(0)
}

type userUsecaseTestSuite struct {
	suite.Suite
	userRepo *mockUserRepository
	userUC user.UserUsecase
}

func (suite *userUsecaseTestSuite) SetupTest() {
	suite.userRepo = new(mockUserRepository)
	suite.userUC = NewUserUsecase(suite.userRepo)
}

func (suite *userUsecaseTestSuite) TestGetAllTrashSuccess() {
	expectedUsers := []userDto.User{
		{
			ID: "9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5",
			Username: "admin",
			Password: "$2a$10$4YY9SUvhhVtqURrsbBbDre5MjimjgajD5KsJZg6QkaJ.75jR.6soq",
			Role: "ADMIN",
		},
	}

	suite.userRepo.On("GetAllTrash").Return(expectedUsers, nil)
	actualUsers, err := suite.userUC.GetAllTrash()
	
	suite.Nil(err)
	suite.Equal(expectedUsers, actualUsers)
}

func (suite *userUsecaseTestSuite) TestGetAllSuccess() {
	expectedUsers := []userDto.User{
		{
			ID: "9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5",
			Username: "admin",
			Password: "$2a$10$4YY9SUvhhVtqURrsbBbDre5MjimjgajD5KsJZg6QkaJ.75jR.6soq",
			Role: "ADMIN",
		},
	}

	suite.userRepo.On("GetAll").Return(expectedUsers, nil)
	actualUsers, err := suite.userUC.GetAll()
	
	suite.Nil(err)
	suite.Equal(expectedUsers, actualUsers)
}

func (suite *userUsecaseTestSuite) TestGetByIDSuccess() {
	expectedUsers := userDto.User{
		ID: "9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5",
		Username: "admin",
		Password: "$2a$10$4YY9SUvhhVtqURrsbBbDre5MjimjgajD5KsJZg6QkaJ.75jR.6soq",
		Role: "ADMIN",
	}

	suite.userRepo.On("GetByID", expectedUsers.ID).Return(expectedUsers, nil)
	actualUsers, err := suite.userUC.GetByID(expectedUsers.ID)
	
	suite.Nil(err)
	suite.Equal(expectedUsers, actualUsers)
}

// Start Register Patient
func (suite *userUsecaseTestSuite) TestPatientRegisterSuccess() {
	request := userDto.AuthRequest{
		Username: "user",
		Password: "user",
	}

	suite.userRepo.On("IsUsernameExists", request.Username).Return(false)
	suite.userRepo.On("Insert", mock.Anything).Return("9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5", nil)
	actualUser, err := suite.userUC.PatientRegister(request)

	suite.Nil(err)
	suite.Equal(request.Username, actualUser.Username)
}

func (suite *userUsecaseTestSuite) TestPatientRegisterErrorUsernameExists() {
	request := userDto.AuthRequest{
		Username: "user",
		Password: "user",
	}

	suite.userRepo.On("IsUsernameExists", request.Username).Return(true)
	actualUser, err := suite.userUC.PatientRegister(request)

	suite.Error(err)
	suite.Empty(actualUser)
}

func (suite *userUsecaseTestSuite) TestPatientRegisterInternalServerError() {
	request := userDto.AuthRequest{
		Username: "user",
		Password: "user",
	}

	suite.userRepo.On("IsUsernameExists", request.Username).Return(false)
	suite.userRepo.On("Insert", mock.Anything).Return("", sql.ErrConnDone)
	actualUser, err := suite.userUC.PatientRegister(request)

	suite.Error(err)
	suite.Empty(actualUser)
}
// End Register Patient

// Start Register User
func (suite *userUsecaseTestSuite) TestUserRegisterSuccess() {
	request := userDto.RegisterRequest{
		Username: "user",
		Password: "user",
		Role: "DOCTOR",
		Specialization: "Gigi",
	}

	suite.userRepo.On("IsUsernameExists", request.Username).Return(false)
	suite.userRepo.On("Insert", mock.Anything).Return("9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5", nil)
	actualUser, err := suite.userUC.UserRegister(request)

	suite.Nil(err)
	suite.Equal(request.Username, actualUser.Username)
}

func (suite *userUsecaseTestSuite) TestUserRegisterErrorUsernameExists() {
	request := userDto.RegisterRequest{
		Username: "user",
		Password: "user",
		Role: "DOCTOR",
		Specialization: "Gigi",
	}

	suite.userRepo.On("IsUsernameExists", request.Username).Return(true)
	actualUser, err := suite.userUC.UserRegister(request)

	suite.Error(err)
	suite.Empty(actualUser)
}

func (suite *userUsecaseTestSuite) TestUserRegisterErrorBadRequest() {
	request := userDto.RegisterRequest{
		Username: "user",
		Password: "user",
		Role: "DOCTOR",
	}

	suite.userRepo.On("IsUsernameExists", request.Username).Return(false)
	actualUser, err := suite.userUC.UserRegister(request)

	suite.Error(err)
	suite.Empty(actualUser)
}

func (suite *userUsecaseTestSuite) TestUserRegisterInternalServerError() {
	request := userDto.RegisterRequest{
		Username: "user",
		Password: "user",
		Role: "DOCTOR",
		Specialization: "Gigi",
	}

	suite.userRepo.On("IsUsernameExists", request.Username).Return(false)
	suite.userRepo.On("Insert", mock.Anything).Return("", sql.ErrConnDone)
	actualUser, err := suite.userUC.UserRegister(request)

	suite.Error(err)
	suite.Empty(actualUser)
}
// End Register User

// Start Login
func (suite *userUsecaseTestSuite) TestLoginSuccess() {
	request := userDto.AuthRequest{
		Username: "admin",
		Password: "admin",
	}
	suite.userRepo.On("IsUsernameExists", request.Username).Return(true)

	expectedUser := userDto.User{
		ID: "9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5",
		Username: "admin",
		Password: "$2a$10$4YY9SUvhhVtqURrsbBbDre5MjimjgajD5KsJZg6QkaJ.75jR.6soq",
		Role: "ADMIN",
	}
	suite.userRepo.On("GetByUsername", request.Username).Return(expectedUser, nil)
	expectedToken, _ := utils.GenerateJWT(expectedUser.ID, expectedUser.Username, expectedUser.Role)
	actualToken, err := suite.userUC.Login(request)

	suite.Nil(err)
	suite.Equal(expectedToken, actualToken)
}

func (suite *userUsecaseTestSuite) TestLoginErrorUsernameNotExists() {
	suite.userRepo.On("IsUsernameExists", "user").Return(false)

	request := userDto.AuthRequest{
		Username: "user",
		Password: "rahasia",
	}
	actualToken, err := suite.userUC.Login(request)

	suite.Error(err)
	suite.Empty(actualToken)
}

func (suite *userUsecaseTestSuite) TestLoginErrorWrongPassword() {
	expectedUser := userDto.User{
		ID: "9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5",
		Username: "admin",
		Password: "$2a$10$4YY9SUvhhVtqURrsbBbDre5MjimjgajD5KsJZg6QkaJ.75jR.6soq",
		Role: "ADMIN",
	}

	suite.userRepo.On("IsUsernameExists", "admin").Return(true)
	suite.userRepo.On("GetByUsername", "admin").Return(expectedUser, nil)
	
	request := userDto.AuthRequest{
		Username: "admin",
		Password: "rahasia",
	}
	actualToken, err := suite.userUC.Login(request)

	suite.Error(err)
	suite.Empty(actualToken)
}
// End Login

// Start Update
func (suite *userUsecaseTestSuite) TestUpdateSuccess() {
	request := userDto.UpdateRequest{
		ID: "9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5",
		Username: "user",
	}

	expectUser := userDto.User{
		ID: "9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5",
		Username: "admin",
		Role: "ADMIN",
	}

	suite.userRepo.On("GetByID", request.ID).Return(expectUser, nil)
	suite.userRepo.On("IsUsernameExists", request.Username).Return(false)
	suite.userRepo.On("Update", mock.Anything).Return(nil)
	actualUser, err := suite.userUC.Update(request)

	suite.Nil(err)
	suite.Equal(request.Username, actualUser.Username)
}

func (suite *userUsecaseTestSuite) TestUpdateErrorUserNotFound() {
	request := userDto.UpdateRequest{
		ID: "9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5",
		Username: "user",
	}

	suite.userRepo.On("GetByID", request.ID).Return(userDto.User{}, sql.ErrNoRows)
	actualUser, err := suite.userUC.Update(request)

	suite.Error(err)
	suite.Empty(actualUser)
}

func (suite *userUsecaseTestSuite) TestUpdateErrorUsernameExists() {
	request := userDto.UpdateRequest{
		ID: "9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5",
		Username: "user",
	}

	expectUser := userDto.User{
		ID: "9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5",
		Username: "admin",
		Role: "ADMIN",
	}

	suite.userRepo.On("GetByID", request.ID).Return(expectUser, nil)
	suite.userRepo.On("IsUsernameExists", request.Username).Return(true)
	actualUser, err := suite.userUC.Update(request)

	suite.Error(err)
	suite.Empty(actualUser)
}

func (suite *userUsecaseTestSuite) TestUpdateInternalServerError() {
	request := userDto.UpdateRequest{
		ID: "9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5",
		Username: "user",
	}

	expectUser := userDto.User{
		ID: "9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5",
		Username: "admin",
		Role: "ADMIN",
	}

	suite.userRepo.On("GetByID", request.ID).Return(expectUser, nil)
	suite.userRepo.On("IsUsernameExists", request.Username).Return(false)
	suite.userRepo.On("Update", mock.Anything).Return(sql.ErrConnDone)
	actualUser, err := suite.userUC.Update(request)

	suite.Error(err)
	suite.Empty(actualUser)
}
// End Update

// Start Update Password
func (suite *userUsecaseTestSuite) TestUpdatePasswordSuccess() {
	request := userDto.UpdatePasswordRequest{
		ID: "9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5",
		CurrentPassword: "admin",
		NewPassword: "secret",
		ConfirmationPassword: "secret",
	}

	expectedUser := userDto.User{
		ID: "9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5",
		Username: "admin",
		Password: "$2a$10$4YY9SUvhhVtqURrsbBbDre5MjimjgajD5KsJZg6QkaJ.75jR.6soq",
		Role: "ADMIN",
	}

	suite.userRepo.On("GetByID", request.ID).Return(expectedUser, nil)
	suite.userRepo.On("UpdatePassword", request.ID, mock.Anything).Return(nil)
	err := suite.userUC.UpdatePassword(request)

	suite.Nil(err)
}

func (suite *userUsecaseTestSuite) TestUpdatePasswordErrorUserNotFound() {
	request := userDto.UpdatePasswordRequest{
		ID: "9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5",
		CurrentPassword: "admin",
		NewPassword: "secret",
		ConfirmationPassword: "secret",
	}

	suite.userRepo.On("GetByID", request.ID).Return(userDto.User{}, sql.ErrNoRows)
	err := suite.userUC.UpdatePassword(request)

	suite.Error(err)
}

func (suite *userUsecaseTestSuite) TestUpdatePasswordErrorWrongPassword() {
	request := userDto.UpdatePasswordRequest{
		ID: "9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5",
		CurrentPassword: "admin1",
		NewPassword: "secret",
		ConfirmationPassword: "secret",
	}

	expectedUser := userDto.User{
		ID: "9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5",
		Username: "admin",
		Password: "$2a$10$4YY9SUvhhVtqURrsbBbDre5MjimjgajD5KsJZg6QkaJ.75jR.6soq",
		Role: "ADMIN",
	}

	suite.userRepo.On("GetByID", request.ID).Return(expectedUser, nil)
	err := suite.userUC.UpdatePassword(request)

	suite.Error(err)
}

func (suite *userUsecaseTestSuite) TestUpdatePasswordErrorWrongConfirmation() {
	request := userDto.UpdatePasswordRequest{
		ID: "9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5",
		CurrentPassword: "admin",
		NewPassword: "secret",
		ConfirmationPassword: "secret1",
	}

	expectedUser := userDto.User{
		ID: "9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5",
		Username: "admin",
		Password: "$2a$10$4YY9SUvhhVtqURrsbBbDre5MjimjgajD5KsJZg6QkaJ.75jR.6soq",
		Role: "ADMIN",
	}

	suite.userRepo.On("GetByID", request.ID).Return(expectedUser, nil)
	err := suite.userUC.UpdatePassword(request)

	suite.Error(err)
}

func (suite *userUsecaseTestSuite) TestUpdatePasswordInternalServerError() {
	request := userDto.UpdatePasswordRequest{
		ID: "9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5",
		CurrentPassword: "admin",
		NewPassword: "secret",
		ConfirmationPassword: "secret",
	}

	expectedUser := userDto.User{
		ID: "9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5",
		Username: "admin",
		Password: "$2a$10$4YY9SUvhhVtqURrsbBbDre5MjimjgajD5KsJZg6QkaJ.75jR.6soq",
		Role: "ADMIN",
	}

	suite.userRepo.On("GetByID", request.ID).Return(expectedUser, nil)
	suite.userRepo.On("UpdatePassword", request.ID, mock.Anything).Return(sql.ErrConnDone)
	err := suite.userUC.UpdatePassword(request)

	suite.Error(err)
}
// End Update Password

// Start Delete
func (suite *userUsecaseTestSuite) TestDeleteSuccess() {
	userID := "9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5"
	
	expectUser := userDto.User{
		ID: "9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5",
	}

	suite.userRepo.On("GetUserByID", userID).Return(expectUser, nil)
	suite.userRepo.On("Delete", userID).Return(nil)
	err := suite.userUC.Delete(userID)

	suite.Nil(err)
}

func (suite *userUsecaseTestSuite) TestDeleteErrorUserNotFound() {
	userID := "9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5"

	expectUser := userDto.User{}

	suite.userRepo.On("GetUserByID", userID).Return(expectUser, sql.ErrNoRows)
	err := suite.userUC.Delete(userID)

	suite.Error(err)
}

func (suite *userUsecaseTestSuite) TestDeleteInternalServerError() {
	userID := "9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5"

	expectUser := userDto.User{
		ID: "9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5",
	}

	suite.userRepo.On("GetUserByID", userID).Return(expectUser, nil)
	suite.userRepo.On("Delete", userID).Return(sql.ErrConnDone)
	err := suite.userUC.Delete(userID)

	suite.Error(err)
}
// End Delete

// Start Soft Delete
func (suite *userUsecaseTestSuite) TestSoftDeleteSuccess() {
	userID := "9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5"
	
	expectUser := userDto.User{
		ID: "9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5",
	}

	suite.userRepo.On("GetByID", userID).Return(expectUser, nil)
	suite.userRepo.On("SoftDelete", userID).Return(nil)
	err := suite.userUC.SoftDelete(userID)

	suite.Nil(err)
}

func (suite *userUsecaseTestSuite) TestSoftDeleteErrorUserNotFound() {
	userID := "9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5"

	expectUser := userDto.User{}

	suite.userRepo.On("GetByID", userID).Return(expectUser, sql.ErrNoRows)
	err := suite.userUC.SoftDelete(userID)

	suite.Error(err)
}

func (suite *userUsecaseTestSuite) TestSoftDeleteInternalServerError() {
	userID := "9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5"

	expectUser := userDto.User{
		ID: "9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5",
	}

	suite.userRepo.On("GetByID", userID).Return(expectUser, nil)
	suite.userRepo.On("SoftDelete", userID).Return(sql.ErrConnDone)
	err := suite.userUC.SoftDelete(userID)

	suite.Error(err)
}
// End Soft Delete

// Start Restore
func (suite *userUsecaseTestSuite) TestRestoreSuccess() {
	userID := "9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5"
	
	expectUser := userDto.User{
		ID: "9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5",
	}

	suite.userRepo.On("GetTrashByID", userID).Return(expectUser, nil)
	suite.userRepo.On("Restore", userID).Return(nil)
	err := suite.userUC.Restore(userID)

	suite.Nil(err)
}

func (suite *userUsecaseTestSuite) TestRestoreErrorUserNotFound() {
	userID := "9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5"

	expectUser := userDto.User{}

	suite.userRepo.On("GetTrashByID", userID).Return(expectUser, sql.ErrNoRows)
	err := suite.userUC.Restore(userID)

	suite.Error(err)
}

func (suite *userUsecaseTestSuite) TestRestoreInternalServerError() {
	userID := "9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5"

	expectUser := userDto.User{
		ID: "9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5",
	}

	suite.userRepo.On("GetTrashByID", userID).Return(expectUser, nil)
	suite.userRepo.On("Restore", userID).Return(sql.ErrConnDone)
	err := suite.userUC.Restore(userID)

	suite.Error(err)
}
// End Restore

func TestUserUsecase(t *testing.T) {
	suite.Run(t, new(userUsecaseTestSuite))
}