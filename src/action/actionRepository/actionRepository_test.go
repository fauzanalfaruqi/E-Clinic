package actionRepository

import (
	"avengers-clinic/model/dto/actionDto"
	"avengers-clinic/src/action"
	"database/sql"
	"database/sql/driver"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
)

type actionRepositoryTestSuite struct {
	suite.Suite
	actionRepo  action.ActionRepository
	mock sqlmock.Sqlmock
}

func (suite *actionRepositoryTestSuite) SetupTest() {
	db, mock, _ := sqlmock.New()

	suite.actionRepo = NewActionRepository(db)
	suite.mock = mock
}

// Start Get All
func (suite *actionRepositoryTestSuite) TestGetAllSuccess() {
	rows := sqlmock.NewRows([]string{"id", "name", "price", "description", "created_at", "updated_at"})

	suite.mock.ExpectQuery("SELECT (.+), (.+), (.+), (.+), (.+), (.+) FROM actions").
		WillReturnRows(rows.AddRow("1", "Konsultasi", 20000, nil, "2024-03-12T05:20:00Z", "2024-03-12T05:20:00Z"))

	actualActions, err := suite.actionRepo.GetAll()

	suite.Nil(err)
	suite.NotEmpty(actualActions)
}

func (suite *actionRepositoryTestSuite) TestGetAllError() {
	suite.mock.ExpectQuery("SELECT (.+), (.+), (.+), (.+), (.+), (.+) FROM actions").
		WillReturnError(sql.ErrConnDone)

	actualActions, err := suite.actionRepo.GetAll()

	suite.Error(err)
	suite.Empty(actualActions)
}
// End Get All

func (suite *actionRepositoryTestSuite) TestGetByID() {
	actionID := "1"
	rows := sqlmock.NewRows([]string{"id", "name", "price", "description", "created_at", "updated_at"})

	suite.mock.ExpectQuery("SELECT (.+), (.+), (.+), (.+), (.+), (.+) FROM actions").
		WillReturnRows(rows.AddRow("1", "Konsultasi", 20000, nil, "2024-03-12T05:20:00Z", "2024-03-12T05:20:00Z"))

	actualAction, err := suite.actionRepo.GetByID(actionID)

	suite.Nil(err)
	suite.NotEmpty(actualAction)
}

func (suite *actionRepositoryTestSuite) TestGetTrashByID() {
	actionID := "1"
	rows := sqlmock.NewRows([]string{"id", "name", "price", "description", "created_at", "updated_at"})

	suite.mock.ExpectQuery("SELECT (.+), (.+), (.+), (.+), (.+), (.+) FROM actions").
		WillReturnRows(rows.AddRow("1", "Konsultasi", 20000, nil, "2024-03-12T05:20:00Z", "2024-03-12T05:20:00Z"))

	actualAction, err := suite.actionRepo.GetTrashByID(actionID)

	suite.Nil(err)
	suite.NotEmpty(actualAction)
}

func (suite *actionRepositoryTestSuite) TestInsert() {
	args := []driver.Value{"Konsultasi", 20000, nil, "2024-03-12T05:20:00Z", "2024-03-12T05:20:00Z"}

	suite.mock.ExpectQuery("INSERT INTO actions").
		WithArgs(args...).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("1"))

	action := actionDto.Action{
		Name: "Konsultasi",
		Price: 20000,
		CreatedAt: "2024-03-12T05:20:00Z",
		UpdatedAt: "2024-03-12T05:20:00Z",
	}
	actionID, err := suite.actionRepo.Insert(action)

	suite.Nil(err)
	suite.NotEmpty(actionID)
}

func (suite *actionRepositoryTestSuite) TestUpdate() {
	args := []driver.Value{"1", "Konsultasi", 25000, nil, "2024-03-12T05:20:00Z"}

	suite.mock.ExpectExec("UPDATE actions").
		WithArgs(args...).
		WillReturnResult(sqlmock.NewResult(1, 1))

	action := actionDto.Action{
		ID: "1",
		Name: "Konsultasi",
		Price: 25000,
		UpdatedAt: "2024-03-12T05:20:00Z",
	}
	err := suite.actionRepo.Update(action)

	suite.Nil(err)
}

func (suite *actionRepositoryTestSuite) TestDelete() {
	actionID := "1"

	suite.mock.ExpectExec("DELETE FROM actions").
		WithArgs(actionID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := suite.actionRepo.Delete(actionID)

	suite.Nil(err)
}

func (suite *actionRepositoryTestSuite) TestSoftDelete() {
	actionID := "1"

	suite.mock.ExpectExec("UPDATE actions").
		WithArgs(actionID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := suite.actionRepo.SoftDelete(actionID)

	suite.Nil(err)
}

func (suite *actionRepositoryTestSuite) TestRestore() {
	actionID := "1"

	suite.mock.ExpectExec("UPDATE actions").
		WithArgs(actionID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := suite.actionRepo.Restore(actionID)

	suite.Nil(err)
}

func (suite *actionRepositoryTestSuite) TestIsNameExist() {
	name := "Konsultasi"

	suite.mock.ExpectQuery("SELECT (.+) FROM actions").
		WithArgs(name).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))

	actual := suite.actionRepo.IsNameExist(name)
	
	suite.True(actual)
}

func TestActionRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(actionRepositoryTestSuite))
}