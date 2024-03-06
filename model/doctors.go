package model

import "time"

type (
	MedicalRecord struct {
		ID               string
		Booking_ID       string
		Diagnosis_Result string
		Total_Medicine   int
		Total_Action     int
		Total_Amount     int
		Payment_Status   bool
		Medicine_Details []Medical_Record_Medicine_Details
		Action_Details   []Medical_Record_Action_Details
		Created_At       time.Time
		Updated_At       time.Time
		Deleted_At       time.Time
	}

	CreateMedicalRecord struct {
		Booking_ID       string
		Diagnosis_Result string
		Medicine_Details []Medical_Record_Medicine_Details
		Action_Details   []Medical_Record_Action_Details
		//Items            interface{}
	}

	Medical_Record_Medicine_Details struct {
		ID                string
		Medical_Record_ID string
		Medicine_ID       string
		Medicine_Price    string
		Quantity          int
		Created_At        time.Time
		Updated_At        time.Time
		Deleted_At        time.Time
	}

	Medical_Record_Action_Details struct {
		ID                string
		Medical_Record_ID string
		Action_ID         string
		Action_Price      string
		Created_At        string
		Updated_At        string
		Deleted_At        string
	}
)
