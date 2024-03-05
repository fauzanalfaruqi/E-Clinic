package dto

import (
	"github.com/google/uuid"
)

type (
	CreateDoctorSchedule struct {
		DoctorID       uuid.UUID              `json:"doctor_id"`
		ScheduleDetail []DoctorScheduleDetail `json:"schedule_detail" validate:"required"`
	}

	DoctorScheduleDetail struct {
		DayOfWeek int    `json:"day_of_week" validate:"required"` // dow is day of week, represented by integer, start from 0.Sunday--7.Saturday
		StartAt   string `json:"start_at" validate:"required"`
		EndAt     string `json:"end_at" validate:"required"`
	}

	UpdateSchedule struct {
		DayOfWeek int    `json:"day_of_week"`
		StartAt   string `json:"start_at"`
		EndAt     string `json:"end_at"`
	}
)
