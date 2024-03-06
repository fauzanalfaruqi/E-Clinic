package actionDto

type Action struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Price       int    `json:"price,omitempty"`
	Description string `json:"description,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
}