package doctorSchedule

import (
	"avengers-clinic/model/dto"
	"avengers-clinic/model/entity"

	"github.com/google/uuid"
)

type (
	DoctorScheduleRepository interface {
		RetrieveAll(startDate, endDate string) ([]entity.DoctorSchedule, error)
		RetrieveByID(id uuid.UUID) (entity.DoctorSchedule, error)
		InsertSchedule(input dto.CreateDoctorSchedule) (uuid.UUIDs, error)
		GetMySchedule(doctorId uuid.UUID, dayOfWeek []int, startDate, endDate string) ([]entity.DoctorSchedule, error)
		UpdateSchedule(id uuid.UUID, input entity.DoctorSchedule) (error)
		GetByIDs(uuid.UUIDs) ([]entity.DoctorSchedule, error)
		DeleteSchedule(id uuid.UUID) error
		Restore(id uuid.UUID) error
		SearchByDateAndDoctorID(date string, doctorID uuid.UUID) error
	}


	DoctorScheduleUsecase interface {
		GetAll(startDate, endDate string) ([]entity.DoctorSchedule, error)
		GetByID(id uuid.UUID, status string) (entity.DoctorSchedule, error)
		CreateSchedule(input dto.CreateDoctorSchedule) ([]entity.DoctorSchedule, error)
		GetMySchedule(doctorId uuid.UUID, dayOfWeek, status string, startDate, endDate string) ([]entity.DoctorSchedule, error)
		UpdateSchedule(id uuid.UUID, input dto.UpdateSchedule) (entity.DoctorSchedule, error)
		DeleteSchedule(id uuid.UUID) error
		Restore(id uuid.UUID) error
	}
)
