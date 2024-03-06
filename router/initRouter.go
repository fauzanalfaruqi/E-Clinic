package router

import (
	"avengers-clinic/src/medicalRecord/medicalRecordDelivery"
	"avengers-clinic/src/medicalRecord/medicalRecordRepository"
	"avengers-clinic/src/medicalRecord/medicalRecordUsecase"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func InitRoute(v1Group *gin.RouterGroup, db *sql.DB) {
	medicalRecordRepository := medicalRecordRepository.NewMedicalRecordRepository(db)
	medicalRecordUsecase := medicalRecordUsecase.NewMedicalRecordUsecase(medicalRecordRepository)
	medicalRecordDelivery.NewMedicalRecordDelivery(v1Group, medicalRecordUsecase)
}
