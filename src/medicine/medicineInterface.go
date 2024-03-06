package medicine

import "avengers-clinic/model/dto"

type MedicineRepository interface {
	RetrieveAll() ([]dto.Medicine, error)
}

type MedicineUsecase interface {
	GetAll() ([]dto.Medicine, error)
}
