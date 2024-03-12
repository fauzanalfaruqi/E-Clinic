package userDto

type User struct {
	ID             string      `json:"id,omitempty"`
	Username       string      `json:"username,omitempty"`
	Password       string      `json:"-"`
	Role           string      `json:"role,omitempty"`
	Specialization interface{} `json:"specialization,omitempty"`
	CreatedAt      string      `json:"created_at,omitempty"`
	UpdatedAt      string      `json:"updated_at,omitempty"`
	DeletedAt      interface{} `json:"deleted_at,omitempty"`
}

type AuthRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegisterRequest struct {
	Username       string      `json:"username" validate:"required"`
	Password       string      `json:"password" validate:"required"`
	Role           string      `json:"role" validate:"required,enum=ADMIN DOCTOR PATIENT"`
	Specialization interface{} `json:"specialization"`
}

type UpdateRequest struct {
	ID             string      `json:"id"`
	Username       string      `json:"username"`
	Specialization interface{} `json:"specialization"`
}

type UpdatePasswordRequest struct {
	ID                   string `json:"id"`
	CurrentPassword      string `json:"current_password" validate:"required"`
	NewPassword          string `json:"new_password" validate:"required"`
	ConfirmationPassword string `json:"confirmation_password" validate:"required"`
}