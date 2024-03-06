package userUsecase

import (
	"avengers-clinic/model/dto/userDto"
	"avengers-clinic/pkg/utils"
	"avengers-clinic/src/user"
	"errors"
	"fmt"
	"time"
)

type userUsecase struct {
	userRepo user.UserRepository
}

func NewUserUsecase(userRepo user.UserRepository) user.UserUsecase {
	return &userUsecase{userRepo}
}

func (usecase *userUsecase) PatientRegister(req userDto.AuthRequest) (userDto.User, error) {
	if usecase.userRepo.IsUsernameExists(req.Username) {
		return userDto.User{}, errors.New("1")
	}

	hashPassword, err := utils.GenerateHashPassword(req.Password)
	if err != nil {
		return userDto.User{}, err
	}

	now := time.Now().Format("2006-01-02 15:04:05")
	var newUser = userDto.User{
		Username: req.Username,
		Password: hashPassword,
		Role: "PATIENT",
		CreatedAt: now,
		UpdatedAt: now,
	}

	newUser.ID, err = usecase.userRepo.Insert(newUser)
	if err != nil {
		return userDto.User{}, err
	}

	newUser.Password = ""
	return newUser, nil
}

func (usecase *userUsecase) UserRegister(req userDto.RegisterRequest) (userDto.User, error) {
	if usecase.userRepo.IsUsernameExists(req.Username) {
		return userDto.User{}, errors.New("1")
	}

	role := req.Role
	if role != "ADMIN" && role != "DOCTOR" && role != "PATIENT" {
		return userDto.User{}, errors.New("2")
	}

	if role == "DOCTOR" && req.Specialization == nil {
		return userDto.User{}, errors.New("3")
	}

	hashPassword, err := utils.GenerateHashPassword(req.Password)
	if err != nil {
		return userDto.User{}, err
	}

	now := time.Now().Format("2006-01-02 15:04:05")
	var newUser = userDto.User{
		Username: req.Username,
		Password: hashPassword,
		Role: role,
		Specialization: req.Specialization,
		CreatedAt: now,
		UpdatedAt: now,
	}

	newUser.ID, err = usecase.userRepo.Insert(newUser)
	if err != nil {
		return userDto.User{}, err
	}

	newUser.Password = ""
	return newUser, nil
}

func (usecase *userUsecase) Login(req userDto.AuthRequest) (userDto.LoginResponse, error) {
	if !usecase.userRepo.IsUsernameExists(req.Username) {
		return userDto.LoginResponse{}, errors.New("1")
	}

	user, err := usecase.userRepo.GetByUsername(req.Username)
	if err != nil {
		return userDto.LoginResponse{}, err
	}

	if utils.VerifyHashPassword(user.Password, req.Password) {
		return userDto.LoginResponse{}, errors.New("1")
	}

	token, err := utils.GenerateJWT(user.Username, user.Role)
	fmt.Println("Token:", token)
	if err != nil {
		return userDto.LoginResponse{}, err
	}

	return userDto.LoginResponse{Token: token}, nil
}