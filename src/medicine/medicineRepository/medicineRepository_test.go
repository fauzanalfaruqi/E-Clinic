package medicineRepository

import (
	"avengers-clinic/model/dto"
	"avengers-clinic/src/medicine"
	"database/sql"
	"database/sql/driver"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
)

type medicineRepositoryTestSuite struct {
	suite.Suite
	medicineRepo medicine.MedicineRepository
	mock sqlmock.Sqlmock
}

func (suite *medicineRepositoryTestSuite) SetupTest() {
	db, mock, _ := sqlmock.New()

	suite.medicineRepo = NewMedicineRepository(db)
	suite.mock = mock
}

// Start Get All
func (suite *medicineRepositoryTestSuite) TestRetrieveAllSuccess() {
	rows := sqlmock.NewRows([]string{"id", "name", "medicine_type", "price", "stock", "description", "created_at", "updated_at", "deleted_at"})
	row := []driver.Value{"1", "Komik", "CAIR", 5000, 200, nil, "2024-03-12 16:06", "2024-03-12 16:06", ""}

	suite.mock.ExpectQuery("SELECT").
		WillReturnRows(rows.AddRow(row...))
		
	actual, err := suite.medicineRepo.RetrieveAll()

	suite.Nil(err)
	suite.NotEmpty(actual)
}

func (suite *medicineRepositoryTestSuite) TestRetrieveAllError() {
	suite.mock.ExpectQuery("SELECT").
		WillReturnError(sql.ErrConnDone)
		
	actual, err := suite.medicineRepo.RetrieveAll()

	suite.Error(err)
	suite.Empty(actual)
}
// End Get All

// Start Trash
func (suite *medicineRepositoryTestSuite) TestTrashSuccess() {
	rows := sqlmock.NewRows([]string{"id", "name", "medicine_type", "price", "stock", "description", "created_at", "updated_at", "deleted_at"})
	row := []driver.Value{"1", "Komik", "CAIR", 5000, 200, nil, "2024-03-12 16:06", "2024-03-12 16:06", ""}

	suite.mock.ExpectQuery("SELECT").
		WillReturnRows(rows.AddRow(row...))
		
	actual, err := suite.medicineRepo.Trash()

	suite.Nil(err)
	suite.NotEmpty(actual)
}

func (suite *medicineRepositoryTestSuite) TestTrashError() {
	suite.mock.ExpectQuery("SELECT").
		WillReturnError(sql.ErrConnDone)
		
	actual, err := suite.medicineRepo.Trash()

	suite.Error(err)
	suite.Empty(actual)
}
// End Trash

func (suite *medicineRepositoryTestSuite) TestRetrieveByIdSuccess() {
	rows := sqlmock.NewRows([]string{"id", "name", "medicine_type", "price", "stock", "description", "created_at", "updated_at"})
	row := []driver.Value{"1", "Komik", "CAIR", 5000, 200, nil, "2024-03-12 16:06", "2024-03-12 16:06"}

	suite.mock.ExpectQuery("SELECT").
		WithArgs("1").
		WillReturnRows(rows.AddRow(row...))
		
	actual, err := suite.medicineRepo.RetrieveById("1")

	suite.Nil(err)
	suite.NotEmpty(actual)
}

func (suite *medicineRepositoryTestSuite) TestCreateSuccess() {
	rows := sqlmock.NewRows([]string{"id"})
	args := []driver.Value{"Komik", "CAIR", 5000, 200, nil, "2024-03-12 16:06", "2024-03-12 16:06"}
	row := []driver.Value{"1"}

	suite.mock.ExpectQuery("INSERT INTO medicines").
		WithArgs(args...).
		WillReturnRows(rows.AddRow(row...))
		
	request := dto.MedicineRequest{
		Id: "1",
		Name: "Komik",
		MedicineType: "CAIR",
		Price: 5000,
		Stock: 200,
		CreatedAt: "2024-03-12 16:06",
		UpdatedAt: "2024-03-12 16:06",
	}
	
	actual, err := suite.medicineRepo.Create(request)

	suite.Nil(err)
	suite.NotEmpty(actual)
}

func (suite *medicineRepositoryTestSuite) TestUpdateSuccess() {
	rows := sqlmock.NewRows([]string{"id", "name", "medicine_type", "price", "stock", "description", "created_at", "updated_at"})
	args := []driver.Value{"1", "Komik", 5000, 200, nil, "2024-03-12 16:06", "CAIR"}
	row := []driver.Value{"1", "Komik", "CAIR", 5000, 200, nil, "2024-03-12 16:06", "2024-03-12 16:06"}

	suite.mock.ExpectQuery("UPDATE medicines").
		WithArgs(args...).
		WillReturnRows(rows.AddRow(row...))
		
	request := dto.MedicineRequest{
		Id: "1",
		Name: "Komik",
		MedicineType: "CAIR",
		Price: 5000,
		Stock: 200,
		CreatedAt: "2024-03-12 16:06",
		UpdatedAt: "2024-03-12 16:06",
	}
	
	actual, err := suite.medicineRepo.Update(request)

	suite.Nil(err)
	suite.NotEmpty(actual)
}

func (suite *medicineRepositoryTestSuite) TestDeleteSuccess() {
	args := []driver.Value{"2024-03-12 16:06", "1"}

	suite.mock.ExpectExec("UPDATE medicines").
		WithArgs(args...).
		WillReturnResult(sqlmock.NewResult(1, 1))
	
	err := suite.medicineRepo.Delete("1", "2024-03-12 16:06")

	suite.Nil(err)
}

func (suite *medicineRepositoryTestSuite) TestRestoreSuccess() {
	args := []driver.Value{"1"}

	suite.mock.ExpectExec("UPDATE medicines").
		WithArgs(args...).
		WillReturnResult(sqlmock.NewResult(1, 1))
	
	err := suite.medicineRepo.Restore("1")

	suite.Nil(err)
}

func TestMedicineRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(medicineRepositoryTestSuite))
}