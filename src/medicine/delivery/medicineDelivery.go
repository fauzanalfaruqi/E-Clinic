package delivery

import (
	"avengers-clinic/model/dto"
	"avengers-clinic/model/dto/json"
	"avengers-clinic/src/medicine"

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
	}
}

func (m *medicineDelivery) getAll(ctx *gin.Context) {
	getAll, err := m.medicineUC.GetAll()
	if err != nil {
		json.NewResponseError(ctx, err.Error(), "01", "01")
	}
	json.NewResponseSuccess(ctx, getAll, "success", "01", "01")
}

func (m *medicineDelivery) getById(ctx *gin.Context) {
	id := (ctx.Param("id"))
	getById, err := m.medicineUC.GetById(id)
	if err != nil {
		json.NewResponseError(ctx, err.Error(), "01", "01")
	}
	json.NewResponseSuccess(ctx, getById, "success", "01", "01")
}

func (m *medicineDelivery) create(ctx *gin.Context) {
	var medicine dto.Medicine
	if err := ctx.ShouldBindJSON(&medicine); err != nil {
		json.NewResponseBadRequest(ctx, []json.ValidationField{}, "error insert new medicine", "01", "01")
		return
	}
	insert, err := m.medicineUC.CreateRecord(medicine)
	if err != nil {
		json.NewResponseError(ctx, err.Error(), "01", "01")
		return
	}
	json.NewResponseSuccess(ctx, insert, "succesfuly insert new medicine", "01", "01")
}
