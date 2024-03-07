package medicalRecordUsecase

import (
	"avengers-clinic/model/dto/medicalRecordDTO"
	"avengers-clinic/src/medicalRecord"
)

type medicalRecordUsecase struct {
	medicalRecordRepo medicalRecord.MedicalRecordRepository
}

func NewMedicalRecordUsecase(medicalRecordRepo medicalRecord.MedicalRecordRepository) medicalRecord.MedicalRecordUsecase {
	return &medicalRecordUsecase{medicalRecordRepo}
}

func (du *medicalRecordUsecase) CreateMedicalRecord(mrr medicalRecordDTO.Medical_Record_Request) (medicalRecordDTO.Medical_Record, error) {
	medicalRecord, err := du.medicalRecordRepo.AddMedicalRecord(mrr)
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
