package usecase

import (
	"avengers-clinic/model/dto"
	"avengers-clinic/src/medicine"
)

type medicineUC struct {
	medicineRepo medicine.MedicineRepository
}

func NewMedicineUsecase(medicineRepo medicine.MedicineRepository) medicine.MedicineUsecase {
	return &medicineUC{medicineRepo}
}

func (m *medicineUC) GetAll() ([]dto.Medicine, error) {
	all, err := m.medicineRepo.RetrieveAll()
	if err != nil {
		return nil, err

	}
	return all, nil
}
