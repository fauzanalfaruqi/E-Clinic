package userDelivery

import (
	"avengers-clinic/model/dto/json"
	"avengers-clinic/model/dto/userDto"
	"avengers-clinic/pkg/middleware"
	"avengers-clinic/pkg/utils"
	"avengers-clinic/src/user"
	"database/sql"

	"github.com/gin-gonic/gin"
)

type userDelivery struct {
	userUC user.UserUsecase
}

func NewUserDelivery(v1Group *gin.RouterGroup, userUC user.UserUsecase) {
	handler := userDelivery{userUC}	

	userGroup := v1Group.Group("/users")
	{
		userGroup.GET("/trash", middleware.JwtAuth("ADMIN"), handler.GetAllTrash)
		userGroup.GET("", middleware.JwtAuth("ADMIN"), handler.GetAll)
		userGroup.GET("/:id", middleware.JwtAuth("ADMIN", "DOCTOR", "PATIENT"), handler.GetByID)
		userGroup.POST("/register", middleware.JwtAuth("PATIENT"), handler.PatientRegister)
		userGroup.POST("", middleware.JwtAuth("ADMIN"), handler.UserRegister)
		userGroup.POST("/login", handler.Login)
		userGroup.PUT("/:id", middleware.JwtAuth("ADMIN", "DOCTOR", "PATIENT"), handler.Update)
		userGroup.PUT("/:id/password", middleware.JwtAuth("ADMIN", "DOCTOR", "PATIENT"), handler.UpdatePassword)
		userGroup.DELETE("/:id", middleware.JwtAuth("ADMIN"), handler.Delete)
		userGroup.DELETE("/:id/trash", middleware.JwtAuth("ADMIN"), handler.SoftDelete)
		userGroup.PUT("/:id/restore", middleware.JwtAuth("ADMIN"), handler.Restore)
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
			json.NewResponseBadRequest(c, []json.ValidationField{{FieldName:"username", Message:"Username is already registered"}}, "Bad request", "01", "03")
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
		json.NewResponseError(c, err.Error(), "01", "01")
		return
	}

	if err := utils.Validated(request); err != nil {
		json.NewResponseBadRequest(c, err, "Bad request", "01", "02")
		return
	}

	user, err := delivery.userUC.UserRegister(request)
	if err != nil {
		if err.Error() == "1" {
			json.NewResponseBadRequest(c, []json.ValidationField{{FieldName:"username", Message:"Username is already registered"}}, "Bad request", "01", "03")
			return
		}

		if err.Error() == "2" {
			json.NewResponseBadRequest(c, []json.ValidationField{{FieldName:"specialization", Message:"field is required"}}, "Bad request", "01", "05")
			return
		}

		json.NewResponseError(c, err.Error(), "01", "06")
		return
	}

	json.NewResponseCreated(c, user, "User created successfully.", "01", "01")
}

func (delivery *userDelivery) Login(c *gin.Context) {
	var request userDto.AuthRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		json.NewResponseError(c, err.Error(), "01", "01")
		return
	}

	if err := utils.Validated(request); err != nil {
		json.NewResponseBadRequest(c, err, "Bad request", "01", "02")
		return
	}

	response, err := delivery.userUC.Login(request)
	if err != nil {
		if err.Error() == "1" {
			json.NewResponseError(c, "Incorrect username or password", "01", "03")
			return
		}

		json.NewResponseError(c, err.Error(), "01", "04")
		return
	}

	json.NewResponseSuccess(c, response, "Login successfully", "01", "01")
}

func (delivery *userDelivery) GetAllTrash(c *gin.Context) {
	users, err := delivery.userUC.GetAllTrash()
	if err != nil {
		json.NewResponseError(c, err.Error(), "01", "01")
		return
	}
	
	if len(users) == 0 {
		json.NewResponseForbidden(c, "Users not found", "01", "02")
		return
	}

	json.NewResponseSuccess(c, users, "Users retrieved successfully", "01", "01")
}

