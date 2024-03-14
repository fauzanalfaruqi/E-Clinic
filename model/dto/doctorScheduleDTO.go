package dto

import (
	"github.com/google/uuid"
)

type (
	CreateDoctorSchedule struct {
		DoctorID       uuid.UUID              `json:"doctor_id" validate:"required,uuid"`
		ScheduleDetail []DoctorScheduleDetail `json:"schedule_detail" validate:"required"`
	}

	DoctorScheduleDetail struct {
		ScheduleDate string `json:"schedule_date" validate:"required"`
		StartAt      int    `json:"start_at" validate:"required"`
		EndAt        int    `json:"end_at" validate:"required,gtfield=StartAt"`
	}

	UpdateSchedule struct {
		ScheduleDate string `json:"schedule_date"`
		StartAt      int    `json:"start_at"`
		EndAt        int    `json:"end_at"`
	}
)
