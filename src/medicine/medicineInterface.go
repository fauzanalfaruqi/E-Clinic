package medicine

import "avengers-clinic/model/dto"

type MedicineRepository interface {
	RetrieveAll() ([]dto.MedicineResponse, error)
	RetrieveById(id string) ([]dto.MedicineResponse, error)
	Create(medicine dto.Medicine) (dto.MedicineResponse, error)
}

type MedicineUsecase interface {
	GetAll() ([]dto.MedicineResponse, error)
	GetById(id string) ([]dto.MedicineResponse, error)

	CreateRecord(medicine dto.Medicine) (dto.MedicineResponse, error)
}
