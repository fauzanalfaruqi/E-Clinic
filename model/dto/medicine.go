package dto

type Medicine struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	MedicineType string `json:"medicine_type"`
	Price        int    `json:"price"`
	Stock        int    `json:"stock"`
}
