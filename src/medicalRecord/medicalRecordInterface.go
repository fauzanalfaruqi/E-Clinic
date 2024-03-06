package medicalRecord

import (
	"avengers-clinic/model/dto/medicalRecordDTO"

	"github.com/gin-gonic/gin"
)

type MedicalRecordRepository interface {
	AddMedicalRecord(*medicalRecordDTO.CreateMedicalRecord) (*medicalRecordDTO.CreateMedicalRecord, error)
	RetrieveMedicalRecords() ([]medicalRecordDTO.MedicalRecord, error)
	RetrieveMedicalRecordByID(id string) (*medicalRecordDTO.MedicalRecord, error)
}

type MedicalRecordUsecase interface {
	CreateMedicalRecord(mr *medicalRecordDTO.CreateMedicalRecord) (*medicalRecordDTO.CreateMedicalRecord, error)
	GetMedicalRecords(ctx *gin.Context) ([]medicalRecordDTO.MedicalRecord, error)
	GetMedicalRecordByID(ctx *gin.Context) (medicalRecordDTO.MedicalRecord, error)
}
