package doctorSchedule

import (
	"avengers-clinic/model/dto"
	"avengers-clinic/model/entity"

	"github.com/google/uuid"
)

type (
	DoctorScheduleRepository interface {
		RetrieveAll(where string) ([]entity.DoctorSchedule, error)
		RetrieveByID(id uuid.UUID) (entity.DoctorSchedule, error)
		InsertSchedule(input dto.CreateDoctorSchedule) ([]entity.DoctorSchedule, error)
		UpdateSchedule(id uuid.UUID, input entity.DoctorSchedule) (error)
		DeleteSchedule(id uuid.UUID) error
		Restore(id uuid.UUID) error
	}

	DoctorScheduleUsecase interface {
		GetAll() ([]entity.DoctorSchedule, error)
		GetByID(id uuid.UUID) (entity.DoctorSchedule, error)
		CreateSchedule(input dto.CreateDoctorSchedule) ([]entity.DoctorSchedule, error)
		UpdateSchedule(id uuid.UUID, input dto.UpdateSchedule) (entity.DoctorSchedule, error)
		DeleteSchedule(id uuid.UUID) error
		Restore(id uuid.UUID) error
	}
)
