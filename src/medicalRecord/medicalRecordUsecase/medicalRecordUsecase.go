package medicalRecordUsecase

import (
	"avengers-clinic/model/dto/medicalRecordDTO"
	"avengers-clinic/src/medicalRecord"
	"time"
)

type medicalRecordUsecase struct {
	medicalRecordRepo medicalRecord.MedicalRecordRepository
}

func NewMedicalRecordUsecase(medicalRecordRepo medicalRecord.MedicalRecordRepository) medicalRecord.MedicalRecordUsecase {
	return &medicalRecordUsecase{medicalRecordRepo}
}

func (du *medicalRecordUsecase) CreateMedicalRecord(req medicalRecordDTO.Medical_Record_Request) (medicalRecordDTO.Medical_Record, error) {

	// if req.Booking_ID == "" || req.Diagnosis_Result == "" {
	// 	return medicalRecordDTO.Medical_Record{}, errors.New("err1")
	// }

	// for i := range req.Medicine_Details {
	// 	if req.Medicine_Details[i].Medicine_ID == "" || req.Medicine_Details[i].Quantity >= 0 {
	// 		return medicalRecordDTO.Medical_Record{}, errors.New("err1")
	// 	}
	// }

	// for i := range req.Action_Details {
	// 	if req.Action_Details[i].Action_ID == "" {
	// 		return medicalRecordDTO.Medical_Record{}, errors.New("err1")
	// 	}
	// }

	if req.Created_At == "" {
		req.Created_At = time.Now().Format("2006-01-02 15:04:05")
	}

	if req.Updated_At == "" {
		req.Updated_At = time.Now().Format("2006-01-02 15:04:05")
	}

	medicalRecord, err := du.medicalRecordRepo.AddMedicalRecord(req)
	if err != nil {
		return medicalRecordDTO.Medical_Record{}, err
	}

	return medicalRecord, nil
}

func (du *medicalRecordUsecase) GetMedicalRecords() ([]medicalRecordDTO.Medical_Record, error) {
	var medicalRecords []medicalRecordDTO.Medical_Record
	var err error

	medicalRecords, err = du.medicalRecordRepo.RetrieveMedicalRecords()
	if err != nil {
		return []medicalRecordDTO.Medical_Record{}, err
	}

	return medicalRecords, nil
}

func (du *medicalRecordUsecase) GetMedicalRecordByID(id string) (medicalRecordDTO.Medical_Record, error) {
	var medicalRecord medicalRecordDTO.Medical_Record
	var err error

	if medicalRecord, err = du.medicalRecordRepo.RetrieveMedicalRecordByID(id); err != nil {
		return medicalRecordDTO.Medical_Record{}, err
	}

	return medicalRecord, nil
}

func (du *medicalRecordUsecase) UpdatePaymentStatus(id string) (medicalRecordDTO.Medical_Record, error) {
	var medicalRecord medicalRecordDTO.Medical_Record
	var err error

	if medicalRecord, err = du.medicalRecordRepo.UpdatePaymentToDone(id); err != nil {
		return medicalRecordDTO.Medical_Record{}, err
	}

	return medicalRecord, nil
}
