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

func (dr *medicalRecordRepository) AddMedicalRecord(cmr *medicalRecordDTO.CreateMedicalRecord) (*medicalRecordDTO.CreateMedicalRecord, error) {
	tx, err := dr.db.Begin()
	if err != nil {
		tx.Rollback()
		return &medicalRecordDTO.CreateMedicalRecord{}, err
	}

	// Inserting medical record values
	var mrID string
	query := "INSERT INTO medical_records (booking_id, diagnosis_results, total_medicine, total_action, total_amount, payment_status) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	if err := tx.QueryRow(query, cmr.Booking_ID, cmr.Diagnosis_Result, 0, 0, 0, false).Scan(&mrID); err != nil {
		tx.Rollback()
		return &medicalRecordDTO.CreateMedicalRecord{}, err
	}

	var totalMedicine int

	// Inserting medicine details values into medical_record_medicine_details table
	for _, md := range cmr.Medicine_Details {
		// TODO: get medicine price
		// Get price value from medicine table
		query = "SELECT price from medicines WHERE id = $1 AND deleted_at IS null"
		if err = tx.QueryRow(query, md.Medicine_ID).Scan(&md.Medicine_Price); err != nil {
			tx.Rollback()
			return &medicalRecordDTO.CreateMedicalRecord{}, err
		}

		// Insert medicine details
		var mdID string
		query := "INSERT INTO medical_record_medicine_details (medical_record_id, medicine_id, medicine_price, quantity) VALUES ($1, $2, $3, $4) RETURNING id"
		if err := tx.QueryRow(query, mrID, md.Medicine_ID, md.Medicine_Price, md.Quantity).Scan(&mdID); err != nil {
			tx.Rollback()
			return &medicalRecordDTO.CreateMedicalRecord{}, err
		}

		md.ID = mdID
		md.Medicine_Price *= md.Quantity
		totalMedicine += md.Medicine_Price
	}

	var totalAction int

	// Inserting medicine details values into medical_record_action_details table
	for _, ad := range cmr.Action_Details {
		// Get price value from medicine table
		query = "SELECT price from actions WHERE id = $1 AND deleted_at IS null"
		if err = tx.QueryRow(query, ad.Action_ID).Scan(&ad.Action_Price); err != nil {
			tx.Rollback()
			return &medicalRecordDTO.CreateMedicalRecord{}, err
		}

		var adID string
		query := "INSERT INTO medical_record_action_details (medical_record_id, action_id, action_price) VALUES ($1, $2, $3) RETURNING id"
		// TODO: Get medical record id to insert here
		if err := tx.QueryRow(query, mrID, ad.Action_ID, ad.Action_Price).Scan(&adID); err != nil {
			tx.Rollback()
			return &medicalRecordDTO.CreateMedicalRecord{}, err
		}

		ad.ID = adID
		totalAction += ad.Action_Price
	}

	totalAmount := totalMedicine + totalAction
	query = "UPDATE medical_records SET total_medicine = $1, total_action = $2, total_amount = $3 WHERE id = $4"
	_, err = tx.Exec(query, totalMedicine, totalAction, totalAmount, mrID)
	if err != nil {
		tx.Rollback()
		return &medicalRecordDTO.CreateMedicalRecord{}, err
	}

	tx.Commit()
	fmt.Println("transaction commited")

	fmt.Println(cmr)
	return cmr, nil
}

