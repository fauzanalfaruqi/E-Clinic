package dto

import (
	"github.com/google/uuid"
)

type (
	CreateBooking struct {
		DoctorScheduleID uuid.UUID `json:"doctor_schedule_id" validate:"required"`
		PatientID        uuid.UUID `json:"patient_id" validate:"required"`
		MstScheduleID    int       `json:"mst_schedule_id" validate:"required"` //refer to mst_schedule id
		Complaint        string    `json:"complaint" validate:"required"`
	}

	UpdateBookingSchedule struct {
		DoctorScheduleID uuid.UUID `json:"doctor_schedule_id"`
		MstScheduleID    int       `json:"mst_schedule_id"` //refer to mst_schedule id
		// ScheduleID uuid.UUID `json:"schedule_id" validate:"required"`
		// PatientID  uuid.UUID `json:"patient_id" validate:"required"`
		// // StartAt    string    `json:"start_at" validate:"required,regex=^((07|08|09|1[0-6]):[0-2][0-9]:[0-5][0-9]|16:30:00)$"`
		// // EndAt      string    `json:"end_at" validate:"required,regex=^((08|09|1[0-6]):[0-5][0-9]:[0-5][0-9]|17:00:00)$"`
		Complaint  string    `json:"complaint" validate:"required"`
	}
)
