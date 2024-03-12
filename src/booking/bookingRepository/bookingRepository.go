package bookingRepository

import (
	"avengers-clinic/model/entity"
	"avengers-clinic/pkg/constants"
	"avengers-clinic/src/booking"
	"database/sql"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type bookingRepository struct {
	db *sql.DB
}

func NewBookingRepository(db *sql.DB) booking.BookingRepository {
	return &bookingRepository{
		db,
	}
}


func (br bookingRepository) GetAllBooking() ([]entity.Bookings, error) {
	sqlstat := `
		SELECT 
				b.id, 
				b.doctor_schedule_id, 
				b.patient_id, 
				b.mst_schedule_id, 
				b.complaint, 
				b.status, 
				s.id, 
				to_char(s.start_at, 'HH24:MI:SS'), 
				to_char(s.end_at, 'HH24:MI:SS')
		FROM bookings b 
		LEFT JOIN mst_schedule_time s ON s.id = b.mst_schedule_id 
		WHERE b.deleted_at IS NULL ORDER BY b.created_at, b.mst_schedule_id;`


	rows, err := br.db.Query(sqlstat)
	if err != nil {
		return nil, err
	}
	data, err := scanBookingRows(rows)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (br bookingRepository) GetOneByID(id uuid.UUID) (entity.Bookings, error) {
	var book entity.Bookings
	sqlstat := `
		SELECT 
				b.id, 
				b.doctor_schedule_id, 
				b.patient_id, 
				b.mst_schedule_id, 
				b.complaint, 
				b.status, 
				s.id, 
				to_char(s.start_at, 'HH24:MI:SS'), 
				to_char(s.end_at, 'HH24:MI:SS')
		FROM bookings b 
			LEFT JOIN mst_schedule_time s ON s.id = b.mst_schedule_id 
		WHERE b.id = $1 AND b.deleted_at IS NULL;`

	err := br.db.QueryRow(sqlstat, id).Scan(
		&book.ID,
		&book.DoctorScheduleID,
		&book.PatientID,
		&book.MstScheduleID,
		&book.Complaint,
		&book.Status,
		&book.ScheduleTime.ID,
		&book.ScheduleTime.StartAt,
		&book.ScheduleTime.EndAt,
	)
	if err != nil {
		return book, err
	}

	return book, nil
}



func (br bookingRepository) GetBookingByScheduleID(scheduleID uuid.UUID, status []string) ([]entity.Bookings, error) {
	var rows *sql.Rows
	var err error

	sqlstat := `
		SELECT 
				b.id, 
				b.doctor_schedule_id, 
				b.patient_id,
				b.mst_schedule_id, 
				b.complaint, 
				b.status, 
				s.id, 
				to_char(s.start_at, 'HH24:MI:SS'), 
				to_char(s.end_at, 'HH24:MI:SS')
		FROM bookings b 
		LEFT JOIN mst_schedule_time s ON s.id = b.mst_schedule_id 
		WHERE b.doctor_schedule_id = $1 `

	orderStmt := "ORDER BY b.mst_schedule_id ASC;"
	if len(status) > 0 {
		sqlstat += "AND b.status = ANY($2) " + orderStmt
		rows, err = br.db.Query(sqlstat, scheduleID, pq.Array(status))
	}else {
		rows, err = br.db.Query(sqlstat+orderStmt, scheduleID)
	}
	if err != nil {
		return nil, err
	}
	data, err := scanBookingRows(rows)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (br bookingRepository) GetAllBookingByDoctorID(doctorId uuid.UUID) ([]entity.Bookings, error) {
	return nil, nil
}

func (br bookingRepository) CreateBooking(input entity.Bookings) (entity.Bookings, error) {
	sqlstat := `
	
	INSERT INTO bookings(doctor_schedule_id, patient_id, mst_schedule_id, complaint, status)
		SELECT $1, $2, $3, $4, $5 WHERE NOT EXISTS(
			SELECT 1 FROM bookings 
				WHERE doctor_schedule_id = $1 AND mst_schedule_id = $3 AND status = 'WAITING'
		)
	RETURNING id;`

	err := br.db.QueryRow(sqlstat, 
		input.DoctorScheduleID, 
		input.PatientID, 
		input.MstScheduleID, 
		input.Complaint, 
		input.Status, 
		).Scan(&input.ID)
	if err != nil {
		return input, err
	}

	return input, nil
}

func (br bookingRepository) EditSchedule(id uuid.UUID, input entity.Bookings) error {
	sqlstat := "UPDATE bookings SET doctor_schedule_id = $1, mst_schedule_id = $2, complaint = $3 WHERE id = $4"
	_, err := br.db.Exec(sqlstat, 
		input.DoctorScheduleID, 
		input.MstScheduleID, 
		input.Complaint, 
		id,
		)		
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

func (br bookingRepository) CheckExist(doctorScheduleID uuid.UUID, mstScheduleID int) bool {
	exist := false
	sqlstat := "SELECT true FROM bookings WHERE doctor_schedule_id = $1 AND mst_schedule_id = $3;"

	_ = br.db.QueryRow(sqlstat, doctorScheduleID, mstScheduleID).Scan(&exist)
	return exist
}

func scanBookingRows(rows *sql.Rows) ([]entity.Bookings, error) {
	var bookings []entity.Bookings

	defer rows.Close()
	for rows.Next() {
		book := entity.Bookings{}
		err := rows.Scan(
			&book.ID,
			&book.DoctorScheduleID,
			&book.PatientID,
			&book.MstScheduleID,
			&book.Complaint,
			&book.Status,
			&book.ScheduleTime.ID,
			&book.ScheduleTime.StartAt,
			&book.ScheduleTime.EndAt,
		)
		if err != nil {
			return nil, err
		}
		bookings = append(bookings, book)
	}

	return bookings, nil
}
