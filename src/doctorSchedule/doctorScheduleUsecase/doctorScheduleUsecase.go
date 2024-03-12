package doctorScheduleUsecase

import (
	"avengers-clinic/model/dto"
	"avengers-clinic/model/entity"
	"avengers-clinic/pkg/constants"
	"avengers-clinic/src/booking"
	"avengers-clinic/src/doctorSchedule"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

type doctorScheduleUsecase struct {
	scheduleRepo doctorSchedule.DoctorScheduleRepository
	bookingRepo  booking.BookingRepository
}

func NewDoctorScheduleUsecase(scheduleRepo doctorSchedule.DoctorScheduleRepository, bookingRepo booking.BookingRepository) doctorSchedule.DoctorScheduleUsecase {
	return &doctorScheduleUsecase{
		scheduleRepo,
		bookingRepo,
	}
}

func (du doctorScheduleUsecase) GetAll(startDate, endDate string) ([]entity.DoctorSchedule, error) {
	var err error

	startDate, endDate, err =  validateStartEndDate(startDate, endDate)
	if err != nil {
		return nil, err
	}

	data, err := du.scheduleRepo.RetrieveAll(startDate, endDate)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (du doctorScheduleUsecase) GetByID(id uuid.UUID, status string) (entity.DoctorSchedule, error) {
	data, err := du.scheduleRepo.RetrieveByID(id)
	if err != nil {
		return data, err
	}

	arrStatus := sanitizeStatusQuery(status)
	data.Schedules, _ = du.bookingRepo.GetBookingByScheduleID(data.ID, arrStatus)

	return data, nil
}

func (du doctorScheduleUsecase) CreateSchedule(input dto.CreateDoctorSchedule) ([]entity.DoctorSchedule, error) {
	// TODO : Add validation for input.doctor_id
	//TODO : Add validation for input.schedule_date
	var err error

	for i, v := range input.ScheduleDetail {

		input.ScheduleDetail[i].ScheduleDate, err = formatDate(v.ScheduleDate)
		if err != nil {
			return nil, fmt.Errorf("invalid date format at index %v", i)
		}

	}

	data, err := du.scheduleRepo.InsertSchedule(input)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (du doctorScheduleUsecase) GetMySchedule(doctorId uuid.UUID, dayOfWeek, status string, startDate, endDate string) ([]entity.DoctorSchedule, error) {
	var err error

	startDate, endDate, err = validateStartEndDate(startDate, endDate)
	if err != nil {
		return nil, err
	}

	//Sanitizing dow query param, replacing string, sql inject etc
	arrDow := sanitizeDowQuery(dayOfWeek)
	arrStatus := sanitizeStatusQuery(status)

	sched, err := du.scheduleRepo.GetMySchedule(doctorId, arrDow, startDate, endDate)
	if err != nil {
		return sched, err
	}

	for i, v := range sched {
		sched[i].Schedules, _ = du.bookingRepo.GetBookingByScheduleID(v.ID, arrStatus)
	}

	return sched, nil
}

func (du doctorScheduleUsecase) UpdateSchedule(id uuid.UUID, input dto.UpdateSchedule) (entity.DoctorSchedule, error) {
	schedule, err := du.scheduleRepo.RetrieveByID(id)
	if err != nil {
		return schedule, err
	}

	if input.ScheduleDate != "" && input.ScheduleDate != schedule.ScheduleDate {

		//TODO : move this validation to validate helper
		sd, err := formatDate(input.ScheduleDate)
		if err != nil {
			return entity.DoctorSchedule{}, err
		}
		schedule.ScheduleDate = sd

		err = du.scheduleRepo.SearchByDateAndDoctorID(input.ScheduleDate, schedule.DoctorID)
		if err == nil {
			return schedule, fmt.Errorf(constants.ErrScheduleDateExist)
		}
	}

	if input.StartAt > 0 {
		schedule.StartAt = input.StartAt
	}

	if input.EndAt > 0 {
		schedule.EndAt = input.EndAt
	}

	//if new updated startAt is greater than endAt, swap the value
	if schedule.StartAt > schedule.EndAt {
		schedule.StartAt, schedule.EndAt = schedule.EndAt, schedule.StartAt
	}

	now := getNow()
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

func sanitizeDowQuery(dayOfWeek string) []int {
	var arrDow []int

	if dayOfWeek != "" {
		strarr := strings.Split(dayOfWeek, "#")
		for _, v := range strarr {
			dow, err := strconv.Atoi(v)
			if err != nil {
				continue
			}
			arrDow = append(arrDow, dow)
		}
	}

	return arrDow

}

func sanitizeStatusQuery(status string) []string {
	var arrStatus []string

	fmt.Println(status)
	if status != "" {
		arrStr := strings.Split(status, "#")
		for _, v := range arrStr {
			v = strings.ToUpper(v)
			switch v {
			case "DONE":
				arrStatus = append(arrStatus, v)
				continue
			case "WAITING":
				arrStatus = append(arrStatus, v)
				continue
			case "CANCELED":
				arrStatus = append(arrStatus, v)
				continue
			default:
				continue
			}
		}
	}

	fmt.Println(arrStatus)
	return arrStatus
}

func validateStartEndDate(startDate, endDate string) (string, string, error) {
	var err error

	if startDate == "" {
		startDate = getDefaultStartDate()
	}else {
		startDate, err = formatDate(startDate)
		if err != nil {
			return "", "", err
		}
	}

	if endDate == "" {
		endDate = getDefaultEndDate()
	}else {
		endDate, err = formatDate(endDate)
		if err != nil {
			return "", "", err
		}
	}

	return startDate, endDate, nil
}

func formatDate(date string) (string, error) {
	d, err := time.Parse("2006-01-02", date)
	if err != nil {
		return "", fmt.Errorf(constants.ErrDateFormat)
	}

	return d.Format("2006-01-02"), nil
}


func getDefaultStartDate() string {
	return getNow()
}

func getDefaultEndDate() string {
	return time.Now().AddDate(0, 1, 0).Format("2006-01-02")
}

func getNow() string {
	return time.Now().Format("2006-01-02 15:04:05")
}