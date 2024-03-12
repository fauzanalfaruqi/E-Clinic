package medicineDelivery

import (
	"avengers-clinic/model/dto"
	"avengers-clinic/model/dto/json"
	"avengers-clinic/pkg/constants"
	"avengers-clinic/pkg/middleware"
	"avengers-clinic/pkg/utils"
	"avengers-clinic/src/medicine"
	"database/sql"

	"github.com/gin-gonic/gin"
)

type medicineDelivery struct {
	medicineUC medicine.MedicineUsecase
}

func NewMedicineDelivery(v1Group *gin.RouterGroup, medicineUC medicine.MedicineUsecase) {
	handler := &medicineDelivery{medicineUC: medicineUC}
	medicineGroup := v1Group.Group("/medicines")
	{
		medicineGroup.GET("", middleware.JwtAuth("ADMIN"), handler.getAll)
		medicineGroup.GET("/:id", middleware.JwtAuth("ADMIN"), handler.getById)
		medicineGroup.POST("", middleware.JwtAuth("ADMIN"), handler.create)
		medicineGroup.PUT("/:id", middleware.JwtAuth("ADMIN"), handler.update)
		medicineGroup.DELETE("/:id", middleware.JwtAuth("ADMIN"), handler.delete)
		medicineGroup.GET("/trash", middleware.JwtAuth("ADMIN"), handler.trash)
		medicineGroup.PUT("/:id/restore", middleware.JwtAuth("ADMIN"), handler.restore)
	}
}

func (m *medicineDelivery) getAll(ctx *gin.Context) {
	getAll, err := m.medicineUC.GetAll()
	if err != nil {
		json.NewResponseError(ctx, err.Error(), constants.MedicineService, "01")
		return
	}
	
	if len(getAll) == 0 {
		json.NewResponseForbidden(ctx, "medicines not found", constants.MedicineService, "01")
		return
	}
	json.NewResponseSuccess(ctx, getAll, "success", constants.MedicineService, "01")
}

func (m *medicineDelivery) getById(ctx *gin.Context) {
	id := (ctx.Param("id"))
	getById, err := m.medicineUC.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			json.NewResponseForbidden(ctx, "medicine not found", constants.MedicineService, "01")
			return
		}
		json.NewResponseError(ctx, err.Error(), constants.MedicineService, "01")
		return
	}
	json.NewResponseSuccess(ctx, getById, "success", constants.MedicineService, "01")
}

func (m *medicineDelivery) create(ctx *gin.Context) {
	var medicine dto.MedicineRequest
	if err := ctx.ShouldBindJSON(&medicine); err != nil {
		json.NewResponseError(ctx, err.Error(), constants.MedicineService, "01")
		return
	}
	if err := utils.Validated(medicine); err != nil {
		json.NewResponseBadRequest(ctx, err, "bad request", constants.MedicineService, "01")
		return
	}
	insert, err := m.medicineUC.CreateRecord(medicine)
	if err != nil {
		json.NewResponseError(ctx, err.Error(), constants.MedicineService, "01")
		return
	}
	json.NewResponseCreated(ctx, insert, "succesfully insert new medicine", constants.MedicineService, "01")
}

func (m *medicineDelivery) update(ctx *gin.Context) {
	id := ctx.Param("id")
	var medicine dto.UpdateRequest
	medicine.Id = id
	if err := ctx.ShouldBindJSON(&medicine); err != nil {
		json.NewResponseError(ctx, err.Error(), constants.MedicineService, "01")
		return
	}
	insert, err := m.medicineUC.UpdateRecord(medicine)
	if err != nil {
		if err == sql.ErrNoRows {
			json.NewResponseForbidden(ctx, "medicine not found", constants.MedicineService, "01")
			return
		}

		json.NewResponseError(ctx, err.Error(), constants.MedicineService, "01")
		return
	}
	json.NewResponseSuccess(ctx, insert, "success update medicine", constants.MedicineService, "01")

}

func (m *medicineDelivery) delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := m.medicineUC.DeleteRecord(id)
	if err != nil {
		if err == sql.ErrNoRows {
			json.NewResponseForbidden(ctx, "medicine not found", constants.MedicineService, "01")
			return
		}
		
		json.NewResponseError(ctx, err.Error(), constants.MedicineService, "01")
		return
	}

	json.NewResponseSuccess(ctx, id, "success delete medicine", constants.MedicineService, "01")
}

func (m *medicineDelivery) trash(ctx *gin.Context) {
	getAll, err := m.medicineUC.TrashRecord()
	if err != nil {
		json.NewResponseError(ctx, err.Error(), constants.MedicineService, "01")
		return
	}

	if len(getAll) == 0 {
		json.NewResponseForbidden(ctx, "medicines not found", constants.MedicineService, "01")
		return
	}

	json.NewResponseSuccess(ctx, getAll, "success", constants.MedicineService, "01")
}
func (m *medicineDelivery) restore(ctx *gin.Context) {
	id := ctx.Param("id")
	err := m.medicineUC.RestoreRecord(id)
	if err != nil {
		json.NewResponseError(ctx, err.Error(), constants.MedicineService, "01")
		return
	}

	json.NewResponseSuccess(ctx, id, "success restore medicine", constants.MedicineService, "01")
}
