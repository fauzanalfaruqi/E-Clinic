package booking

import (
	"avengers-clinic/model/dto"
	"avengers-clinic/model/entity"

	"github.com/google/uuid"
)

type (
	BookingRepository interface {
		GetAllBooking() ([]entity.Bookings, error)
		GetOneByID(id uuid.UUID) (entity.Bookings, error)
		GetBookingByScheduleID(scheduleId uuid.UUID, status []string) ([]entity.Bookings, error)
		CreateBooking(input entity.Bookings) (entity.Bookings, error)
		CheckExist(doctorScheduleID uuid.UUID, mstScheduleID int) bool
		EditSchedule(id uuid.UUID, input entity.Bookings) error
		CancelBooking(id uuid.UUID) error
		FinishBooking(id uuid.UUID) error
	}

	BookingUsecase interface {
		GetAll() ([]entity.Bookings, error)
		GetOneByID(id uuid.UUID) (entity.Bookings, error)
		GetBookingByScheduleID(scheduleId uuid.UUID, status string) ([]entity.Bookings, error)
		Create(input dto.CreateBooking) (entity.Bookings, error)
		EditSchedule(id uuid.UUID, input dto.UpdateBookingSchedule) (entity.Bookings, error)
		Cancel(id uuid.UUID) (entity.Bookings, error)
		FinishBooking(id uuid.UUID) (entity.Bookings, error)
	}
)