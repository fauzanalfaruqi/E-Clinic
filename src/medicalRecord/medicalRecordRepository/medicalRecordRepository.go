package medicalRecordRepository

import (
	"avengers-clinic/model/dto/medicalRecordDTO"
	"avengers-clinic/src/medicalRecord"
	"database/sql"
	"errors"
	"fmt"
)

type medicalRecordRepository struct {
	db *sql.DB
}

func NewMedicalRecordRepository(db *sql.DB) medicalRecord.MedicalRecordRepository {
	return &medicalRecordRepository{db}
}

func (dr *medicalRecordRepository) AddMedicalRecord(req medicalRecordDTO.Medical_Record_Request) (medicalRecordDTO.Medical_Record, error) {
	var medicalRecord medicalRecordDTO.Medical_Record
	medicalRecord.Diagnosis_Result = req.Diagnosis_Result

	tx, err := dr.db.Begin()
	if err != nil {
		tx.Rollback()
		return medicalRecordDTO.Medical_Record{}, err
	}

	// Inserting medical record values
	query := "INSERT INTO medical_records (booking_id, diagnosis_results, total_medicine, total_action, total_amount, payment_status) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, created_at"
	if err := tx.QueryRow(query, req.Booking_ID, req.Diagnosis_Result, 0, 0, 0, req.Payment_Status).Scan(&medicalRecord.ID, &medicalRecord.Created_At); err != nil {
		tx.Rollback()
		return medicalRecordDTO.Medical_Record{}, err
	}

	var totalMedicine int

	//var medicineDetails []medicalRecordDTO.Medical_Record_Medicine_Details

	// Inserting medicine details values into medical_record_medicine_details table
	for _, md := range req.Medicine_Details {
		var medicineDetail medicalRecordDTO.Medical_Record_Medicine_Details
		medicineDetail.Medicine_ID = md.Medicine_ID
		medicineDetail.Quantity = md.Quantity

		// TODO: get medicine price
		// Read and assign values from medicines tables into medicine_details struct
		query = "SELECT name, stock, price from medicines WHERE id = $1"
		err = tx.QueryRow(query, md.Medicine_ID).Scan(&medicineDetail.Medicine_Name, &medicineDetail.Medicine_Stock, &medicineDetail.Medicine_Price)
		if err != nil {
			tx.Rollback()
			return medicalRecordDTO.Medical_Record{}, err
		}

		// Insert medicine details
		query := "INSERT INTO medical_record_medicine_details (medical_record_id, medicine_id, medicine_price, quantity) VALUES ($1, $2, $3, $4) RETURNING id"
		if err := tx.QueryRow(query, medicalRecord.ID, md.Medicine_ID, medicineDetail.Medicine_Price, md.Quantity).Scan(&medicineDetail.ID); err != nil {
			tx.Rollback()
			return medicalRecordDTO.Medical_Record{}, err
		}

		// In case the payment status set to true in the request body
		if req.Payment_Status {
			fmt.Printf("medicineDetail.Stock %d - md.Quantity %d", medicineDetail.Medicine_Stock, md.Quantity)
			medicineDetail.Medicine_Stock -= md.Quantity

			// Update stock
			query = "UPDATE medicines SET stock = $1 WHERE id = $2 RETURNING stock"
			err = tx.QueryRow(query, medicineDetail.Medicine_Stock, md.Medicine_ID).Scan(&medicineDetail.Medicine_Stock)
			if err != nil {
				tx.Rollback()
				return medicalRecordDTO.Medical_Record{}, err
			}
		}

		medicineDetail.Medicine_Price *= md.Quantity
		totalMedicine += medicineDetail.Medicine_Price
		medicalRecord.Medicine_Details = append(medicalRecord.Medicine_Details, medicineDetail)
	}

	var totalAction int

	// Inserting medicine details values into medical_record_action_details table
	for _, ad := range req.Action_Details {
		var actionDetail medicalRecordDTO.Medical_Record_Action_Details
		actionDetail.Action_ID = ad.Action_ID

		// Get price value from medicine table
		query = "SELECT name, price, description from actions WHERE id = $1 AND deleted_at IS null"
		if err = tx.QueryRow(query, ad.Action_ID).Scan(&actionDetail.Action_Name, &actionDetail.Action_Price, &actionDetail.Action_Description); err != nil {
			tx.Rollback()
			return medicalRecordDTO.Medical_Record{}, err
		}

		query := "INSERT INTO medical_record_action_details (medical_record_id, action_id, action_price) VALUES ($1, $2, $3) RETURNING id"
		// TODO: Get medical record id to insert here
		if err := tx.QueryRow(query, medicalRecord.ID, ad.Action_ID, actionDetail.Action_Price).Scan(&actionDetail.ID); err != nil {
			tx.Rollback()
			return medicalRecordDTO.Medical_Record{}, err
		}

		totalAction += actionDetail.Action_Price
		medicalRecord.Action_Details = append(medicalRecord.Action_Details, actionDetail)
	}

	totalAmount := totalMedicine + totalAction
	query = "UPDATE medical_records SET total_medicine = $1, total_action = $2, total_amount = $3 WHERE id = $4"
	_, err = tx.Exec(query, totalMedicine, totalAction, totalAmount, medicalRecord.ID)
	if err != nil {
		tx.Rollback()
		return medicalRecordDTO.Medical_Record{}, err
	}

	tx.Commit()

	//mr.Medicine_Details = append(mr.Medicine_Details, medicineDetails...)
	return medicalRecord, nil
}

