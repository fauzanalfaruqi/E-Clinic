package medicalRecordDelivery

import (
	myjson "avengers-clinic/model/dto/json"
	"avengers-clinic/model/dto/medicalRecordDTO"
	"avengers-clinic/pkg/constants"
	"avengers-clinic/pkg/utils"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type mockMedicalRecordUsecase struct {
	mock.Mock
}

func (m *mockMedicalRecordUsecase) CreateMedicalRecord(mr medicalRecordDTO.Medical_Record_Request) (medicalRecordDTO.Medical_Record, error) {
	args := m.Called(mr)
	return args.Get(0).(medicalRecordDTO.Medical_Record), args.Error(1)
}

func (m *mockMedicalRecordUsecase) GetMedicalRecords() ([]medicalRecordDTO.Medical_Record, error) {
	args := m.Called()
	return args.Get(0).([]medicalRecordDTO.Medical_Record), args.Error(1)
}

func (m *mockMedicalRecordUsecase) GetMedicalRecordByID(id string) (medicalRecordDTO.Medical_Record, error) {
	args := m.Called(id)
	return args.Get(0).(medicalRecordDTO.Medical_Record), args.Error(1)
}

func (m *mockMedicalRecordUsecase) UpdatePaymentStatus(id string) (medicalRecordDTO.Medical_Record, error) {
	args := m.Called(id)
	return args.Get(0).(medicalRecordDTO.Medical_Record), args.Error(1)
}

type MedicalRecordDeliverySuite struct {
	suite.Suite
	router              *gin.Engine
	medicalRecordUCMock *mockMedicalRecordUsecase
}

func (suite *MedicalRecordDeliverySuite) SetupTest() {
	suite.router = gin.Default()
	suite.medicalRecordUCMock = new(mockMedicalRecordUsecase)

	v1Group := suite.router.Group("/api/v1")
	NewMedicalRecordDelivery(v1Group, suite.medicalRecordUCMock)
}

func (suite *MedicalRecordDeliverySuite) TestCreateMedicalRecord_Success() {
	requestPayload := medicalRecordDTO.Medical_Record_Request{
		Booking_ID:       "ea1c7e2c-3799-4ef7-a8e7-4ecf6c413151",
		Diagnosis_Result: "Test diagnosis",
		Payment_Status:   true,
		Medicine_Details: []medicalRecordDTO.Medicine_Details_Request{
			{
				Medicine_ID: "5ad34dce-d1bc-408e-9f82-e5c370cc01f5",
				Quantity:    1,
			},
			{
				Medicine_ID: "83803a11-1388-4beb-b06b-b22f1c98edaf",
				Quantity:    2,
			},
		},
		Action_Details: []medicalRecordDTO.Action_Details_Request{
			{
				Action_ID: "e6ba6dcc-9c95-4f20-9477-6104c07fef2b",
			},
		},
	}

	expectedMedicalRecord := medicalRecordDTO.Medical_Record{
		ID:               "1",
		Booking_ID:       requestPayload.Booking_ID,
		Diagnosis_Result: requestPayload.Diagnosis_Result,
		Payment_Status:   true,
		Total_Medicine:   0,
		Total_Action:     0,
		Total_Amount:     0,
		Created_At:       time.Now().Format("2006-01-02 15:04:05"),
		Updated_At:       time.Now().Format("2006-01-02 15:04:05"),
		Medicine_Details: []medicalRecordDTO.Medical_Record_Medicine_Details{
			{
				Medicine_ID: "5ad34dce-d1bc-408e-9f82-e5c370cc01f5",
				Quantity:    1,
			},
			{
				Medicine_ID: "83803a11-1388-4beb-b06b-b22f1c98edaf",
				Quantity:    2,
			},
		},
		Action_Details: []medicalRecordDTO.Medical_Record_Action_Details{
			{
				Action_ID: "e6ba6dcc-9c95-4f20-9477-6104c07fef2b",
			},
		},
	}
	suite.medicalRecordUCMock.On("CreateMedicalRecord", requestPayload).Return(expectedMedicalRecord, nil)

	w := httptest.NewRecorder()
	reqBody, _ := json.Marshal(requestPayload)
	req, _ := http.NewRequest("POST", "/api/v1/medical-records", bytes.NewBuffer(reqBody))
	token, _ := utils.GenerateJWT("2cfde543-ea6a-469f-b332-4e630a1cad8c", "doctor", "DOCTOR")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusCreated, w.Code)

	var response struct {
		ResponseCode    string                          `json:"responseCode"`
		ResponseMessage string                          `json:"responseMessage"`
		Data            medicalRecordDTO.Medical_Record `json:"data"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	suite.NoError(err)

	suite.Equal("2010601", response.ResponseCode)
	suite.Equal("data created", response.ResponseMessage)
	suite.Equal(expectedMedicalRecord, response.Data)
}

func (suite *MedicalRecordDeliverySuite) TestCreateMedicalRecord_BadRequest() {
	requestPayload := medicalRecordDTO.Medical_Record_Request{}

	w := httptest.NewRecorder()
	reqBody, _ := json.Marshal(requestPayload)
	req, _ := http.NewRequest("POST", "/api/v1/medical-records", bytes.NewBuffer(reqBody))
	token, _ := utils.GenerateJWT("2cfde543-ea6a-469f-b332-4e630a1cad8c", "doctor", "DOCTOR")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusBadRequest, w.Code)

	var response struct {
		ResponseCode     string                   `json:"responseCode"`
		ResponseMessage  string                   `json:"responseMessage"`
		ErrorDescription []myjson.ValidationField `json:"error_description"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	suite.NoError(err)

	suite.Equal("4000602", response.ResponseCode)
	suite.Equal("bad request. required fields cannot be empty", response.ResponseMessage)
	suite.NotEmpty(response.ErrorDescription)
}

func (suite *MedicalRecordDeliverySuite) TestCreateMedicalRecord_Error() {
	requestPayload := medicalRecordDTO.Medical_Record_Request{
		Booking_ID:       "ea1c7e2c-3799-4ef7-a8e7-4ecf6c413151",
		Diagnosis_Result: "Test diagnosis",
		Payment_Status:   true,
		Medicine_Details: []medicalRecordDTO.Medicine_Details_Request{
			{
				Medicine_ID: "5ad34dce-d1bc-408e-9f82-e5c370cc01f5",
				Quantity:    1,
			},
			{
				Medicine_ID: "83803a11-1388-4beb-b06b-b22f1c98edaf",
				Quantity:    2,
			},
		},
		Action_Details: []medicalRecordDTO.Action_Details_Request{
			{
				Action_ID: "e6ba6dcc-9c95-4f20-9477-6104c07fef2b",
			},
		},
	}

	expectedError := errors.New("mocked error")
	suite.medicalRecordUCMock.On("CreateMedicalRecord", requestPayload).Return(medicalRecordDTO.Medical_Record{}, expectedError)

	w := httptest.NewRecorder()
	reqBody, _ := json.Marshal(requestPayload)
	req, _ := http.NewRequest("POST", "/api/v1/medical-records", bytes.NewBuffer(reqBody))
	token, _ := utils.GenerateJWT("2cfde543-ea6a-469f-b332-4e630a1cad8c", "doctor", "DOCTOR")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusInternalServerError, w.Code)

	var response struct {
		ResponseCode    string `json:"responseCode"`
		ResponseMessage string `json:"responseMessage"`
		Error           string `json:"error,omitempty"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	suite.NoError(err)

	suite.Equal("5000605", response.ResponseCode)
	suite.Equal("internal server error", response.ResponseMessage)
	suite.Equal("mocked error", response.Error)
}

func (suite *MedicalRecordDeliverySuite) TestCreateMedicalRecord_NoStockAvailable() {
	requestPayload := medicalRecordDTO.Medical_Record_Request{
		Booking_ID:       "ea1c7e2c-3799-4ef7-a8e7-4ecf6c413151",
		Diagnosis_Result: "Test diagnosis",
		Payment_Status:   true,
		Medicine_Details: []medicalRecordDTO.Medicine_Details_Request{
			{
				Medicine_ID: "5ad34dce-d1bc-408e-9f82-e5c370cc01f5",
				Quantity:    1,
			},
		},
		Action_Details: []medicalRecordDTO.Action_Details_Request{
			{
				Action_ID: "e6ba6dcc-9c95-4f20-9477-6104c07fef2b",
			},
		},
	}

	expectedError := errors.New(constants.ErrNoStockAvailable)
	suite.medicalRecordUCMock.On("CreateMedicalRecord", requestPayload).Return(medicalRecordDTO.Medical_Record{}, expectedError)

	w := httptest.NewRecorder()
	reqBody, _ := json.Marshal(requestPayload)
	req, _ := http.NewRequest("POST", "/api/v1/medical-records", bytes.NewBuffer(reqBody))
	token, _ := utils.GenerateJWT("2cfde543-ea6a-469f-b332-4e630a1cad8c", "doctor", "DOCTOR")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusBadRequest, w.Code)

	var response struct {
		ResponseCode     string                   `json:"responseCode"`
		ResponseMessage  string                   `json:"responseMessage"`
		ErrorDescription []myjson.ValidationField `json:"error_description"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	suite.NoError(err)

	suite.Equal("4000603", response.ResponseCode)
	suite.Equal(constants.ErrNoStockAvailable, response.ResponseMessage)
}

func (suite *MedicalRecordDeliverySuite) TestCreateMedicalRecord_QuantityGreaterThanStock() {
	requestPayload := medicalRecordDTO.Medical_Record_Request{
		Booking_ID:       "ea1c7e2c-3799-4ef7-a8e7-4ecf6c413151",
		Diagnosis_Result: "Test diagnosis",
		Payment_Status:   true,
		Medicine_Details: []medicalRecordDTO.Medicine_Details_Request{
			{
				Medicine_ID: "5ad34dce-d1bc-408e-9f82-e5c370cc01f5",
				Quantity:    10,
			},
		},
		Action_Details: []medicalRecordDTO.Action_Details_Request{
			{
				Action_ID: "e6ba6dcc-9c95-4f20-9477-6104c07fef2b",
			},
		},
	}

	expectedError := errors.New(constants.ErrQuantityGreaterThanStock)
	suite.medicalRecordUCMock.On("CreateMedicalRecord", requestPayload).Return(medicalRecordDTO.Medical_Record{}, expectedError)

	w := httptest.NewRecorder()
	reqBody, _ := json.Marshal(requestPayload)
	req, _ := http.NewRequest("POST", "/api/v1/medical-records", bytes.NewBuffer(reqBody))
	token, _ := utils.GenerateJWT("2cfde543-ea6a-469f-b332-4e630a1cad8c", "doctor", "DOCTOR")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusBadRequest, w.Code)

	var response struct {
		ResponseCode     string                   `json:"responseCode"`
		ResponseMessage  string                   `json:"responseMessage"`
		ErrorDescription []myjson.ValidationField `json:"error_description"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	suite.NoError(err)

	suite.Equal("4000604", response.ResponseCode)
	suite.Equal(constants.ErrQuantityGreaterThanStock, response.ResponseMessage)
}

func (suite *MedicalRecordDeliverySuite) TestGetMedicalRecords_Success() {
	mockMedicalRecords := []medicalRecordDTO.Medical_Record{
		{
			ID:               "1",
			Booking_ID:       "ea1c7e2c-3799-4ef7-a8e7-4ecf6c413151",
			Diagnosis_Result: "Test diagnosis",
			Payment_Status:   true,
			Created_At:       "2022-03-10 12:00:00",
			Updated_At:       "2022-03-10 12:00:00",
			Medicine_Details: []medicalRecordDTO.Medical_Record_Medicine_Details{
				{
					Medicine_ID: "5ad34dce-d1bc-408e-9f82-e5c370cc01f5",
					Quantity:    1,
				},
				{
					Medicine_ID: "83803a11-1388-4beb-b06b-b22f1c98edaf",
					Quantity:    2,
				},
			},
			Action_Details: []medicalRecordDTO.Medical_Record_Action_Details{
				{
					Action_ID: "e6ba6dcc-9c95-4f20-9477-6104c07fef2b",
				},
			},
		},
		{
			ID:               "2",
			Booking_ID:       "ea1c7e2c-3799-4ef7-a8e7-4ecf6c413151",
			Diagnosis_Result: "Test diagnosis 2",
			Payment_Status:   true,
			Created_At:       "2022-03-10 12:00:00",
			Updated_At:       "2022-03-10 12:00:00",
			Medicine_Details: []medicalRecordDTO.Medical_Record_Medicine_Details{
				{
					Medicine_ID: "5ad34dce-d1bc-408e-9f82-e5c370cc01f5",
					Quantity:    1,
				},
				{
					Medicine_ID: "83803a11-1388-4beb-b06b-b22f1c98edaf",
					Quantity:    2,
				},
			},
			Action_Details: []medicalRecordDTO.Medical_Record_Action_Details{
				{
					Action_ID: "e6ba6dcc-9c95-4f20-9477-6104c07fef2b",
				},
			},
		},
	}

	suite.medicalRecordUCMock.On("GetMedicalRecords").Return(mockMedicalRecords, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/medical-records", nil)
	token, _ := utils.GenerateJWT("2cfde543-ea6a-469f-b332-4e630a1cad8c", "hello", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)

	var response struct {
		ResponseMessage string                            `json:"responseMessage"`
		Data            []medicalRecordDTO.Medical_Record `json:"data"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	suite.NoError(err)

	suite.Equal("data received", response.ResponseMessage)

	suite.Equal(mockMedicalRecords, response.Data)
}

func (suite *MedicalRecordDeliverySuite) TestGetMedicalRecordByID_Success() {
	mockMedicalRecord := medicalRecordDTO.Medical_Record{
		ID:               "1",
		Booking_ID:       "booking1",
		Diagnosis_Result: "diagnosis1",
	}

	suite.medicalRecordUCMock.On("GetMedicalRecordByID", "1").Return(mockMedicalRecord, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/medical-records/1", nil)
	token, _ := utils.GenerateJWT("2cfde543-ea6a-469f-b332-4e630a1cad8c", "hello", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)

	var response struct {
		ResponseMessage string                          `json:"responseMessage"`
		Data            medicalRecordDTO.Medical_Record `json:"data"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	suite.NoError(err)

	suite.Equal("data received", response.ResponseMessage)

	suite.Equal(mockMedicalRecord, response.Data)
}

func (suite *MedicalRecordDeliverySuite) TestGetMedicalRecordByID_NotFound() {
	expectedError := errors.New("data not found")
	suite.medicalRecordUCMock.On("GetMedicalRecordByID", "non_existent_id").Return(medicalRecordDTO.Medical_Record{}, expectedError)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/medical-records/non_existent_id", nil)
	token, _ := utils.GenerateJWT("2cfde543-ea6a-469f-b332-4e630a1cad8c", "hello", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusBadRequest, w.Code)

	var response struct {
		ResponseMessage string                          `json:"responseMessage"`
		Data            medicalRecordDTO.Medical_Record `json:"data"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	suite.NoError(err)

	suite.Equal("data not found", response.ResponseMessage)

	suite.Equal(medicalRecordDTO.Medical_Record{}, response.Data)
}

func (suite *MedicalRecordDeliverySuite) TestUpdatePaymentStatus_Success() {
	mockMedicalRecord := medicalRecordDTO.Medical_Record{
		ID:               "1",
		Booking_ID:       "booking1",
		Diagnosis_Result: "diagnosis1",
		Payment_Status:   true,
	}

	suite.medicalRecordUCMock.On("UpdatePaymentStatus", "1").Return(mockMedicalRecord, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/v1/medical-records/1", nil)
	token, _ := utils.GenerateJWT("2cfde543-ea6a-469f-b332-4e630a1cad8c", "hello", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)

	var response struct {
		ResponseMessage string                          `json:"responseMessage"`
		Data            medicalRecordDTO.Medical_Record `json:"data"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	suite.NoError(err)

	suite.Equal("data updated", response.ResponseMessage)

	suite.Equal(mockMedicalRecord, response.Data)
}

func (suite *MedicalRecordDeliverySuite) TestUpdatePaymentStatus_PaymentAlreadyTrue() {
	mockError := errors.New(constants.ErrPaymentAlreadyTrue)
	suite.medicalRecordUCMock.On("UpdatePaymentStatus", "1").Return(medicalRecordDTO.Medical_Record{}, mockError)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/v1/medical-records/1", nil)
	token, _ := utils.GenerateJWT("2cfde543-ea6a-469f-b332-4e630a1cad8c", "hello", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusBadRequest, w.Code)

	var response struct {
		ResponseMessage  string                          `json:"responseMessage"`
		Data             medicalRecordDTO.Medical_Record `json:"data"`
		ErrorDescription []myjson.ValidationField        `json:"error_description"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	suite.NoError(err)

	suite.Equal(constants.ErrPaymentAlreadyTrue, response.ResponseMessage)
	suite.Empty(response.Data)
}

func (suite *MedicalRecordDeliverySuite) TestUpdatePaymentStatus_NoStockAvailable() {
	mockError := errors.New(constants.ErrNoStockAvailable)
	suite.medicalRecordUCMock.On("UpdatePaymentStatus", "1").Return(medicalRecordDTO.Medical_Record{}, mockError)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/v1/medical-records/1", nil)
	token, _ := utils.GenerateJWT("2cfde543-ea6a-469f-b332-4e630a1cad8c", "hello", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusBadRequest, w.Code)

	var response struct {
		ResponseMessage  string                          `json:"responseMessage"`
		Data             medicalRecordDTO.Medical_Record `json:"data"`
		ErrorDescription []myjson.ValidationField        `json:"error_description"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	suite.NoError(err)

	suite.Equal(constants.ErrNoStockAvailable, response.ResponseMessage)
	suite.Empty(response.Data)
}

func (suite *MedicalRecordDeliverySuite) TestUpdatePaymentStatus_QuantityGreaterThanStock() {
	mockError := errors.New(constants.ErrQuantityGreaterThanStock)
	suite.medicalRecordUCMock.On("UpdatePaymentStatus", "1").Return(medicalRecordDTO.Medical_Record{}, mockError)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/v1/medical-records/1", nil)
	token, _ := utils.GenerateJWT("2cfde543-ea6a-469f-b332-4e630a1cad8c", "hello", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusBadRequest, w.Code)

	var response struct {
		ResponseMessage  string                          `json:"responseMessage"`
		Data             medicalRecordDTO.Medical_Record `json:"data"`
		ErrorDescription []myjson.ValidationField        `json:"error_description"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	suite.NoError(err)

	suite.Equal(constants.ErrQuantityGreaterThanStock, response.ResponseMessage)
	suite.Empty(response.Data)
}

func (suite *MedicalRecordDeliverySuite) TestUpdatePaymentStatus_DataNotFound() {
	mockError := errors.New("data not found")
	suite.medicalRecordUCMock.On("UpdatePaymentStatus", "1").Return(medicalRecordDTO.Medical_Record{}, mockError)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/v1/medical-records/1", nil)
	token, _ := utils.GenerateJWT("2cfde543-ea6a-469f-b332-4e630a1cad8c", "hello", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusBadRequest, w.Code)

	var response struct {
		ResponseMessage  string                          `json:"responseMessage"`
		Data             medicalRecordDTO.Medical_Record `json:"data"`
		ErrorDescription []myjson.ValidationField        `json:"error_description"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	suite.NoError(err)

	suite.Equal("data not found", response.ResponseMessage)
	suite.Empty(response.Data)
}

func (suite *medicalRecordDelivery) SetupSuite() {
	gin.SetMode(gin.TestMode)
}

func TestMedicalRecordDeliverySuite(t *testing.T) {
	suite.Run(t, new(MedicalRecordDeliverySuite))
}
