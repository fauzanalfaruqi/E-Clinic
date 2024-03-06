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

func (m *medicineRepository) RetrieveAll() ([]dto.Medicine, error) {
	sqlstatement := "SELECT * from medicines where deleted_at=null"
	rows, err := m.db.Query(sqlstatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	expenses, err := scan(rows)
	if err != nil {
		return nil, err

	}

	return expenses, nil
}

func scan(rows *sql.Rows) ([]dto.Medicine, error) {
	Exp := []dto.Medicine{}
	var err error
	for rows.Next() {
		medicine := dto.Medicine{}
		fmt.Println(medicine.Id)
		err := rows.Scan(&medicine.Id, &medicine.Name, &medicine.MedicineType, &medicine.Price, &medicine.Stock)
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
