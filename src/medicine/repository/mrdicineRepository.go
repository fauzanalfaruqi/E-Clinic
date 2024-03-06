package repository

import (
	"avengers-clinic/model/dto"
	"avengers-clinic/src/medicine"
	"database/sql"
	"fmt"
)

type medicineRepository struct {
	db *sql.DB
}

func NewMedicineRepository(db *sql.DB) medicine.MedicineRepository {
	return &medicineRepository{db}
}

func (m *medicineRepository) Create(medicine dto.Medicine) (dto.MedicineResponse, error) {
	queryStatement := "INSERT INTO medicines (name,medicine_type,price,stock,description,created_at)VALUES($1,$2,$3,$4,$5,$6) returning id"

	var returning dto.Medicine
	err := m.db.QueryRow(queryStatement, medicine.Name, medicine.MedicineType, medicine.Price, medicine.Stock, medicine.Description, medicine.CreatedAt).Scan(&returning.Id)

	newMedicine := dto.MedicineResponse{Id: returning.Id, Name: medicine.Name, MedicineType: medicine.MedicineType, Price: medicine.Price, Stock: medicine.Stock, Description: medicine.Description}
	return newMedicine, err
}

func (m *medicineRepository) RetrieveAll() ([]dto.MedicineResponse, error) {
	sqlstatement := "SELECT id,name,medicine_type,price,stock,description from medicines where deleted_at is null"
	rows, err := m.db.Query(sqlstatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	expenses, err := scan(rows)
	if err != nil {
		return nil, err

	}

	return expenses, err
}

func (m *medicineRepository) RetrieveById(id string) ([]dto.MedicineResponse, error) {
	queryStatement := "SELECT id,name,medicine_type,price,stock,description from medicines where id=$1 and deleted_at is null"
	rows, err := m.db.Query(queryStatement, id)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer rows.Close()
	medicine, err := scan(rows)
	return medicine, err
}

func scan(rows *sql.Rows) ([]dto.MedicineResponse, error) {
	Exp := []dto.MedicineResponse{}
	var err error
	for rows.Next() {
		medicine := dto.MedicineResponse{}
		fmt.Println(medicine.Id)
		err := rows.Scan(&medicine.Id, &medicine.Name, &medicine.MedicineType, &medicine.Price, &medicine.Stock, &medicine.Description)
		if err != nil {
			fmt.Println("error scanning")
			return nil, err

		}
		Exp = append(Exp, medicine)
	}
	err = rows.Err()
	if err != nil {
		panic(err)

	}
	return Exp, err
}
