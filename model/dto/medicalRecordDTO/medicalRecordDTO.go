package medicalRecordDTO

type (
	MedicalRecord struct {
		ID               string                            `json:"id,omitempty"`
		Booking_ID       string                            `json:"booking_id,omitempty"`
		Diagnosis_Result string                            `json:"diagnosis_result,omitempty"`
		Total_Medicine   int                               `json:"total_medicine,omitempty"`
		Total_Action     int                               `json:"total_action,omitempty"`
		Total_Amount     int                               `json:"total_amount,omitempty"`
		Payment_Status   bool                              `json:"payment_status,omitempty"`
		Medicine_Details []Medical_Record_Medicine_Details `json:"medicine_details,omitempty"`
		Action_Details   []Medical_Record_Action_Details   `json:"action_details,omitempty"`
		Created_At       string                            `json:"created_at,omitempty"`
		Updated_At       string                            `json:"updated_at,omitempty"`
		Deleted_At       string                            `json:"deleted_at,omitempty"`
	}

	CreateMedicalRecord struct {
		Booking_ID       string                            `json:"booking_id,omitempty" validate:"required"`
		Diagnosis_Result string                            `json:"diagnosis_result,omitempty" validate:"required"`
		Medicine_Details []Medical_Record_Medicine_Details `json:"medicine_details,omitempty" validate:"required"`
		Action_Details   []Medical_Record_Action_Details   `json:"action_details,omitempty" validate:"required"`
	}

	Medical_Record_Medicine_Details struct {
		ID                string `json:"id,omitempty"`
		Medical_Record_ID string `json:"medical_record_id,omitempty"`
		Medicine_ID       string `json:"medicine_id,omitempty"`
		Medicine_Price    int    `json:"medicine_price,omitempty"`
		Quantity          int    `json:"quantity,omitempty"`
		Created_At        string `json:"created_at,omitempty"`
		Updated_At        string `json:"updated_at,omitempty"`
		Deleted_At        string `json:"deleted_at,omitempty"`
	}

	Medical_Record_Action_Details struct {
		ID                string `json:"id,omitempty"`
		Medical_Record_ID string `json:"medical_record_id,omitempty"`
		Action_ID         string `json:"action_id,omitempty"`
		Action_Price      int    `json:"action_price,omitempty"`
		Created_At        string `json:"created_at,omitempty"`
		Updated_At        string `json:"updated_at,omitempty"`
		Deleted_At        string `json:"deleted_at,omitempty"`
	}
)
