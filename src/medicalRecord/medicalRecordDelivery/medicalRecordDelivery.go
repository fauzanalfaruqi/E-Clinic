package medicalRecordDelivery

import (
	"avengers-clinic/model/dto/json"
	"avengers-clinic/model/dto/medicalRecordDTO"
	"avengers-clinic/pkg/constants"
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
		medicalRecordGoup.PUT("/:id", handler.updatePaymentStatus)
	}
}

func (dd *medicalRecordDelivery) createMedicalRecord(ctx *gin.Context) {
	var req medicalRecordDTO.Medical_Record_Request
	var err error

	if err = ctx.ShouldBindJSON(&req); err != nil {
		json.NewResponseBadRequest(ctx, []json.ValidationField{}, "error when binding JSON.", constants.MedicalRecordService, "01")
		return
	}

	if errV := utils.Validated(req); errV != nil {
		json.NewResponseBadRequest(ctx, errV, "bad request. required fields cannot be empty", constants.MedicalRecordService, "02")
		return
	}

	medicalRecord, err := dd.medicalRecordUC.CreateMedicalRecord(req)
	if err != nil {
		if err.Error() == constants.ErrNoStockAvailable {
			json.NewResponseBadRequest(ctx, []json.ValidationField{}, constants.ErrNoStockAvailable, constants.MedicalRecordService, "03")
			return
		}

		if err.Error() == constants.ErrQuantityGreaterThanStock {
			json.NewResponseBadRequest(ctx, []json.ValidationField{}, constants.ErrQuantityGreaterThanStock, constants.MedicalRecordService, "04")
			return
		}

		json.NewResponseError(ctx, err.Error(), constants.MedicalRecordService, "05")
		return
	}

	json.NewResponseCreated(ctx, medicalRecord, "data created", constants.MedicalRecordService, "01")
}

func (dd *medicalRecordDelivery) getMedicalRecords(ctx *gin.Context) {
	var mrs []medicalRecordDTO.Medical_Record
	var err error

	mrs, err = dd.medicalRecordUC.GetMedicalRecords()
	if err != nil {
		json.NewResponseError(ctx, err.Error(), constants.MedicalRecordService, "01")
		return
	}

	json.NewResponseSuccess(ctx, mrs, "data received", constants.MedicalRecordService, "01")
}

func (dd *medicalRecordDelivery) getMedicalRecordByID(ctx *gin.Context) {
	var mr medicalRecordDTO.Medical_Record
	var err error

	id := ctx.Param("id")

	mr, err = dd.medicalRecordUC.GetMedicalRecordByID(id)
	if err != nil {
		json.NewResponseBadRequest(ctx, []json.ValidationField{}, "data not found", constants.MedicalRecordService, "01")
		return
	}

	json.NewResponseSuccess(ctx, mr, "data received", constants.MedicalRecordService, "01")
}

func (dd *medicalRecordDelivery) updatePaymentStatus(ctx *gin.Context) {
	var mr medicalRecordDTO.Medical_Record
	var err error

	id := ctx.Param("id")

	mr, err = dd.medicalRecordUC.UpdatePaymentStatus(id)
	if err != nil {
		if err.Error() == constants.ErrPaymentAlreadyTrue {
			json.NewResponseBadRequest(ctx, []json.ValidationField{}, constants.ErrPaymentAlreadyTrue, constants.MedicalRecordService, "01")
			return
		}

		if err.Error() == constants.ErrNoStockAvailable {
			json.NewResponseBadRequest(ctx, []json.ValidationField{}, constants.ErrNoStockAvailable, constants.MedicalRecordService, "02")
			return
		}

		if err.Error() == constants.ErrQuantityGreaterThanStock {
			json.NewResponseBadRequest(ctx, []json.ValidationField{}, constants.ErrQuantityGreaterThanStock, constants.MedicalRecordService, "03")
			return
		}
		json.NewResponseBadRequest(ctx, []json.ValidationField{}, "data not found", constants.MedicalRecordService, "02")
		return
	}

	json.NewResponseSuccess(ctx, mr, "data updated", constants.MedicalRecordService, "01")
}
