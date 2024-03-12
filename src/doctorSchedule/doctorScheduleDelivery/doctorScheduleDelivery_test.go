package doctorScheduleDelivery

import (
	"avengers-clinic/model/dto"
	"avengers-clinic/model/entity"
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

func TestDoctorScheduleDelivery(t *testing.T) {
	suite.Run(t, new(doctorScheduleDeliveryTestSuite))
}
