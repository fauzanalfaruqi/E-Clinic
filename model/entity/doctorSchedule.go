package entity

import (
	"github.com/google/uuid"
)

type DoctorSchedule struct {
	ID           uuid.UUID  `json:"id,omitempty"`
	DoctorID     uuid.UUID  `json:"doctor_id,omitempty"`
	ScheduleDate string     `json:"schedule_date,omitempty"`
	StartAt      int        `json:"start_at,omitempty"`
	EndAt        int        `json:"end_at,omitempty"`
	CreatedAt    string     `json:"created_at,omitempty"`
	UpdatedAt    *string    `json:"updated_at,omitempty"`
	DeletedAt    *string    `json:"deleted_at,omitempty"`
	Schedules    []Bookings `json:"schedule,omitempty"`
}
