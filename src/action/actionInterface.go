package action

import "avengers-clinic/model/dto/actionDto"

type ActionRepository interface {
	GetAll() ([]actionDto.Action, error)
	GetByID(actionID string) (actionDto.Action, error)
	GetTrashByID(actionID string) (actionDto.Action, error)
	Insert(action actionDto.Action) (string, error)
	Update(action actionDto.Action) error
	Delete(actionID string) error
	SoftDelete(actionID string) error
	Restore(actionID string) error
	IsNameExist(name string) bool
}

type ActionUsecase interface {
	GetAll() ([]actionDto.Action, error)
	GetByID(actionID string) (actionDto.Action, error)
	Create(req actionDto.CreateRequest) (actionDto.Action, error)
	Update(req actionDto.UpdateRequest) (actionDto.Action, error)
	Delete(actionID string) error
	SoftDelete(actionID string) error
	Restore(actionID string) error
}