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
}

type MedicineUsecase interface {
	GetAll() ([]dto.MedicineResponse, error)
	GetById(id string) (dto.MedicineResponse, error)
	CreateRecord(medicine dto.MedicineRequest) (dto.MedicineResponse, error)
	UpdateRecord(medicine dto.MedicineRequest) (dto.MedicineResponse, error)
	DeleteRecord(id string) error
}
