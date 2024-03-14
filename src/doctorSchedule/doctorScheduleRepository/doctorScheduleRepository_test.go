package doctorScheduleRepository

import (
	"avengers-clinic/model/dto"
	"avengers-clinic/model/entity"
	"avengers-clinic/src/doctorSchedule"
	"database/sql"
	"database/sql/driver"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type doctorRepositoryTestSuite struct {
	suite.Suite
	doctorRepo doctorSchedule.DoctorScheduleRepository
	mock       sqlmock.Sqlmock
}

func (suite *doctorRepositoryTestSuite) SetupTest() {
	db, mock, _ := sqlmock.New()

	suite.doctorRepo = NewDoctorScheduleRepo(db)
	suite.mock = mock
}

func (suite *doctorRepositoryTestSuite) TestRetrieveAllSuccess() {
	startDate := "2024-03-11"
	endDate := "2024-03-17"

	rows := sqlmock.NewRows([]string{
		"id",
		"doctor_id",
		"schedule_date",
		"start_at",
		"end_at",
		"created_at",
		"updated_at",
	})
	suite.mock.ExpectQuery(`SELECT (.+), (.+), (.+), (.+), (.+), (.+), (.+) FROM doctor_schedules`).
		WillReturnRows(rows.AddRow(
			"74d93144-6f2e-4bbc-9f89-973c62d3ac54",
			"5bc18dd0-58cb-4612-8dc3-5fc2419b7f29",
			"2024-03-14",
			1,
			9,
			"2024-03-12 22:39:22.245736",
			"2024-03-12 22:39:22.245736",
		))

	data, err := suite.doctorRepo.RetrieveAll(startDate, endDate)
	suite.NoError(err)
	suite.NotEmpty(data)
}

func (suite *doctorRepositoryTestSuite) TestRetrieveAllFail() {
	startDate := "2024-03-11"
	endDate := "2024-03-17"

	suite.mock.ExpectQuery(`SELECT (.+), (.+), (.+), (.+), (.+), (.+), (.+) FROM doctor_schedules`).
		WillReturnError(sql.ErrConnDone)

	data, err := suite.doctorRepo.RetrieveAll(startDate, endDate)
	suite.Error(err)
	suite.Empty(data)
}

func (suite *doctorRepositoryTestSuite) TestRetrieveByID() {

	id, _ := uuid.Parse("74d93144-6f2e-4bbc-9f89-973c62d3ac54")
	rows := sqlmock.NewRows([]string{
		"id",
		"doctor_id",
		"schedule_date",
		"start_at",
		"end_at",
	})
	suite.mock.ExpectQuery(`SELECT (.+), (.+), (.+), (.+), (.+) FROM doctor_schedules`).
		WillReturnRows(rows.AddRow(
			"74d93144-6f2e-4bbc-9f89-973c62d3ac54",
			"5bc18dd0-58cb-4612-8dc3-5fc2419b7f29",
			"2024-03-14",
			1,
			9,
		))

	data, err := suite.doctorRepo.RetrieveByID(id)
	suite.NoError(err)
	suite.NotEmpty(data)
}

func (suite *doctorRepositoryTestSuite) TestGetMySchedule() {
	doctorID, _ := uuid.Parse("5bc18dd0-58cb-4612-8dc3-5fc2419b7f29")
	dayOfWeeks := []int{}
	startDate := "2024-03-11"
	endDate := "2024-03-17"

	rows := sqlmock.NewRows([]string{
		"id",
		"doctor_id",
		"schedule_date",
		"start_at",
		"end_at",
		"created_at",
		"updated_at",
	})
	suite.mock.ExpectQuery(`SELECT (.+), (.+), (.+), (.+), (.+), (.+), (.+) FROM doctor_schedules`).
		WillReturnRows(rows.AddRow(
			"74d93144-6f2e-4bbc-9f89-973c62d3ac54",
			"5bc18dd0-58cb-4612-8dc3-5fc2419b7f29",
			"2024-03-14",
			1,
			9,
			"2024-03-12 22:39:22.245736",
			"2024-03-12 22:39:22.245736",
		))

	data, err := suite.doctorRepo.GetMySchedule(doctorID, dayOfWeeks, startDate, endDate)
	suite.NoError(err)
	suite.NotEmpty(data)
}

func (suite *doctorRepositoryTestSuite) TestInsertSchedules() {
	doctorID, _ := uuid.Parse("5bc18dd0-58cb-4612-8dc3-5fc2419b7f29")
	args := []driver.Value{"5bc18dd0-58cb-4612-8dc3-5fc2419b7f29", "2024-03-14", 1, 9}

	rows := sqlmock.NewRows([]string{
		"id",
	})
	suite.mock.ExpectQuery(`INSERT INTO doctor_schedules`).
		WithArgs(args...).
		WillReturnRows(rows.AddRow(
			"74d93144-6f2e-4bbc-9f89-973c62d3ac54",
		))

	input := dto.CreateDoctorSchedule{
		DoctorID: doctorID,
		ScheduleDetail: []dto.DoctorScheduleDetail{
			{
				ScheduleDate: "2024-03-14",
				StartAt:      1,
				EndAt:        9,
			},
		},
	}

	ids, err := suite.doctorRepo.InsertSchedule(input)
	suite.NoError(err)
	suite.NotEmpty(ids)
}

func (suite *doctorRepositoryTestSuite) TestGetByIDs() {
	var IDs uuid.UUIDs
	ID, _ := uuid.Parse("5bc18dd0-58cb-4612-8dc3-5fc2419b7f29")
	IDs = append(IDs, ID)

	rows := sqlmock.NewRows([]string{
		"id",
		"doctor_id",
		"schedule_date",
		"start_at",
		"end_at",
		"created_at",
		"updated_at",
	})
	suite.mock.ExpectQuery(`SELECT (.+), (.+), (.+), (.+), (.+), (.+), (.+) FROM doctor_schedules`).
		WillReturnRows(rows.AddRow(
			"74d93144-6f2e-4bbc-9f89-973c62d3ac54",
			"5bc18dd0-58cb-4612-8dc3-5fc2419b7f29",
			"2024-03-14",
			1,
			9,
			"2024-03-12 22:39:22.245736",
			"2024-03-12 22:39:22.245736",
		))

	ids, err := suite.doctorRepo.GetByIDs(IDs)
	suite.NoError(err)
	suite.NotEmpty(ids)
}



func (suite *doctorRepositoryTestSuite) TestUpdateSchedules() {
	id, _ := uuid.Parse("74d93144-6f2e-4bbc-9f89-973c62d3ac54")
	args := []driver.Value{"2024-03-14", 1, 9, "2024-03-12 22:39:22.245736", "74d93144-6f2e-4bbc-9f89-973c62d3ac54"}
	updatedAt := "2024-03-12 22:39:22.245736"

	suite.mock.ExpectExec(`UPDATE doctor_schedules`).
		WithArgs(args...).
		WillReturnResult(sqlmock.NewResult(1,1))

	input := entity.DoctorSchedule{
				ScheduleDate: "2024-03-14",
				StartAt:      1,
				EndAt:        9,
				UpdatedAt: &updatedAt,
		}

	err := suite.doctorRepo.UpdateSchedule(id, input)
	suite.NoError(err)
}

func (suite *doctorRepositoryTestSuite) TestDeleteSchedules() {
	id, _ := uuid.Parse("74d93144-6f2e-4bbc-9f89-973c62d3ac54")

	suite.mock.ExpectExec(`UPDATE doctor_schedules`).
		WithArgs(id).
		WillReturnResult(sqlmock.NewResult(1,1))

	err := suite.doctorRepo.DeleteSchedule(id)
	suite.NoError(err)
}

func (suite *doctorRepositoryTestSuite) TestRestoreSchedules() {
	id, _ := uuid.Parse("74d93144-6f2e-4bbc-9f89-973c62d3ac54")

	suite.mock.ExpectExec(`UPDATE doctor_schedules`).
		WithArgs(id).
		WillReturnResult(sqlmock.NewResult(1,1))

	err := suite.doctorRepo.Restore(id)
	suite.NoError(err)
}




func (suite *doctorRepositoryTestSuite) TestSearchByDateAndDoctorID() {
	id, _ := uuid.Parse("74d93144-6f2e-4bbc-9f89-973c62d3ac54")
	args := []driver.Value{"74d93144-6f2e-4bbc-9f89-973c62d3ac54", "2024-03-12"}
	date := "2024-03-12"

	suite.mock.ExpectQuery(`SELECT true FROM doctor_schedules`).
		WithArgs(args...).
		WillReturnRows(sqlmock.NewRows([]string{"column"}).AddRow(true))

	err := suite.doctorRepo.SearchByDateAndDoctorID(date, id)
	suite.NoError(err)
}



func TestDoctorSchedule(t *testing.T) {
	suite.Run(t, new(doctorRepositoryTestSuite))
}
