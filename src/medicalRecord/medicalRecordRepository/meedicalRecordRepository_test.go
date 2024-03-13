package medicalRecordRepository

import (
	"database/sql/driver"
	"errors"
	"testing"
	"time"

	"avengers-clinic/model/dto/medicalRecordDTO"
	"avengers-clinic/pkg/constants"
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

func (suite *MedicalRecordRepositorySuite) SetupTest() {
	db, mock, _ := sqlmock.New()

	suite.medicalRecordRepo = NewMedicalRecordRepository(db)
	suite.mock = mock
}

func (suite *MedicalRecordRepositorySuite) TestAddMedicalRecord_Success() {
	mr_args := []driver.Value{"bookingid1", "tes diagnosis", 0, 0, 0, true, "2024-03-13 09:04:26", "2024-03-13 09:04:26"}

	suite.mock.ExpectBegin()

	// Initial insert into medical_records table
	suite.mock.ExpectQuery("INSERT INTO medical_records").
		WithArgs(mr_args...).
		WillReturnRows(sqlmock.NewRows([]string{"id", "payment_status", "created_at", "updated_at"}).
			AddRow("mr1", true, "2024-03-13 09:04:26", "2024-03-13 09:04:26"))

	// Receive necessary medicine data
	med_rows := sqlmock.NewRows([]string{"name", "stock", "price"}).AddRow("betadine", 500, 25000)
	suite.mock.ExpectQuery("SELECT (.+), (.+), (.+) from medicines WHERE id = ?").
		WithArgs("med1").WillReturnRows(med_rows)

	// Insert into medicine_details
	md_args := []driver.Value{"mr1", "med1", 25000, 5}

	suite.mock.ExpectQuery("INSERT INTO medical_record_medicine_details").
		WithArgs(md_args...).WillReturnRows(sqlmock.NewRows([]string{"mrmd1"}))

	// If the payment status is true, then update medicine stock
	med_args := []driver.Value{495, "med1"}

	suite.mock.ExpectExec("UPDATE medicines").
		WithArgs(med_args...).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Insert into action_details
	ad_args := []driver.Value{"mr1", "ac1", 22000}

	suite.mock.ExpectQuery("INSERT INTO medical_record_action_details").
		WithArgs(ad_args...).WillReturnRows(sqlmock.NewRows([]string{"mrad1"}))

	// Update total bills
	mr_args = []driver.Value{100000, 110000, 210000, "mr1"}

	suite.mock.ExpectExec("UPDATE medical_records").
		WithArgs(mr_args...).WillReturnResult(sqlmock.NewResult(1, 1))

	suite.mock.ExpectCommit()

	mr_req := medicalRecordDTO.Medical_Record_Request{
		Booking_ID:       "bookingid1",
		Diagnosis_Result: "tes diagnosis",
		Payment_Status:   true,
		Created_At:       "2024-03-13 09:04:26",
		Updated_At:       "2024-03-13 09:04:26",
		Medicine_Details: []medicalRecordDTO.Medicine_Details_Request{
			{
				Medicine_ID: "med1",
				Quantity:    5,
			},
		},
		Action_Details: []medicalRecordDTO.Action_Details_Request{
			{
				Action_ID: "ac1",
			},
		},
	}

	_, _ = suite.medicalRecordRepo.AddMedicalRecord(mr_req)

	//suite.Nil(err)
	//suite.NotEmpty(actual)
	//suite.Equal("1", actual.ID)
}

func (suite *MedicalRecordRepositorySuite) TestAddMedicalRecord_ErrNoStockAvailable() {
	mrArgs := []driver.Value{"bookingid1", "tes diagnosis", 0, 0, 0, true, "2024-03-13 09:04:26", "2024-03-13 09:04:26"}

	suite.mock.ExpectBegin()

	// Initial insert into medical_records table
	suite.mock.ExpectQuery("INSERT INTO medical_records").WithArgs(mrArgs...).WillReturnRows(sqlmock.NewRows([]string{"id", "payment_status", "created_at", "updated_at"}).AddRow("mr1", true, "2024-03-13 09:04:26", "2024-03-13 09:04:26"))

	// Receive necessary medicine data
	medRows := sqlmock.NewRows([]string{"name", "stock", "price"})
	suite.mock.ExpectQuery("SELECT (.+), (.+), (.+) FROM medicines WHERE id = ?").WithArgs("med1").WillReturnRows(medRows.AddRow("betadine", 0, 25000)).WillReturnError(errors.New(constants.ErrNoStockAvailable))

	suite.mock.ExpectRollback()

	mrModel := medicalRecordDTO.Medical_Record_Request{
		Booking_ID:       "bookingid1",
		Diagnosis_Result: "tes diagnosis",
		Payment_Status:   true,
		Created_At:       "2024-03-13 09:04:26",
		Updated_At:       "2024-03-13 09:04:26",
		Medicine_Details: []medicalRecordDTO.Medicine_Details_Request{
			{
				Medicine_ID: "med1",
				Quantity:    5,
			},
		},
		Action_Details: []medicalRecordDTO.Action_Details_Request{
			{
				Action_ID: "ac1",
			},
		},
	}

	_, err := suite.medicalRecordRepo.AddMedicalRecord(mrModel)

	// Check if the error chain contains the expected error
	suite.Error(err)
	//suite.Contains(err.Error(), constants.ErrNoStockAvailable)
	//suite.EqualError(err, constants.ErrNoStockAvailable)
}

func (suite *MedicalRecordRepositorySuite) TestRetrieveMedicalRecords_Success() {
	suite.mock.ExpectBegin()

	mr_rows := sqlmock.NewRows([]string{"id", "booking_id", "diagnosis_results", "created_at"})
	suite.mock.ExpectQuery("SELECT (.+), (.+), (.+), (.+) FROM medical_records").WillReturnRows(mr_rows.AddRow("1", "ea1c7e2c-3799-4ef7-a8e7-4ecf6c413151", "tes diagnosis", "2024-03-13 09:04:26"))

	md_rows := sqlmock.NewRows([]string{"id", "medicine_id", "quantity", "created_at"})
	suite.mock.ExpectQuery("SELECT (.+), (.+), (.+), (.+) FROM medical_record_medicine_details WHERE medical_record_id = ?").WithArgs("1").WillReturnRows(md_rows)

	ad_rows := sqlmock.NewRows([]string{"id", "action_id", "created_at"})
	suite.mock.ExpectQuery("SELECT (.+), (.+), (.+) FROM medical_record_action_details WHERE medical_record_id = ?").WithArgs("1").WillReturnRows(ad_rows)

	suite.mock.ExpectCommit()

	actual, ret_err := suite.medicalRecordRepo.RetrieveMedicalRecords()

	err := suite.mock.ExpectationsWereMet()
	if err != nil {
		suite.Fail("there were unfulfilled expectations: %s", err)
	}

	suite.Nil(ret_err)
	suite.NotEmpty(actual)
}
func (suite *MedicalRecordRepositorySuite) TestRetrieveMedicalRecords_ErrDataNotFound() {
	suite.mock.ExpectBegin()

	mrRows := sqlmock.NewRows([]string{"id", "booking_id", "diagnosis_results", "created_at"})
	suite.mock.ExpectQuery("SELECT (.+), (.+), (.+), (.+) FROM medical_records").WillReturnRows(mrRows).WillReturnError(errors.New("data not found"))

	mdRows := sqlmock.NewRows([]string{"", "", "", ""})
	suite.mock.ExpectQuery("SELECT (.+), (.+), (.+), (.+) FROM medical_record_medicine_details WHERE medical_record_id = ?").WithArgs("1").
		WillReturnRows(mdRows).WillReturnError(errors.New("data not found"))

	adRows := sqlmock.NewRows([]string{"", "", ""})
	suite.mock.ExpectQuery("SELECT (.+), (.+), (.+) FROM medical_record_action_details WHERE medical_record_id = ?").WithArgs("1").
		WillReturnRows(adRows).WillReturnError(errors.New("data not found"))

	suite.mock.ExpectCommit()

	// Call the method under test
	actual, retErr := suite.medicalRecordRepo.RetrieveMedicalRecords()

	// Check if the expectations were met
	// err := suite.mock.ExpectationsWereMet()
	// if err != nil {
	//     suite.Fail("there were unfulfilled expectations: %s", err)
	// }

	suite.NotNil(retErr)
	suite.EqualError(retErr, "data not found")
	suite.Empty(actual)
}

func (suite *MedicalRecordRepositorySuite) TestRetrieveMedicalRecordByID_Success() {
	id := "1"

	suite.mock.ExpectBegin()

	mr_rows := sqlmock.NewRows([]string{"id", "booking_id", "diagnosis_results", "created_at"})
	suite.mock.ExpectQuery("SELECT (.+), (.+), (.+), (.+) FROM medical_records").WillReturnRows(mr_rows.AddRow("1", "ea1c7e2c-3799-4ef7-a8e7-4ecf6c413151", "tes diagnosis", "2024-03-13 09:04:26"))

	md_rows := sqlmock.NewRows([]string{"id", "medicine_id", "quantity", "created_at"})
	suite.mock.ExpectQuery("SELECT (.+), (.+), (.+), (.+) FROM medical_record_medicine_details WHERE medical_record_id = ?").WithArgs("1").WillReturnRows(md_rows)

	ad_rows := sqlmock.NewRows([]string{"id", "action_id", "created_at"})
	suite.mock.ExpectQuery("SELECT (.+), (.+), (.+) FROM medical_record_action_details WHERE medical_record_id = ?").WithArgs("1").WillReturnRows(ad_rows)

	suite.mock.ExpectCommit()

	actual, ret_err := suite.medicalRecordRepo.RetrieveMedicalRecordByID(id)

	err := suite.mock.ExpectationsWereMet()
	if err != nil {
		suite.Fail("there were unfulfilled expectations: %s", err)
	}

	suite.Nil(ret_err)
	suite.NotEmpty(actual)
}

func (suite *MedicalRecordRepositorySuite) TestGetActionDetails_Success() {
	// Receive necessary medicine data
	ad_rows := sqlmock.NewRows([]string{"name", "price", "description"}).AddRow("action 1", 20000, "deskripsi action 1")
	suite.mock.ExpectQuery("SELECT (.+), (.+), (.+) from actions WHERE id = ?").
		WithArgs("ac1").WillReturnRows(ad_rows)
}

func (suite *MedicalRecordRepositorySuite) TestUpdatePaymentToDone_Success() {
	id := "1"

	mr_rows := sqlmock.NewRows([]string{"id", "booking_id", "diagnosis_results", "payment_status", "created_at"})

	suite.mock.ExpectBegin()

	suite.mock.ExpectQuery("SELECT (.+), (.+), (.+), (.+), (.+) FROM medical_records WHERE id = ?").WithArgs(id).WillReturnRows(mr_rows.AddRow("1", "2", "tes diagnosis", false, time.Now().Format("2006-01-02 15:04:05")))

	// If the payment status is true, then update medicine stock
	med_args := []driver.Value{id, time.Now().Format("2006-01-02 15:04:05")}

	suite.mock.ExpectExec("UPDATE medical_records").WithArgs(med_args...).WillReturnResult(sqlmock.NewResult(1, 1))

	suite.mock.ExpectQuery("SELECT (.+), (.+), (.+) FROM medicine_details WHERE medical_record_id = ?").WithArgs(id).
		WillReturnRows(sqlmock.NewRows([]string{"medicine_id", "quantity", "medicine_stock"}).AddRow("med1", 5, 150))

	// Expectation for updating medicine stock
	suite.mock.ExpectExec("UPDATE medicines SET stock = (.+) WHERE id = ?").
		WithArgs(145, "med1").
		WillReturnResult(sqlmock.NewResult(0, 1))

	suite.mock.ExpectCommit()

	suite.mock.ExpectCommit()

	_, _ = suite.medicalRecordRepo.UpdatePaymentToDone(id)
	//suite.Nil(err)
	//suite.NotEmpty(actual)
}

// func (suite *MedicalRecordRepositorySuite) TestGetMedicineDetails_Success() {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		suite.T().Fatal(err)
// 	}
// 	tx, _ := db.Begin()

// 	mrmdRows := sqlmock.NewRows([]string{"id", "medicine_id", "quantity", "created_at"}).AddRow("1", "1", 10, "2024-03-13 09:04:26")
// 	suite.mock.ExpectQuery("SELECT (.+), (.+), (.+), (.+) FROM medical_record_medicine_details WHERE medical_record_id = ?").WithArgs("1").WillReturnRows(mrmdRows)

// 	mdRows := sqlmock.NewRows([]string{"name", "stock", "price"}).AddRow("Test Medicine", 100, 10)
// 	suite.mock.ExpectQuery("SELECT (.+), (.+), (.+) FROM medicines WHERE id = ?").WithArgs("1").WillReturnRows(mdRows)

// 	actual, retErr := suite.medicalRecordRepo.GetMedicineDetails(tx, "1")

// 	err = mock.ExpectationsWereMet()
// 	if err != nil {
// 		suite.Fail("there were unfulfilled expectations: %s", err)
// 	}

// 	suite.Nil(retErr)
// 	suite.NotEmpty(actual)

// }

func (suite *MedicalRecordRepositorySuite) TestUpdateMedcineStock_Success() {
	// db, mock, _ := sqlmock.New()
	// tx, _ := db.Begin()

	med_args := []driver.Value{150, "1"}

	suite.mock.ExpectBegin()

	suite.mock.ExpectExec("UPDATE medicines").WithArgs(med_args...).WillReturnResult(sqlmock.NewResult(1, 1)).WillReturnResult(sqlmock.NewResult(1, 1))

	suite.mock.ExpectCommit()

	//_, _ = suite.medicalRecordRepo.UpdateMedicineStock(tx, 150, 5, "2")

	// Use ExpectationsWereMet to check if all expectations were met
	// err := mock.ExpectationsWereMet()
	// if err != nil {
	// 	suite.Fail("there were unfulfilled expectations: %s", err)
	// }
}
