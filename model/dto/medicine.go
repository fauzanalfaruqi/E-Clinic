package dto

type MedicineRequest struct {
	Id           string      `json:"id"`
	Name         string      `json:"name" validate:"required"`
	MedicineType string      `json:"medicine_type" validate:"required,enum=TABLET KAPSUL OLES CAIR TETES"`
	Price        int         `json:"price" validate:"required"`
	Stock        int         `json:"stock"`
	Description  interface{} `json:"description"`
	CreatedAt    string      `json:"created_ad"`
	UpdatedAt    string      `json:"updated_ad"`
	DeletedAt    string      `json:"deleted_ad"`
}

type UpdateRequest struct {
	Id           string      `json:"id"`
	Name         string      `json:"name"`
	MedicineType string      `json:"medicine_type" validate:"enum=TABLET KAPSUL OLES CAIR TETES "`
	Price        int         `json:"price"`
	Stock        int         `json:"stock"`
	Description  interface{} `json:"description"`
	CreatedAt    string      `json:"created_ad"`
	UpdatedAt    string      `json:"updated_ad"`
	DeletedAt    string      `json:"deleted_ad"`
}

type MedicineResponse struct {
	Id           string      `json:"id"`
	Name         string      `json:"name"`
	MedicineType string      `json:"medicine_type"`
	Price        int         `json:"price"`
	Stock        int         `json:"stock,omitempty"`
	Description  interface{} `json:"description,omitempty"`
	CreatedAt    string      `json:"created_ad,omitempty"`
	UpdatedAt    string      `json:"updated_ad,omitempty"`
	DeletedAt    string      `json:"deleted_ad,omitempty"`
}
