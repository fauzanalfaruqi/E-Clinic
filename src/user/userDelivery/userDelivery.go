package userDelivery

import (
	"avengers-clinic/model/dto/json"
	"avengers-clinic/model/dto/userDto"
	"avengers-clinic/pkg/middleware"
	"avengers-clinic/pkg/utils"
	"avengers-clinic/src/user"

	"github.com/gin-gonic/gin"
)

type userDelivery struct {
	userUC user.UserUsecase
}

func NewUserDelivery(v1Group *gin.RouterGroup, userUC user.UserUsecase) {
	handler := userDelivery{userUC}	

	userGroup := v1Group.Group("/users")
	{
		userGroup.POST("/register",  handler.PatientRegister)
		userGroup.POST("", middleware.JwtAuth("ADMIN"), handler.UserRegister)
		userGroup.POST("/login", handler.Login)
	}
}

func (delivery *userDelivery) PatientRegister(c *gin.Context) {
	var request userDto.AuthRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		json.NewResponseError(c, err.Error(), "01", "01")
		return
	}

	if err := utils.Validated(request); err != nil {
		json.NewResponseBadRequest(c, err, "Bad request", "01", "02")
		return
	}

	user, err := delivery.userUC.PatientRegister(request)
	if err != nil {
		if err.Error() == "1" {
			json.NewResponseBadRequest(c, []json.ValidationField{{FieldName:"username", Message:"Username is already registered"}}, "Bad request", "02", "03")
			return
		}

		json.NewResponseError(c, err.Error(), "01", "04")
		return
	}

	json.NewResponseCreated(c, user, "Patient created successfully.", "01", "01")
}

func (delivery *userDelivery) UserRegister(c *gin.Context) {
	var request userDto.RegisterRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		json.NewResponseError(c, err.Error(), "02", "01")
		return
	}

	if err := utils.Validated(request); err != nil {
		json.NewResponseBadRequest(c, err, "Bad request", "02", "02")
		return
	}

	user, err := delivery.userUC.UserRegister(request)
	if err != nil {
		if err.Error() == "1" {
			json.NewResponseBadRequest(c, []json.ValidationField{{FieldName:"username", Message:"Username is already registered"}}, "Bad request", "02", "03")
			return
		}

		if err.Error() == "2" {
			json.NewResponseBadRequest(c, []json.ValidationField{{FieldName:"role", Message:"Invalid role"}}, "Bad request", "02", "04")
			return
		}

		if err.Error() == "3" {
			json.NewResponseBadRequest(c, []json.ValidationField{{FieldName:"specialization", Message:"field is required"}}, "Bad request", "02", "05")
			return
		}

		json.NewResponseError(c, err.Error(), "02", "06")
		return
	}

	json.NewResponseCreated(c, user, "User created successfully.", "02", "01")
}

func (delivery *userDelivery) Login(c *gin.Context) {
	var request userDto.AuthRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		json.NewResponseError(c, err.Error(), "03", "01")
		return
	}

	if err := utils.Validated(request); err != nil {
		json.NewResponseBadRequest(c, err, "Bad request", "03", "02")
		return
	}

	response, err := delivery.userUC.Login(request)
	if err != nil {
		if err.Error() == "1" {
			json.NewResponseError(c, "Incorrect username or password", "03", "03")
			return
		}

		json.NewResponseError(c, err.Error(), "03", "04")
		return
	}

	json.NewResponseSuccess(c, response, "Login successfully", "03", "01")
}