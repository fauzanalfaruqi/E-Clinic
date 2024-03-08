package medicalRecord

import (
	"avengers-clinic/model/dto/medicalRecordDTO"
	"database/sql"
)

type MedicalRecordRepository interface {
	AddMedicalRecord(medicalRecordDTO.Medical_Record_Request) (medicalRecordDTO.Medical_Record, error)
	RetrieveMedicalRecords() ([]medicalRecordDTO.Medical_Record, error)
	RetrieveMedicalRecordByID(id string) (medicalRecordDTO.Medical_Record, error)
	GetMedicineDetails(db *sql.Tx, mrID string) ([]medicalRecordDTO.Medical_Record_Medicine_Details, error)
	GetActionDetails(db *sql.Tx, mrID string) ([]medicalRecordDTO.Medical_Record_Action_Details, error)
	UpdatePaymentToDone(id string) (medicalRecordDTO.Medical_Record, error)
	UpdateMedicineStock(tx *sql.Tx, stock, quantity int, medicineID string) (int, error)
}

type MedicalRecordUsecase interface {
	CreateMedicalRecord(mr medicalRecordDTO.Medical_Record_Request) (medicalRecordDTO.Medical_Record, error)
	GetMedicalRecords() ([]medicalRecordDTO.Medical_Record, error)
	GetMedicalRecordByID(id string) (medicalRecordDTO.Medical_Record, error)
	UpdatePaymentStatus(id string) (medicalRecordDTO.Medical_Record, error)
}
