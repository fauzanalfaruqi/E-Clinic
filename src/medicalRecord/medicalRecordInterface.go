package medicalRecord

import (
	"avengers-clinic/model/dto/medicalRecordDTO"
)

type MedicalRecordRepository interface {
	AddMedicalRecord(*medicalRecordDTO.CreateMedicalRecord) (*medicalRecordDTO.CreateMedicalRecord, error)
	RetrieveMedicalRecords() ([]medicalRecordDTO.MedicalRecord, error)
	RetrieveMedicalRecordByID(id string) (medicalRecordDTO.MedicalRecord, error)
}

type MedicalRecordUsecase interface {
	CreateMedicalRecord(mr *medicalRecordDTO.CreateMedicalRecord) (*medicalRecordDTO.CreateMedicalRecord, error)
	GetMedicalRecords() ([]medicalRecordDTO.MedicalRecord, error)
	GetMedicalRecordByID(id string) (medicalRecordDTO.MedicalRecord, error)
}
