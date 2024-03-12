package medicineUsecase

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

func (m *medicineUC) GetById(id string) (dto.MedicineResponse, error) {
	all, err := m.medicineRepo.RetrieveById(id)
	if err != nil {
		return all, err
	}
	return all, nil
}

func (m *medicineUC) CreateRecord(medicine dto.MedicineRequest) (dto.MedicineResponse, error) {
	var new dto.MedicineResponse
	var err error
	newCreatedAt := time.Now().Format("2006-01-02 15:04:05")
	newUpdatedAt := time.Now().Format("2006-01-02 15:04:05")
	newMedicine := dto.MedicineRequest{Name: medicine.Name, MedicineType: medicine.MedicineType, Price: medicine.Price, Stock: medicine.Stock, Description: medicine.Description, CreatedAt: newCreatedAt, UpdatedAt: newUpdatedAt}
	new, err = m.medicineRepo.Create(newMedicine)
	if err != nil {
		return new, err
	}

	return new, err
}

func (m *medicineUC) UpdateRecord(Updated dto.UpdateRequest) (dto.MedicineResponse, error) {
	action, err := m.medicineRepo.RetrieveById(Updated.Id)
	if err != nil {
		return dto.MedicineResponse{}, err
	}

	if Updated.Name == "" {
		Updated.Name = action.Name
	}

	if Updated.MedicineType == "" {
		Updated.MedicineType = action.MedicineType
	}

	if Updated.Price == 0 {
		Updated.Price = action.Price
	}

	if Updated.Stock == 0 {
		Updated.Stock = action.Stock
	}

	if Updated.Description == nil {
		Updated.Description = action.Description
	}
	var all dto.MedicineResponse
	Updated.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	product := dto.MedicineRequest{Id: Updated.Id, Name: Updated.Name, MedicineType: Updated.MedicineType, Price: Updated.Price, Stock: Updated.Stock, Description: Updated.Description, UpdatedAt: Updated.UpdatedAt}
	fmt.Println(product)
	all, err = m.medicineRepo.Update(product)
	if err != nil {
		return all, err
	}
	return all, err
}

func (m *medicineUC) DeleteRecord(id string) error {
	_, err :=  m.medicineRepo.RetrieveById(id)
	if err != nil {
		return err
	}
	deletedAt := time.Now().Format("2006-01-02 15:04:05")
	err = m.medicineRepo.Delete(id, deletedAt)
	if err != nil {
		return err
	}
	return err
}

func (m *medicineUC) TrashRecord() ([]dto.MedicineResponse, error) {
	all, err := m.medicineRepo.Trash()
	if err != nil {
		return nil, err

	}
	return all, nil
}

func (m *medicineUC) RestoreRecord(id string) error {
	err := m.medicineRepo.Restore(id)
	if err != nil {
		return err
	}
	return err
}
