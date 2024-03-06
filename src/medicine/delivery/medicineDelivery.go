package delivery

import (
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
		medicineGroup.GET("/", handler.getAll)
	}
}

func (m *medicineDelivery) getAll(ctx *gin.Context) {
	getAll, err := m.medicineUC.GetAll()
	if err != nil {
		json.NewResponseError(ctx, err.Error(), "01", "01")
	}
	json.NewResponseSuccess(ctx, getAll, "success", "01", "01")
}
