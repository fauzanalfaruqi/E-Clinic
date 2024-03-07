package repository

import (
	"avengers-clinic/model/dto"
	"avengers-clinic/src/medicine"
	"database/sql"
)

type medicineRepository struct {
	db *sql.DB
}

func NewMedicineRepository(db *sql.DB) medicine.MedicineRepository {
	return &medicineRepository{db}
}

func (m *medicineRepository) Create(medicine dto.MedicineRequest) (dto.MedicineResponse, error) {
	queryStatement := "INSERT INTO medicines (name,medicine_type,price,stock,description,created_at)VALUES($1,$2,$3,$4,$5,$6) returning id"

	var returning dto.MedicineRequest
	err := m.db.QueryRow(queryStatement, medicine.Name, medicine.MedicineType, medicine.Price, medicine.Stock, medicine.Description, medicine.CreatedAt).Scan(&returning.Id)

	newMedicine := dto.MedicineResponse{Id: returning.Id, Name: medicine.Name, MedicineType: medicine.MedicineType, Price: medicine.Price, Stock: medicine.Stock, Description: medicine.Description, CreatedAt: medicine.CreatedAt}
	return newMedicine, err
}

func (m *medicineRepository) Update(medicine dto.MedicineRequest) (dto.MedicineResponse, error) {
	sqlstament := "UPDATE medicines SET name=$2,price=$3,stock=$4,description=$5,updated_at=$6,medicine_type=$7 where id=$1 and deleted_at is null returning id,name,medicine_type,price,stock,description,created_at,updated_at;"
	var out dto.MedicineResponse
	err := m.db.QueryRow(sqlstament, medicine.Id, medicine.Name, medicine.Price, medicine.Stock, medicine.Description, medicine.UpdatedAt, medicine.MedicineType).Scan(&out.Id, &out.Name, &out.MedicineType, &out.Price, &out.Stock, &out.Description, &out.CreatedAt, &out.UpdatedAt)
	if err != nil {
		return out, err
	}
	return out, err
}

func (m *medicineRepository) Delete(id string, deletedAt string) error {
	sqlstament := "UPDATE medicines SET deleted_at=$1 where id=$2;"
	_, err := m.db.Exec(sqlstament, deletedAt, id)
	return err
}

func (m *medicineRepository) RetrieveAll() ([]dto.MedicineResponse, error) {
	sqlstatement := "SELECT id, name, medicine_type, price, stock, description,created_at,COALESCE(TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS'), '') AS formatted_updated_at FROM medicines WHERE deleted_at IS NULL;"
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

func (m *medicineRepository) RetrieveById(id string) (dto.MedicineResponse, error) {
	queryStatement := "SELECT id,name,medicine_type,price,stock,description from medicines where id=$1 and deleted_at is null"
	var medicine dto.MedicineResponse
	rows := m.db.QueryRow(queryStatement, id)
	err := rows.Scan(&medicine.Id, &medicine.Name, &medicine.MedicineType, &medicine.Price, &medicine.Stock, &medicine.Description)

	return medicine, err
}

func scan(rows *sql.Rows) ([]dto.MedicineResponse, error) {
	Exp := []dto.MedicineResponse{}
	var err error
	for rows.Next() {
		medicine := dto.MedicineResponse{}
		err := rows.Scan(&medicine.Id, &medicine.Name, &medicine.MedicineType, &medicine.Price, &medicine.Stock, &medicine.Description, &medicine.CreatedAt, &medicine.UpdatedAt)
		if err != nil {
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
