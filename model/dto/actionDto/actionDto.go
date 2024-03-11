package actionDto

type Action struct {
	ID          string      `json:"id,omitempty"`
	Name        string      `json:"name,omitempty"`
	Price       int         `json:"price,omitempty"`
	Description interface{} `json:"description,omitempty"`
	CreatedAt   string      `json:"created_at,omitempty"`
	UpdatedAt   string      `json:"updated_at,omitempty"`
	DeletedAt   string      `json:"deleted_at,omitempty"`
}

type CreateRequest struct {
	Name        string      `json:"name" validate:"required"`
	Price       int         `json:"price" validate:"required,number"`
	Description interface{} `json:"description"`
}

type UpdateRequest struct {
	ID          string
	Name        string `json:"name"`
	Price       int    `json:"price" validate:"number"`
	Description string `json:"description"`
}