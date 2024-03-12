package doctorScheduleDelivery

import (
	"avengers-clinic/model/dto"
	"avengers-clinic/model/dto/json"
	"avengers-clinic/pkg/constants"
	"avengers-clinic/pkg/utils"
	"avengers-clinic/src/doctorSchedule"
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type doctorScheduleDelivery struct {
	scheduleUC doctorSchedule.DoctorScheduleUsecase
}

func NewDoctorScheduleDelivery(v1Group *gin.RouterGroup, scheduleUC doctorSchedule.DoctorScheduleUsecase) {
	handler := doctorScheduleDelivery{
		scheduleUC,
	}

	doctorScheduleGroup := v1Group.Group("/doctor-schedule")
	{
		doctorScheduleGroup.GET("", handler.GetAll)
		doctorScheduleGroup.GET("/:id", handler.GetByID)
		doctorScheduleGroup.POST("", handler.CreateSchedule)
		doctorScheduleGroup.PUT("/:id", handler.UpdateSchedule)
		doctorScheduleGroup.DELETE("/:id", handler.DeleteSchedule)
		doctorScheduleGroup.GET("/restore/:id", handler.RestoreSchedule)
		doctorScheduleGroup.GET("/my-schedule/:doctor-id", handler.GetMySchedule)
	}
}

func (dd doctorScheduleDelivery) GetAll(ctx *gin.Context) {

	startDate := ctx.Query("sd")
	endDate := ctx.Query("ed")
	data, err := dd.scheduleUC.GetAll(startDate, endDate)
	if err != nil {
		json.NewResponseError(ctx, err.Error(), "04", "01")
		return
	}

	json.NewResponseSuccess(ctx, data, "success", "04", "01")
}

func (dd doctorScheduleDelivery) GetByID(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		json.NewResponseBadRequest(ctx, nil, err.Error(), constants.DoctorScheduleService, "01")
		return
	}
	status := ctx.Query("status")


	data, err := dd.scheduleUC.GetByID(id, status)
	if err != nil && err == sql.ErrNoRows {
		json.NewResponseBadRequest(ctx, nil, "data not found", constants.DoctorScheduleService, "01")
		return
	} else if err != nil {
		json.NewResponseError(ctx, err.Error(), constants.DoctorScheduleService, "01")
		return
	}

	json.NewResponseSuccess(ctx, data, "success", constants.DoctorScheduleService, "01")

}

func (dd doctorScheduleDelivery) GetMySchedule(ctx *gin.Context) {
	doctorId, err := uuid.Parse(ctx.Param("doctor-id"))
	if err != nil {
		json.NewResponseBadRequest(ctx, nil, err.Error(), constants.DoctorScheduleService, "01")
		return
	}
	dayOfWeeks := ctx.Query("dow")
	status := ctx.Query("status")
	startDate := ctx.Query("sd")
	endDate := ctx.Query("ed")

	data, err := dd.scheduleUC.GetMySchedule(doctorId, dayOfWeeks, status, startDate, endDate)
	if err != nil && err == sql.ErrNoRows {
		json.NewResponseBadRequest(ctx, nil, "data not found", constants.DoctorScheduleService, "01")
		return
	}else if err != nil {
		json.NewResponseError(ctx, err.Error(), constants.DoctorScheduleService, "01")
		return
	}
	json.NewResponseSuccess(ctx, data, "success", constants.DoctorScheduleService, "01")
}

func (dd doctorScheduleDelivery) CreateSchedule(ctx *gin.Context) {
	var input dto.CreateDoctorSchedule

	if err := ctx.ShouldBindJSON(&input); err != nil {
		json.NewResponseError(ctx, err.Error(), constants.DoctorScheduleService, "01")
		return
	}

	if err := utils.Validated(input); err != nil {
		json.NewResponseBadRequest(ctx, err, "Bad request", constants.DoctorScheduleService, "01")
		return
	}

	for _, v := range input.ScheduleDetail {
		if err := utils.Validated(v); err != nil {
			json.NewResponseBadRequest(ctx, err, "Bad request", constants.DoctorScheduleService, "01")
			return
		}
	}

	data, err := dd.scheduleUC.CreateSchedule(input)
	if err != nil {
		json.NewResponseError(ctx, err.Error(), constants.DoctorScheduleService, "01")
		return
	}

	json.NewResponseCreated(ctx, data, "success", constants.DoctorScheduleService, "01")
}

func (dd doctorScheduleDelivery) UpdateSchedule(ctx *gin.Context) {
	var input dto.UpdateSchedule

	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		json.NewResponseBadRequest(ctx, nil, err.Error(), constants.DoctorScheduleService, "01")
		return
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		json.NewResponseError(ctx, err.Error(), constants.DoctorScheduleService, "01")
		return
	}

	if err := utils.Validated(input); err != nil {
		json.NewResponseBadRequest(ctx, err, "Bad request", constants.DoctorScheduleService, "01")
		return
	}

	data, err := dd.scheduleUC.UpdateSchedule(id, input)
	if err != nil && (err == sql.ErrNoRows || err.Error() == constants.ErrScheduleDateExist) {
		json.NewResponseBadRequest(ctx, nil, err.Error(), constants.DoctorScheduleService, "01")
		return
	}else if err != nil {
		json.NewResponseError(ctx, err.Error(), constants.DoctorScheduleService, "01")
		return
	}

	json.NewResponseCreated(ctx, data, "success", constants.DoctorScheduleService, "01")
}

func (dd doctorScheduleDelivery) DeleteSchedule(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		json.NewResponseBadRequest(ctx, nil, err.Error(), constants.DoctorScheduleService, "01")
		return
	}

	err = dd.scheduleUC.DeleteSchedule(id)
	if err != nil && err == sql.ErrNoRows {
		json.NewResponseBadRequest(ctx, nil, "data not found", constants.DoctorScheduleService, "01")
		return
	}else if err != nil {
		json.NewResponseError(ctx, err.Error(), constants.DoctorScheduleService, "01")
		return
	}
	json.NewResponseCreated(ctx, nil, "deleted", constants.DoctorScheduleService, "01")
}


func (dd doctorScheduleDelivery) RestoreSchedule(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		json.NewResponseBadRequest(ctx, nil, err.Error(), constants.DoctorScheduleService, "01")
		return
	}

	err = dd.scheduleUC.Restore(id)
	if err != nil && err == sql.ErrNoRows {
		json.NewResponseBadRequest(ctx, nil, "data not found", constants.DoctorScheduleService, "01")
		return
	}else if err != nil {
		json.NewResponseError(ctx, err.Error(), constants.DoctorScheduleService, "01")
		return
	}
	json.NewResponseCreated(ctx, nil, "restored", constants.DoctorScheduleService, "01")
}