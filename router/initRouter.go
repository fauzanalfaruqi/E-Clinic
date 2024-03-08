package router

import (
	"avengers-clinic/src/booking/bookingDelivery"
	"avengers-clinic/src/booking/bookingRepository"
	"avengers-clinic/src/booking/bookingUsecase"
	"avengers-clinic/src/doctorSchedule/doctorScheduleDelivery"
	"avengers-clinic/src/doctorSchedule/doctorScheduleRepository"
	"avengers-clinic/src/doctorSchedule/doctorScheduleUsecase"
	"avengers-clinic/src/user/userDelivery"
	"avengers-clinic/src/user/userRepository"
	"avengers-clinic/src/user/userUsecase"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func InitRoute(v1Group *gin.RouterGroup, db *sql.DB) {
	scheduleRepo := doctorScheduleRepository.NewDoctorScheduleRepo(db)
	scheduleUC := doctorScheduleUsecase.NewDoctorScheduleUsecase(scheduleRepo)
	doctorScheduleDelivery.NewDoctorScheduleDelivery(v1Group, scheduleUC)

	bookingRepo := bookingRepository.NewBookingRepository(db)
	bookingUC := bookingUsecase.NewBookingUsecase(bookingRepo)
	bookingDelivery.NewBookingDelivery(v1Group, bookingUC)

	userRepository := userRepository.NewUserRepository(db)
	userUsecase := userUsecase.NewUserUsecase(userRepository)
	userDelivery.NewUserDelivery(v1Group, userUsecase)
}
