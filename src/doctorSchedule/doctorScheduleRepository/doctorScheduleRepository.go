package doctorScheduleRepository

import (
	"avengers-clinic/model/dto"
	"avengers-clinic/model/entity"
	"avengers-clinic/src/doctorSchedule"
	"database/sql"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

type doctorScheduleRepository struct {
	db *sql.DB
}

func NewDoctorScheduleRepo(db *sql.DB) doctorSchedule.DoctorScheduleRepository {
	return &doctorScheduleRepository{
		db,
	}
}

func (ds doctorScheduleRepository) RetrieveAll(startDate, endDate string) ([]entity.DoctorSchedule, error) {
	sqlstat := `
		SELECT 
				id, 
				doctor_id, 
				to_char(schedule_date, 'YYYY-MM-DD'), 
				start_at, 
				end_at, 
				created_at, 
				updated_at
		FROM doctor_schedules 
		WHERE deleted_at IS NULL AND schedule_date BETWEEN $1 AND $2;`
	rows, err := ds.db.Query(sqlstat, startDate, endDate)
	if err != nil {
		return nil, err
	}
	return scanDoctorSchedules(rows)
}

func (ds doctorScheduleRepository) GetMySchedule(doctorId uuid.UUID, daysOfWeek []int, startDate, endDate string) ([]entity.DoctorSchedule, error) {
	var rows *sql.Rows
	var err error
	sqlstat := `
		SELECT 
				id, 
				doctor_id, 
				to_char(schedule_date, 'YYYY-MM-DD'), 
				start_at, 
				end_at, 
				created_at, 
				updated_at
		FROM doctor_schedules 
		WHERE doctor_id = $1 AND schedule_date BETWEEN $2 AND $3`

	if len(daysOfWeek) > 0 {
		//Search by day of week, represented on int
		//Start from SUNDAY = 0 .... SATURDAY = 6
		sqlstat += " AND EXTRACT(dow from date (schedule_date)) = ANY($4)"
		rows, err = ds.db.Query(sqlstat, doctorId, startDate, endDate, pq.Array(daysOfWeek))
	} else {
		rows, err = ds.db.Query(sqlstat, doctorId, startDate, endDate)
	}

	if err != nil {
		return nil, err
	}

	return scanDoctorSchedules(rows)
}

func (ds doctorScheduleRepository) RetrieveByID(id uuid.UUID) (entity.DoctorSchedule, error) {
	var schedule entity.DoctorSchedule
	sqlstat := `
		SELECT 
				id, 
				doctor_id, 
				to_char(schedule_date, 'YYYY-MM-DD'), 
				start_at, 
				end_at 
		FROM doctor_schedules 
		WHERE id = $1 AND deleted_at IS NULL;`
	err := ds.db.QueryRow(sqlstat, id).Scan(
		&schedule.ID,
		&schedule.DoctorID,
		&schedule.ScheduleDate,
		&schedule.StartAt,
		&schedule.EndAt,
	)
	if err != nil {
		return entity.DoctorSchedule{}, err
	}

	schedule.Schedules = nil

	return schedule, nil

}

func (ds doctorScheduleRepository) InsertSchedule(input dto.CreateDoctorSchedule) (uuid.UUIDs, error) {

	insertQuery := "INSERT INTO doctor_schedules(doctor_id, schedule_date, start_at, end_at) VALUES"
	returnIDQ := " RETURNING id;"

	var inserts []string
	vals := []interface{}{}
	idx := 1

	//Prepare sql statement with $1..., prevent sql injection
	for _, v := range input.ScheduleDetail {
		row := fmt.Sprintf(" ($%v, $%v, $%v, $%v)", idx, idx+1, idx+2, idx+3)
		inserts = append(inserts, row)
		vals = append(vals, input.DoctorID, v.ScheduleDate, v.StartAt, v.EndAt)
		idx += 4
	}

	sqlstat := insertQuery + strings.Join(inserts, ",") + returnIDQ

	//format all vals at once
	rows, err := ds.db.Query(sqlstat, vals...)
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}

	var ids uuid.UUIDs
	for rows.Next() {
		var id uuid.UUID
		err := rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	rows.Close()

	return ids, nil
}

func (ds doctorScheduleRepository) GetByIDs(ids uuid.UUIDs) ([]entity.DoctorSchedule, error) {
	var vals []interface{}

	//Setup placeholders $1... refer by length id
	placeholders := make([]string, len(ids))
	for i, v := range ids {
		placeholders[i] = fmt.Sprintf("$%v", i+1)
		vals = append(vals, v)
	}

	query := fmt.Sprintf("SELECT id, doctor_id, to_char(schedule_date, 'YYYY-MM-DD'), start_at, end_at, created_at, updated_at FROM doctor_schedules WHERE id IN (%s)",
		strings.Join(placeholders, ","))

	rows, err := ds.db.Query(query, vals...)
	if err != nil {
		return nil, err
	}

	return scanDoctorSchedules(rows)
}

func (ds doctorScheduleRepository) UpdateSchedule(id uuid.UUID, data entity.DoctorSchedule) error {

	sqlStat := "UPDATE doctor_schedules SET schedule_date = $1, start_at = $2, end_at = $3, updated_at = $4 WHERE id = $5;"

	_, err := ds.db.Exec(sqlStat, data.ScheduleDate, data.StartAt, data.EndAt, data.UpdatedAt, id)
	if err != nil {
		return err
	}

	return nil
}

func (ds doctorScheduleRepository) DeleteSchedule(id uuid.UUID) error {
	sqlStat := "UPDATE doctor_schedules SET deleted_at = CURRENT_TIMESTAMP WHERE id = $1"
	_, err := ds.db.Exec(sqlStat, id)
	if err != nil {
		return err
	}

	return nil
}

func (ds doctorScheduleRepository) Restore(id uuid.UUID) error {
	sqlStat := "UPDATE doctor_schedules SET deleted_at = NULL WHERE id = $1"
	_, err := ds.db.Exec(sqlStat, id)
	if err != nil {
		return err
	}

	return nil
}

func (ds doctorScheduleRepository) SearchByDateAndDoctorID(date string, doctorID uuid.UUID) error {

	tr := false
	sqlStat := "SELECT true FROM doctor_schedules WHERE doctor_id = $1 AND schedule_date = $2"
	err := ds.db.QueryRow(sqlStat, doctorID, date).Scan(&tr)
	fmt.Println("HERE :", tr)
	if err != nil {
		return err
	}

	return nil
}

func scanDoctorSchedules(rows *sql.Rows) ([]entity.DoctorSchedule, error) {
	var datas []entity.DoctorSchedule
	defer rows.Close()
	for rows.Next() {
		var dt entity.DoctorSchedule
		err := rows.Scan(
			&dt.ID,
			&dt.DoctorID,
			&dt.ScheduleDate,
			&dt.StartAt,
			&dt.EndAt,
			&dt.CreatedAt,
			&dt.UpdatedAt,
		)
		if err != nil {
			return datas, err
		}
		datas = append(datas, dt)
	}

	return datas, nil
}
