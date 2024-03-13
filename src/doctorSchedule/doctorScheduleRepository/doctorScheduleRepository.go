package doctorScheduleRepository

import (
	"avengers-clinic/model/dto"
	"avengers-clinic/model/entity"
	"avengers-clinic/src/doctorSchedule"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
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

func (ds doctorScheduleRepository) RetrieveAll(where string) ([]entity.DoctorSchedule, error) {
	var datas []entity.DoctorSchedule
	sqlstat := "SELECT id, doctor_id, day_of_week, to_char(start_at, 'HH24:MI:SS'), to_char(end_at, 'HH24:MI:SS'), created_at, updated_at, deleted_at FROM doctor_schedules WHERE deleted_at IS NULL;"
	rows, err := ds.db.Query(sqlstat)
	if err != nil {
		return datas, err
	}
	defer rows.Close()
	for rows.Next() {
		var dt entity.DoctorSchedule
		err := rows.Scan(&dt.ID, &dt.DoctorID, &dt.DayOfWeek, &dt.StartAt, &dt.EndAt, &dt.CreatedAt, &dt.UpdatedAt, &dt.DeletedAt)
		if err != nil {
			return datas, err
		}
		datas = append(datas, dt)
	}

	return datas, nil
}

func (ds doctorScheduleRepository) RetrieveByID(id uuid.UUID) (entity.DoctorSchedule, error) {
	var schedule entity.DoctorSchedule
	sqlstat := "SELECT id, doctor_id, day_of_week, to_char(start_at, 'HH24:MI:SS'), to_char(end_at, 'HH24:MI:SS') FROM doctor_schedules WHERE id = $1 AND deleted_at IS NULL"
	err := ds.db.QueryRow(sqlstat, id).Scan(
		&schedule.ID,
		&schedule.DoctorID,
		&schedule.DayOfWeek,
		&schedule.StartAt,
		&schedule.EndAt,
	)
	if err != nil {
		return entity.DoctorSchedule{}, err
	}

	schedule.Schedules = nil

	return schedule, nil

}

func (ds doctorScheduleRepository) InsertSchedule(input dto.CreateDoctorSchedule) ([]entity.DoctorSchedule, error) {

	insertQuery := "INSERT INTO doctor_schedules(doctor_id, day_of_week, start_at, end_at) VALUES"
	returnIDQ := " RETURNING id;"
	
	var inserts []string
	vals := []interface{}{}
	idx := 1
	for _, v := range input.ScheduleDetail {
		row := fmt.Sprintf(" ($%v, $%v, $%v, $%v)", idx, idx+1, idx+2, idx+3)
		inserts = append(inserts, row)
		vals = append(vals, input.DoctorID, v.DayOfWeek, v.StartAt, v.EndAt)
		idx +=4
	}

	sqlstat := insertQuery + strings.Join(inserts, ",") + returnIDQ

	fmt.Println("VALS : ", vals)
	fmt.Println("STAT : ", sqlstat)

	//prepare the statement
	stmt, err := ds.db.Prepare(sqlstat)
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}

	//close stmt after use
	defer stmt.Close()

	//format all vals at once
	rows, err := stmt.Query(vals...)

	var ids uuid.UUIDs
	for rows.Next() {
		var id uuid.UUID
		err := rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		fmt.Println("id - ", id)
		ids = append(ids, id)
	}
	rows.Close()

	if err != nil {
		return nil, err
	}

	return ds.GetByIDs(ids)
}

// func (ds doctorScheduleRepository) InsertSchedule(input dto.CreateDoctorSchedule) ([]entity.DoctorSchedule, error) {
// 	insertQuery := "INSERT INTO doctor_schedules(doctor_id, day_of_week, start_at, end_at) VALUES"
// 	row := " (?, ?, ?, ?)"
// 	returnIDQ := " RETURNING id;"

// 	var inserts []string
// 	vals := []interface{}{}
// 	for _, v := range input.ScheduleDetail {
// 		inserts = append(inserts, row)
// 		vals = append(vals, []interface{}{input.DoctorID, v.DayOfWeek, v.StartAt, v.EndAt})
// 	}

// 	sqlstat := insertQuery + strings.Join(inserts, ",") + returnIDQ
// 	fmt.Println("STATEMENT : ", sqlstat)

// 	// Prepare the statement
// 	stmt, _ := ds.db.Prepare(sqlstat)
// 	// if err != nil {
// 	// 	log.Error().Msg(err.Error())
// 	// 	return nil, err
// 	// }
// 	defer stmt.Close()

// 	// Execute the statement
// 	var ids []uuid.UUID
// 	for _, val := range vals {
// 		var id uuid.UUID
// 		_, err := stmt.Exec(vals...)
// 		if err != nil {
// 			log.Error().Msg(err.Error())
// 			return nil, err
// 		}
// 		ids = append(ids, id)
// 	}

// 	return ds.GetByIDs(ids)
// }

func (ds doctorScheduleRepository) GetByIDs(ids uuid.UUIDs) ([]entity.DoctorSchedule, error) {
	var vals []interface{}
	var schedules []entity.DoctorSchedule

	placeholders := make([]string, len(ids))
	for i, v := range ids {
		placeholders[i] = fmt.Sprintf("$%v", i+1)
		vals = append(vals, v)
	}

	query := fmt.Sprintf("SELECT id, doctor_id, day_of_week, to_char(start_at, 'HH24:MI:SS'), to_char(end_at, 'HH24:MI:SS'), created_at, updated_at, deleted_at FROM doctor_schedules WHERE id IN (%s)",
		strings.Join(placeholders, ","))

	rows, err := ds.db.Query(query, vals...)
	if err != nil {
		fmt.Println(query)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		schedule := entity.DoctorSchedule{}
		err = rows.Scan(&schedule.ID, &schedule.DoctorID, &schedule.DayOfWeek, &schedule.StartAt, &schedule.EndAt, &schedule.CreatedAt, &schedule.UpdatedAt, &schedule.DeletedAt)
		if err != nil {
			return nil, err
		}
		schedules = append(schedules, schedule)
	}

	return schedules, nil
}

func (ds doctorScheduleRepository) UpdateSchedule(id uuid.UUID, data entity.DoctorSchedule) error {
	sqlStat := "UPDATE doctor_schedules SET day_of_week = $1, start_at = $2, end_at = $3, updated_at = $4 WHERE id = $5;"

	_, err := ds.db.Exec(sqlStat, data.DayOfWeek, data.StartAt, data.EndAt, data.UpdatedAt, id)
	if err != nil {
		return err
	}

	return nil
}

func (ds doctorScheduleRepository) DeleteSchedule(id uuid.UUID) error {
	now := time.Now().Format("2006-01-02 15:04:05")
	sqlStat := "UPDATE doctor_schedules SET deleted_at = $1 WHERE id = $2"
	_, err := ds.db.Exec(sqlStat, now, id)
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
