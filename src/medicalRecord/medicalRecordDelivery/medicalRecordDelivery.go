package medicalRecordDelivery

import (
	"avengers-clinic/model/dto/json"
	"avengers-clinic/model/dto/medicalRecordDTO"
	"avengers-clinic/pkg/utils"
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
	var req medicalRecordDTO.Medical_Record_Request
	var err error

	if err = ctx.ShouldBindJSON(&req); err != nil {
		json.NewResponseBadRequest(ctx, nil, "Error when binding JSON.", "01", "01")
		return
	}

	if errV := utils.Validated(req); errV != nil {
		json.NewResponseBadRequest(ctx, errV, "Bad request", "01", "01")
		return
	}

	medicalRecord, err := dd.medicalRecordUC.CreateMedicalRecord(req)
	if err != nil {
		// if err.Error() == "err1" {
		// 	json.NewResponseBadRequest(ctx, nil, "required fields cannot be empty", "01", "01")
		// 	return
		// }
		json.NewResponseError(ctx, err.Error(), "01", "02")
		return
	}

	json.NewResponseCreated(ctx, medicalRecord, "data created", "01", "01")
}

func (dd *medicalRecordDelivery) getMedicalRecords(ctx *gin.Context) {
	var mrs []medicalRecordDTO.Medical_Record
	var err error

	mrs, err = dd.medicalRecordUC.GetMedicalRecords()
	if err != nil {
		json.NewResponseError(ctx, err.Error(), "01", "01")
		return
	}

	json.NewResponseSuccess(ctx, mrs, "data received", "01", "01")
}

func (dd *medicalRecordDelivery) getMedicalRecordByID(ctx *gin.Context) {
	var mr medicalRecordDTO.Medical_Record
	var err error

	id := ctx.Param("id")

	mr, err = dd.medicalRecordUC.GetMedicalRecordByID(id)
	if err != nil {
		json.NewResponseBadRequest(ctx, nil, "data not found", "01", "01")
		return
	}

	json.NewResponseSuccess(ctx, mr, "data received", "01", "01")
}