func (delivery *userDelivery) GetAll(c *gin.Context) {
	users, err := delivery.userUC.GetAll()
	if err != nil {
		json.NewResponseError(c, err.Error(), "01", "01")
		return
	}
	
	if len(users) == 0 {
		json.NewResponseSuccess(c, nil, "Users not found", "01", "02")
		return
	}

	json.NewResponseSuccess(c, users, "Users retrieved successfully", "01", "01")
}

func (delivery *userDelivery) GetByID(c *gin.Context) {
	userID := c.Param("id")
	user, err := delivery.userUC.GetByID(userID)
	if err != nil {
		if err == sql.ErrNoRows {
			json.NewResponseForbidden(c, "User not found", "01", "01")
			return
		}

		json.NewResponseError(c, err.Error(), "01", "02")
		return
	}

	json.NewResponseSuccess(c, user, "User retrieved successfully", "01", "01")
}

func (delivery *userDelivery) Update(c *gin.Context) {
	var request userDto.UpdateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		json.NewResponseError(c, err.Error(), "01", "01")
		return 
	}

	if err := utils.Validated(request); err != nil {
		json.NewResponseBadRequest(c, err, "Bad request", "01", "02")
		return
	}
	request.ID = c.Param("id")

	response, err := delivery.userUC.Update(request)
	if err != nil {
		if err == sql.ErrNoRows {
			json.NewResponseForbidden(c, "User not found", "01", "03")
			return
		}

		if err.Error() == "1" {
			json.NewResponseBadRequest(c, []json.ValidationField{{FieldName:"username",Message:"Username is already registered"}}, "Bad request", "01", "04")
			return
		}

		json.NewResponseError(c, err.Error(), "01", "05")
		return
	}

	json.NewResponseSuccess(c, response, "User updeted successfully", "01", "01")
}

func (delivery *userDelivery) UpdatePassword(c *gin.Context) {
	var request userDto.UpdatePasswordRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		json.NewResponseError(c, err.Error(), "01", "01")
		return 
	}

	if err := utils.Validated(request); err != nil {
		json.NewResponseBadRequest(c, err, "Bad request", "01", "02")
		return
	}
	request.ID = c.Param("id")

	err := delivery.userUC.UpdatePassword(request)
	if err != nil {
		if err == sql.ErrNoRows {
			json.NewResponseForbidden(c, "User not found", "01", "03")
			return
		}

		if err.Error() == "1" {
			json.NewResponseBadRequest(c, []json.ValidationField{{FieldName:"current_password",Message:"Current password is incorrect"}}, "Bad request", "01", "04")
			return
		}

		if err.Error() == "2" {
			json.NewResponseBadRequest(c, []json.ValidationField{{FieldName:"new_password",Message:"Password do not match"}}, "Bad request", "01", "05")
			return
		}

		json.NewResponseError(c, err.Error(), "01", "06")
		return
	}

	json.NewResponseSuccess(c, nil, "Password updated successfully", "01", "01")
}

func (delivery *userDelivery) Delete(c *gin.Context) {
	userID := c.Param("id")
	err := delivery.userUC.Delete(userID)
	if err != nil {
		if err == sql.ErrNoRows {
			json.NewResponseForbidden(c, "User not found", "01", "01")
			return
		}

		json.NewResponseError(c, err.Error(), "01", "02")
		return
	}

	json.NewResponseSuccess(c, nil, "User deleted successfully", "01", "01")
}

func (delivery *userDelivery) SoftDelete(c *gin.Context) {
	userID := c.Param("id")
	err := delivery.userUC.SoftDelete(userID)
	if err != nil {
		if err == sql.ErrNoRows {
			json.NewResponseForbidden(c, "User not found", "01", "01")
			return
		}

		json.NewResponseError(c, err.Error(), "01", "02")
		return
	}

	json.NewResponseSuccess(c, nil, "User deleted successfully", "01", "01")
}

func (delivery *userDelivery) Restore(c *gin.Context) {
	userID := c.Param("id")
	err := delivery.userUC.Restore(userID)
	if err != nil {
		if err == sql.ErrNoRows {
			json.NewResponseForbidden(c, "User not found", "01", "01")
			return
		}

		json.NewResponseError(c, err.Error(), "01", "02")
		return
	}

	json.NewResponseSuccess(c, nil, "User restored successfully", "01", "01")
}