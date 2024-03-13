package medicineDelivery

import (
	"avengers-clinic/model/dto"
	"avengers-clinic/pkg/utils"
	"bytes"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type mockMedicineUsecase struct {
	mock.Mock
}

func (mock *mockMedicineUsecase) GetAll() ([]dto.MedicineResponse, error) {
	args := mock.Called()
	return args.Get(0).([]dto.MedicineResponse), args.Error(1)
}

func (mock *mockMedicineUsecase) GetById(id string) (dto.MedicineResponse, error) {
	args := mock.Called(id)
	return args.Get(0).(dto.MedicineResponse), args.Error(1)
}

func (mock *mockMedicineUsecase) CreateRecord(medicine dto.MedicineRequest) (dto.MedicineResponse, error) {
	args := mock.Called(medicine)
	return args.Get(0).(dto.MedicineResponse), args.Error(1)
}

func (mock *mockMedicineUsecase) UpdateRecord(medicine dto.UpdateRequest) (dto.MedicineResponse, error) {
	args := mock.Called(medicine)
	return args.Get(0).(dto.MedicineResponse), args.Error(1)
}

func (mock *mockMedicineUsecase) DeleteRecord(id string) error {
	args := mock.Called(id)
	return args.Error(0)
}

func (mock *mockMedicineUsecase) TrashRecord() ([]dto.MedicineResponse, error) {
	args := mock.Called()
	return args.Get(0).([]dto.MedicineResponse), args.Error(1)
}

func (mock *mockMedicineUsecase) RestoreRecord(id string) error {
	args := mock.Called(id)
	return args.Error(0)
}

type medicineDeliveryTestSuite struct {
	suite.Suite
	medicineUC *mockMedicineUsecase
	router *gin.Engine
}

func (suite *medicineDeliveryTestSuite) SetupTest() {
	suite.router = gin.New()
	suite.medicineUC = new(mockMedicineUsecase)

	v1Group := suite.router.Group("/api/v1")
	NewMedicineDelivery(v1Group, suite.medicineUC)
}

// Start Get All
func (suite *medicineDeliveryTestSuite) TestGetAllSuccess() {
	medicines := []dto.MedicineResponse{{Id: "1", Name: "Komik", MedicineType: "CAIR"}}

	suite.medicineUC.On("GetAll").Return(medicines, nil)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/medicines", nil)

	token, _ := utils.GenerateJWT("1", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expected := `{"responseCode":"2000301","responseMessage":"success","data":[{"id":"1","name":"Komik","medicine_type":"CAIR","price":0}]}`

	suite.Equal(http.StatusOK, res.Code)
	suite.JSONEq(expected, res.Body.String())
}

func (suite *medicineDeliveryTestSuite) TestGetAllErrorNotFound() {
	medicines := []dto.MedicineResponse{}

	suite.medicineUC.On("GetAll").Return(medicines, nil)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/medicines", nil)

	token, _ := utils.GenerateJWT("1", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expected := `{"responseCode":"4030301","responseMessage":"medicines not found"}`
	
	suite.Equal(http.StatusForbidden, res.Code)
	suite.JSONEq(expected, res.Body.String())
}

func (suite *medicineDeliveryTestSuite) TestGetAllInternalServerError() {
	medicines := []dto.MedicineResponse{}

	suite.medicineUC.On("GetAll").Return(medicines, sql.ErrConnDone)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/medicines", nil)

	token, _ := utils.GenerateJWT("1", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expected := `{"responseCode":"5000301","responseMessage":"internal server error","error":"sql: connection is already closed"}`
	
	suite.Equal(http.StatusInternalServerError, res.Code)
	suite.JSONEq(expected, res.Body.String())
}
// End Get All

// Start Get By Id
func (suite *medicineDeliveryTestSuite) TestGetByIdSuccess() {
	medicine := dto.MedicineResponse{Id: "1", Name: "Komik", MedicineType: "CAIR"}

	suite.medicineUC.On("GetById", mock.Anything).Return(medicine, nil)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/medicines/1", nil)

	token, _ := utils.GenerateJWT("1", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expected := `{"responseCode":"2000301","responseMessage":"success","data":{"id":"1","name":"Komik","medicine_type":"CAIR","price":0}}`

	suite.Equal(http.StatusOK, res.Code)
	suite.JSONEq(expected, res.Body.String())
}

func (suite *medicineDeliveryTestSuite) TestGetByIdErrorNotFound() {
	medicine := dto.MedicineResponse{}

	suite.medicineUC.On("GetById", mock.Anything).Return(medicine, sql.ErrNoRows)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/medicines/1", nil)

	token, _ := utils.GenerateJWT("1", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expected := `{"responseCode":"4030301","responseMessage":"medicine not found"}`
	
	suite.Equal(http.StatusForbidden, res.Code)
	suite.JSONEq(expected, res.Body.String())
}

func (suite *medicineDeliveryTestSuite) TestGetByIdInternalServerError() {
	medicine := dto.MedicineResponse{}

	suite.medicineUC.On("GetById", mock.Anything).Return(medicine, sql.ErrConnDone)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/medicines/1", nil)

	token, _ := utils.GenerateJWT("1", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expected := `{"responseCode":"5000301","responseMessage":"internal server error","error":"sql: connection is already closed"}`
	
	suite.Equal(http.StatusInternalServerError, res.Code)
	suite.JSONEq(expected, res.Body.String())
}
// End Get By Id

// Start Create
func (suite *medicineDeliveryTestSuite) TestCreateSuccess() {
	requestBody := []byte(`{"name":"komik","medicine_type":"CAIR","price":5000,"stock":200}`)

	medicine := dto.MedicineResponse{Id: "1", Name: "Komik", MedicineType: "CAIR", Price: 5000, Stock: 200}

	suite.medicineUC.On("CreateRecord", mock.Anything).Return(medicine, nil)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/medicines", bytes.NewBuffer(requestBody))

	token, _ := utils.GenerateJWT("1", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expected := `{"responseCode":"2010301","responseMessage":"succesfully insert new medicine","data":{"id":"1","name":"Komik","medicine_type":"CAIR","price":5000,"stock":200}}`
	
	suite.Equal(http.StatusCreated, res.Code)
	suite.JSONEq(expected, res.Body.String())
}

func (suite *medicineDeliveryTestSuite) TestCreateErrorJSON() {
	requestBody := []byte(`{"name":"komik","medicine_type":"CAIR","price":5000,"stock":200,}`)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/medicines", bytes.NewBuffer(requestBody))

	token, _ := utils.GenerateJWT("1", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expected := `{"responseCode":"5000301","responseMessage":"internal server error","error":"invalid character '}' looking for beginning of object key string"}`
	
	suite.Equal(http.StatusInternalServerError, res.Code)
	suite.JSONEq(expected, res.Body.String())
}

func (suite *medicineDeliveryTestSuite) TestCreateErrorBadRequest() {
	requestBody := []byte(`{}`)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/medicines", bytes.NewBuffer(requestBody))

	token, _ := utils.GenerateJWT("1", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expected := `{"responseCode":"4000301","responseMessage":"bad request","error_description":[{"field":"Name","message":"Field is required"},{"field":"MedicineType","message":"Field is required"},{"field":"Price","message":"Field is required"}]}`
	
	suite.Equal(http.StatusBadRequest, res.Code)
	suite.JSONEq(expected, res.Body.String())
}

func (suite *medicineDeliveryTestSuite) TestCreateInternalServerError() {
	requestBody := []byte(`{"name":"komik","medicine_type":"CAIR","price":5000,"stock":200}`)

	medicine := dto.MedicineResponse{}

	suite.medicineUC.On("CreateRecord", mock.Anything).Return(medicine, sql.ErrConnDone)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/medicines", bytes.NewBuffer(requestBody))

	token, _ := utils.GenerateJWT("1", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expected := `{"responseCode":"5000301","responseMessage":"internal server error","error":"sql: connection is already closed"}`
	
	suite.Equal(http.StatusInternalServerError, res.Code)
	suite.JSONEq(expected, res.Body.String())
}
// End Create

// Start Update
func (suite *medicineDeliveryTestSuite) TestUpdateSuccess() {
	requestBody := []byte(`{"name":"komik"}`)

	medicine := dto.MedicineResponse{Id: "1", Name: "Komik", MedicineType: "CAIR", Price: 5000, Stock: 200}

	suite.medicineUC.On("UpdateRecord", mock.Anything).Return(medicine, nil)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/medicines/1", bytes.NewBuffer(requestBody))

	token, _ := utils.GenerateJWT("1", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expected := `{"responseCode":"2000301","responseMessage":"success update medicine","data":{"id":"1","name":"Komik","medicine_type":"CAIR","price":5000,"stock":200}}`
	
	suite.Equal(http.StatusOK, res.Code)
	suite.JSONEq(expected, res.Body.String())
}

func (suite *medicineDeliveryTestSuite) TestUpdateErrorJSON() {
	requestBody := []byte(`{"name":"komik",}`)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/medicines/1", bytes.NewBuffer(requestBody))

	token, _ := utils.GenerateJWT("1", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expected := `{"responseCode":"5000301","responseMessage":"internal server error","error":"invalid character '}' looking for beginning of object key string"}`
	
	suite.Equal(http.StatusInternalServerError, res.Code)
	suite.JSONEq(expected, res.Body.String())
}

func (suite *medicineDeliveryTestSuite) TestUpdateErrorNotFound() {
	requestBody := []byte(`{"name":"komik"}`)

	medicine := dto.MedicineResponse{}

	suite.medicineUC.On("UpdateRecord", mock.Anything).Return(medicine, sql.ErrNoRows)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/medicines/1", bytes.NewBuffer(requestBody))

	token, _ := utils.GenerateJWT("1", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expected := `{"responseCode":"4030301","responseMessage":"medicine not found"}`
	
	suite.Equal(http.StatusForbidden, res.Code)
	suite.JSONEq(expected, res.Body.String())
}

func (suite *medicineDeliveryTestSuite) TestUpdateInternalServerError() {
	requestBody := []byte(`{"name":"komik"}`)

	medicine := dto.MedicineResponse{}

	suite.medicineUC.On("UpdateRecord", mock.Anything).Return(medicine, sql.ErrConnDone)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/medicines/1", bytes.NewBuffer(requestBody))

	token, _ := utils.GenerateJWT("1", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expected := `{"responseCode":"5000301","responseMessage":"internal server error","error":"sql: connection is already closed"}`
	
	suite.Equal(http.StatusInternalServerError, res.Code)
	suite.JSONEq(expected, res.Body.String())
}
// End Update

// Start Delete
func (suite *medicineDeliveryTestSuite) TestDeleteSuccess() {
	suite.medicineUC.On("DeleteRecord", mock.Anything).Return(nil)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/medicines/1", nil)

	token, _ := utils.GenerateJWT("1", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expected := `{"responseCode":"2000301","responseMessage":"success delete medicine","data":"1"}`
	
	suite.Equal(http.StatusOK, res.Code)
	suite.JSONEq(expected, res.Body.String())
}

func (suite *medicineDeliveryTestSuite) TestDeleteErrorNotFound() {
	suite.medicineUC.On("DeleteRecord", mock.Anything).Return(sql.ErrNoRows)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/medicines/1", nil)

	token, _ := utils.GenerateJWT("1", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expected := `{"responseCode":"4030301","responseMessage":"medicine not found"}`
	
	suite.Equal(http.StatusForbidden, res.Code)
	suite.JSONEq(expected, res.Body.String())
}

func (suite *medicineDeliveryTestSuite) TestDeleteInternalServerError() {
	suite.medicineUC.On("DeleteRecord", mock.Anything).Return(sql.ErrConnDone)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/medicines/1", nil)

	token, _ := utils.GenerateJWT("1", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expected := `{"responseCode":"5000301","responseMessage":"internal server error","error":"sql: connection is already closed"}`
	
	suite.Equal(http.StatusInternalServerError, res.Code)
	suite.JSONEq(expected, res.Body.String())
}
// End Delete

// Start Trash
func (suite *medicineDeliveryTestSuite) TestTrashSuccess() {
	medicines := []dto.MedicineResponse{{Id: "1", Name: "Komik", MedicineType: "CAIR"}}

	suite.medicineUC.On("TrashRecord").Return(medicines, nil)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/medicines/trash", nil)

	token, _ := utils.GenerateJWT("1", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expected := `{"responseCode":"2000301","responseMessage":"success","data":[{"id":"1","name":"Komik","medicine_type":"CAIR","price":0}]}`
	
	suite.Equal(http.StatusOK, res.Code)
	suite.JSONEq(expected, res.Body.String())
}

func (suite *medicineDeliveryTestSuite) TestTrashErrorNotFound() {
	medicines := []dto.MedicineResponse{}

	suite.medicineUC.On("TrashRecord").Return(medicines, nil)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/medicines/trash", nil)

	token, _ := utils.GenerateJWT("1", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expected := `{"responseCode":"4030301","responseMessage":"medicines not found"}`
	
	suite.Equal(http.StatusForbidden, res.Code)
	suite.JSONEq(expected, res.Body.String())
}

func (suite *medicineDeliveryTestSuite) TestTrashInternalServerError() {
	medicines := []dto.MedicineResponse{}

	suite.medicineUC.On("TrashRecord").Return(medicines, sql.ErrConnDone)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/medicines/trash", nil)

	token, _ := utils.GenerateJWT("1", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expected := `{"responseCode":"5000301","responseMessage":"internal server error","error":"sql: connection is already closed"}`
	
	suite.Equal(http.StatusInternalServerError, res.Code)
	suite.JSONEq(expected, res.Body.String())
}
// End Trash

// Start Restore
func (suite *medicineDeliveryTestSuite) TestRestoreSuccess() {
	suite.medicineUC.On("RestoreRecord", mock.Anything).Return(nil)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/medicines/1/restore", nil)

	token, _ := utils.GenerateJWT("1", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expected := `{"responseCode":"2000301","responseMessage":"success restore medicine","data":"1"}`
	
	suite.Equal(http.StatusOK, res.Code)
	suite.JSONEq(expected, res.Body.String())
}

func (suite *medicineDeliveryTestSuite) TestRestoreInternalServerError() {
	suite.medicineUC.On("RestoreRecord", mock.Anything).Return(sql.ErrConnDone)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/medicines/1/restore", nil)

	token, _ := utils.GenerateJWT("1", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expected := `{"responseCode":"5000301","responseMessage":"internal server error","error":"sql: connection is already closed"}`
	
	suite.Equal(http.StatusInternalServerError, res.Code)
	suite.JSONEq(expected, res.Body.String())
}
// End Restore

func (suite *medicineDeliveryTestSuite) SetupSuite() {
	gin.SetMode(gin.TestMode)
}

func TestMedicineDeliveryTestSuite(t *testing.T) {
	suite.Run(t, new(medicineDeliveryTestSuite))
}