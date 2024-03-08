package delivery

import (
	"avengers-clinic/model/dto"
	"avengers-clinic/model/dto/json"
	"avengers-clinic/pkg/utils"
	"avengers-clinic/src/medicine"
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
)

type medicineDelivery struct {
	medicineUC medicine.MedicineUsecase
}

func NewMedicineDelivery(v1Group *gin.RouterGroup, medicineUC medicine.MedicineUsecase) {
	handler := &medicineDelivery{medicineUC: medicineUC}
	medicineGroup := v1Group.Group("/medicines")
	{
		medicineGroup.GET("", handler.getAll)
		medicineGroup.GET("/:id", handler.getById)
		medicineGroup.POST("", handler.create)
		medicineGroup.PUT("/:id", handler.update)
		medicineGroup.DELETE("", handler.delete)
	}
}

func (m *medicineDelivery) getAll(ctx *gin.Context) {
	getAll, err := m.medicineUC.GetAll()
	if err != nil {
		json.NewResponseError(ctx, err.Error(), "01", "01")
		return
	}
	json.NewResponseSuccess(ctx, getAll, "success", "01", "01")
}

func (m *medicineDelivery) getById(ctx *gin.Context) {
	id := (ctx.Param("id"))
	getById, err := m.medicineUC.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			json.NewResponseForbidden(ctx, "Id Not Found", "01", "01")
			return
		}
		json.NewResponseError(ctx, err.Error(), "01", "01")
		return
	}
	json.NewResponseSuccess(ctx, getById, "success", "01", "01")
}

func (m *medicineDelivery) create(ctx *gin.Context) {
	var medicine dto.MedicineRequest
	if err := ctx.ShouldBindJSON(&medicine); err != nil {
		json.NewResponseBadRequest(ctx, []json.ValidationField{}, "error insert new medicine", "01", "01")
		return
	}
	if err := utils.Validated(medicine); err != nil {
		json.NewResponseBadRequest(ctx, err, "bad request", "01", "01")
		return
	}
	fmt.Println(medicine)
	insert, err := m.medicineUC.CreateRecord(medicine)
	if err != nil {
		json.NewResponseError(ctx, err.Error(), "01", "01")
		return
	}
	json.NewResponseCreated(ctx, insert, "succesfully insert new medicine", "01", "01")
}

func (m *medicineDelivery) update(ctx *gin.Context) {
	id := ctx.Param("id")
	var medicine dto.MedicineRequest
	medicine.Id = id
	if err := ctx.ShouldBindJSON(&medicine); err != nil {
		json.NewResponseBadRequest(ctx, []json.ValidationField{}, "bad request", "01", "01")
		return
	}

	insert, err := m.medicineUC.UpdateRecord(medicine)
	if err != nil {
		json.NewResponseError(ctx, err.Error(), "01", "01")
		return
	}
	json.NewResponseSuccess(ctx, insert, "success update medicine", "01", "01")

}

func (m *medicineDelivery) delete(ctx *gin.Context) {

	var prod dto.MedicineRequest
	if err := ctx.ShouldBindJSON(&prod); err != nil {
		json.NewResponseBadRequest(ctx, []json.ValidationField{}, "bad request", "01", "01")
		fmt.Println(err.Error())
		return
	}
	err := m.medicineUC.DeleteRecord(prod.Id)
	if err != nil {
		json.NewResponseError(ctx, err.Error(), "01", "01")
		return
	}

	json.NewResponseSuccess(ctx, prod.Id, "success delete product", "01", "01")
}
