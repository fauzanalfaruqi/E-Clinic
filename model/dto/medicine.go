package dto

import "time"

type Medicine struct {
	Id           string    `json:"id"`
	Name         string    `json:"name"`
	MedicineType string    `json:"medicine_type"`
	Price        int       `json:"price"`
	Stock        int       `json:"stock"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_ad"`
	UpdatedAt    time.Time `json:"updated_ad"`
	DeletedAt    time.Time `json:"deleted_ad"`
}

type MedicineResponse struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	MedicineType string `json:"medicine_type"`
	Price        int    `json:"price"`
	Stock        int    `json:"stock"`
	Description  string `json:"description"`
	CreatedAt    string `json:"created_ad"`
	UpdatedAt    string `json:"updated_ad"`
	DeletedAt    string `json:"deleted_ad"`
}
