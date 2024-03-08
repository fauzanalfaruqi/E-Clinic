package booking

import (
	"avengers-clinic/model/dto"
	"avengers-clinic/model/entity"

	"github.com/google/uuid"
)

type (
	BookingRepository interface {
		GetAllBooking(date string) ([]entity.Bookings, error)
		GetAllBookingByDoctorID(doctorId uuid.UUID) ([]entity.Bookings, error)
		GetOneByID(id uuid.UUID) (entity.Bookings, error)
		GetBookingByScheduleID()
		CreateBooking(input entity.Bookings) (entity.Bookings, error)
		CheckExist(doctorScheduleID uuid.UUID, bookingDate string, mstScheduleID int) bool
		EditSchedule(id uuid.UUID, input entity.Bookings) error
		CancelBooking(id uuid.UUID) error
		FinishBooking(id uuid.UUID) error
	}

	BookingUsecase interface {
		GetAll(date string) ([]entity.Bookings, error)
		GetAllByDoctorID(doctorId uuid.UUID) ([]entity.Bookings, error)
		GetOneByID(id uuid.UUID) (entity.Bookings, error)
		GetByScheduleID()
		Create(input dto.CreateBooking) (entity.Bookings, error)
		EditSchedule(id uuid.UUID, input dto.UpdateBookingSchedule) (entity.Bookings, error)
		Cancel(id uuid.UUID) (entity.Bookings, error)
		FinishBooking(id uuid.UUID) (entity.Bookings, error)
	}
)