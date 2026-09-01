package usecase

import (
	"donasi-app/entity"
	"donasi-app/repositories"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	Register(name, email, password, role string) (*entity.User, error)
	Login(email, password string) (*entity.User, error)
}

type userUsecase struct {
	userRepo repositories.UserRepository
}

func NewUserUsecase(userRepo repositories.UserRepository) UserUsecase {
	return &userUsecase{userRepo: userRepo}
}

func (u *userUsecase) Register(name, email, password, role string) (*entity.User, error) {
	// Cek apakah email sudah terdaftar
	existingUser, _ := u.userRepo.FindByEmail(email)
	if existingUser != nil {
		return nil, errors.New("email already registered")
	}

	// Hash password sebelum disimpan
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	newUser := &entity.User{
		Name:         name,
		Email:        email,
		PasswordHash: string(hashedPassword),
		Role:         role,
	}

	err = u.userRepo.Create(newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (u *userUsecase) Login(email, password string) (*entity.User, error) {
	user, err := u.userRepo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// Verifikasi password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, errors.New("invalid password")
	}

	return user, nil
}
