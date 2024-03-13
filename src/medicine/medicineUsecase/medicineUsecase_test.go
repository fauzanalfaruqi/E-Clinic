package medicineUsecase

import (
	"avengers-clinic/model/dto"
	"avengers-clinic/src/medicine"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type mockMedicineRepository struct {
	mock.Mock
}

func (mock *mockMedicineRepository) RetrieveAll() ([]dto.MedicineResponse, error) {
	args := mock.Called()
	return args.Get(0).([]dto.MedicineResponse), args.Error(1)
}

func (mock *mockMedicineRepository) RetrieveById(id string) (dto.MedicineResponse, error) {
	args := mock.Called(id)
	return args.Get(0).(dto.MedicineResponse), args.Error(1)
}

func (mock *mockMedicineRepository) Create(medicine dto.MedicineRequest) (dto.MedicineResponse, error) {
	args := mock.Called(medicine)
	return args.Get(0).(dto.MedicineResponse), args.Error(1)
}

func (mock *mockMedicineRepository) Update(medicine dto.MedicineRequest) (dto.MedicineResponse, error) {
	args := mock.Called(medicine)
	return args.Get(0).(dto.MedicineResponse), args.Error(1)
}

func (mock *mockMedicineRepository) Delete(id string, deletedAt string) error {
	args := mock.Called(id, deletedAt)
	return args.Error(0)
}

func (mock *mockMedicineRepository) Trash() ([]dto.MedicineResponse, error) {
	args := mock.Called()
	return args.Get(0).([]dto.MedicineResponse), args.Error(1)
}

func (mock *mockMedicineRepository) Restore(id string) error {
	args := mock.Called(id)
	return args.Error(0)
}

type medicineUsecaseTestSuite struct {
	suite.Suite
	medicineRepo *mockMedicineRepository
	medicineUC medicine.MedicineUsecase
}

func (suite *medicineUsecaseTestSuite) SetupTest() {
	suite.medicineRepo = new(mockMedicineRepository)
	suite.medicineUC = NewMedicineUsecase(suite.medicineRepo)
}

func (suite *medicineUsecaseTestSuite) TestGetAllSuccess() {
	expected := []dto.MedicineResponse{
		{Id: "1", Name: "Komik", MedicineType: "CAIR"},
	}

	suite.medicineRepo.On("RetrieveAll").Return(expected, nil)
	actual, err := suite.medicineUC.GetAll()

	suite.Nil(err)
	suite.Equal(expected, actual)
}

func (suite *medicineUsecaseTestSuite) TestGetByIdSuccess() {
	expected := dto.MedicineResponse{Id: "1", Name: "Komik", MedicineType: "CAIR"}

	suite.medicineRepo.On("RetrieveById", mock.Anything).Return(expected, nil)
	actual, err := suite.medicineUC.GetById("1")

	suite.Nil(err)
	suite.Equal(expected, actual)
}

func (suite *medicineUsecaseTestSuite) TestCreateRecordSuccess() {
	now := time.Now().Format("2006-01-02 15:04:05")
	expected := dto.MedicineResponse{Id: "1", Name: "Komik", MedicineType: "CAIR", CreatedAt: now, UpdatedAt: now}

	request := dto.MedicineRequest{
		Name: "Komik",
		MedicineType: "CAIR",
	}

	suite.medicineRepo.On("Create", mock.Anything).Return(expected, nil)
	actual, err := suite.medicineUC.CreateRecord(request)

	suite.Nil(err)
	suite.Equal(expected, actual)
}

// Start Update
func (suite *medicineUsecaseTestSuite) TestUpdateRecordSuccess() {
	now := time.Now().Format("2006-01-02 15:04:05")
	expected := dto.MedicineResponse{Id: "1", Name: "Komik 1", MedicineType: "CAIR", UpdatedAt: now}

	request := dto.UpdateRequest{
		Name: "Komik 1",
	}

	suite.medicineRepo.On("RetrieveById", mock.Anything).Return(expected, nil)
	suite.medicineRepo.On("Update", mock.Anything).Return(expected, nil)
	actual, err := suite.medicineUC.UpdateRecord(request)

	suite.Nil(err)
	suite.Equal(expected, actual)
}

func (suite *medicineUsecaseTestSuite) TestUpdateRecordErrorNotFound() {
	expected := dto.MedicineResponse{}

	request := dto.UpdateRequest{
		Name: "Komik 1",
	}

	suite.medicineRepo.On("RetrieveById", mock.Anything).Return(expected, sql.ErrNoRows)
	actual, err := suite.medicineUC.UpdateRecord(request)

	suite.Error(err)
	suite.Equal(expected, actual)
}
// End Update

// Start Update
func (suite *medicineUsecaseTestSuite) TestDeleteRecordSuccess() {
	expected := dto.MedicineResponse{Id: "1", Name: "Komik 1", MedicineType: "CAIR"}

	suite.medicineRepo.On("RetrieveById", mock.Anything).Return(expected, nil)
	suite.medicineRepo.On("Delete", mock.Anything, mock.Anything).Return(nil)
	err := suite.medicineUC.DeleteRecord("1")

	suite.Nil(err)
}

func (suite *medicineUsecaseTestSuite) TestDeleteRecordErrorNotFound() {
	expected := dto.MedicineResponse{}

	suite.medicineRepo.On("RetrieveById", mock.Anything).Return(expected, sql.ErrNoRows)
	err := suite.medicineUC.DeleteRecord("1")

	suite.Error(err)
}
// End Update

func (suite *medicineUsecaseTestSuite) TestTrashRecordSuccess() {
	expected := []dto.MedicineResponse{{Id: "1", Name: "Komik 1", MedicineType: "CAIR"}}

	suite.medicineRepo.On("Trash").Return(expected, nil)
	actual, err := suite.medicineUC.TrashRecord()

	suite.Nil(err)
	suite.Equal(expected, actual)
}

func (suite *medicineUsecaseTestSuite) TestRestoreRecordSuccess() {
	suite.medicineRepo.On("Restore", mock.Anything).Return(nil)
	err := suite.medicineUC.RestoreRecord("1")

	suite.Nil(err)
}

func TestMedicineUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(medicineUsecaseTestSuite))
}
