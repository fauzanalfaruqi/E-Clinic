package router

import (
	"avengers-clinic/src/medicine/delivery"
	"avengers-clinic/src/medicine/repository"
	"avengers-clinic/src/medicine/usecase"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func InitRoute(v1Group *gin.RouterGroup, db *sql.DB) {
	medicineR := repository.NewMedicineRepository(db)
	medicineUC := usecase.NewMedicineUsecase(medicineR)
	delivery.NewMedicineDelivery(v1Group, medicineUC)
}
