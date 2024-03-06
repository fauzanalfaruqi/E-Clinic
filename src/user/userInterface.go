package user

import "avengers-clinic/model/dto/userDto"

type UserRepository interface {
	GetByID(userID string) (userDto.User, error)
	GetByUsername(username string) (userDto.User, error)
	Insert(user userDto.User) (string, error)
	IsUsernameExists(username string) bool
}

type UserUsecase interface {
	PatientRegister(req userDto.AuthRequest) (userDto.User, error)
	UserRegister(req userDto.RegisterRequest) (userDto.User, error)
	Login(req userDto.AuthRequest) (userDto.LoginResponse, error)
}