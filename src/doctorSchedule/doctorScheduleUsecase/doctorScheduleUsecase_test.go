package doctorScheduleUsecase

import (
	"avengers-clinic/model/dto"
	"avengers-clinic/model/entity"
	"avengers-clinic/pkg/utils"
	"avengers-clinic/src/doctorSchedule"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type mockDoctorScheduleRepo struct {
	mock.Mock
}

func (mr *mockDoctorScheduleRepo) RetrieveAll(startDate string, endDate string) ([]entity.DoctorSchedule, error) {
	args := mr.Called()
	return args.Get(0).([]entity.DoctorSchedule), args.Error(1)
}

func (mr *mockDoctorScheduleRepo) RetrieveByID(id uuid.UUID) (entity.DoctorSchedule, error) {
	args := mr.Called()
	return args.Get(0).(entity.DoctorSchedule), args.Error(1)
}

func (mr *mockDoctorScheduleRepo) GetByIDs(id uuid.UUIDs) ([]entity.DoctorSchedule, error) {
	args := mr.Called()
	return args.Get(0).([]entity.DoctorSchedule), args.Error(1)
}

func (mr *mockDoctorScheduleRepo) InsertSchedule(input dto.CreateDoctorSchedule) (uuid.UUIDs, error) {
	args := mr.Called()
	return args.Get(0).(uuid.UUIDs), args.Error(1)
}

func (mr *mockDoctorScheduleRepo) GetMySchedule(doctorId uuid.UUID, dayOfWeek []int, startDate string, endDate string) ([]entity.DoctorSchedule, error) {
	args := mr.Called()
	return args.Get(0).([]entity.DoctorSchedule), args.Error(1)
}

func (mr *mockDoctorScheduleRepo) UpdateSchedule(id uuid.UUID, input entity.DoctorSchedule) error {
	args := mr.Called()
	return args.Error(0)
}

func (mr *mockDoctorScheduleRepo) DeleteSchedule(id uuid.UUID) error {
	args := mr.Called()
	return args.Error(0)
}

func (mr *mockDoctorScheduleRepo) Restore(id uuid.UUID) error {
	args := mr.Called()
	return args.Error(0)
}

func (mr *mockDoctorScheduleRepo) SearchByDateAndDoctorID(date string, doctorID uuid.UUID) error {
	args := mr.Called()
	return args.Error(0)
}

type mockBookingRepo struct {
	mock.Mock
}

func (mb *mockBookingRepo) GetAllBooking() ([]entity.Bookings, error) {
	args := mb.Called()
	return args.Get(0).([]entity.Bookings), args.Error(1)
}

func (mb *mockBookingRepo) GetAllBookingByDoctorID(doctorId uuid.UUID) ([]entity.Bookings, error) {
	args := mb.Called()
	return args.Get(0).([]entity.Bookings), args.Error(1)
}

func (mb *mockBookingRepo) GetOneByID(id uuid.UUID) (entity.Bookings, error) {
	args := mb.Called()
	return args.Get(0).(entity.Bookings), args.Error(1)
}

func (mb *mockBookingRepo) GetBookingByScheduleID(scheduleId uuid.UUID, status []string) ([]entity.Bookings, error) {
	args := mb.Called()
	return args.Get(0).([]entity.Bookings), args.Error(1)
}

func (mb *mockBookingRepo) CreateBooking(input entity.Bookings) (entity.Bookings, error) {
	args := mb.Called()
	return args.Get(0).(entity.Bookings), args.Error(1)
}

func (mb *mockBookingRepo) CheckExist(doctorScheduleID uuid.UUID, mstScheduleID int) bool {
	args := mb.Called()
	return args.Get(0).(bool)
}

func (mb *mockBookingRepo) EditSchedule(id uuid.UUID, input entity.Bookings) error {
	args := mb.Called()
	return args.Error(0)
}

func (mb *mockBookingRepo) CancelBooking(id uuid.UUID) error {
	args := mb.Called()
	return args.Error(0)
}

func (mb *mockBookingRepo) FinishBooking(id uuid.UUID) error {
	args := mb.Called()
	return args.Error(0)
}

type doctorUcTestSuite struct {
	suite.Suite
	doctorRepo  *mockDoctorScheduleRepo
	bookingRepo *mockBookingRepo
	doctorUC    doctorSchedule.DoctorScheduleUsecase
}

func (suite *doctorUcTestSuite) SetupTest() {
	suite.doctorRepo = new(mockDoctorScheduleRepo)
	suite.bookingRepo = new(mockBookingRepo)
	suite.doctorUC = NewDoctorScheduleUsecase(suite.doctorRepo, suite.bookingRepo)
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
		Schedules:    bookings,
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

func (suite *doctorUcTestSuite) TestGetAll() {
	suite.doctorRepo.On("RetrieveAll").Return(arrExpected, nil)
	actual, err := suite.doctorUC.GetAll(startDate, endDate)
	suite.Nil(err)
	suite.Equal(arrExpected, actual)
}

func (suite *doctorUcTestSuite) TestGetByID() {
	suite.doctorRepo.On("RetrieveByID").Return(expected, nil)
	suite.bookingRepo.On("GetBookingByScheduleID").Return(bookings, nil)
	actual, err := suite.doctorUC.GetByID(id, status)
	suite.Nil(err)
	suite.Equal(expected, actual)
}

func (suite *doctorUcTestSuite) TestCreate() {
	suite.doctorRepo.On("InsertSchedule").Return(uuids, nil)
	suite.doctorRepo.On("GetByIDs").Return(arrExpected, nil)
	actual, err := suite.doctorUC.CreateSchedule(dto.CreateDoctorSchedule{})
	suite.Nil(err)
	suite.Equal(arrExpected, actual)
}

func (suite *doctorUcTestSuite) TestGetMySchedule() {
	suite.doctorRepo.On("GetMySchedule").Return(arrExpected, nil)
	suite.bookingRepo.On("GetBookingByScheduleID").Return(bookings, nil)
	actual, err := suite.doctorUC.GetMySchedule(doctorID, dayOfWeeks, status, startDate, endDate)
	suite.Nil(err)
	suite.Equal(arrExpected, actual)
}

func (suite *doctorUcTestSuite) TestUpdate() {

	expected := entity.DoctorSchedule{
		ID:           id,
		DoctorID:     doctorID,
		ScheduleDate: "2024-03-14",
		StartAt:      1,
		EndAt:        9,
		CreatedAt:    "2024-03-12 22:39:22.245736",
		UpdatedAt:    &updatedAt,
	}

	suite.doctorRepo.On("RetrieveByID").Return(expected, nil)
	suite.doctorRepo.On("UpdateSchedule").Return(nil)

	actual, err := suite.doctorUC.UpdateSchedule(id, dto.UpdateSchedule{})
	suite.Nil(err)
	suite.Equal(expected, actual)
}

func (suite *doctorUcTestSuite) TestDelete() {
	suite.doctorRepo.On("RetrieveByID").Return(expected, nil)
	suite.doctorRepo.On("DeleteSchedule").Return(nil)
	err := suite.doctorUC.DeleteSchedule(id)
	suite.Nil(err)
}

func (suite *doctorUcTestSuite) TestRestore() {
	suite.doctorRepo.On("Restore").Return(nil)
	err := suite.doctorUC.Restore(id)
	suite.Nil(err)
}

func TestDoctorUsecase(t *testing.T) {
	suite.Run(t, new(doctorUcTestSuite))
}
