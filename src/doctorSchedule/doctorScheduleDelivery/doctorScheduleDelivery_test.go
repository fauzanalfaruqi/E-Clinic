package doctorScheduleDelivery

import (
	"avengers-clinic/model/dto"
	"avengers-clinic/model/entity"
	"avengers-clinic/pkg/utils"
	"bytes"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type mockDoctorScheduleUC struct {
	mock.Mock
}

func (du *mockDoctorScheduleUC) GetAll(startDate string, endDate string) ([]entity.DoctorSchedule, error) {
	args := du.Called()
	return args.Get(0).([]entity.DoctorSchedule), args.Error(1)
}

func (du *mockDoctorScheduleUC) GetByID(id uuid.UUID, status string) (entity.DoctorSchedule, error) {
	args := du.Called()
	return args.Get(0).(entity.DoctorSchedule), args.Error(1)
}

func (du *mockDoctorScheduleUC) CreateSchedule(input dto.CreateDoctorSchedule) ([]entity.DoctorSchedule, error) {
	args := du.Called()
	return args.Get(0).([]entity.DoctorSchedule), args.Error(1)
}

func (du *mockDoctorScheduleUC) GetMySchedule(doctorId uuid.UUID, dayOfWeek string, status string, startDate string, endDate string) ([]entity.DoctorSchedule, error) {
	args := du.Called()
	return args.Get(0).([]entity.DoctorSchedule), args.Error(1)
}

func (du *mockDoctorScheduleUC) UpdateSchedule(id uuid.UUID, input dto.UpdateSchedule) (entity.DoctorSchedule, error) {
	args := du.Called()
	return args.Get(0).(entity.DoctorSchedule), args.Error(1)
}

func (du *mockDoctorScheduleUC) DeleteSchedule(id uuid.UUID) error {
	args := du.Called()
	return args.Error(0)
}

func (du *mockDoctorScheduleUC) Restore(id uuid.UUID) error {
	args := du.Called()
	return args.Error(0)
}

type doctorScheduleDeliveryTestSuite struct {
	suite.Suite
	router           *gin.Engine
	doctorScheduleUC *mockDoctorScheduleUC
}

func (suite *doctorScheduleDeliveryTestSuite) SetupTest() {
	suite.router = gin.New()
	suite.doctorScheduleUC = new(mockDoctorScheduleUC)

	v1Group := suite.router.Group("/api/v1")
	NewDoctorScheduleDelivery(v1Group, suite.doctorScheduleUC)
}

var (
	id, _       = uuid.Parse("74d93144-6f2e-4bbc-9f89-973c62d3ac54")
	doctorID, _ = uuid.Parse("5bc18dd0-58cb-4612-8dc3-5fc2419b7f29")
	dayOfWeeks  = "1#2"
	uuids       = uuid.UUIDs{}
	startDate   = "2024-03-11"
	endDate     = "2024-03-17"
	status      = "waiting"
	updatedAt   = utils.GetNow()
	arrExpected = []entity.DoctorSchedule{
		{
			ID:           id,
			DoctorID:     doctorID,
			ScheduleDate: "2024-03-14",
			StartAt:      1,
			EndAt:        9,
			CreatedAt:    "2024-03-12 22:39:22.245736",
		},
	}
	expected = entity.DoctorSchedule{
		ID:           id,
		DoctorID:     doctorID,
		ScheduleDate: "2024-03-14",
		StartAt:      1,
		EndAt:        9,
		CreatedAt:    "2024-03-12 22:39:22.245736",
	}
	bookings = []entity.Bookings{
		{
			ID:               id,
			DoctorScheduleID: id,
			PatientID:        id,
			MstScheduleID:    1,
			Complaint:        "test",
			Status:           status,
			CreatedAt:        "2024-03-12 22:39:22.245736",
		},
	}
)

func (suite *doctorScheduleDeliveryTestSuite) TestGetAll() {
	suite.doctorScheduleUC.On("GetAll").Return(arrExpected, nil)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/doctor-schedule", nil)

	token, _ := utils.GenerateJWT("1", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expected := `{"responseCode":"2000401","responseMessage":"success","data":[{"id":"74d93144-6f2e-4bbc-9f89-973c62d3ac54","doctor_id":"5bc18dd0-58cb-4612-8dc3-5fc2419b7f29","schedule_date":"2024-03-14","start_at":1,"end_at":9,"created_at":"2024-03-12 22:39:22.245736"}]}`

	suite.Equal(http.StatusOK, res.Code)
	suite.JSONEq(expected, res.Body.String())

}

func (suite *doctorScheduleDeliveryTestSuite) TestGetByIDSuccess() {
	suite.doctorScheduleUC.On("GetByID").Return(expected, nil)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/doctor-schedule/5bc18dd0-58cb-4612-8dc3-5fc2419b7f29", nil)

	token, _ := utils.GenerateJWT("1", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expected := `{"responseCode":"2000401","responseMessage":"success","data":{"id":"74d93144-6f2e-4bbc-9f89-973c62d3ac54","doctor_id":"5bc18dd0-58cb-4612-8dc3-5fc2419b7f29","schedule_date":"2024-03-14","start_at":1,"end_at":9,"created_at":"2024-03-12 22:39:22.245736"}}`

	suite.Equal(http.StatusOK, res.Code)
	suite.JSONEq(expected, res.Body.String())

}

func (suite *doctorScheduleDeliveryTestSuite) TestGetByIDError() {
	suite.doctorScheduleUC.On("GetByID").Return(entity.DoctorSchedule{}, sql.ErrNoRows)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/doctor-schedule/5bc18dd0-58cb-4612-8dc3-5fc2419b7f29", nil)

	token, _ := utils.GenerateJWT("1", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expected := `{"responseCode":"4000401","responseMessage":"data not found"}`

	suite.Equal(http.StatusBadRequest, res.Code)
	suite.JSONEq(expected, res.Body.String())

}

func (suite *doctorScheduleDeliveryTestSuite) TestGetMySchedule() {
	suite.doctorScheduleUC.On("GetMySchedule").Return(arrExpected, nil)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/doctor-schedule/my-schedule/5bc18dd0-58cb-4612-8dc3-5fc2419b7f29", nil)

	token, _ := utils.GenerateJWT("1", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expected := `{"responseCode":"2000401","responseMessage":"success","data":[{"id":"74d93144-6f2e-4bbc-9f89-973c62d3ac54","doctor_id":"5bc18dd0-58cb-4612-8dc3-5fc2419b7f29","schedule_date":"2024-03-14","start_at":1,"end_at":9,"created_at":"2024-03-12 22:39:22.245736"}]}`

	suite.Equal(http.StatusOK, res.Code)
	suite.JSONEq(expected, res.Body.String())

}

func (suite *doctorScheduleDeliveryTestSuite) TestCreate() {
	suite.doctorScheduleUC.On("CreateSchedule").Return(arrExpected, nil)

	reqBody := []byte(`{"doctor_id":"5bc18dd0-58cb-4612-8dc3-5fc2419b7f29","schedule_detail":[{"schedule_date":"2024-03-19","start_at":1,"end_at":9}]}`)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/doctor-schedule", bytes.NewBuffer(reqBody))

	token, _ := utils.GenerateJWT("1", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expected := `{"responseCode":"2010401","responseMessage":"success","data":[{"id":"74d93144-6f2e-4bbc-9f89-973c62d3ac54","doctor_id":"5bc18dd0-58cb-4612-8dc3-5fc2419b7f29","schedule_date":"2024-03-14","start_at":1,"end_at":9,"created_at":"2024-03-12 22:39:22.245736"}]}`

	suite.Equal(http.StatusCreated, res.Code)
	suite.JSONEq(expected, res.Body.String())

}

func (suite *doctorScheduleDeliveryTestSuite) TestUpdate() {
	suite.doctorScheduleUC.On("UpdateSchedule").Return(expected, nil)

	reqBody := []byte(`{"schedule_date":"2024-03-19","start_at":1,"end_at":9}`)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/doctor-schedule/74d93144-6f2e-4bbc-9f89-973c62d3ac54", bytes.NewBuffer(reqBody))

	token, _ := utils.GenerateJWT("1", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expected := `{"responseCode":"2010401","responseMessage":"success","data":{"id":"74d93144-6f2e-4bbc-9f89-973c62d3ac54","doctor_id":"5bc18dd0-58cb-4612-8dc3-5fc2419b7f29","schedule_date":"2024-03-14","start_at":1,"end_at":9,"created_at":"2024-03-12 22:39:22.245736"}}`

	suite.Equal(http.StatusCreated, res.Code)
	suite.JSONEq(expected, res.Body.String())

}

func (suite *doctorScheduleDeliveryTestSuite) TestDelete() {
	suite.doctorScheduleUC.On("DeleteSchedule").Return(nil)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/doctor-schedule/5bc18dd0-58cb-4612-8dc3-5fc2419b7f29", nil)

	token, _ := utils.GenerateJWT("1", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expected := `{"responseCode":"2000401","responseMessage":"deleted"}`

	suite.Equal(http.StatusOK, res.Code)
	suite.JSONEq(expected, res.Body.String())

}

func (suite *doctorScheduleDeliveryTestSuite) TestRestore() {
	suite.doctorScheduleUC.On("Restore").Return(nil)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/doctor-schedule/restore/5bc18dd0-58cb-4612-8dc3-5fc2419b7f29", nil)

	token, _ := utils.GenerateJWT("1", "admin", "ADMIN")
	req.Header.Set("Authorization", "Bearer "+token)
	suite.router.ServeHTTP(res, req)

	expected := `{"responseCode":"2000401","responseMessage":"restored"}`

	suite.Equal(http.StatusOK, res.Code)
	suite.JSONEq(expected, res.Body.String())

}

func TestDoctorScheduleDelivery(t *testing.T) {
	suite.Run(t, new(doctorScheduleDeliveryTestSuite))
}
