package medicalRecordRepository

import (
	"testing"

	"avengers-clinic/src/medicalRecord"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
)

func TestMedicalRecordRepositorySuite(t *testing.T) {
	suite.Run(t, new(MedicalRecordRepositorySuite))
}

type MedicalRecordRepositorySuite struct {
	suite.Suite
	medicalRecordRepo medicalRecord.MedicalRecordRepository
	mock              sqlmock.Sqlmock
}

// func (suite *MedicalRecordRepositorySuite) SetupTest() {

// }

func (suite *MedicalRecordRepositorySuite) TestRetrieveMedicalRecords() {
	// Create a new mock DB connection
	db, mock, err := sqlmock.New()
	if err != nil {
		suite.T().Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	suite.mock = mock
	suite.medicalRecordRepo = NewMedicalRecordRepository(db)

	// Expected query for retrieving medical records
	queryMedicalRecords := "SELECT id, booking_id, diagnosis_results, created_at FROM medical_records WHERE deleted_at IS null"
	rowsMedicalRecords := sqlmock.NewRows([]string{"id", "booking_id", "diagnosis_results", "created_at"}).
		AddRow("1", "ea1c7e2c-3799-4ef7-a8e7-4ecf6c413151", "Test Diagnosis", "2024-03-12T23:00:00Z")
	suite.mock.ExpectQuery(queryMedicalRecords).WillReturnRows(rowsMedicalRecords)

	// Expect the Begin transaction
	suite.mock.ExpectBegin()

	// Expect the commit transaction
	suite.mock.ExpectCommit()

	// Call the method under test
	actual, err := suite.medicalRecordRepo.RetrieveMedicalRecords()

	// Assert that there's no error
	suite.Nil(err)
	// Assert that the result is not empty
	suite.NotEmpty(actual)
	// Assert that the expected queries were executed
	suite.NoError(suite.mock.ExpectationsWereMet())
}
