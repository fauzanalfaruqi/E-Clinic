package medicine

import (
	"avengers-clinic/model/dto"
)

type MedicineRepository interface {
	RetrieveAll() ([]dto.MedicineResponse, error)
	RetrieveById(id string) (dto.MedicineResponse, error)
	Create(medicine dto.MedicineRequest) (dto.MedicineResponse, error)
	Update(medicine dto.MedicineRequest) (dto.MedicineResponse, error)
	Delete(id string, deletedAt string) error
	Trash() ([]dto.MedicineResponse, error)
	Restore(id string) error
}

type MedicineUsecase interface {
	GetAll() ([]dto.MedicineResponse, error)
	GetById(id string) (dto.MedicineResponse, error)
	CreateRecord(medicine dto.MedicineRequest) (dto.MedicineResponse, error)
	UpdateRecord(medicine dto.UpdateRequest) (dto.MedicineResponse, error)
	DeleteRecord(id string) error
	TrashRecord() ([]dto.MedicineResponse, error)
	RestoreRecord(id string) error
}
