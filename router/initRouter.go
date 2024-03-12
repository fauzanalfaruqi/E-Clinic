package router

import (
	"avengers-clinic/src/action/actionDelivery"
	"avengers-clinic/src/action/actionRepository"
	"avengers-clinic/src/action/actionUsecase"
	"avengers-clinic/src/booking/bookingDelivery"
	"avengers-clinic/src/booking/bookingRepository"
	"avengers-clinic/src/booking/bookingUsecase"
	"avengers-clinic/src/doctorSchedule/doctorScheduleDelivery"
	"avengers-clinic/src/doctorSchedule/doctorScheduleRepository"
	"avengers-clinic/src/doctorSchedule/doctorScheduleUsecase"
	"avengers-clinic/src/medicalRecord/medicalRecordDelivery"
	"avengers-clinic/src/medicalRecord/medicalRecordRepository"
	"avengers-clinic/src/medicalRecord/medicalRecordUsecase"
	"avengers-clinic/src/medicine/medicineDelivery"
	"avengers-clinic/src/medicine/medicineRepository"
	"avengers-clinic/src/medicine/medicineUsecase"
	"avengers-clinic/src/user/userDelivery"
	"avengers-clinic/src/user/userRepository"
	"avengers-clinic/src/user/userUsecase"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func InitRoute(v1Group *gin.RouterGroup, db *sql.DB) {
	userRepository := userRepository.NewUserRepository(db)
	userUsecase := userUsecase.NewUserUsecase(userRepository)
	userDelivery.NewUserDelivery(v1Group, userUsecase)

	actionRepository := actionRepository.NewActionRepository(db)
	actionUsecase := actionUsecase.NewActionUsecase(actionRepository)
	actionDelivery.NewActionDelivery(v1Group, actionUsecase)

	medicineRepo := medicineRepository.NewMedicineRepository(db)
	medicineUC := medicineUsecase.NewMedicineUsecase(medicineRepo)
	medicineDelivery.NewMedicineDelivery(v1Group, medicineUC)
	
	scheduleRepo := doctorScheduleRepository.NewDoctorScheduleRepo(db)
	scheduleUC := doctorScheduleUsecase.NewDoctorScheduleUsecase(scheduleRepo)
	doctorScheduleDelivery.NewDoctorScheduleDelivery(v1Group, scheduleUC)

	bookingRepo := bookingRepository.NewBookingRepository(db)
	bookingUC := bookingUsecase.NewBookingUsecase(bookingRepo, scheduleRepo)
	bookingDelivery.NewBookingDelivery(v1Group, bookingUC)

	medicalRecordRepository := medicalRecordRepository.NewMedicalRecordRepository(db)
	medicalRecordUsecase := medicalRecordUsecase.NewMedicalRecordUsecase(medicalRecordRepository)
	medicalRecordDelivery.NewMedicalRecordDelivery(v1Group, medicalRecordUsecase)
}
