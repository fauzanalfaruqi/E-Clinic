package usecase

import (
	"avengers-clinic/model/dto"
	"avengers-clinic/src/medicine"
	"fmt"
	"time"
)

type medicineUC struct {
	medicineRepo medicine.MedicineRepository
}

func NewMedicineUsecase(medicineRepo medicine.MedicineRepository) medicine.MedicineUsecase {
	return &medicineUC{medicineRepo}
}

func (m *medicineUC) GetAll() ([]dto.MedicineResponse, error) {
	all, err := m.medicineRepo.RetrieveAll()
	if err != nil {
		return nil, err

	}
	return all, nil
}

func (m *medicineUC) GetById(id string) ([]dto.MedicineResponse, error) {
	all, err := m.medicineRepo.RetrieveById(id)
	if err != nil {
		return nil, err
	}
	return all, nil
}

func (m *medicineUC) CreateRecord(medicine dto.Medicine) (dto.MedicineResponse, error) {
	var new dto.MedicineResponse
	var err error
	newCreatedAt := time.Now()
	fmt.Println(newCreatedAt)
	newMedicine := dto.Medicine{Name: medicine.Name, MedicineType: medicine.MedicineType, Price: medicine.Price, Stock: medicine.Stock, Description: medicine.Description, CreatedAt: newCreatedAt}
	new, err = m.medicineRepo.Create(newMedicine)
	if err != nil {
		return new, err
	}

	return new, err
}
