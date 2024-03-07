package entity

import (
	"time"

	"github.com/google/uuid"
)

type Bookings struct {
	ID         uuid.UUID `json:"id,omitempty"`
	ScheduleID uuid.UUID `json:"schedule_id,omitempty"`
	PatientID  uuid.UUID `json:"patient_id,omitempty"`
	StartAt    time.Time `json:"start_at,omitempty"`
	EndAt      time.Time `json:"end_at,omitempty"`
	Complaint  string    `json:"complaint,omitempty"`
	Status     string    `json:"status,omitempty"` //(waiting, done, canceled)
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
	DeletedAt  time.Time `json:"deleted_at,omitempty"`
}
