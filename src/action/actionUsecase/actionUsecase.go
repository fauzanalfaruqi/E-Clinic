package actionUsecase

import (
	"avengers-clinic/model/dto/actionDto"
	"avengers-clinic/src/action"
	"database/sql"
	"errors"
	"time"
)

type actionUsecase struct {
	actionRepo action.ActionRepository
}

func NewActionUsecase(actionRepo action.ActionRepository) action.ActionUsecase {
	return &actionUsecase{actionRepo}
}

func (usecase *actionUsecase) GetAll() ([]actionDto.Action, error) {
	actions, err := usecase.actionRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return actions, nil
}

func (usecase *actionUsecase) GetByID(actionID string) (actionDto.Action, error) {
	action, err := usecase.actionRepo.GetByID(actionID)
	if err != nil {
		if err == sql.ErrNoRows {
			return actionDto.Action{}, errors.New("1")
		}
		return actionDto.Action{}, err
	}
	return action, nil
}

func (usecase *actionUsecase) Create(req actionDto.CreateRequest) (actionDto.Action, error) {
	if usecase.actionRepo.IsNameExist(req.Name) {
		return actionDto.Action{}, errors.New("1")
	}

	now := time.Now().Format("2006-01-02 15:04:05")
	action := actionDto.Action{
		Name: req.Name,
		Price: req.Price,
		Description: req.Description,
		CreatedAt: now,
		UpdatedAt: now,
	}

	var err error
	action.ID, err = usecase.actionRepo.Insert(action)
	if err != nil {
		return actionDto.Action{}, err
	}

	return action, err
}

func (usecase *actionUsecase) Update(req actionDto.UpdateRequest) (actionDto.Action, error) {
	action, err := usecase.actionRepo.GetByID(req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return actionDto.Action{}, errors.New("1")
		}
		return actionDto.Action{}, err
	}

	if req.Name != "" {
		if usecase.actionRepo.IsNameExist(req.Name) && req.Name != action.Name {
			return actionDto.Action{}, errors.New("2")
		}
		action.Name = req.Name
	}

	if req.Price != 0 {
		action.Price = req.Price
	}

	if req.Description != "" {
		action.Description = req.Description
	}
	action.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	err = usecase.actionRepo.Update(action)
	if err != nil {
		return actionDto.Action{}, err
	}
	return action, nil
}

func (usecase *actionUsecase) Delete(actionID string) error {
	err := usecase.actionRepo.Delete(actionID)
	if err != nil {
		return err
	}
	return nil
}

func (usecase *actionUsecase) SoftDelete(actionID string) error {
	err := usecase.actionRepo.SoftDelete(actionID)
	if err != nil {
		return err
	}
	return nil
}

func (usecase *actionUsecase) Restore(actionID string) error {
	err := usecase.actionRepo.Restore(actionID)
	if err != nil {
		return err
	}
	return nil
}