package userDelivery

import (
	"avengers-clinic/model/dto/userDto"
	"avengers-clinic/pkg/utils"
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type mockUserUsecase struct {
	mock.Mock
}

func (mock *mockUserUsecase)GetAllTrash() ([]userDto.User, error) {
	args := mock.Called()
	return args.Get(0).([]userDto.User), args.Error(1)
}

func (mock *mockUserUsecase)GetAll() ([]userDto.User, error) {
	args := mock.Called()
	return args.Get(0).([]userDto.User), args.Error(1)
}

func (mock *mockUserUsecase)GetByID(userID string) (userDto.User, error) {
	args := mock.Called(userID)
	return args.Get(0).(userDto.User), args.Error(1)
}

func (mock *mockUserUsecase)PatientRegister(req userDto.AuthRequest) (userDto.User, error) {
	args := mock.Called(req)
	return args.Get(0).(userDto.User), args.Error(1)
}

func (mock *mockUserUsecase)UserRegister(req userDto.RegisterRequest) (userDto.User, error) {
	args := mock.Called(req)
	return args.Get(0).(userDto.User), args.Error(1)
}

func (mock *mockUserUsecase)Login(req userDto.AuthRequest) (string, error) {
	args := mock.Called(req)
	return args.String(0), args.Error(1)
}

func (mock *mockUserUsecase)Update(req userDto.UpdateRequest) (userDto.User, error) {
	args := mock.Called(req)
	return args.Get(0).(userDto.User), args.Error(1)
}

func (mock *mockUserUsecase)UpdatePassword(req userDto.UpdatePasswordRequest) error {
	args := mock.Called(req)
	return args.Error(0)
}

func (mock *mockUserUsecase)Delete(userID string) error {
	args := mock.Called(userID)
	return args.Error(0)
}

func (mock *mockUserUsecase)SoftDelete(userID string) error {
	args := mock.Called(userID)
	return args.Error(0)
}

func (mock *mockUserUsecase)Restore(userID string) error {
	args := mock.Called(userID)
	return args.Error(0)
}

type userDeliveryTestSuite struct {
	suite.Suite
	router *gin.Engine
	userUC *mockUserUsecase
}

func (suite *userDeliveryTestSuite) SetupTest() {
	suite.router = gin.New()
	suite.userUC = new(mockUserUsecase)

	v1Group := suite.router.Group("/api/v1")
	NewUserDelivery(v1Group, suite.userUC)
}

// Start Get All Trash
func (suite *userDeliveryTestSuite) TestGetAllTrashSuccess() {
	users := []userDto.User{
		{
			ID: "9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5",
			Username: "admin",
			Password: "$2a$10$4YY9SUvhhVtqURrsbBbDre5MjimjgajD5KsJZg6QkaJ.75jR.6soq",
			Role: "ADMIN",
		},
	}
	suite.userUC.On("GetAllTrash").Return(users, nil)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/users/trash", nil)

	token, _ := utils.GenerateJWT("9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expectedResponse := `{"responseCode":"2000101","responseMessage":"Users retrieved successfully","data":[{"id":"9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5","username":"admin","role":"ADMIN"}]}`
	
	suite.Equal(http.StatusOK, res.Code)
	suite.JSONEq(expectedResponse, res.Body.String())
}

func (suite *userDeliveryTestSuite) TestGetAllTrashErrorUserNotFound() {
	users := []userDto.User{}
	suite.userUC.On("GetAllTrash").Return(users, nil)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/users/trash", nil)

	token, _ := utils.GenerateJWT("9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	fmt.Println("Body:", res.Body.String())

	expectedResponse := `{"responseCode":"4030102","responseMessage":"Users not found"}`

	suite.Equal(http.StatusForbidden, res.Code)
	suite.JSONEq(expectedResponse, res.Body.String())
}

func (suite *userDeliveryTestSuite) TestGetAllTrashInternalServerError() {
	users := []userDto.User{}
	suite.userUC.On("GetAllTrash").Return(users, sql.ErrConnDone)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/users/trash", nil)

	token, _ := utils.GenerateJWT("9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expectedResponse := `{"responseCode":"5000101","responseMessage":"internal server error","error":"sql: connection is already closed"}`

	suite.Equal(http.StatusInternalServerError, res.Code)
	suite.JSONEq(expectedResponse, res.Body.String())
}
// End Get All Trash

// Start Get All
func (suite *userDeliveryTestSuite) TestGetAllSuccess() {
	users := []userDto.User{
		{
			ID: "9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5",
			Username: "admin",
			Password: "$2a$10$4YY9SUvhhVtqURrsbBbDre5MjimjgajD5KsJZg6QkaJ.75jR.6soq",
			Role: "ADMIN",
		},
	}
	suite.userUC.On("GetAll").Return(users, nil)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/users", nil)

	token, _ := utils.GenerateJWT("9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expectedResponse := `{"responseCode":"2000101","responseMessage":"Users retrieved successfully","data":[{"id":"9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5","username":"admin","role":"ADMIN"}]}`
	
	suite.Equal(http.StatusOK, res.Code)
	suite.JSONEq(expectedResponse, res.Body.String())
}

func (suite *userDeliveryTestSuite) TestGetAllErrorUserNotFound() {
	users := []userDto.User{}
	suite.userUC.On("GetAll").Return(users, nil)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/users", nil)

	token, _ := utils.GenerateJWT("9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expectedResponse := `{"responseCode":"4030102","responseMessage":"Users not found"}`

	suite.Equal(http.StatusForbidden, res.Code)
	suite.JSONEq(expectedResponse, res.Body.String())
}

func (suite *userDeliveryTestSuite) TestGetAllInternalServerError() {
	users := []userDto.User{}
	suite.userUC.On("GetAll").Return(users, sql.ErrConnDone)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/users", nil)

	token, _ := utils.GenerateJWT("9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expectedResponse := `{"responseCode":"5000101","responseMessage":"internal server error","error":"sql: connection is already closed"}`

	suite.Equal(http.StatusInternalServerError, res.Code)
	suite.JSONEq(expectedResponse, res.Body.String())
}
// End Get All

// Start Get By ID
func (suite *userDeliveryTestSuite) TestGetByIDSuccess() {
	user := userDto.User{
		ID: "9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5",
		Username: "admin",
		Password: "$2a$10$4YY9SUvhhVtqURrsbBbDre5MjimjgajD5KsJZg6QkaJ.75jR.6soq",
		Role: "ADMIN",
	}
	suite.userUC.On("GetByID", mock.Anything).Return(user, nil)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/users/9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5", nil)

	token, _ := utils.GenerateJWT("9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expectedResponse := `{"responseCode":"2000101","responseMessage":"User retrieved successfully","data":{"id":"9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5","username":"admin","role":"ADMIN"}}`
	
	suite.Equal(http.StatusOK, res.Code)
	suite.JSONEq(expectedResponse, res.Body.String())
}

func (suite *userDeliveryTestSuite) TestGetByIDErrorUserNotFound() {
	user := userDto.User{}
	suite.userUC.On("GetByID", mock.Anything).Return(user, sql.ErrNoRows)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/users/9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5", nil)

	token, _ := utils.GenerateJWT("9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expectedResponse := `{"responseCode":"4030101","responseMessage":"User not found"}`

	suite.Equal(http.StatusForbidden, res.Code)
	suite.JSONEq(expectedResponse, res.Body.String())
}

func (suite *userDeliveryTestSuite) TestGetByIDInternalServerError() {
	user := userDto.User{}
	suite.userUC.On("GetByID", mock.Anything).Return(user, sql.ErrConnDone)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/users/9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5", nil)

	token, _ := utils.GenerateJWT("9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expectedResponse := `{"responseCode":"5000102","responseMessage":"internal server error","error":"sql: connection is already closed"}`

	suite.Equal(http.StatusInternalServerError, res.Code)
	suite.JSONEq(expectedResponse, res.Body.String())
}
// End Get By ID

// Start Patient Register
func (suite *userDeliveryTestSuite) TestPatientRegisterSuccess() {
	requestBody := []byte(`{"username":"user","password":"user"}`)

	user := userDto.User{
		ID: "9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5",
		Username: "user",
		Password: "user",
	}
	suite.userUC.On("PatientRegister", mock.Anything).Return(user, nil)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/users/register", bytes.NewBuffer(requestBody))
	suite.router.ServeHTTP(res, req)

	expectedResponse := `{"responseCode":"2010101","responseMessage":"Patient created successfully.","data":{"id":"9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5","username":"user"}}`

	suite.Equal(http.StatusCreated, res.Code)
	suite.JSONEq(expectedResponse, res.Body.String())
}

func (suite *userDeliveryTestSuite) TestPatientRegisterErrorJSON() {
	requestBody := []byte(`{"username":"user","password":"user",}`)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/users/register", bytes.NewBuffer(requestBody))
	suite.router.ServeHTTP(res, req)

	expectedResponse := `{"responseCode":"5000101","responseMessage":"internal server error","error":"invalid character '}' looking for beginning of object key string"}`

	suite.Equal(http.StatusInternalServerError, res.Code)
	suite.JSONEq(expectedResponse, res.Body.String())
}

func (suite *userDeliveryTestSuite) TestPatientRegisterErrorBadRequest() {
	requestBody := []byte(`{"username":"","password":""}`)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/users/register", bytes.NewBuffer(requestBody))
	suite.router.ServeHTTP(res, req)

	expectedResponse := `{"responseCode":"4000102","responseMessage":"Bad request","error_description":[{"field":"Username","message":"Field is required"},{"field":"Password","message":"Field is required"}]}`

	suite.Equal(http.StatusBadRequest, res.Code)
	suite.JSONEq(expectedResponse, res.Body.String())
}

func (suite *userDeliveryTestSuite) TestPatientRegisterErrorUsernameExists() {
	requestBody := []byte(`{"username":"user","password":"user"}`)

	user := userDto.User{}
	suite.userUC.On("PatientRegister", mock.Anything).Return(user, errors.New("1"))

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/users/register", bytes.NewBuffer(requestBody))
	suite.router.ServeHTTP(res, req)

	expectedResponse := `{"responseCode":"4000103","responseMessage":"Bad request","error_description":[{"field":"username","message":"Username is already registered"}]}`

	suite.Equal(http.StatusBadRequest, res.Code)
	suite.JSONEq(expectedResponse, res.Body.String())
}

func (suite *userDeliveryTestSuite) TestPatientRegisterInternalServerError() {
	requestBody := []byte(`{"username":"user","password":"user"}`)

	user := userDto.User{}
	suite.userUC.On("PatientRegister", mock.Anything).Return(user, sql.ErrConnDone)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/users/register", bytes.NewBuffer(requestBody))
	suite.router.ServeHTTP(res, req)

	expectedResponse := `{"responseCode":"5000104","responseMessage":"internal server error","error":"sql: connection is already closed"}`

	suite.Equal(http.StatusInternalServerError, res.Code)
	suite.JSONEq(expectedResponse, res.Body.String())
}
// End Patient Register

// Start User Register
func (suite *userDeliveryTestSuite) TestUserRegisterSuccess() {
	requestBody := []byte(`{"username":"user","password":"user","role":"PATIENT"}`)

	user := userDto.User{
		ID: "9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5",
		Username: "user",
		Password: "user",
		Role: "PATIENT",
	}
	suite.userUC.On("UserRegister", mock.Anything).Return(user, nil)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(requestBody))
	
	token, _ := utils.GenerateJWT("9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expectedResponse := `{"responseCode":"2010101","responseMessage":"User created successfully.","data":{"id":"9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5","username":"user","role":"PATIENT"}}`

	suite.Equal(http.StatusCreated, res.Code)
	suite.JSONEq(expectedResponse, res.Body.String())
}

func (suite *userDeliveryTestSuite) TestUserRegisterErrorJSON() {
	requestBody := []byte(`{"username":"user","password":"user","role":"PATIENT",}`)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(requestBody))
	
	token, _ := utils.GenerateJWT("9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expectedResponse := `{"responseCode":"5000101","responseMessage":"internal server error","error":"invalid character '}' looking for beginning of object key string"}`

	suite.Equal(http.StatusInternalServerError, res.Code)
	suite.JSONEq(expectedResponse, res.Body.String())
}

func (suite *userDeliveryTestSuite) TestUserRegisterBadRequest() {
	requestBody := []byte(`{"username":"","password":"","role":""}`)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(requestBody))
	
	token, _ := utils.GenerateJWT("9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expectedResponse := `{"responseCode":"4000102","responseMessage":"Bad request","error_description":[{"field":"Username","message":"Field is required"},{"field":"Password","message":"Field is required"},{"field":"Role","message":"Field is required"}]}`

	suite.Equal(http.StatusBadRequest, res.Code)
	suite.JSONEq(expectedResponse, res.Body.String())
}

func (suite *userDeliveryTestSuite) TestUserRegisterErrorUsernameExists() {
	requestBody := []byte(`{"username":"user","password":"user","role":"PATIENT"}`)

	user := userDto.User{}
	suite.userUC.On("UserRegister", mock.Anything).Return(user, errors.New("1"))

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(requestBody))
	
	token, _ := utils.GenerateJWT("9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expectedResponse := `{"responseCode":"4000103","responseMessage":"Bad request","error_description":[{"field":"username","message":"Username is already registered"}]}`

	suite.Equal(http.StatusBadRequest, res.Code)
	suite.JSONEq(expectedResponse, res.Body.String())
}

func (suite *userDeliveryTestSuite) TestUserRegisterErrorSpecialization() {
	requestBody := []byte(`{"username":"user","password":"user","role":"DOCTOR"}`)

	user := userDto.User{}
	suite.userUC.On("UserRegister", mock.Anything).Return(user, errors.New("2"))

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(requestBody))
	
	token, _ := utils.GenerateJWT("9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expectedResponse := `{"responseCode":"4000105","responseMessage":"Bad request","error_description":[{"field":"specialization","message":"field is required"}]}`

	suite.Equal(http.StatusBadRequest, res.Code)
	suite.JSONEq(expectedResponse, res.Body.String())
}

func (suite *userDeliveryTestSuite) TestUserRegisterInternalServerError() {
	requestBody := []byte(`{"username":"user","password":"user","role":"PATIENT"}`)

	user := userDto.User{}
	suite.userUC.On("UserRegister", mock.Anything).Return(user, sql.ErrConnDone)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(requestBody))
	
	token, _ := utils.GenerateJWT("9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expectedResponse := `{"responseCode":"5000106","responseMessage":"internal server error","error":"sql: connection is already closed"}`

	suite.Equal(http.StatusInternalServerError, res.Code)
	suite.JSONEq(expectedResponse, res.Body.String())
}
// End User Register

// Start Login
func (suite *userDeliveryTestSuite) TestLoginSuccess() {
	requestBody := []byte(`{"username":"user","password":"user"}`)

	suite.userUC.On("Login", mock.Anything).Return("token", nil)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/users/login", bytes.NewBuffer(requestBody))
	suite.router.ServeHTTP(res, req)

	expectedResponse := `{"responseCode":"2000101","responseMessage":"Login successfully","data":"token"}`

	suite.Equal(http.StatusOK, res.Code)
	suite.JSONEq(expectedResponse, res.Body.String())
}

func (suite *userDeliveryTestSuite) TestLoginErrorJSON() {
	requestBody := []byte(`{"username":"user","password":"user",}`)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/users/login", bytes.NewBuffer(requestBody))
	suite.router.ServeHTTP(res, req)

	expectedResponse := `{"responseCode":"5000101","responseMessage":"internal server error","error":"invalid character '}' looking for beginning of object key string"}`

	suite.Equal(http.StatusInternalServerError, res.Code)
	suite.JSONEq(expectedResponse, res.Body.String())
}

func (suite *userDeliveryTestSuite) TestLoginBadRequest() {
	requestBody := []byte(`{"username":"","password":""}`)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/users/login", bytes.NewBuffer(requestBody))
	suite.router.ServeHTTP(res, req)

	expectedResponse := `{"responseCode":"4000102","responseMessage":"Bad request","error_description":[{"field":"Username","message":"Field is required"},{"field":"Password","message":"Field is required"}]}`

	suite.Equal(http.StatusBadRequest, res.Code)
	suite.JSONEq(expectedResponse, res.Body.String())
}

func (suite *userDeliveryTestSuite) TestLoginErrorWrongPassword() {
	requestBody := []byte(`{"username":"user","password":"user"}`)

	suite.userUC.On("Login", mock.Anything).Return("", errors.New("1"))

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/users/login", bytes.NewBuffer(requestBody))
	suite.router.ServeHTTP(res, req)

	expectedResponse := `{"responseCode":"5000103","responseMessage":"internal server error","error":"Incorrect username or password"}`

	suite.Equal(http.StatusInternalServerError, res.Code)
	suite.JSONEq(expectedResponse, res.Body.String())
}

func (suite *userDeliveryTestSuite) TestLoginInternalServerError() {
	requestBody := []byte(`{"username":"user","password":"user"}`)

	suite.userUC.On("Login", mock.Anything).Return("token", sql.ErrConnDone)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/users/login", bytes.NewBuffer(requestBody))
	suite.router.ServeHTTP(res, req)

	expectedResponse := `{"responseCode":"5000104","responseMessage":"internal server error","error":"sql: connection is already closed"}`

	suite.Equal(http.StatusInternalServerError, res.Code)
	suite.JSONEq(expectedResponse, res.Body.String())
}
// End Login

// Start Update
func (suite *userDeliveryTestSuite) TestUpdateSuccess() {
	requestBody := []byte(`{"username":"user"}`)
	user := userDto.User{
		ID: "9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5",
		Username: "admin",
		Role: "ADMIN",
	}

	suite.userUC.On("Update", mock.Anything).Return(user, nil)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/users/9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5", bytes.NewBuffer(requestBody))

	token, _ := utils.GenerateJWT("9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+ token)
	suite.router.ServeHTTP(res, req)

	expectedResponse := `{"responseCode":"2000101","responseMessage":"User updeted successfully","data":{"id":"9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5","username":"admin","role":"ADMIN"}}`
	
	suite.Equal(http.StatusOK, res.Code)
	suite.JSONEq(expectedResponse, res.Body.String())
}

func (suite *userDeliveryTestSuite) TestUpdateErrorJSON() {
	requestBody := []byte(`{"username":"user",}`)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/users/9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5", bytes.NewBuffer(requestBody))

	token, _ := utils.GenerateJWT("9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+ token)
	suite.router.ServeHTTP(res, req)

	expectedResponse := `{"responseCode":"5000101","responseMessage":"internal server error","error":"invalid character '}' looking for beginning of object key string"}`
	
	suite.Equal(http.StatusInternalServerError, res.Code)
	suite.JSONEq(expectedResponse, res.Body.String())
}

func (suite *userDeliveryTestSuite) TestUpdateErrorUserNotFound() {
	requestBody := []byte(`{"username":"user"}`)
	user := userDto.User{}

	suite.userUC.On("Update", mock.Anything).Return(user, sql.ErrNoRows)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/users/9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5", bytes.NewBuffer(requestBody))

	token, _ := utils.GenerateJWT("9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+ token)
	suite.router.ServeHTTP(res, req)

	expectedResponse := `{"responseCode":"4030103","responseMessage":"User not found"}`
	
	suite.Equal(http.StatusForbidden, res.Code)
	suite.JSONEq(expectedResponse, res.Body.String())
}

func (suite *userDeliveryTestSuite) TestUpdateErrorUsernameExists() {
	requestBody := []byte(`{"username":"user"}`)
	user := userDto.User{}

	suite.userUC.On("Update", mock.Anything).Return(user, errors.New("1"))

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/users/9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5", bytes.NewBuffer(requestBody))

	token, _ := utils.GenerateJWT("9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+ token)
	suite.router.ServeHTTP(res, req)

	expectedResponse := `{"responseCode":"4000104","responseMessage":"Bad request","error_description":[{"field":"username","message":"Username is already registered"}]}`
	
	suite.Equal(http.StatusBadRequest, res.Code)
	suite.JSONEq(expectedResponse, res.Body.String())
}

func (suite *userDeliveryTestSuite) TestUpdateInternalServerError() {
	requestBody := []byte(`{"username":"user"}`)
	user := userDto.User{}

	suite.userUC.On("Update", mock.Anything).Return(user, sql.ErrConnDone)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/users/9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5", bytes.NewBuffer(requestBody))

	token, _ := utils.GenerateJWT("9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+ token)
	suite.router.ServeHTTP(res, req)

	expectedResponse := `{"responseCode":"5000105","responseMessage":"internal server error","error":"sql: connection is already closed"}`
	
	suite.Equal(http.StatusInternalServerError, res.Code)
	suite.JSONEq(expectedResponse, res.Body.String())
}
// End Update

// Start Update Password
func (suite *userDeliveryTestSuite) TestUpdatePasswordSuccess() {
	request := []byte(`{"current_password":"admin","new_password":"secret","confirmation_password":"secret"}`)

	suite.userUC.On("UpdatePassword", mock.Anything).Return(nil)
	
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/users/9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5/password", bytes.NewBuffer(request))

	token, _ := utils.GenerateJWT("9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expectedResponse := `{"responseCode":"2000101","responseMessage":"Password updated successfully"}`
	
	suite.Equal(http.StatusOK, res.Code)
	suite.Equal(expectedResponse, res.Body.String())
}

func (suite *userDeliveryTestSuite) TestUpdatePasswordErrorJSON() {
	request := []byte(`{"current_password":"admin","new_password":"secret","confirmation_password":"secret",}`)
	
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/users/9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5/password", bytes.NewBuffer(request))

	token, _ := utils.GenerateJWT("9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expectedResponse := `{"responseCode":"5000101","responseMessage":"internal server error","error":"invalid character '}' looking for beginning of object key string"}`
	
	suite.Equal(http.StatusInternalServerError, res.Code)
	suite.Equal(expectedResponse, res.Body.String())
}

func (suite *userDeliveryTestSuite) TestUpdatePasswordBadRequest() {
	request := []byte(`{"current_password":"","new_password":"","confirmation_password":""}`)
	
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/users/9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5/password", bytes.NewBuffer(request))

	token, _ := utils.GenerateJWT("9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expectedResponse := `{"responseCode":"4000102","responseMessage":"Bad request","error_description":[{"field":"CurrentPassword","message":"Field is required"},{"field":"NewPassword","message":"Field is required"},{"field":"ConfirmationPassword","message":"Field is required"}]}`
	
	suite.Equal(http.StatusBadRequest, res.Code)
	suite.Equal(expectedResponse, res.Body.String())
}

func (suite *userDeliveryTestSuite) TestUpdatePasswordErrorUserNotFound() {
	request := []byte(`{"current_password":"admin","new_password":"secret","confirmation_password":"secret"}`)

	suite.userUC.On("UpdatePassword", mock.Anything).Return(sql.ErrNoRows)
	
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/users/9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5/password", bytes.NewBuffer(request))

	token, _ := utils.GenerateJWT("9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expectedResponse := `{"responseCode":"4030103","responseMessage":"User not found"}`
	
	suite.Equal(http.StatusForbidden, res.Code)
	suite.Equal(expectedResponse, res.Body.String())
}

func (suite *userDeliveryTestSuite) TestUpdatePasswordErrorWrongPassword() {
	request := []byte(`{"current_password":"admin1","new_password":"secret","confirmation_password":"secret"}`)

	suite.userUC.On("UpdatePassword", mock.Anything).Return(errors.New("1"))
	
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/users/9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5/password", bytes.NewBuffer(request))

	token, _ := utils.GenerateJWT("9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expectedResponse := `{"responseCode":"4000104","responseMessage":"Bad request","error_description":[{"field":"current_password","message":"Current password is incorrect"}]}`
	
	suite.Equal(http.StatusBadRequest, res.Code)
	suite.Equal(expectedResponse, res.Body.String())
}

func (suite *userDeliveryTestSuite) TestUpdatePasswordErrorConfirmationPassword() {
	request := []byte(`{"current_password":"admin","new_password":"secret1","confirmation_password":"secret"}`)

	suite.userUC.On("UpdatePassword", mock.Anything).Return(errors.New("2"))
	
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/users/9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5/password", bytes.NewBuffer(request))

	token, _ := utils.GenerateJWT("9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expectedResponse := `{"responseCode":"4000105","responseMessage":"Bad request","error_description":[{"field":"new_password","message":"Password do not match"}]}`
	
	suite.Equal(http.StatusBadRequest, res.Code)
	suite.Equal(expectedResponse, res.Body.String())
}

func (suite *userDeliveryTestSuite) TestUpdatePasswordInternalServerError() {
	request := []byte(`{"current_password":"admin","new_password":"secret","confirmation_password":"secret"}`)

	suite.userUC.On("UpdatePassword", mock.Anything).Return(sql.ErrConnDone)
	
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/users/9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5/password", bytes.NewBuffer(request))

	token, _ := utils.GenerateJWT("9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expectedResponse := `{"responseCode":"5000106","responseMessage":"internal server error","error":"sql: connection is already closed"}`
	
	suite.Equal(http.StatusInternalServerError, res.Code)
	suite.Equal(expectedResponse, res.Body.String())
}
// End Update Password

// Start Delete
func (suite *userDeliveryTestSuite) TestDeleteSuccess() {
	suite.userUC.On("Delete", mock.Anything).Return(nil)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/users/9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5", nil)

	token, _ := utils.GenerateJWT("9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expectedResponse := `{"responseCode":"2000101","responseMessage":"User deleted successfully"}`
	
	suite.Equal(http.StatusOK, res.Code)
	suite.Equal(expectedResponse, res.Body.String())
}

func (suite *userDeliveryTestSuite) TestDeleteErrorUserNotFound() {
	suite.userUC.On("Delete", mock.Anything).Return(sql.ErrNoRows)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/users/9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5", nil)

	token, _ := utils.GenerateJWT("9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expectedResponse := `{"responseCode":"4030101","responseMessage":"User not found"}`
	
	suite.Equal(http.StatusForbidden, res.Code)
	suite.Equal(expectedResponse, res.Body.String())
}

func (suite *userDeliveryTestSuite) TestDeleteInternalServerError() {
	suite.userUC.On("Delete", mock.Anything).Return(sql.ErrConnDone)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/users/9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5", nil)

	token, _ := utils.GenerateJWT("9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expectedResponse := `{"responseCode":"5000102","responseMessage":"internal server error","error":"sql: connection is already closed"}`
	
	suite.Equal(http.StatusInternalServerError, res.Code)
	suite.Equal(expectedResponse, res.Body.String())
}
// End Delete

// Start Soft Delete
func (suite *userDeliveryTestSuite) TestSoftDeleteSuccess() {
	suite.userUC.On("SoftDelete", mock.Anything).Return(nil)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/users/9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5/trash", nil)

	token, _ := utils.GenerateJWT("9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expectedResponse := `{"responseCode":"2000101","responseMessage":"User deleted successfully"}`
	
	suite.Equal(http.StatusOK, res.Code)
	suite.Equal(expectedResponse, res.Body.String())
}

func (suite *userDeliveryTestSuite) TestSoftDeleteErrorUserNotFound() {
	suite.userUC.On("SoftDelete", mock.Anything).Return(sql.ErrNoRows)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/users/9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5/trash", nil)

	token, _ := utils.GenerateJWT("9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expectedResponse := `{"responseCode":"4030101","responseMessage":"User not found"}`
	
	suite.Equal(http.StatusForbidden, res.Code)
	suite.Equal(expectedResponse, res.Body.String())
}

func (suite *userDeliveryTestSuite) TestSoftDeleteInternalServerError() {
	suite.userUC.On("SoftDelete", mock.Anything).Return(sql.ErrConnDone)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/users/9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5/trash", nil)

	token, _ := utils.GenerateJWT("9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expectedResponse := `{"responseCode":"5000102","responseMessage":"internal server error","error":"sql: connection is already closed"}`
	
	suite.Equal(http.StatusInternalServerError, res.Code)
	suite.Equal(expectedResponse, res.Body.String())
}
// End Soft Delete

// Start Restore
func (suite *userDeliveryTestSuite) TestRestoreSuccess() {
	suite.userUC.On("Restore", mock.Anything).Return(nil)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/users/9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5/restore", nil)

	token, _ := utils.GenerateJWT("9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expectedResponse := `{"responseCode":"2000101","responseMessage":"User restored successfully"}`
	
	suite.Equal(http.StatusOK, res.Code)
	suite.Equal(expectedResponse, res.Body.String())
}

func (suite *userDeliveryTestSuite) TestRestoreErrorUserNotFound() {
	suite.userUC.On("Restore", mock.Anything).Return(sql.ErrNoRows)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/users/9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5/restore", nil)

	token, _ := utils.GenerateJWT("9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expectedResponse := `{"responseCode":"4030101","responseMessage":"User not found"}`
	
	suite.Equal(http.StatusForbidden, res.Code)
	suite.Equal(expectedResponse, res.Body.String())
}

func (suite *userDeliveryTestSuite) TestRestoreInternalServerError() {
	suite.userUC.On("Restore", mock.Anything).Return(sql.ErrConnDone)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/users/9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5/restore", nil)

	token, _ := utils.GenerateJWT("9d3cd7b1-ade2-4f8c-b215-9e74f0c87bf5", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expectedResponse := `{"responseCode":"5000102","responseMessage":"internal server error","error":"sql: connection is already closed"}`
	
	suite.Equal(http.StatusInternalServerError, res.Code)
	suite.Equal(expectedResponse, res.Body.String())
}
// End Restore

func (suite *userDeliveryTestSuite) SetupSuite() {
	gin.SetMode(gin.TestMode)
}

func TestUserDeliveryTestSuite(t *testing.T)  {
	suite.Run(t, new(userDeliveryTestSuite))}