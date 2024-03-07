package medicalRecordDelivery

import (
	"avengers-clinic/model/dto/json"
	"avengers-clinic/model/dto/medicalRecordDTO"
	"avengers-clinic/src/medicalRecord"

	"github.com/gin-gonic/gin"
)

type medicalRecordDelivery struct {
	medicalRecordUC medicalRecord.MedicalRecordUsecase
}

func NewMedicalRecordDelivery(v1Group *gin.RouterGroup, medicalRecordUC medicalRecord.MedicalRecordUsecase) {
	handler := medicalRecordDelivery{
		medicalRecordUC: medicalRecordUC,
	}

	medicalRecordGoup := v1Group.Group("/medical-records")
	{
		medicalRecordGoup.POST("", handler.createMedicalRecord)
		medicalRecordGoup.GET("", handler.getMedicalRecords)
		medicalRecordGoup.GET("/:id", handler.getMedicalRecordByID)
	}
}

func (dd *medicalRecordDelivery) createMedicalRecord(ctx *gin.Context) {
	var cmr medicalRecordDTO.CreateMedicalRecord
	var err error

	if err = ctx.ShouldBindJSON(&cmr); err != nil {
		json.NewResponseBadRequest(ctx, nil, "Error when binding JSON.", "01", "01")
		return
	}

	medicalRecord, err := dd.medicalRecordUC.CreateMedicalRecord(&cmr)
	if err != nil {
		json.NewResponseError(ctx, err.Error(), "01", "01")
		return
	}

	json.NewResponseCreated(ctx, medicalRecord, "data created", "01", "01")
}

func (dd *medicalRecordDelivery) getMedicalRecords(ctx *gin.Context) {
	var mrs []medicalRecordDTO.MedicalRecord
	var err error

	mrs, err = dd.medicalRecordUC.GetMedicalRecords()
	if err != nil {
		json.NewResponseError(ctx, err.Error(), "01", "01")
		return
	}

	json.NewResponseSuccess(ctx, mrs, "data received", "01", "01")
}

func (dd *medicalRecordDelivery) getMedicalRecordByID(ctx *gin.Context) {
	var mr medicalRecordDTO.MedicalRecord
	var err error

	id := ctx.Param("id")

	mr, err = dd.medicalRecordUC.GetMedicalRecordByID(id)
	if err != nil {
		json.NewResponseBadRequest(ctx, nil, "Error when try to get get medical record by id", "01", "01")
		return
	}

	json.NewResponseSuccess(ctx, mr, "data received", "01", "01")
}
