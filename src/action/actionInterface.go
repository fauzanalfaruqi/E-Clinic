package action

import "avengers-clinic/model/dto/actionDto"

type ActionRepository interface {
	GetAll() ([]actionDto.Action, error)
	GetByID(actionID string) (actionDto.Action, error)
	Insert(action actionDto.Action) (string, error)
	Update(action actionDto.Action) error
	Delete(actionID string) error
	SoftDelete(actionID string) error
	Restore(actionID string) error
	IsNameExist(name string) bool
}