func (dr *medicalRecordRepository) RetrieveMedicalRecords() ([]medicalRecordDTO.Medical_Record, error) {
	var mrs []medicalRecordDTO.Medical_Record

	// Getting medical_record values
	query := "SELECT id, booking_id, diagnosis_results, created_at FROM medical_records WHERE deleted_at IS null"
	row, err := dr.db.Query(query)
	if err != nil {
		return []medicalRecordDTO.Medical_Record{}, err
	}
	defer row.Close()

	// Assign for each received medical_record values into mr variable
	for row.Next() {
		var mr medicalRecordDTO.Medical_Record
		if err := row.Scan(&mr.ID, &mr.Booking_ID, &mr.Diagnosis_Result, &mr.Created_At); err != nil {
			return []medicalRecordDTO.Medical_Record{}, err
		}

		// Assign received medical record values into mr slice
		mrs = append(mrs, mr)
	}

	// Here we try to assign medicine_details and action_details for each medical record in the mr array
	for i := range mrs {
		// Get and assign medical record medicine details into medical record struct
		if mrs[i].Medicine_Details, err = dr.GetMedicineDetails(dr.db, mrs[i].ID); err != nil {
			return []medicalRecordDTO.Medical_Record{}, err
		}

		// Get and assign medical record action details into medical record struct
		if mrs[i].Action_Details, err = dr.GetActionDetails(dr.db, mrs[i].ID); err != nil {
			return []medicalRecordDTO.Medical_Record{}, err
		}
	}

	// If received medical_record data in db is not empty, return the data
	if len(mrs) > 0 {
		return mrs, nil
	}

	// return data not found as error if data is empty on the db
	return []medicalRecordDTO.Medical_Record{}, errors.New("data not found")
}

func (dr *medicalRecordRepository) RetrieveMedicalRecordByID(id string) (medicalRecordDTO.Medical_Record, error) {
	var mr medicalRecordDTO.Medical_Record

	// Getting medical record values
	query := "SELECT id, booking_id, diagnosis_results, created_at FROM medical_records WHERE id = $1 AND deleted_at IS null"
	err := dr.db.QueryRow(query, id).Scan(&mr.ID, &mr.Booking_ID, &mr.Diagnosis_Result, &mr.Created_At)
	if err != nil {
		return medicalRecordDTO.Medical_Record{}, err
	}

	// Get and assign medical record medicine details into medical record struct
	if mr.Medicine_Details, err = dr.GetMedicineDetails(dr.db, mr.ID); err != nil {
		return medicalRecordDTO.Medical_Record{}, err
	}

	// Get and assign medical record action details into medical record struct
	if mr.Action_Details, err = dr.GetActionDetails(dr.db, mr.ID); err != nil {
		return medicalRecordDTO.Medical_Record{}, err
	}

	// Assign medical record medicine details into medical record struct at the current iteration
	//mr.Action_Details = mrads

	return mr, nil
}

func (dr *medicalRecordRepository) GetMedicineDetails(db *sql.DB, mrID string) ([]medicalRecordDTO.Medical_Record_Medicine_Details, error) {
	var medicineDetails []medicalRecordDTO.Medical_Record_Medicine_Details
	var query string

	query = "SELECT id, medicine_id, quantity, created_at FROM medical_record_medicine_details WHERE medical_record_id = $1"
	rows, err := db.Query(query, mrID)
	if err != nil {
		return []medicalRecordDTO.Medical_Record_Medicine_Details{}, err
	}

	for rows.Next() {
		var md medicalRecordDTO.Medical_Record_Medicine_Details
		if err := rows.Scan(&md.ID, &md.Medicine_ID, &md.Quantity, &md.Created_At); err != nil {
			return []medicalRecordDTO.Medical_Record_Medicine_Details{}, err
		}

		medicineDetails = append(medicineDetails, md)
	}

	for i := range medicineDetails {
		query = "SELECT name, stock, price from medicines WHERE id = $1"
		err = db.QueryRow(query, medicineDetails[i].Medicine_ID).Scan(&medicineDetails[i].Medicine_Name, &medicineDetails[i].Medicine_Stock, &medicineDetails[i].Medicine_Price)
		if err != nil {
			return []medicalRecordDTO.Medical_Record_Medicine_Details{}, err
		}
	}

	return medicineDetails, nil
}

func (dr *medicalRecordRepository) GetActionDetails(db *sql.DB, mrID string) ([]medicalRecordDTO.Medical_Record_Action_Details, error) {
	var actionDetails []medicalRecordDTO.Medical_Record_Action_Details
	var query string

	query = "SELECT id, action_id, created_at FROM medical_record_action_details WHERE medical_record_id = $1"
	rows, err := db.Query(query, mrID)
	if err != nil {
		return []medicalRecordDTO.Medical_Record_Action_Details{}, err
	}

	for rows.Next() {
		var ad medicalRecordDTO.Medical_Record_Action_Details
		if err := rows.Scan(&ad.ID, &ad.Action_ID, &ad.Created_At); err != nil {
			return []medicalRecordDTO.Medical_Record_Action_Details{}, err
		}

		actionDetails = append(actionDetails, ad)
	}

	for i := range actionDetails {
		query = "SELECT name, price, description from actions WHERE id = $1"
		err = db.QueryRow(query, actionDetails[i].Action_ID).Scan(&actionDetails[i].Action_Name, &actionDetails[i].Action_Price, &actionDetails[i].Action_Description)
		if err != nil {
			return []medicalRecordDTO.Medical_Record_Action_Details{}, err
		}
	}

	return actionDetails, nil
}
