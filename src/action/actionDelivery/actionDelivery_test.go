package actionDelivery

import (
	"avengers-clinic/model/dto/actionDto"
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

type mockActionUsecase struct {
	mock.Mock
}

func (mock *mockActionUsecase) GetAll() ([]actionDto.Action, error) {
	args := mock.Called()
	return args.Get(0).([]actionDto.Action), args.Error(1)
}

func (mock *mockActionUsecase) GetByID(actionID string) (actionDto.Action, error) {
	args := mock.Called(actionID)
	return args.Get(0).(actionDto.Action), args.Error(1)
}

func (mock *mockActionUsecase) Create(req actionDto.CreateRequest) (actionDto.Action, error) {
	args := mock.Called(req)
	return args.Get(0).(actionDto.Action), args.Error(1)
}

func (mock *mockActionUsecase) Update(req actionDto.UpdateRequest) (actionDto.Action, error) {
	args := mock.Called(req)
	return args.Get(0).(actionDto.Action), args.Error(1)
}

func (mock *mockActionUsecase) Delete(actionID string) error {
	args := mock.Called(actionID)
	return args.Error(0)
}

func (mock *mockActionUsecase) SoftDelete(actionID string) error {
	args := mock.Called(actionID)
	return args.Error(0)
}

func (mock *mockActionUsecase) Restore(actionID string) error {
	args := mock.Called(actionID)
	return args.Error(0)
}

type actionDeliveryTestSuite struct {
	suite.Suite
	router *gin.Engine
	actionUC *mockActionUsecase
}

func (suite *actionDeliveryTestSuite) SetupTest() {
	suite.router = gin.New()
	suite.actionUC = new(mockActionUsecase)
	
	v1Group := suite.router.Group("/api/v1")
	NewActionDelivery(v1Group, suite.actionUC)
}

// Start Get All
func (suite *actionDeliveryTestSuite) TestGetAllSuccess() {
	actions := []actionDto.Action{{ID: "1", Name: "Konsultasi", Price: 20000}}

	suite.actionUC.On("GetAll").Return(actions, nil)
	
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/actions", nil)

	token, _ := utils.GenerateJWT("1", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expected := `{"responseCode":"2000201","responseMessage":"actions successfully retrieved","data":[{"id":"1","name":"Konsultasi","price":20000}]}`

	suite.Equal(http.StatusOK, res.Code)
	suite.JSONEq(expected, res.Body.String())
}

func (suite *actionDeliveryTestSuite) TestGetAllNotFound() {
	actions := []actionDto.Action{}

	suite.actionUC.On("GetAll").Return(actions, nil)
	
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/actions", nil)

	token, _ := utils.GenerateJWT("1", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expected := `{"responseCode":"4030202","responseMessage":"Actions not found"}`
	
	suite.Equal(http.StatusForbidden, res.Code)
	suite.JSONEq(expected, res.Body.String())
}

func (suite *actionDeliveryTestSuite) TestGetAllInternalServerError() {
	actions := []actionDto.Action{}

	suite.actionUC.On("GetAll").Return(actions, sql.ErrConnDone)
	
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/actions", nil)

	token, _ := utils.GenerateJWT("1", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expected := `{"responseCode":"5000201","responseMessage":"internal server error","error":"sql: connection is already closed"}`
	
	suite.Equal(http.StatusInternalServerError, res.Code)
	suite.JSONEq(expected, res.Body.String())
}
// End Get All

// Start Get By ID
func (suite *actionDeliveryTestSuite) TestGetByIDSuccess() {
	action := actionDto.Action{ID: "1", Name: "Konsultasi", Price: 20000}

	suite.actionUC.On("GetByID", mock.Anything).Return(action, nil)
	
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/actions/1", nil)

	token, _ := utils.GenerateJWT("1", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expected := `{"responseCode":"2000201","responseMessage":"action successfully retrieved","data":{"id":"1","name":"Konsultasi","price":20000}}`

	suite.Equal(http.StatusOK, res.Code)
	suite.JSONEq(expected, res.Body.String())
}

func (suite *actionDeliveryTestSuite) TestGetByIDNotFound() {
	action := actionDto.Action{}

	suite.actionUC.On("GetByID", mock.Anything).Return(action, sql.ErrNoRows)
	
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/actions/1", nil)

	token, _ := utils.GenerateJWT("1", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expected := `{"responseCode":"4030201","responseMessage":"Action not found"}`
	
	suite.Equal(http.StatusForbidden, res.Code)
	suite.JSONEq(expected, res.Body.String())
}

func (suite *actionDeliveryTestSuite) TestGetByIDInternalServerError() {
	actions := actionDto.Action{}

	suite.actionUC.On("GetByID", mock.Anything).Return(actions, sql.ErrConnDone)
	
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/actions/1", nil)

	token, _ := utils.GenerateJWT("1", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expected := `{"responseCode":"5000202","responseMessage":"internal server error","error":"sql: connection is already closed"}`
	
	suite.Equal(http.StatusInternalServerError, res.Code)
	suite.JSONEq(expected, res.Body.String())
}
// End Get By ID

// Start Create
func (suite *actionDeliveryTestSuite) TestCreateSuccess() {
	requestBody := []byte(`{"name":"Konsultasi","price":20000}`)
	action := actionDto.Action{ID: "1", Name: "Konsultasi", Price: 20000}

	suite.actionUC.On("Create", mock.Anything).Return(action, nil)
	
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/actions", bytes.NewBuffer(requestBody))

	token, _ := utils.GenerateJWT("1", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expected := `{"responseCode":"2010201","responseMessage":"Action created successfully","data":{"id":"1","name":"Konsultasi","price":20000}}`
	
	suite.Equal(http.StatusCreated, res.Code)
	suite.JSONEq(expected, res.Body.String())
}

func (suite *actionDeliveryTestSuite) TestCreateErrorJSON() {
	requestBody := []byte(`{"name":"Konsultasi","price":20000,}`)
	
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/actions", bytes.NewBuffer(requestBody))

	token, _ := utils.GenerateJWT("1", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expected := `{"responseCode":"5000201","responseMessage":"internal server error","error":"invalid character '}' looking for beginning of object key string"}`
	
	suite.Equal(http.StatusInternalServerError, res.Code)
	suite.JSONEq(expected, res.Body.String())
}

func (suite *actionDeliveryTestSuite) TestCreateErrorBadRequest() {
	requestBody := []byte(`{"name":"","price":0}`)
	
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/actions", bytes.NewBuffer(requestBody))

	token, _ := utils.GenerateJWT("1", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expected := `{"responseCode":"4000202","responseMessage":"Bad request","error_description":[{"field":"Name","message":"Field is required"},{"field":"Price","message":"Field is required"}]}`
	
	suite.Equal(http.StatusBadRequest, res.Code)
	suite.JSONEq(expected, res.Body.String())
}

func (suite *actionDeliveryTestSuite) TestCreateErrorNameExist() {
	requestBody := []byte(`{"name":"Konsultasi","price":20000}`)
	action := actionDto.Action{}

	suite.actionUC.On("Create", mock.Anything).Return(action, errors.New("1"))
	
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/actions", bytes.NewBuffer(requestBody))

	token, _ := utils.GenerateJWT("1", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expected := `{"responseCode":"4000203","responseMessage":"Bad request","error_description":[{"field":"name","message":"Name is already registered"}]}`
	
	suite.Equal(http.StatusBadRequest, res.Code)
	suite.JSONEq(expected, res.Body.String())
}

func (suite *actionDeliveryTestSuite) TestCreateErrorInternalServerError() {
	requestBody := []byte(`{"name":"Konsultasi","price":20000}`)
	action := actionDto.Action{}

	suite.actionUC.On("Create", mock.Anything).Return(action, sql.ErrConnDone)
	
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/actions", bytes.NewBuffer(requestBody))

	token, _ := utils.GenerateJWT("1", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expected := `{"responseCode":"5000204","responseMessage":"internal server error","error":"sql: connection is already closed"}`

	suite.Equal(http.StatusInternalServerError, res.Code)
	suite.JSONEq(expected, res.Body.String())
}
// End Create

// Start Update
func (suite *actionDeliveryTestSuite) TestUpdateSuccess() {
	requestBody := []byte(`{"name":"Konsultasi"}`)
	action := actionDto.Action{ID: "1", Name: "Konsultasi", Price: 20000}

	suite.actionUC.On("Update", mock.Anything).Return(action, nil)
	
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/actions/1", bytes.NewBuffer(requestBody))

	token, _ := utils.GenerateJWT("1", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expected := `{"responseCode":"2000201","responseMessage":"Action updated successfully","data":{"id":"1","name":"Konsultasi","price":20000}}`
	
	suite.Equal(http.StatusOK, res.Code)
	suite.JSONEq(expected, res.Body.String())
}

func (suite *actionDeliveryTestSuite) TestUpdateErrorJSON() {
	requestBody := []byte(`{"name":"Konsultasi",}`)
	
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/actions/1", bytes.NewBuffer(requestBody))

	token, _ := utils.GenerateJWT("1", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expected := `{"responseCode":"5000201","responseMessage":"internal server error","error":"invalid character '}' looking for beginning of object key string"}`
	
	suite.Equal(http.StatusInternalServerError, res.Code)
	suite.JSONEq(expected, res.Body.String())
}

func (suite *actionDeliveryTestSuite) TestUpdateErrorActionNotFound() {
	requestBody := []byte(`{"name":"Konsultasi"}`)
	action := actionDto.Action{}

	suite.actionUC.On("Update", mock.Anything).Return(action, sql.ErrNoRows)
	
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/actions/1", bytes.NewBuffer(requestBody))

	token, _ := utils.GenerateJWT("1", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expected := `{"responseCode":"4030202","responseMessage":"Action not found"}`
	fmt.Println(res.Body.String())
	suite.Equal(http.StatusForbidden, res.Code)
	suite.JSONEq(expected, res.Body.String())
}

func (suite *actionDeliveryTestSuite) TestUpdateErrorNameExist() {
	requestBody := []byte(`{"name":"Konsultasi"}`)
	action := actionDto.Action{}

	suite.actionUC.On("Update", mock.Anything).Return(action, errors.New("1"))
	
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/actions/1", bytes.NewBuffer(requestBody))

	token, _ := utils.GenerateJWT("1", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expected := `{"responseCode":"4000203","responseMessage":"Bad request","error_description":[{"field":"name","message":"Name is already registered"}]}`
	
	suite.Equal(http.StatusBadRequest, res.Code)
	suite.JSONEq(expected, res.Body.String())
}

func (suite *actionDeliveryTestSuite) TestUpdateErrorInternalServerError() {
	requestBody := []byte(`{"name":"Konsultasi"}`)
	action := actionDto.Action{}

	suite.actionUC.On("Update", mock.Anything).Return(action, sql.ErrConnDone)
	
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/actions/1", bytes.NewBuffer(requestBody))

	token, _ := utils.GenerateJWT("1", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expected := `{"responseCode":"5000204","responseMessage":"internal server error","error":"sql: connection is already closed"}`
	
	suite.Equal(http.StatusInternalServerError, res.Code)
	suite.JSONEq(expected, res.Body.String())
}
// End Update

// Start Delete
func (suite *actionDeliveryTestSuite) TestDeleteSuccess() {
	suite.actionUC.On("Delete", mock.Anything).Return(nil)
	
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/actions/1", nil)

	token, _ := utils.GenerateJWT("1", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expected := `{"responseCode":"2000201","responseMessage":"Action deleted successfully"}`
	
	suite.Equal(http.StatusOK, res.Code)
	suite.JSONEq(expected, res.Body.String())
}

func (suite *actionDeliveryTestSuite) TestDeleteErrorActionNotFound() {
	suite.actionUC.On("Delete", mock.Anything).Return(sql.ErrNoRows)
	
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/actions/1", nil)

	token, _ := utils.GenerateJWT("1", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expected := `{"responseCode":"4030201","responseMessage":"Action not found"}`
	
	suite.Equal(http.StatusForbidden, res.Code)
	suite.JSONEq(expected, res.Body.String())
}

func (suite *actionDeliveryTestSuite) TestDeleteInternalServerError() {
	suite.actionUC.On("Delete", mock.Anything).Return(sql.ErrConnDone)
	
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/actions/1", nil)

	token, _ := utils.GenerateJWT("1", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expected := `{"responseCode":"5000202","responseMessage":"internal server error","error":"sql: connection is already closed"}`
	
	suite.Equal(http.StatusInternalServerError, res.Code)
	suite.JSONEq(expected, res.Body.String())
}
// End Delete

// Start Soft Delete
func (suite *actionDeliveryTestSuite) TestSoftDeleteSuccess() {
	suite.actionUC.On("SoftDelete", mock.Anything).Return(nil)
	
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/actions/1/trash", nil)

	token, _ := utils.GenerateJWT("1", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expected := `{"responseCode":"2000201","responseMessage":"Action deleted successfully"}`
	
	suite.Equal(http.StatusOK, res.Code)
	suite.JSONEq(expected, res.Body.String())
}

func (suite *actionDeliveryTestSuite) TestSoftDeleteErrorActionNotFound() {
	suite.actionUC.On("SoftDelete", mock.Anything).Return(sql.ErrNoRows)
	
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/actions/1/trash", nil)

	token, _ := utils.GenerateJWT("1", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expected := `{"responseCode":"4030201","responseMessage":"Action not found"}`
	
	suite.Equal(http.StatusForbidden, res.Code)
	suite.JSONEq(expected, res.Body.String())
}

func (suite *actionDeliveryTestSuite) TestSoftDeleteInternalServerError() {
	suite.actionUC.On("SoftDelete", mock.Anything).Return(sql.ErrConnDone)
	
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/actions/1/trash", nil)

	token, _ := utils.GenerateJWT("1", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expected := `{"responseCode":"5000202","responseMessage":"internal server error","error":"sql: connection is already closed"}`
	
	suite.Equal(http.StatusInternalServerError, res.Code)
	suite.JSONEq(expected, res.Body.String())
}
// End Soft Delete

// Start Restore
func (suite *actionDeliveryTestSuite) TestRestoreSuccess() {
	suite.actionUC.On("Restore", mock.Anything).Return(nil)
	
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/actions/1/restore", nil)

	token, _ := utils.GenerateJWT("1", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expected := `{"responseCode":"2000201","responseMessage":"Action restored successfully"}`
	
	suite.Equal(http.StatusOK, res.Code)
	suite.JSONEq(expected, res.Body.String())
}

func (suite *actionDeliveryTestSuite) TestRestoreErrorActionNotFound() {
	suite.actionUC.On("Restore", mock.Anything).Return(sql.ErrNoRows)
	
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/actions/1/restore", nil)

	token, _ := utils.GenerateJWT("1", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expected := `{"responseCode":"4030201","responseMessage":"Action not found"}`
	
	suite.Equal(http.StatusForbidden, res.Code)
	suite.JSONEq(expected, res.Body.String())
}

func (suite *actionDeliveryTestSuite) TestRestoreInternalServerError() {
	suite.actionUC.On("Restore", mock.Anything).Return(sql.ErrConnDone)
	
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/actions/1/restore", nil)

	token, _ := utils.GenerateJWT("1", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expected := `{"responseCode":"5000202","responseMessage":"internal server error","error":"sql: connection is already closed"}`
	
	suite.Equal(http.StatusInternalServerError, res.Code)
	suite.JSONEq(expected, res.Body.String())
}
// End Restore

func (suite *actionDeliveryTestSuite) SetupSuite() {
	gin.SetMode(gin.TestMode)
}

func TestActionDeliveryTestSuite(t *testing.T) {
	suite.Run(t, new(actionDeliveryTestSuite))
}