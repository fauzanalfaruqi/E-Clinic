package actionUsecase

import (
	"avengers-clinic/model/dto/actionDto"
	"avengers-clinic/src/action"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type mockActionRepository struct {
	mock.Mock
}

func (mock *mockActionRepository) GetAll() ([]actionDto.Action, error) {
	args := mock.Called()	
	return args.Get(0).([]actionDto.Action), args.Error(1)
}

func (mock *mockActionRepository) GetByID(actionID string) (actionDto.Action, error) {
	args := mock.Called(actionID)
	return args.Get(0).(actionDto.Action), args.Error(1)
}

func (mock *mockActionRepository) GetTrashByID(actionID string) (actionDto.Action, error) {
	args := mock.Called(actionID)
	return args.Get(0).(actionDto.Action), args.Error(1)
}

func (mock *mockActionRepository) Insert(action actionDto.Action) (string, error) {
	args := mock.Called(action)
	return args.String(0), args.Error(1)
}

func (mock *mockActionRepository)  Update(action actionDto.Action) error {
	args := mock.Called(action)
	return args.Error(0)
}

func (mock *mockActionRepository) Delete(actionID string) error {
	args := mock.Called(actionID)
	return args.Error(0)
}

func (mock *mockActionRepository) SoftDelete(actionID string) error {
	args := mock.Called(actionID)
	return args.Error(0)
}

func (mock *mockActionRepository) Restore(actionID string) error {
	args := mock.Called(actionID)
	return args.Error(0)
}

func (mock *mockActionRepository) IsNameExist(name string) bool {
	args := mock.Called(name)
	return args.Bool(0)
}

type actionUsecaseTestSuite struct {
	suite.Suite
	actionRepo *mockActionRepository
	actionUC action.ActionUsecase
}

func (suite *actionUsecaseTestSuite) SetupTest() {
	suite.actionRepo = new(mockActionRepository)
	suite.actionUC = NewActionUsecase(suite.actionRepo)
}

func (suite *actionUsecaseTestSuite) TestGetAllSuccess() {
	expected := []actionDto.Action{
		{ID: "1", Name: "Konsultasi", Price: 20000, CreatedAt: "2024-03-12T15:04:05Z", UpdatedAt: "2024-03-12T15:04:05Z"},
	}

	suite.actionRepo.On("GetAll").Return(expected, nil)
	actual, err := suite.actionUC.GetAll()

	suite.Nil(err)
	suite.Equal(expected, actual)
}

// Start Get By ID
func (suite *actionUsecaseTestSuite) TestGetByIDSuccess() {
	expected := actionDto.Action{ID: "1", Name: "Konsultasi", Price: 20000, CreatedAt: "2024-03-12T15:04:05Z", UpdatedAt: "2024-03-12T15:04:05Z"}

	suite.actionRepo.On("GetByID", mock.Anything).Return(expected, nil)
	actual, err := suite.actionUC.GetByID(expected.ID)

	suite.Nil(err)
	suite.Equal(expected, actual)
}

func (suite *actionUsecaseTestSuite) TestGetByIDError() {
	expected := actionDto.Action{}

	suite.actionRepo.On("GetByID", mock.Anything).Return(expected, sql.ErrConnDone)
	actual, err := suite.actionUC.GetByID(expected.ID)

	suite.Error(err)
	suite.Equal(expected, actual)
}
// End Get By ID

// Start Create
func (suite *actionUsecaseTestSuite) TestCreateSuccess() {
	now := time.Now().Format("2006-01-02 15:04:05")
	expected := actionDto.Action{ID: "1", Name: "Konsultasi", Price: 20000, CreatedAt: now, UpdatedAt: now}
	request := actionDto.CreateRequest{Name: "Konsultasi", Price: 20000}

	suite.actionRepo.On("IsNameExist", mock.Anything).Return(false)
	suite.actionRepo.On("Insert", mock.Anything).Return(expected.ID, nil)
	actual, err := suite.actionUC.Create(request)

	suite.Nil(err)
	suite.Equal(expected, actual)
}

func (suite *actionUsecaseTestSuite) TestCreateErrorActionExist() {
	expected := actionDto.Action{}
	request := actionDto.CreateRequest{Name: "Konsultasi", Price: 20000}

	suite.actionRepo.On("IsNameExist", mock.Anything).Return(true)
	actual, err := suite.actionUC.Create(request)

	suite.Error(err)
	suite.Equal(expected, actual)
}

func (suite *actionUsecaseTestSuite) TestCreateInternalServerError() {
	expected := actionDto.Action{}
	request := actionDto.CreateRequest{Name: "Konsultasi", Price: 20000}

	suite.actionRepo.On("IsNameExist", mock.Anything).Return(false)
	suite.actionRepo.On("Insert", mock.Anything).Return("", sql.ErrConnDone)
	actual, err := suite.actionUC.Create(request)

	suite.Error(err)
	suite.Equal(expected, actual)
}
// End Create

// Start Update
func (suite *actionUsecaseTestSuite) TestUpdateSuccess() {
	now := time.Now().Format("2006-01-02 15:04:05")
	expected := actionDto.Action{ID: "1", Name: "Konsultasi", Price: 20000, CreatedAt: now, UpdatedAt: now}
	request := actionDto.UpdateRequest{ID: "1",Name: "Konsultasi Dokter"}

	suite.actionRepo.On("GetByID", mock.Anything).Return(expected, nil)
	suite.actionRepo.On("IsNameExist", mock.Anything).Return(false)
	
	expected.Name = request.Name
	suite.actionRepo.On("Update", mock.Anything).Return(nil)

	actual, err := suite.actionUC.Update(request)

	suite.Nil(err)
	suite.Equal(expected, actual)
}

func (suite *actionUsecaseTestSuite) TestUpdateNotFound() {
	expected := actionDto.Action{}
	request := actionDto.UpdateRequest{ID: "1",Name: "Konsultasi Dokter"}

	suite.actionRepo.On("GetByID", mock.Anything).Return(expected, sql.ErrNoRows)

	actual, err := suite.actionUC.Update(request)

	suite.Error(err)
	suite.Equal(expected, actual)
}

func (suite *actionUsecaseTestSuite) TestUpdateErrorActionExist() {
	expected := actionDto.Action{}
	request := actionDto.UpdateRequest{ID: "1",Name: "Konsultasi Dokter"}

	suite.actionRepo.On("GetByID", mock.Anything).Return(expected, nil)
	suite.actionRepo.On("IsNameExist", mock.Anything).Return(true)

	actual, err := suite.actionUC.Update(request)

	suite.Error(err)
	suite.Equal(expected, actual)
}

func (suite *actionUsecaseTestSuite) TestUpdateInternalServerError() {
	expected := actionDto.Action{}
	request := actionDto.UpdateRequest{ID: "1",Name: "Konsultasi Dokter"}

	suite.actionRepo.On("GetByID", mock.Anything).Return(expected, nil)
	suite.actionRepo.On("IsNameExist", mock.Anything).Return(false)
	suite.actionRepo.On("Update", mock.Anything).Return(sql.ErrConnDone)

	actual, err := suite.actionUC.Update(request)

	suite.Error(err)
	suite.Equal(expected, actual)
}
// End Update

// Start Delete
func (suite *actionUsecaseTestSuite) TestDeleteSuccess()  {
	actionID := "1"
	action := actionDto.Action{ID: "1", Name: "Konsultasi"}

	suite.actionRepo.On("GetByID", mock.Anything).Return(action, nil)
	suite.actionRepo.On("Delete", mock.Anything).Return(nil)
	err := suite.actionUC.Delete(actionID)

	suite.Nil(err)
}

func (suite *actionUsecaseTestSuite) TestDeleteErrorNotFound()  {
	actionID := "1"
	action := actionDto.Action{}

	suite.actionRepo.On("GetByID", mock.Anything).Return(action, sql.ErrNoRows)
	err := suite.actionUC.Delete(actionID)

	suite.Error(err)
}

func (suite *actionUsecaseTestSuite) TestDeleteInternalServerError()  {
	actionID := "1"
	action := actionDto.Action{ID: "1", Name: "Konsultasi"}

	suite.actionRepo.On("GetByID", mock.Anything).Return(action, nil)
	suite.actionRepo.On("Delete", mock.Anything).Return(sql.ErrConnDone)
	err := suite.actionUC.Delete(actionID)

	suite.Error(err)
}
// End Delete

// Start Soft Delete
func (suite *actionUsecaseTestSuite) TestSoftDeleteSuccess()  {
	actionID := "1"
	action := actionDto.Action{ID: "1", Name: "Konsultasi"}

	suite.actionRepo.On("GetByID", mock.Anything).Return(action, nil)
	suite.actionRepo.On("SoftDelete", mock.Anything).Return(nil)
	err := suite.actionUC.SoftDelete(actionID)

	suite.Nil(err)
}

func (suite *actionUsecaseTestSuite) TestSoftDeleteErrorNotFound()  {
	actionID := "1"
	action := actionDto.Action{}

	suite.actionRepo.On("GetByID", mock.Anything).Return(action, sql.ErrNoRows)
	err := suite.actionUC.SoftDelete(actionID)

	suite.Error(err)
}

func (suite *actionUsecaseTestSuite) TestSoftDeleteInternalServerError()  {
	actionID := "1"
	action := actionDto.Action{ID: "1", Name: "Konsultasi"}

	suite.actionRepo.On("GetByID", mock.Anything).Return(action, nil)
	suite.actionRepo.On("SoftDelete", mock.Anything).Return(sql.ErrConnDone)
	err := suite.actionUC.SoftDelete(actionID)

	suite.Error(err)
}
// End Soft Delete

// Start Restore
func (suite *actionUsecaseTestSuite) TestRestoreSuccess()  {
	actionID := "1"
	action := actionDto.Action{ID: "1", Name: "Konsultasi"}

	suite.actionRepo.On("GetTrashByID", mock.Anything).Return(action, nil)
	suite.actionRepo.On("Restore", mock.Anything).Return(nil)
	err := suite.actionUC.Restore(actionID)

	suite.Nil(err)
}

func (suite *actionUsecaseTestSuite) TestRestoreErrorNotFound()  {
	actionID := "1"
	action := actionDto.Action{}

	suite.actionRepo.On("GetTrashByID", mock.Anything).Return(action, sql.ErrNoRows)
	err := suite.actionUC.Restore(actionID)

	suite.Error(err)
}

func (suite *actionUsecaseTestSuite) TestRestoreInternalServerError()  {
	actionID := "1"
	action := actionDto.Action{ID: "1", Name: "Konsultasi"}

	suite.actionRepo.On("GetTrashByID", mock.Anything).Return(action, nil)
	suite.actionRepo.On("Restore", mock.Anything).Return(sql.ErrConnDone)
	err := suite.actionUC.Restore(actionID)

	suite.Error(err)
}
// End Restore

func TestActionUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(actionUsecaseTestSuite))
}