package userRepository

import (
	"avengers-clinic/model/dto/userDto"
	"avengers-clinic/src/user"
	"database/sql"
	"database/sql/driver"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
)

type userRepositoryTestSuite struct {
	suite.Suite
	userRepo user.UserRepository
	mock sqlmock.Sqlmock
}

func (suite *userRepositoryTestSuite) SetupTest() {
	db, mock, _ := sqlmock.New()

	suite.mock = mock
	suite.userRepo = NewUserRepository(db)
}

// Start Get ALl Trash
func (suite *userRepositoryTestSuite) TestGetAllTrashSuccess() {
	query := "SELECT (.+), (.+), (.+), (.+), (.+), (.+), (.+), (.+)"
	
	rows := sqlmock.NewRows([]string{"id", "username", "password", "role", "specialization", "created_at", "updated_at", "deleted_at"})

	suite.mock.ExpectQuery(query).WillReturnRows(rows.AddRow("1", "admin", "admin", "ADMIN", nil, "2024-03-12T23:00:00Z", "2024-03-12T23:00:00Z", "2024-03-12T23:00:00Z"))
	actialUsers, err := suite.userRepo.GetAllTrash()
	
	suite.Nil(err)
	suite.NotEmpty(actialUsers)
}

func (suite *userRepositoryTestSuite) TestGetAllTrashError() {
	query := "SELECT (.+), (.+), (.+), (.+), (.+), (.+), (.+), (.+)"

	suite.mock.ExpectQuery(query).WillReturnError(sql.ErrConnDone)
	expectedUser, err := suite.userRepo.GetAllTrash()
	
	suite.Error(err)
	suite.Empty(expectedUser)
}
// End Get ALl Trash

// Start Get ALl
func (suite *userRepositoryTestSuite) TestGetAllSuccess() {
	query := "SELECT (.+), (.+), (.+), (.+), (.+), (.+), (.+), (.+)"
	
	rows := sqlmock.NewRows([]string{"id", "username", "password", "role", "specialization", "created_at", "updated_at", "deleted_at"})

	suite.mock.ExpectQuery(query).WillReturnRows(rows.AddRow("1", "admin", "admin", "ADMIN", nil, "2024-03-12T23:00:00Z", "2024-03-12T23:00:00Z", nil))
	actialUsers, err := suite.userRepo.GetAll()
	
	suite.Nil(err)
	suite.NotEmpty(actialUsers)
}

func (suite *userRepositoryTestSuite) TestGetAllError() {
	query := "SELECT (.+), (.+), (.+), (.+), (.+), (.+), (.+), (.+)"

	suite.mock.ExpectQuery(query).WillReturnError(sql.ErrConnDone)
	expectedUser, err := suite.userRepo.GetAll()
	
	suite.Error(err)
	suite.Empty(expectedUser)
}
// End Get ALl

func (suite *userRepositoryTestSuite) TestGetUserByID() {
	query := "SELECT (.+), (.+), (.+), (.+), (.+), (.+), (.+), (.+) FROM users"

	rows := sqlmock.NewRows([]string{"id", "username", "password", "role", "specialization", "created_at", "updated_at", "deleted_at"})

	suite.mock.ExpectQuery(query).WithArgs("1").WillReturnRows(rows.AddRow("1", "admin", "admin", "ADMIN", nil, "2024-03-12T23:00:00Z", "2024-03-12T23:00:00Z", nil))
	actialUser, err := suite.userRepo.GetUserByID("1")

	suite.Nil(err)
	suite.NotEmpty(actialUser)
}

func (suite *userRepositoryTestSuite) TestGetTrashByID() {
	query := "SELECT (.+), (.+), (.+), (.+), (.+), (.+), (.+), (.+) FROM users"

	rows := sqlmock.NewRows([]string{"id", "username", "password", "role", "specialization", "created_at", "updated_at", "deleted_at"})

	suite.mock.ExpectQuery(query).WithArgs("1").WillReturnRows(rows.AddRow("1", "admin", "admin", "ADMIN", nil, "2024-03-12T23:00:00Z", "2024-03-12T23:00:00Z", nil))
	actialUser, err := suite.userRepo.GetTrashByID("1")

	suite.Nil(err)
	suite.NotEmpty(actialUser)
}

