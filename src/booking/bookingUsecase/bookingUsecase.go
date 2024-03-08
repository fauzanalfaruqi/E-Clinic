package bookingUsecase

import (
	"avengers-clinic/model/dto"
	"avengers-clinic/model/entity"
	"avengers-clinic/pkg/constants"
	"avengers-clinic/src/booking"
	"avengers-clinic/src/doctorSchedule"
	"errors"
	"time"

	"github.com/google/uuid"
)

type bookingUsecase struct {
	bookingRepo  booking.BookingRepository
	scheduleRepo doctorSchedule.DoctorScheduleRepository
}

func NewBookingUsecase(bookingRepo booking.BookingRepository, scheduleRepo doctorSchedule.DoctorScheduleRepository) booking.BookingUsecase {
	return &bookingUsecase{
		bookingRepo,
		scheduleRepo,
	}
}

func (bu bookingUsecase) GetAll(date string) ([]entity.Bookings, error) {
	data, err := bu.bookingRepo.GetAllBooking(date)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (bu bookingUsecase) GetOneByID(id uuid.UUID) (entity.Bookings, error) {
	data, err := bu.bookingRepo.GetOneByID(id)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (bu bookingUsecase) GetAllByDoctorID(doctorId uuid.UUID) ([]entity.Bookings, error) {
	panic("not implemented") // TODO: Implement
}

func (bu bookingUsecase) GetByScheduleID() {
	panic("not implemented") // TODO: Implement
}

func (bu bookingUsecase) Create(input dto.CreateBooking) (entity.Bookings, error) {
	book := entity.Bookings{
		DoctorScheduleID: input.DoctorScheduleID,
		PatientID:        input.PatientID,
		BookingDate:      input.BookingDate,
		MstScheduleID:    input.MstScheduleID,
		Complaint:        input.Complaint,
		Status:           constants.Waiting,
	}

	valid, err := bu.validateDay(book.BookingDate, book.DoctorScheduleID)
	if !valid {
		return book, err
	}

	data, err := bu.bookingRepo.CreateBooking(book)
	if err != nil {
		return data, err
	}

	return data, nil

}

func (bu bookingUsecase) EditSchedule(id uuid.UUID, input dto.UpdateBookingSchedule) (entity.Bookings, error) {
	//Find data
	data, err := bu.bookingRepo.GetOneByID(id)
	if err != nil {
		return data, err
	}

	if input.DoctorScheduleID != uuid.Nil {
		data.DoctorScheduleID = input.DoctorScheduleID
	}
	if input.BookingDate != "" {
		data.BookingDate = input.BookingDate
	}
	if input.MstScheduleID > 0 {
		data.MstScheduleID = input.MstScheduleID
	}
	if input.Complaint != "" {
		data.Complaint = input.Complaint
	}

	existUpdate := bu.bookingRepo.CheckExist(data.DoctorScheduleID, data.BookingDate, data.MstScheduleID)

	if existUpdate {
		return data, errors.New(constants.ErrScheduleTaken)
	}

	err = bu.bookingRepo.EditSchedule(id, data)
	if err != nil {
		return data, err
	}
	return data, nil
}

func (bu bookingUsecase) Cancel(id uuid.UUID) (entity.Bookings, error) {
	data, err := bu.bookingRepo.GetOneByID(id)
	if err != nil {
		return data, err
	}
	err = bu.bookingRepo.CancelBooking(id)
	if err != nil {
		return data, err
	}
	data.Status = constants.Canceled
	return data, nil
}

func (bu bookingUsecase) FinishBooking(id uuid.UUID) (entity.Bookings, error) {
	data, err := bu.bookingRepo.GetOneByID(id)
	if err != nil {
		return data, err
	}
	err = bu.bookingRepo.FinishBooking(id)
	if err != nil {
		return data, err
	}
	data.Status = constants.Done
	return data, nil
}

func (bu bookingUsecase) validateDay(bookingDate string, doctorScheduleID uuid.UUID) (bool, error) {
	docSched, err := bu.scheduleRepo.RetrieveByID(doctorScheduleID)
	if err != nil {
		return false, err
	}

	date, _ := time.Parse("2006-01-02", bookingDate)
	dOWDate := date.Weekday()
	if dOWDate != time.Weekday(docSched.DayOfWeek) {
		return false, errors.New(constants.ErrScheduleNotMatch)
	}

	return true, nil
}
