package userDto

type User struct {
	ID             string      `json:"id,omitempty"`
	Username       string      `json:"username,omitempty"`
	Password       string      `json:"password,omitempty"`
	Role           string      `json:"role,omitempty"`
	Specialization interface{} `json:"specialization,omitempty"`
	CreatedAt      string      `json:"created_at,omitempty"`
	UpdatedAt      string      `json:"updated_at,omitempty"`
	DeletedAt      string      `json:"deleted_at,omitempty"`
}

type AuthRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegisterRequest struct {
	Username       string      `json:"username" validate:"required"`
	Password       string      `json:"password" validate:"required"`
	Role           string      `json:"role" validate:"required"`
	Specialization interface{} `json:"specialization"`
}

type LoginResponse struct {
	Token string `json:"token"`
}