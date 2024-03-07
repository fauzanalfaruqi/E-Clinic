package medicalRecordDTO

type (
	Medical_Record struct {
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

	Medical_Record_Request struct {
		ID               string                     `json:"id,omitempty"`
		Booking_ID       string                     `json:"booking_id,omitempty" validate:"require"`
		Diagnosis_Result string                     `json:"diagnosis_result,omitempty" validate:"required"`
		Payment_Status   bool                       `json:"payment_status,omitempty"`
		Medicine_Details []Medicine_Details_Request `json:"medicine_details,omitempty" validate:"required"`
		Action_Details   []Action_Details_Request   `json:"action_details,omitempty" validate:"required"`
	}

	Medical_Record_Medicine_Details struct {
		ID                string `json:"id,omitempty"`
		Medical_Record_ID string `json:"medical_record_id,omitempty"`
		Medicine_ID       string `json:"medicine_id,omitempty"`
		Medicine_Name     string `json:"medicine_name,omitempty"`
		Medicine_Price    int    `json:"medicine_price,omitempty"`
		Quantity          int    `json:"quantity,omitempty"`
		Medicine_Stock    int    `json:"medicine_stock,omitempty"`
		Created_At        string `json:"created_at,omitempty"`
		Updated_At        string `json:"updated_at,omitempty"`
		Deleted_At        string `json:"deleted_at,omitempty"`
	}

	Medicine_Details_Request struct {
		Medicine_ID string `json:"medicine_id,omitempty" validate:"required"`
		Quantity    int    `json:"quantity,omitempty" validate:"required"`
	}

	Medical_Record_Action_Details struct {
		ID                 string `json:"id,omitempty"`
		Medical_Record_ID  string `json:"medical_record_id,omitempty"`
		Action_ID          string `json:"action_id,omitempty"`
		Action_Name        string `json:"action_name,omitempty"`
		Action_Price       int    `json:"action_price,omitempty"`
		Action_Description string `json:"action_description,omitempty"`
		Created_At         string `json:"created_at,omitempty"`
		Updated_At         string `json:"updated_at,omitempty"`
		Deleted_At         string `json:"deleted_at,omitempty"`
	}

	Action_Details_Request struct {
		Medical_Record_ID string `json:"medical_record_id,omitempty" validate:"required"`
		Action_ID         string `json:"action_id,omitempty" validate:"required"`
	}
)
