package user

import "avengers-clinic/model/dto/userDto"

type UserRepository interface {
	GetAllTrash() ([]userDto.User, error)
	GetAll() ([]userDto.User, error)
	GetByID(userID string) (userDto.User, error)
	GetTrashByID(userID string) (userDto.User, error)
	GetByUsername(username string) (userDto.User, error)
	Insert(user userDto.User) (string, error)
	Update(user userDto.User) error
	UpdatePassword(userId, hashPassword string) error
	Delete(userID string) error
	SoftDelete(userID string) error
	Restore(userID string) error
	IsUsernameExists(username string) bool
}

type UserUsecase interface {
	GetAllTrash() ([]userDto.User, error)
	GetAll() ([]userDto.User, error)
	GetByID(userID string) (userDto.User, error)
	PatientRegister(req userDto.AuthRequest) (userDto.User, error)
	UserRegister(req userDto.RegisterRequest) (userDto.User, error)
	Login(req userDto.AuthRequest) (userDto.LoginResponse, error)
	Update(req userDto.UpdateRequest) (userDto.User, error)
	UpdatePassword(req userDto.UpdatePasswordRequest) error
	Delete(userID string) error
	SoftDelete(userID string) error
	Restore(userID string) error
}