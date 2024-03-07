package router

import (
	"avengers-clinic/src/doctorSchedule/doctorScheduleDelivery"
	"avengers-clinic/src/doctorSchedule/doctorScheduleRepository"
	"avengers-clinic/src/doctorSchedule/doctorScheduleUsecase"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func InitRoute(v1Group *gin.RouterGroup, db *sql.DB) {
	scheduleRepo := doctorScheduleRepository.NewDoctorScheduleRepo(db)
	scheduleUC := doctorScheduleUsecase.NewDoctorScheduleUsecase(scheduleRepo)
	doctorScheduleDelivery.NewDoctorScheduleDelivery(v1Group, scheduleUC)
}