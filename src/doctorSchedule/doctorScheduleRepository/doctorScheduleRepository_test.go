package doctorScheduleRepository

import (
	"avengers-clinic/src/doctorSchedule"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
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

func TestDoctorSchedule(t *testing.T) {
	suite.Run(t, new(doctorRepositoryTestSuite))
}
