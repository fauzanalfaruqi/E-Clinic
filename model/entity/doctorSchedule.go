package entity

import (
	"github.com/google/uuid"
)

type DoctorSchedule struct {
	ID        uuid.UUID  `json:"id,omitempty"`
	DoctorID  uuid.UUID  `json:"doctor_id,omitempty"`
	DayOfWeek int        `json:"day_of_week,omitempty"` //day of week represented by integer, start from 0.Sunday--7.Saturday
	StartAt   string     `json:"start_at,omitempty"`
	EndAt     string     `json:"end_at,omitempty"`
	CreatedAt string     `json:"created_at,omitempty"`
	UpdatedAt *string     `json:"updated_at,omitempty"`
	DeletedAt *string     `json:"deleted_at,omitempty"`
	Schedules []Bookings `json:"schedule,omitempty"`
}
