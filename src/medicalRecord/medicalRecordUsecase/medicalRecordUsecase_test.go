package medicalRecordUsecase

import (
	"avengers-clinic/model/dto/medicalRecordDTO"
	"avengers-clinic/src/medicalRecord"
	"database/sql"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type mockMedicalRecordRepository struct {
	mock.Mock
}

func (m *mockMedicalRecordRepository) AddMedicalRecord(req medicalRecordDTO.Medical_Record_Request) (medicalRecordDTO.Medical_Record, error) {
	args := m.Called(req)
	// TODO: find out what to returns
	return args.Get(0).(medicalRecordDTO.Medical_Record), args.Error(1)
}

func (m *mockMedicalRecordRepository) RetrieveMedicalRecords() ([]medicalRecordDTO.Medical_Record, error) {
	args := m.Called()
	return args.Get(0).([]medicalRecordDTO.Medical_Record), args.Error(1)
}

func (m *mockMedicalRecordRepository) RetrieveMedicalRecordByID(id string) (medicalRecordDTO.Medical_Record, error) {
	args := m.Called(id)
	return args.Get(0).(medicalRecordDTO.Medical_Record), args.Error(1)
}

func (m *mockMedicalRecordRepository) GetMedicineDetails(db *sql.Tx, mrID string) ([]medicalRecordDTO.Medical_Record_Medicine_Details, error) {
	args := m.Called(db, mrID)
	return args.Get(0).([]medicalRecordDTO.Medical_Record_Medicine_Details), args.Error(1)
}

func (m *mockMedicalRecordRepository) GetActionDetails(db *sql.Tx, mrID string) ([]medicalRecordDTO.Medical_Record_Action_Details, error) {
	args := m.Called(db, mrID)
	return args.Get(0).([]medicalRecordDTO.Medical_Record_Action_Details), args.Error(1)
}

func (m *mockMedicalRecordRepository) UpdatePaymentToDone(id string) (medicalRecordDTO.Medical_Record, error) {
	args := m.Called(id)
	return args.Get(0).(medicalRecordDTO.Medical_Record), args.Error(1)
}

func (m *mockMedicalRecordRepository) UpdateMedicineStock(tx *sql.Tx, stock, quantity int, medicineID string) (int, error) {
	args := m.Called(tx, stock, quantity, medicineID)
	return args.Int(0), args.Error(1)
}

type MedicalRecordUsecaseSuite struct {
	suite.Suite
	medicalRecordUsecase  medicalRecord.MedicalRecordUsecase
	medicalRecordRepoMock *mockMedicalRecordRepository
}

func (suite *MedicalRecordUsecaseSuite) SetupTest() {
	suite.medicalRecordRepoMock = new(mockMedicalRecordRepository)
	suite.medicalRecordUsecase = NewMedicalRecordUsecase(suite.medicalRecordRepoMock)
}

func (suite *MedicalRecordUsecaseSuite) TestCreateMedicalRecord_Success() {
	// Create a mock medical record request
	mockRequest := medicalRecordDTO.Medical_Record_Request{
		Booking_ID:       "ea1c7e2c-3799-4ef7-a8e7-4ecf6c413151",
		Diagnosis_Result: "Test diagnosis",
		Payment_Status:   true,
		Created_At:       "2022-03-10 12:00:00",
		Updated_At:       "2022-03-10 12:00:00",
		Medicine_Details: []medicalRecordDTO.Medicine_Details_Request{
			{
				Medicine_ID: "5ad34dce-d1bc-408e-9f82-e5c370cc01f5",
				Quantity:    1,
			},
			{
				Medicine_ID: "83803a11-1388-4beb-b06b-b22f1c98edaf",
				Quantity:    2,
			},
		},
		Action_Details: []medicalRecordDTO.Action_Details_Request{
			{Action_ID: "e6ba6dcc-9c95-4f20-9477-6104c07fef2b"},
		},
	}

	var mds []medicalRecordDTO.Medical_Record_Medicine_Details
	for i := range mockRequest.Medicine_Details {
		var md medicalRecordDTO.Medical_Record_Medicine_Details
		md.Medicine_ID = mockRequest.Medicine_Details[i].Medicine_ID
		md.Quantity = mockRequest.Medicine_Details[i].Quantity
		mds = append(mds, md)
	}

	var ads []medicalRecordDTO.Medical_Record_Action_Details
	for i := range mockRequest.Action_Details {
		var ad medicalRecordDTO.Medical_Record_Action_Details
		ad.Action_ID = mockRequest.Action_Details[i].Action_ID
		ads = append(ads, ad)
	}

	// Define the expected return value from the repository
	expectedMedicalRecord := medicalRecordDTO.Medical_Record{
		ID:               "1",
		Booking_ID:       mockRequest.Booking_ID,
		Diagnosis_Result: mockRequest.Diagnosis_Result,
		Payment_Status:   true,
		Created_At:       mockRequest.Created_At,
		Updated_At:       mockRequest.Updated_At,
		Medicine_Details: mds,
		Action_Details:   ads,
	}

	suite.medicalRecordRepoMock.On("AddMedicalRecord", mockRequest).Return(expectedMedicalRecord, nil)

	createdMedicalRecord, err := suite.medicalRecordUsecase.CreateMedicalRecord(mockRequest)

	suite.NoError(err)

	suite.Equal(expectedMedicalRecord, createdMedicalRecord)

	//suite.medicalRecordRepoMock.AssertCalled(suite.T(), "AddMedicalRecord", mockRequest)
}

func (suite *MedicalRecordUsecaseSuite) TestCreateMedicalRecord_Error() {
	// Define the mock medical record request
	mockRequest := medicalRecordDTO.Medical_Record_Request{
		Booking_ID:       "ea1c7e2c-3799-4ef7-a8e7-4ecf6c413151",
		Diagnosis_Result: "Test diagnosis",
		Payment_Status:   true,
		Created_At:       "2022-03-10 12:00:00",
		Updated_At:       "2022-03-10 12:00:00",
		Medicine_Details: []medicalRecordDTO.Medicine_Details_Request{
			{
				Medicine_ID: "5ad34dce-d1bc-408e-9f82-e5c370cc01f5",
				Quantity:    1,
			},
			{
				Medicine_ID: "83803a11-1388-4beb-b06b-b22f1c98edaf",
				Quantity:    2,
			},
		},
		Action_Details: []medicalRecordDTO.Action_Details_Request{
			{
				Action_ID: "e6ba6dcc-9c95-4f20-9477-6104c07fef2b",
			},
		},
	}

	expectedError := errors.New("repository error")
	suite.medicalRecordRepoMock.On("AddMedicalRecord", mockRequest).Return(medicalRecordDTO.Medical_Record{}, expectedError)

	createdMedicalRecord, err := suite.medicalRecordUsecase.CreateMedicalRecord(mockRequest)

	suite.EqualError(err, expectedError.Error())

	suite.Empty(createdMedicalRecord)

	//suite.medicalRecordRepoMock.AssertCalled(suite.T(), "AddMedicalRecord", mockRequest)
}

func (suite *MedicalRecordUsecaseSuite) TestGetMedicalRecords_Success() {
	expectedMedicalRecords := []medicalRecordDTO.Medical_Record{
		{
			ID:               "1",
			Booking_ID:       "ea1c7e2c-3799-4ef7-a8e7-4ecf6c413151",
			Diagnosis_Result: "Test diagnosis",
			Payment_Status:   true,
			Created_At:       "2022-03-10 12:00:00",
			Updated_At:       "2022-03-10 12:00:00",
			Medicine_Details: []medicalRecordDTO.Medical_Record_Medicine_Details{
				{
					Medicine_ID: "5ad34dce-d1bc-408e-9f82-e5c370cc01f5",
					Quantity:    1,
				},
				{
					Medicine_ID: "83803a11-1388-4beb-b06b-b22f1c98edaf",
					Quantity:    2,
				},
			},
			Action_Details: []medicalRecordDTO.Medical_Record_Action_Details{
				{
					Action_ID: "e6ba6dcc-9c95-4f20-9477-6104c07fef2b",
				},
			},
		},
		{
			ID:               "1",
			Booking_ID:       "ea1c7e2c-3799-4ef7-a8e7-4ecf6c413151",
			Diagnosis_Result: "Test diagnosis 2",
			Payment_Status:   true,
			Created_At:       "2022-03-10 12:00:00",
			Updated_At:       "2022-03-10 12:00:00",
			Medicine_Details: []medicalRecordDTO.Medical_Record_Medicine_Details{
				{
					Medicine_ID: "5ad34dce-d1bc-408e-9f82-e5c370cc01f5",
					Quantity:    1,
				},
				{
					Medicine_ID: "83803a11-1388-4beb-b06b-b22f1c98edaf",
					Quantity:    2,
				},
			},
			Action_Details: []medicalRecordDTO.Medical_Record_Action_Details{
				{
					Action_ID: "e6ba6dcc-9c95-4f20-9477-6104c07fef2b",
				},
			},
		},
	}

	suite.medicalRecordRepoMock.On("RetrieveMedicalRecords").Return(expectedMedicalRecords, nil)

	medicalRecords, err := suite.medicalRecordUsecase.GetMedicalRecords()

	suite.NoError(err)

	suite.Equal(expectedMedicalRecords, medicalRecords)

	//suite.medicalRecordRepoMock.AssertCalled(suite.T(), "RetrieveMedicalRecords")
}

func (suite *MedicalRecordUsecaseSuite) TestGetMedicalRecords_Error() {
	expectedError := errors.New("repository error")

	suite.medicalRecordRepoMock.On("RetrieveMedicalRecords").Return([]medicalRecordDTO.Medical_Record{}, expectedError)

	medicalRecords, err := suite.medicalRecordUsecase.GetMedicalRecords()

	suite.EqualError(err, expectedError.Error())

	suite.Empty(medicalRecords)

	//suite.medicalRecordRepoMock.AssertCalled(suite.T(), "RetrieveMedicalRecords")
}

func (suite *MedicalRecordUsecaseSuite) TestGetMedicalRecordByID_Success() {
	id := "a9a398ce-6c43-473b-a472-055e6c0b5b0c"

	expectedMedicalRecord := medicalRecordDTO.Medical_Record{
		ID:               id,
		Booking_ID:       "ea1c7e2c-3799-4ef7-a8e7-4ecf6c413151",
		Diagnosis_Result: "Test diagnosis",
		Payment_Status:   true,
		Created_At:       "2022-03-10 12:00:00",
		Updated_At:       "2022-03-10 12:00:00",
		Medicine_Details: []medicalRecordDTO.Medical_Record_Medicine_Details{
			{
				Medicine_ID: "5ad34dce-d1bc-408e-9f82-e5c370cc01f5",
				Quantity:    1,
			},
			{
				Medicine_ID: "83803a11-1388-4beb-b06b-b22f1c98edaf",
				Quantity:    2,
			},
		},
		Action_Details: []medicalRecordDTO.Medical_Record_Action_Details{
			{
				Action_ID: "e6ba6dcc-9c95-4f20-9477-6104c07fef2b",
			},
		},
	}

	suite.medicalRecordRepoMock.On("RetrieveMedicalRecordByID", id).Return(expectedMedicalRecord, nil).Once()

	medicalRecord, err := suite.medicalRecordUsecase.GetMedicalRecordByID(id)

	suite.NoError(err)

	suite.Equal(expectedMedicalRecord, medicalRecord)

	//suite.medicalRecordRepoMock.AssertCalled(suite.T(), "RetrieveMedicalRecordByID", id)
}

func (suite *MedicalRecordUsecaseSuite) TestGetMedicalRecordByID_Error() {
	id := "a9a398ce-6c43-473b-a472-055e6c0b5b0c"

	expectedError := errors.New("repository error")

	suite.medicalRecordRepoMock.On("RetrieveMedicalRecordByID", id).Return(medicalRecordDTO.Medical_Record{}, expectedError)

	medicalRecord, err := suite.medicalRecordUsecase.GetMedicalRecordByID(id)

	suite.EqualError(err, expectedError.Error())

	suite.Empty(medicalRecord)

	//suite.medicalRecordRepoMock.AssertCalled(suite.T(), "RetrieveMedicalRecordByID", id)
}

func (suite *MedicalRecordUsecaseSuite) TestUpdatePaymentStatus_Success() {
	id := "a9a398ce-6c43-473b-a472-055e6c0b5b0c"

	expectedMedicalRecord := medicalRecordDTO.Medical_Record{
		ID:               id,
		Booking_ID:       "ea1c7e2c-3799-4ef7-a8e7-4ecf6c413151",
		Diagnosis_Result: "Test diagnosis",
		Payment_Status:   true,
		Created_At:       "2022-03-10 12:00:00",
		Updated_At:       "2022-03-10 12:00:00",
		Medicine_Details: []medicalRecordDTO.Medical_Record_Medicine_Details{
			{
				Medicine_ID: "5ad34dce-d1bc-408e-9f82-e5c370cc01f5",
				Quantity:    1,
			},
			{
				Medicine_ID: "83803a11-1388-4beb-b06b-b22f1c98edaf",
				Quantity:    2,
			},
		},
		Action_Details: []medicalRecordDTO.Medical_Record_Action_Details{
			{
				Action_ID: "e6ba6dcc-9c95-4f20-9477-6104c07fef2b",
			},
		},
	}

	suite.medicalRecordRepoMock.On("UpdatePaymentToDone", id).Return(expectedMedicalRecord, nil)

	updatedMedicalRecord, err := suite.medicalRecordUsecase.UpdatePaymentStatus(id)

	suite.NoError(err)

	suite.Equal(expectedMedicalRecord, updatedMedicalRecord)

	//suite.medicalRecordRepoMock.AssertCalled(suite.T(), "UpdatePaymentToDone", id)
}

func (suite *MedicalRecordUsecaseSuite) TestUpdatePaymentStatus_Error() {
	id := "a9a398ce-6c43-473b-a472-055e6c0b5b0c"

	expectedError := errors.New("repository error")

	suite.medicalRecordRepoMock.On("UpdatePaymentToDone", id).Return(medicalRecordDTO.Medical_Record{}, expectedError)

	updatedMedicalRecord, err := suite.medicalRecordUsecase.UpdatePaymentStatus(id)

	suite.EqualError(err, expectedError.Error())

	suite.Empty(updatedMedicalRecord)

	//suite.medicalRecordRepoMock.AssertCalled(suite.T(), "UpdatePaymentToDone", id)
}

func TestCheckHealthUsecaseSuite(t *testing.T) {
	suite.Run(t, new(MedicalRecordUsecaseSuite))
}
