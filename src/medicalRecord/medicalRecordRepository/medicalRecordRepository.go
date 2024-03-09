package medicalRecordRepository

import (
	"avengers-clinic/model/dto/medicalRecordDTO"
	"avengers-clinic/pkg/constants"
	"avengers-clinic/src/medicalRecord"
	"database/sql"
	"errors"
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
	query := "INSERT INTO medical_records (booking_id, diagnosis_results, total_medicine, total_action, total_amount, payment_status) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, created_at, payment_status"
	if err := tx.QueryRow(query, req.Booking_ID, req.Diagnosis_Result, 0, 0, 0, req.Payment_Status).Scan(&medicalRecord.ID, &medicalRecord.Created_At, &medicalRecord.Payment_Status); err != nil {
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

		// Check if the stock is empty
		if medicineDetail.Medicine_Stock <= 0 {
			tx.Rollback()
			return medicalRecordDTO.Medical_Record{}, errors.New(constants.ErrNoStockAvailable)
		}

		// Check if the quantity amount is greater than stock available
		if md.Quantity > medicineDetail.Medicine_Stock {
			tx.Rollback()
			return medicalRecordDTO.Medical_Record{}, errors.New(constants.ErrQuantityGreaterThanStock)
		}

		// Insert medicine details
		query := "INSERT INTO medical_record_medicine_details (medical_record_id, medicine_id, medicine_price, quantity) VALUES ($1, $2, $3, $4) RETURNING id"
		if err := tx.QueryRow(query, medicalRecord.ID, md.Medicine_ID, medicineDetail.Medicine_Price, md.Quantity).Scan(&medicineDetail.ID); err != nil {
			tx.Rollback()
			return medicalRecordDTO.Medical_Record{}, err
		}

		// In case the payment status set to true in the request body
		if req.Payment_Status {
			// Update medicine stock
			medicineDetail.Medicine_Stock, err = dr.UpdateMedicineStock(tx, medicineDetail.Medicine_Stock, md.Quantity, md.Medicine_ID)
			if err != nil {
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

	if err := tx.Commit(); err != nil {
		return medicalRecordDTO.Medical_Record{}, err
	}

	//mr.Medicine_Details = append(mr.Medicine_Details, medicineDetails...)
	return medicalRecord, nil
}

func (dr *medicalRecordRepository) RetrieveMedicalRecords() ([]medicalRecordDTO.Medical_Record, error) {
	var mrs []medicalRecordDTO.Medical_Record
	tx, err := dr.db.Begin()
	if err != nil {
		return []medicalRecordDTO.Medical_Record{}, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Getting medical_record values
	query := "SELECT id, booking_id, diagnosis_results, created_at FROM medical_records WHERE deleted_at IS null"
	row, err := tx.Query(query)
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
		if mrs[i].Medicine_Details, err = dr.GetMedicineDetails(tx, mrs[i].ID); err != nil {
			return []medicalRecordDTO.Medical_Record{}, err
		}

		// Get and assign medical record action details into medical record struct
		if mrs[i].Action_Details, err = dr.GetActionDetails(tx, mrs[i].ID); err != nil {
			return []medicalRecordDTO.Medical_Record{}, err
		}
	}

	if err = tx.Commit(); err != nil {
		return []medicalRecordDTO.Medical_Record{}, err
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
	tx, err := dr.db.Begin()
	if err != nil {
		return medicalRecordDTO.Medical_Record{}, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Getting medical record values
	query := "SELECT id, booking_id, diagnosis_results, created_at FROM medical_records WHERE id = $1 AND deleted_at IS null"
	err = tx.QueryRow(query, id).Scan(&mr.ID, &mr.Booking_ID, &mr.Diagnosis_Result, &mr.Created_At)
	if err != nil {
		return medicalRecordDTO.Medical_Record{}, err
	}

	// Get and assign medical record medicine details into medical record struct
	if mr.Medicine_Details, err = dr.GetMedicineDetails(tx, mr.ID); err != nil {
		return medicalRecordDTO.Medical_Record{}, err
	}

	// Get and assign medical record action details into medical record struct
	if mr.Action_Details, err = dr.GetActionDetails(tx, mr.ID); err != nil {
		return medicalRecordDTO.Medical_Record{}, err
	}

	// Assign medical record medicine details into medical record struct at the current iteration
	//mr.Action_Details = mrads

	if err = tx.Commit(); err != nil {
		return medicalRecordDTO.Medical_Record{}, err
	}
	return mr, nil
}

func (dr *medicalRecordRepository) GetMedicineDetails(db *sql.Tx, mrID string) ([]medicalRecordDTO.Medical_Record_Medicine_Details, error) {
	var medicineDetails []medicalRecordDTO.Medical_Record_Medicine_Details
	var query string

	query = "SELECT id, medicine_id, quantity, created_at FROM medical_record_medicine_details WHERE medical_record_id = $1"
	rows, err := db.Query(query, mrID)
	if err != nil {
		return []medicalRecordDTO.Medical_Record_Medicine_Details{}, err
	}
	defer rows.Close()

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

func (dr *medicalRecordRepository) GetActionDetails(tx *sql.Tx, mrID string) ([]medicalRecordDTO.Medical_Record_Action_Details, error) {
	var actionDetails []medicalRecordDTO.Medical_Record_Action_Details
	var query string

	query = "SELECT id, action_id, created_at FROM medical_record_action_details WHERE medical_record_id = $1"
	rows, err := tx.Query(query, mrID)
	if err != nil {
		return []medicalRecordDTO.Medical_Record_Action_Details{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var ad medicalRecordDTO.Medical_Record_Action_Details
		if err := rows.Scan(&ad.ID, &ad.Action_ID, &ad.Created_At); err != nil {
			return []medicalRecordDTO.Medical_Record_Action_Details{}, err
		}

		actionDetails = append(actionDetails, ad)
	}

	for i := range actionDetails {
		query = "SELECT name, price, description from actions WHERE id = $1"
		err = tx.QueryRow(query, actionDetails[i].Action_ID).Scan(&actionDetails[i].Action_Name, &actionDetails[i].Action_Price, &actionDetails[i].Action_Description)
		if err != nil {
			return []medicalRecordDTO.Medical_Record_Action_Details{}, err
		}
	}

	return actionDetails, nil
}

func (dr *medicalRecordRepository) UpdatePaymentToDone(id string) (medicalRecordDTO.Medical_Record, error) {
	var medicalRecord medicalRecordDTO.Medical_Record
	medicalRecord.ID = id
	tx, err := dr.db.Begin()
	if err != nil {
		return medicalRecordDTO.Medical_Record{}, err
	}

	// Populate medical record struct fields
	query := "SELECT id, booking_id, diagnosis_results, payment_status, created_at FROM medical_records WHERE id = $1 AND deleted_at IS null"
	err = tx.QueryRow(query, id).Scan(&medicalRecord.ID, &medicalRecord.Booking_ID, &medicalRecord.Diagnosis_Result, &medicalRecord.Payment_Status, &medicalRecord.Created_At)
	if err != nil {
		tx.Rollback()
		return medicalRecordDTO.Medical_Record{}, err
	}

	if medicalRecord.Payment_Status {
		tx.Rollback()
		return medicalRecordDTO.Medical_Record{}, errors.New(constants.ErrPaymentAlreadyTrue)
	}
	// Update payment status
	query = "UPDATE medical_records SET payment_status = true WHERE id = $1"
	_, err = tx.Exec(query, id)
	if err != nil {
		tx.Rollback()
		return medicalRecordDTO.Medical_Record{}, err
	}

	medicalRecord.Payment_Status = true

	//var mds []medicalRecordDTO.Medical_Record_Medicine_Details
	mds, err := dr.GetMedicineDetails(tx, id)
	if err != nil {
		tx.Rollback()
		return medicalRecordDTO.Medical_Record{}, err
	}

	for i := range mds {
		stockResult, err := dr.UpdateMedicineStock(tx, mds[i].Medicine_Stock, mds[i].Quantity, mds[i].Medicine_ID)
		if err != nil {
			tx.Rollback()
			return medicalRecordDTO.Medical_Record{}, err
		}

		mds[i].Medicine_Stock = stockResult
	}

	// Append medicine details into medical record
	medicalRecord.Medicine_Details = mds

	// Populate action details
	var ads []medicalRecordDTO.Medical_Record_Action_Details
	ads, err = dr.GetActionDetails(tx, id)
	if err != nil {
		tx.Rollback()
		return medicalRecordDTO.Medical_Record{}, err
	}

	if err = tx.Commit(); err != nil {
		return medicalRecordDTO.Medical_Record{}, err
	}

	// Append action details into medical record
	medicalRecord.Action_Details = ads

	return medicalRecord, nil

}

func (dr *medicalRecordRepository) UpdateMedicineStock(tx *sql.Tx, stock, quantity int, medicineID string) (int, error) {

	// Check if the stock is empty
	if stock <= 0 {
		tx.Rollback()
		return 0, errors.New(constants.ErrNoStockAvailable)
	}

	// Check if the quantity amount is greater than stock available
	if quantity > stock {
		tx.Rollback()
		return 0, errors.New(constants.ErrQuantityGreaterThanStock)
	}

	stock -= quantity

	// Update stock
	query := "UPDATE medicines SET stock = $1 WHERE id = $2 and deleted_at IS null"
	_, err := tx.Exec(query, stock, medicineID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return stock, nil
}
