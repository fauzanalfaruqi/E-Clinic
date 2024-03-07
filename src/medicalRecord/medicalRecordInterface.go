package medicalRecord

import (
	"avengers-clinic/model/dto/medicalRecordDTO"
	"database/sql"
)

type MedicalRecordRepository interface {
	AddMedicalRecord(medicalRecordDTO.Medical_Record_Request) (medicalRecordDTO.Medical_Record, error)
	RetrieveMedicalRecords() ([]medicalRecordDTO.Medical_Record, error)
	RetrieveMedicalRecordByID(id string) (medicalRecordDTO.Medical_Record, error)
	GetMedicineDetails(db *sql.DB, mrID string) ([]medicalRecordDTO.Medical_Record_Medicine_Details, error)
	GetActionDetails(db *sql.DB, mrID string) ([]medicalRecordDTO.Medical_Record_Action_Details, error)
}

type MedicalRecordUsecase interface {
	CreateMedicalRecord(mr medicalRecordDTO.Medical_Record_Request) (medicalRecordDTO.Medical_Record, error)
	GetMedicalRecords() ([]medicalRecordDTO.Medical_Record, error)
	GetMedicalRecordByID(id string) (medicalRecordDTO.Medical_Record, error)
}
