package doctorScheduleUsecase

import (
	"avengers-clinic/model/dto"
	"avengers-clinic/model/entity"
	"avengers-clinic/src/doctorSchedule"
	"time"

	"github.com/google/uuid"
)

type doctorScheduleUsecase struct {
	scheduleRepo doctorSchedule.DoctorScheduleRepository
}

func NewDoctorScheduleUsecase(scheduleRepo doctorSchedule.DoctorScheduleRepository) doctorSchedule.DoctorScheduleUsecase {
	return &doctorScheduleUsecase{
		scheduleRepo,
	}
}

func (du doctorScheduleUsecase) GetAll() ([]entity.DoctorSchedule, error) {

	data, err := du.scheduleRepo.RetrieveAll("")
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (du doctorScheduleUsecase) GetByID(id uuid.UUID) (entity.DoctorSchedule, error) {
	data, err := du.scheduleRepo.RetrieveByID(id)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (du doctorScheduleUsecase) CreateSchedule(input dto.CreateDoctorSchedule) ([]entity.DoctorSchedule, error) {
	// TODO : Add validation to input.doctor_id
	
	data, err := du.scheduleRepo.InsertSchedule(input)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (du doctorScheduleUsecase) UpdateSchedule(id uuid.UUID, input dto.UpdateSchedule) (entity.DoctorSchedule, error) {
	schedule, err := du.scheduleRepo.RetrieveByID(id)
	if err != nil {
		return schedule, err
	}
	if input.DayOfWeek > 0 {
		schedule.DayOfWeek = input.DayOfWeek
	} else if input.StartAt != "" {
		schedule.StartAt = input.StartAt
	} else if input.EndAt != "" {
		schedule.EndAt = input.EndAt
	}
	now := time.Now().Format("2006-01-02 15:04:05")
	schedule.UpdatedAt = &now

	err = du.scheduleRepo.UpdateSchedule(id, schedule)
	if err != nil {
		return schedule, err
	}

	return schedule, nil

}

func (du doctorScheduleUsecase) DeleteSchedule(id uuid.UUID) error {
	_, err := du.scheduleRepo.RetrieveByID(id)
	if err != nil {
		return err
	}

	err = du.scheduleRepo.DeleteSchedule(id)
	if err != nil {
		return err
	}
	return nil
}

func (du doctorScheduleUsecase) Restore(id uuid.UUID) error {
	err := du.scheduleRepo.Restore(id)
	if err != nil {
		return err
	}
	return nil
}
