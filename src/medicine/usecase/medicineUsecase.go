package usecase

import (
	"avengers-clinic/model/dto"
	"avengers-clinic/src/medicine"
	"database/sql"
	"errors"
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
		return all, err
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

func (m *medicineUC) UpdateRecord(Updated dto.Medicine) (dto.MedicineResponse, error) {
	action, err := m.medicineRepo.RetrieveById(Updated.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return dto.MedicineResponse{}, errors.New("1")
		}
		return dto.MedicineResponse{}, err
	}
	fmt.Println(action)
	if Updated.Name == "" {

	}
	var all dto.MedicineResponse
	Updated.UpdatedAt = time.Now()
	product := dto.Medicine{Id: Updated.Id, Name: Updated.Name, Price: Updated.Price, Stock: Updated.Stock, Description: Updated.Description, UpdatedAt: Updated.UpdatedAt}
	fmt.Println(product)
	all, err = m.medicineRepo.Update(product)
	if err != nil {
		return all, err
	}
	return all, err
}

func (m *medicineUC) DeleteRecord(id string) (string, error) {
	var err error

	fmt.Println(id)
	_, err = m.medicineRepo.Delete(id)
	fmt.Println(id, "delete record Usecacase")
	if err != nil {
		return id, err
	}
	return id, err
}
