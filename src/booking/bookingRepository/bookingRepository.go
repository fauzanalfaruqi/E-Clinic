package bookingRepository

import (
	"avengers-clinic/model/entity"
	"avengers-clinic/pkg/constants"
	"avengers-clinic/src/booking"
	"database/sql"

	"github.com/google/uuid"
)

type bookingRepository struct {
	db *sql.DB
}

func NewBookingRepository(db *sql.DB) booking.BookingRepository {
	return &bookingRepository{
		db,
	}
}

func (br bookingRepository) GetAllBooking(date string) ([]entity.Bookings, error) {
	var bookings []entity.Bookings
	// now := time.Now().Format("2006-01-02")
	sqlstat := "SELECT id, doctor_schedule_id, patient_id, booking_date, mst_schedule_id, complaint, status FROM bookings WHERE booking_date >= $1"

	rows, err := br.db.Query(sqlstat, date)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		book := entity.Bookings{}
		err = rows.Scan(&book.ID, &book.DoctorScheduleID, &book.PatientID, &book.BookingDate, &book.MstScheduleID, &book.Complaint, &book.Status)
		if err != nil {
			return nil, err
		}
		bookings = append(bookings, book)
	}
	return bookings, nil
}

func (br bookingRepository) GetOneByID(id uuid.UUID) (entity.Bookings, error) {
	var book entity.Bookings
	sqlstat := "SELECT id, doctor_schedule_id, patient_id, booking_date, mst_schedule_id, complaint, status FROM bookings WHERE id = $1"

	err := br.db.QueryRow(sqlstat, id).Scan(&book.ID, &book.DoctorScheduleID, &book.PatientID, &book.BookingDate, &book.MstScheduleID, &book.Complaint, &book.Status)
	if err != nil {
		return book, err
	}

	return book, nil
}

func (br bookingRepository) CheckExist(doctorScheduleID uuid.UUID, bookingDate string, mstScheduleID int) bool {
	exist := false
	sqlstat := "SELECT true FROM bookings WHERE doctor_schedule_id = $1 AND booking_date = $2 AND mst_schedule_id = $3"

	_ = br.db.QueryRow(sqlstat, doctorScheduleID, bookingDate, mstScheduleID).Scan(&exist)
	return exist
}

func (br bookingRepository) GetBookingByScheduleID() {}
func (br bookingRepository) GetAllBookingByDoctorID(doctorId uuid.UUID) ([]entity.Bookings, error) {
	return nil, nil
}

func (br bookingRepository) CreateBooking(input entity.Bookings) (entity.Bookings, error) {
	sqlstat := `
	
	INSERT INTO bookings(doctor_schedule_id, patient_id, booking_date, mst_schedule_id, complaint, status)
		SELECT $1, $2, $3, $4, $5, $6 WHERE NOT EXISTS(
			SELECT 1 FROM bookings WHERE doctor_schedule_id = $7 AND booking_date = $8 AND mst_schedule_id = $9
		)
	RETURNING id;`

	err := br.db.QueryRow(sqlstat, input.DoctorScheduleID, input.PatientID, input.BookingDate, input.MstScheduleID, input.Complaint, input.Status, input.DoctorScheduleID, input.BookingDate, input.MstScheduleID).Scan(&input.ID)
	if err != nil {
		return input, err
	}

	return input, nil
}

func (br bookingRepository) EditSchedule(id uuid.UUID, input entity.Bookings) error {
	sqlstat := "UPDATE bookings SET doctor_schedule_id = $1, booking_date= $2, mst_schedule_id = $3, complaint = $4 WHERE id = $5"
	_, err := br.db.Exec(sqlstat, input.DoctorScheduleID, input.BookingDate, input.MstScheduleID, input.Complaint, id)
	if err != nil {
		return err
	}
	return nil
}

func (br bookingRepository) CancelBooking(id uuid.UUID) error {
	sqlstat := "UPDATE bookings SET status = $1 WHERE id = $2;"
	_, err := br.db.Exec(sqlstat, constants.Canceled, id)
	if err != nil {
		return err
	}

	return nil
}

func (br bookingRepository) FinishBooking(id uuid.UUID) error {
	sqlstat := "UPDATE bookings SET status = $1 WHERE id = $2;"
	_, err := br.db.Exec(sqlstat, constants.Done, id)
	if err != nil {
		return err
	}

	return nil
}
