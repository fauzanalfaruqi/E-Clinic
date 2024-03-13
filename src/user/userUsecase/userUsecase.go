package userUsecase

import (
	"avengers-clinic/model/dto/userDto"
	"avengers-clinic/pkg/utils"
	"avengers-clinic/src/user"
	"errors"
	"time"
)

type userUsecase struct {
	userRepo user.UserRepository
}

func NewUserUsecase(userRepo user.UserRepository) user.UserUsecase {
	return &userUsecase{userRepo}
}

func (usecase *userUsecase) GetAllTrash() ([]userDto.User, error) {
	users, err := usecase.userRepo.GetAllTrash()
	return users, err
}

func (usecase *userUsecase) GetAll() ([]userDto.User, error) {
	users, err := usecase.userRepo.GetAll()
	return users, err
}

func (usecase *userUsecase) GetByID(userID string) (userDto.User, error) {
	user, err := usecase.userRepo.GetByID(userID)
	return user, err
}

func (usecase *userUsecase) PatientRegister(req userDto.AuthRequest) (userDto.User, error) {
	if usecase.userRepo.IsUsernameExists(req.Username) {
		return userDto.User{}, errors.New("1")
	}

	hashPassword, _ := utils.GenerateHashPassword(req.Password)
	now := time.Now().Format("2006-01-02T15:04:05Z")
	var newUser = userDto.User{
		Username: req.Username,
		Password: hashPassword,
		Role: "PATIENT",
		CreatedAt: now,
		UpdatedAt: now,
	}

	var err error
	newUser.ID, err = usecase.userRepo.Insert(newUser)
	if err != nil {
		return userDto.User{}, err
	}
	return newUser, nil
}

func (usecase *userUsecase) UserRegister(req userDto.RegisterRequest) (userDto.User, error) {
	if usecase.userRepo.IsUsernameExists(req.Username) {
		return userDto.User{}, errors.New("1")
	}

	if req.Role == "DOCTOR" && req.Specialization == nil || req.Specialization == "" {
		return userDto.User{}, errors.New("2")
	} else if req.Role != "DOCTOR"  {
		req.Specialization = nil
	}

	hashPassword, _ := utils.GenerateHashPassword(req.Password)
	now := time.Now().Format("2006-01-02 15:04:05")
	var newUser = userDto.User{
		Username: req.Username,
		Password: hashPassword,
		Role: req.Role,
		Specialization: req.Specialization,
		CreatedAt: now,
		UpdatedAt: now,
	}

	var err error
	newUser.ID, err = usecase.userRepo.Insert(newUser)
	if err != nil {
		return userDto.User{}, err
	}
	return newUser, nil
}

func (usecase *userUsecase) Login(req userDto.AuthRequest) (string, error) {
	if !usecase.userRepo.IsUsernameExists(req.Username) {
		return "", errors.New("1")
	}

	user, err := usecase.userRepo.GetByUsername(req.Username)
	if err != nil {
		return "", err
	}

	if utils.VerifyHashPassword(user.Password, req.Password) {
		return "", errors.New("1")
	}

	token, _ := utils.GenerateJWT(user.ID, user.Username, user.Role)
	return token, nil
}

func (usecase *userUsecase) Update(req userDto.UpdateRequest) (userDto.User, error) {
	user, err := usecase.userRepo.GetByID(req.ID)
	if err != nil {
		return userDto.User{}, err
	}

	if req.Username != "" {
		if usecase.userRepo.IsUsernameExists(req.Username) && user.Username != req.Username {
			return userDto.User{}, errors.New("1")
		}
		user.Username = req.Username
	}

	if req.Specialization != nil {
		user.Specialization = req.Specialization
	}

	if user.Role != "DOCTOR" {
		user.Specialization = nil
	}
	user.UpdatedAt = time.Now().Format("2006-01-02T15:04:05Z")

	err = usecase.userRepo.Update(user)
	if err != nil {
		return userDto.User{}, err
	}
	return user, nil
}

func (usecase *userUsecase) UpdatePassword(req userDto.UpdatePasswordRequest) error {
	user, err := usecase.userRepo.GetByID(req.ID)
	if err != nil {
		return err
	}

	if utils.VerifyHashPassword(user.Password, req.CurrentPassword) {
		return errors.New("1")
	}

	if req.NewPassword != req.ConfirmationPassword {
		return errors.New("2")
	}

	hashPassword, _ := utils.GenerateHashPassword(req.NewPassword)
	err = usecase.userRepo.UpdatePassword(req.ID, hashPassword)
	if err != nil {
		return err
	}
	return nil
}

func (usecase *userUsecase) Delete(userID string) error {
	_, err := usecase.userRepo.GetUserByID(userID)
	if err != nil {
		return err
	}
	err = usecase.userRepo.Delete(userID)
	return err
}

func (usecase *userUsecase) SoftDelete(userID string) error {
	_, err := usecase.userRepo.GetByID(userID)
	if err != nil {
		return err
	}
	err = usecase.userRepo.SoftDelete(userID)
	return err
}

func (usecase *userUsecase) Restore(userID string) error {
	_, err := usecase.userRepo.GetTrashByID(userID)
	if err != nil {
		return err
	}
	err = usecase.userRepo.Restore(userID)
	return err
}