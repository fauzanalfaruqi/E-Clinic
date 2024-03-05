package entity

import (
	"time"

	"github.com/google/uuid"
)

type DoctorSchedule struct {
	ID        uuid.UUID `json:"id,omitempty"`
	DoctorID  uuid.UUID `json:"doctor_id,omitempty"`
	DayOfWeek int       `json:"day_of_week,omitempty"` //day of week represented by integer, start from 0.Sunday--7.Saturday
	StartAt   time.Time `json:"start_at,omitempty"`
	EndAt     time.Time `json:"end_at,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}