func (suite *userRepositoryTestSuite) TestGetByID() {
	query := "SELECT (.+), (.+), (.+), (.+), (.+), (.+), (.+), (.+) FROM users"
	
	rows := sqlmock.NewRows([]string{"id", "username", "password", "role", "specialization", "created_at", "updated_at", "deleted_at"})

	suite.mock.ExpectQuery(query).WithArgs("1").WillReturnRows(rows.AddRow("1", "admin", "admin", "ADMIN", nil, "2024-03-12T23:00:00Z", "2024-03-12T23:00:00Z", nil))
	actialUser, err := suite.userRepo.GetTrashByID("1")

	suite.Nil(err)
	suite.NotEmpty(actialUser)
}

func (suite *userRepositoryTestSuite) TestGetByUsername() {
	query := "SELECT (.+), (.+), (.+), (.+), (.+), (.+), (.+), (.+) FROM users"
	
	rows := sqlmock.NewRows([]string{"id", "username", "password", "role", "specialization", "created_at", "updated_at", "deleted_at"})

	suite.mock.ExpectQuery(query).WithArgs("admin").WillReturnRows(rows.AddRow("1", "admin", "admin", "ADMIN", nil, "2024-03-12T23:00:00Z", "2024-03-12T23:00:00Z", nil))
	actialUser, err := suite.userRepo.GetTrashByID("admin")

	suite.Nil(err)
	suite.NotEmpty(actialUser)
}

func (suite *userRepositoryTestSuite) TestInsert() {
	args := []driver.Value{"admin", "admin", "ADMIN", nil, "2024-03-12T23:00:00Z", "2024-03-12T23:00:00Z"}
	
	suite.mock.ExpectQuery("INSERT INTO users").
		WithArgs(args...).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("1"))

	user := userDto.User{
		Username: "admin",
		Password: "admin",
		Role: "ADMIN",
		Specialization: nil,
		CreatedAt: "2024-03-12T23:00:00Z",
		UpdatedAt: "2024-03-12T23:00:00Z",
	}
	userID, err := suite.userRepo.Insert(user)

	suite.Nil(err)
	suite.Equal("1", userID)
}

func (suite *userRepositoryTestSuite) TestUpdate() {
	args := []driver.Value{"1", "admin", nil, "2024-03-12T23:00:00Z"}
	
	suite.mock.ExpectExec("UPDATE users").
		WithArgs(args...).
		WillReturnResult(sqlmock.NewResult(1, 1))

	user := userDto.User{
		ID: "1",
		Username: "admin",
		Specialization: nil,
		UpdatedAt: "2024-03-12T23:00:00Z",
	}
	err := suite.userRepo.Update(user)

	suite.Nil(err)
}

func (suite *userRepositoryTestSuite) TestUpdatePassword() {
	args := []driver.Value{"1", "secret"}
	
	suite.mock.ExpectExec("UPDATE users").
		WithArgs(args...).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := suite.userRepo.UpdatePassword("1", "secret")

	suite.Nil(err)
}

func (suite *userRepositoryTestSuite) TestDelete() {
	userID := "1"
	
	suite.mock.ExpectExec("DELETE FROM users").
		WithArgs(userID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := suite.userRepo.Delete(userID)

	suite.Nil(err)
}

func (suite *userRepositoryTestSuite) TestSoftDelete() {
	userID := "1"

	suite.mock.ExpectExec("UPDATE users").
		WithArgs(userID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := suite.userRepo.SoftDelete(userID)

	suite.Nil(err)
}

func (suite *userRepositoryTestSuite) TestRestore() {
	userID := "1"

	suite.mock.ExpectExec("UPDATE users").
		WithArgs(userID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := suite.userRepo.Restore(userID)

	suite.Nil(err)
}

func TestUserDeliveryTestSuite(t *testing.T) {
	suite.Run(t, new(userRepositoryTestSuite))
}