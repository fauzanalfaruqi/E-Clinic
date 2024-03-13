package doctorScheduleUsecase

import (
	"avengers-clinic/model/dto"
	"avengers-clinic/model/entity"
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
	doctorRepo *mockDoctorScheduleRepo
	bookingRepo *mockBookingRepo
	doctorUC   doctorSchedule.DoctorScheduleUsecase
}

func (suite *doctorUcTestSuite) SetupTest() {
	suite.doctorRepo = new(mockDoctorScheduleRepo)
	suite.doctorUC = NewDoctorScheduleUsecase(suite.doctorRepo, suite.bookingRepo)
}


func TestDoctorUsecase(t *testing.T) {
	suite.Run(t, new(doctorUcTestSuite))
}
