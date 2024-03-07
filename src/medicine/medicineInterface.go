package medicine

import "avengers-clinic/model/dto"

type MedicineRepository interface {
	RetrieveAll() ([]dto.MedicineResponse, error)
	RetrieveById(id string) ([]dto.MedicineResponse, error)
	Create(medicine dto.Medicine) (dto.MedicineResponse, error)
	Update(medicine dto.Medicine) (dto.MedicineResponse, error)
	Delete(id string) (string, error)
}

type MedicineUsecase interface {
	GetAll() ([]dto.MedicineResponse, error)
	GetById(id string) ([]dto.MedicineResponse, error)
	CreateRecord(medicine dto.Medicine) (dto.MedicineResponse, error)
	UpdateRecord(medicine dto.Medicine) (dto.MedicineResponse, error)
	DeleteRecord(id string) (string, error)
}
