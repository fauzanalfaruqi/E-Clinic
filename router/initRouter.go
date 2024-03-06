package router

import (
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
}