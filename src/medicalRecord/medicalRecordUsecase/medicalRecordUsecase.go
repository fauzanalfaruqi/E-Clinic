package medicalRecordUsecase

import (
	"avengers-clinic/model/dto/medicalRecordDTO"
	"avengers-clinic/src/medicalRecord"

	"github.com/gin-gonic/gin"
)

type medicalRecordUsecase struct {
	medicalRecordRepo medicalRecord.MedicalRecordRepository
}

func NewMedicalRecordUsecase(medicalRecordRepo medicalRecord.MedicalRecordRepository) medicalRecord.MedicalRecordUsecase {
	return &medicalRecordUsecase{medicalRecordRepo}
}

func (du *medicalRecordUsecase) CreateMedicalRecord(mr *medicalRecordDTO.CreateMedicalRecord) (*medicalRecordDTO.CreateMedicalRecord, error) {
	var medicalRecord *medicalRecordDTO.CreateMedicalRecord
	var err error

	medicalRecord, err = du.medicalRecordRepo.AddMedicalRecord(mr)
	if err != nil {
		return &medicalRecordDTO.CreateMedicalRecord{}, nil
	}
	return medicalRecord, nil
}

func (du *medicalRecordUsecase) GetMedicalRecords(ctx *gin.Context) ([]medicalRecordDTO.MedicalRecord, error) {
	var medicalRecords []medicalRecordDTO.MedicalRecord
	var err error

	medicalRecords, err = du.medicalRecordRepo.RetrieveMedicalRecords()
	if err != nil {
		return []medicalRecordDTO.MedicalRecord{}, err
	}

	return medicalRecords, nil
}

func (du *medicalRecordUsecase) GetMedicalRecordByID(ctx *gin.Context) (medicalRecordDTO.MedicalRecord, error) {
	id := ctx.Param("id")
	var medicalRecord *medicalRecordDTO.MedicalRecord
	var err error

	if medicalRecord, err = du.medicalRecordRepo.RetrieveMedicalRecordByID(id); err != nil {
		return medicalRecordDTO.MedicalRecord{}, nil
	}

	return *medicalRecord, nil
}