func (dr *medicalRecordRepository) RetrieveMedicalRecords() ([]medicalRecordDTO.MedicalRecord, error) {
	var mrs []medicalRecordDTO.MedicalRecord

	// Begin the transaction
	tx, err := dr.db.Begin()
	if err != nil {
		return mrs, err
	}

	// Defer rollback in case of error
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Getting medical_record values
	query := "SELECT id, booking_id, diagnosis_results FROM medical_records WHERE deleted_at IS null"
	row, err := tx.Query(query)
	if err != nil {
		return mrs, err
	}
	defer row.Close()

	// Assign for each received medical_record values into mr variable
	for row.Next() {
		var mr medicalRecordDTO.MedicalRecord
		if err := row.Scan(&mr.ID, &mr.Booking_ID, &mr.Diagnosis_Result); err != nil {
			return mrs, err
		}

		// Getting medical_record_medicine_details values
		var mrmds []medicalRecordDTO.Medical_Record_Medicine_Details
		query = "SELECT id, quantity FROM medical_record_medicine_details WHERE medical_record_id = $1 AND deleted_at IS null"
		mrmdsRow, err := tx.Query(query, mr.ID)
		if err != nil {
			return mrs, err
		}
		defer mrmdsRow.Close()

		for mrmdsRow.Next() {
			var mrmd medicalRecordDTO.Medical_Record_Medicine_Details
			if err := mrmdsRow.Scan(&mrmd.ID, &mrmd.Quantity); err != nil {
				return mrs, err
			}
			mrmds = append(mrmds, mrmd)
		}

		// Assign medical record medicine details into medical record struct at the current iteration
		mr.Medicine_Details = mrmds

		// Getting medical_record_action_details values
		var mrads []medicalRecordDTO.Medical_Record_Action_Details
		query = "SELECT id FROM medical_record_action_details WHERE medical_record_id = $1 AND deleted_at IS null"
		mradsRow, err := tx.Query(query, mr.ID)
		if err != nil {
			return mrs, err
		}
		defer mradsRow.Close()

		for mradsRow.Next() {
			var mrad medicalRecordDTO.Medical_Record_Action_Details
			if err := mradsRow.Scan(&mrad.ID); err != nil {
				return mrs, err
			}
			mrads = append(mrads, mrad)
		}

		// Assign medical record action details into medical record struct at the current iteration
		mr.Action_Details = mrads

		// Assign received medical record values into mr slice
		mrs = append(mrs, mr)
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return mrs, err
	}

	// If received medical_record data in db is not empty, return the data
	if len(mrs) > 0 {
		return mrs, nil
	}

	// return data not found as error if data is empty on the db
	return mrs, errors.New("data not found")
}

func (dr *medicalRecordRepository) RetrieveMedicalRecordByID(id string) (*medicalRecordDTO.MedicalRecord, error) {
	var mr medicalRecordDTO.MedicalRecord
	var err error

	// Begin the transaction
	tx, err := dr.db.Begin()
	if err != nil {
		tx.Rollback()
		return &medicalRecordDTO.MedicalRecord{}, err
	}

	// Getting medical record values
	query := "SELECT id, booking_id, diagnosis_results FROM medical_records WHERE id = S1 AND deleted_at IS null"
	err = tx.QueryRow(query, id).Scan(&mr.ID, &mr.Booking_ID, &mr.Diagnosis_Result)
	if err != nil {
		tx.Rollback()
		return &medicalRecordDTO.MedicalRecord{}, err
	}

	// Getting medical_record_medicine_details values
	var mrmds []medicalRecordDTO.Medical_Record_Medicine_Details

	query = "SELECT id, quantity FROM medical_record_medicine_details WHERE medical_record_id = $1 AND deleted_at IS null"
	mrmdsRow, err := tx.Query(query, mr.ID)
	if err != nil {
		tx.Rollback()
		return &medicalRecordDTO.MedicalRecord{}, err
	}

	for mrmdsRow.Next() {
		var mrmd medicalRecordDTO.Medical_Record_Medicine_Details
		if err := mrmdsRow.Scan(&mrmd.ID, &mrmd.Quantity); err != nil {
			tx.Rollback()
			return &medicalRecordDTO.MedicalRecord{}, err
		}

		mrmds = append(mrmds, mrmd)
	}

	// Assign medical record medicine details into medical record struct at the current iteration
	mr.Medicine_Details = mrmds

	// Getting medical_record_action_details values
	var mrads []medicalRecordDTO.Medical_Record_Action_Details

	query = "SELECT id FROM medical_record_action_details WHERE medical_record_id = $1 AND deleted_at IS null"
	mradsRow, err := tx.Query(query, mr.ID)
	if err != nil {
		tx.Rollback()
		return &medicalRecordDTO.MedicalRecord{}, err
	}

	for mradsRow.Next() {
		var mrad medicalRecordDTO.Medical_Record_Action_Details
		if err := mradsRow.Scan(&mrad.ID); err != nil {
			tx.Rollback()
			return &medicalRecordDTO.MedicalRecord{}, err
		}

		mrads = append(mrads, mrad)
	}

	// Assign medical record medicine details into medical record struct at the current iteration
	mr.Medicine_Details = mrmds

	// Assign medical record medicine details into medical record struct at the current iteration
	mr.Action_Details = mrads

	return &mr, nil
}
