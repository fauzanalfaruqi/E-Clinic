package actionDelivery

import (
	"avengers-clinic/model/dto/actionDto"
	"avengers-clinic/model/dto/json"
	"avengers-clinic/pkg/utils"
	"avengers-clinic/src/action"

	"github.com/gin-gonic/gin"
)

type actionDelivery struct {
	actionUC action.ActionUsecase
}

func NewActionDelivery(v1Group *gin.RouterGroup, actionUC action.ActionUsecase) {
	handler := actionDelivery{actionUC: actionUC}
	actionGroup := v1Group.Group("/actions")
	{
		actionGroup.GET("", handler.GetAll)
		actionGroup.GET("/:id", handler.GetByID)
		actionGroup.POST("/", handler.Create)
		actionGroup.PUT("/:id", handler.Update)
		actionGroup.DELETE("/:id", handler.Delete)
		actionGroup.DELETE("/:id/trash", handler.SoftDelete)
		actionGroup.PUT("/:id/restore", handler.Restore)
	}
}

func (delivery *actionDelivery) GetAll(c *gin.Context) {
	response, err := delivery.actionUC.GetAll()
	if err != nil {
		json.NewResponseError(c, err.Error(), "01", "01")
		return
	}

	if len(response) == 0 {
		json.NewResponseSuccess(c, nil, "Actions not found", "01", "01")
		return
	}

	json.NewResponseSuccess(c, response, "actions successfully retrieved", "01", "01")
}

func (delivery *actionDelivery) GetByID(c *gin.Context) {
	actionID := c.Param("id")
	response, err := delivery.actionUC.GetByID(actionID)
	if err != nil {
		if err.Error() == "1" {
			json.NewResponseSuccess(c, nil, "Action not found", "01", "01")
			return
		}

		json.NewResponseError(c, err.Error(), "01", "02")
		return
	}

	json.NewResponseSuccess(c, response, "actions successfully retrieved", "01", "01")
}

func (delivery *actionDelivery) Create(c *gin.Context) {
	var request actionDto.CreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		json.NewResponseError(c, err.Error(), "01", "01")
		return
	}

	if err := utils.Validated(request); err != nil {
		json.NewResponseBadRequest(c, err, "Bad request", "01", "02")
		return
	}

	response, err := delivery.actionUC.Create(request)
	if err != nil {
		if err.Error() == "1" {
			json.NewResponseBadRequest(c, []json.ValidationField{{FieldName:"name", Message:"Name is already registered"}}, "Bad request", "01", "03")
			return
		}

		json.NewResponseError(c, err.Error(), "01", "04")
		return
	}

	json.NewResponseCreated(c, response, "Action created successfully", "01", "01")
}

func (delivery *actionDelivery) Update(c *gin.Context) {
	var request actionDto.UpdateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		json.NewResponseError(c, err.Error(), "01", "01")
		return
	}
	
	if err := utils.Validated(request); err != nil {
		json.NewResponseBadRequest(c, err, "Bad request", "01", "02")
		return
	}
	request.ID = c.Param("id")

	response, err := delivery.actionUC.Update(request)
	if err != nil {
		if err.Error() == "1" {
			json.NewResponseSuccess(c, nil, "Action not found", "01", "03")
			return
		}

		if err.Error() == "2" {
			json.NewResponseBadRequest(c, []json.ValidationField{{FieldName:"name", Message:"Name is already registered"}}, "Bad request", "01", "04")
			return
		}

		json.NewResponseError(c, err.Error(), "01", "05")
		return
	}

	json.NewResponseSuccess(c, response, "Action updeted successfully", "01", "01")
}

func (delivery *actionDelivery) Delete(c *gin.Context) {
	actionID := c.Param("id")
	err := delivery.actionUC.Delete(actionID)
	if err != nil {
		json.NewResponseError(c, err.Error(), "01", "01")
		return
	}
	json.NewResponseSuccess(c, nil, "Action deleted successfully", "01", "01")
}

func (delivery *actionDelivery) SoftDelete(c *gin.Context) {
	actionID := c.Param("id")
	err := delivery.actionUC.SoftDelete(actionID)
	if err != nil {
		json.NewResponseError(c, err.Error(), "01", "01")
		return
	}
	json.NewResponseSuccess(c, nil, "Action deleted successfully", "01", "01")
}

func (delivery *actionDelivery) Restore(c *gin.Context) {
	actionID := c.Param("id")
	err := delivery.actionUC.Restore(actionID)
	if err != nil {
		json.NewResponseError(c, err.Error(), "01", "01")
		return
	}
	json.NewResponseSuccess(c, nil, "Action restored successfully", "01", "01")
}