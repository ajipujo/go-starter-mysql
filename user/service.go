package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	LoginUser(input LoginUserInput) (User, error)
	GetUserByID(userID int) (User, error)
	UpdateUser(userID int, input UpdateUserInput) (User, error)
	GetEmailAvailability(input CheckEmailAvailabilityInput) (bool, error)
	SaveAvatar(ID int, pathFile string) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	createUser := User{}

	createUser.Name = input.Name
	createUser.Email = input.Email

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		return createUser, err
	}

	createUser.Password = string(passwordHash)
	createUser.RoleID = 2

	newUser, err := s.repository.Save(createUser)

	if err != nil {
		return newUser, err
	}

	return newUser, nil

}

func (s *service) LoginUser(input LoginUserInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindUserByEmail(email)

	if err != nil {
		return user, errors.New("failed to find user")
	}

	if user.ID == 0 {
		return user, errors.New("no user found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return user, errors.New("wrong password")
	}

	return user, nil
}

func (s *service) GetUserByID(userID int) (User, error) {
	user, err := s.repository.FindUserByID(userID)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("user not found")
	}

	return user, nil
}

func (s *service) UpdateUser(userID int, input UpdateUserInput) (User, error) {
	existUser, err := s.repository.FindUserByID(userID)

	if err != nil {
		return existUser, err
	}

	if existUser.ID == 0 {
		return existUser, errors.New("user not found")
	}

	existUser.Name = input.Name
	existUser.Email = input.Email

	newUser, err := s.repository.Update(existUser)

	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) GetEmailAvailability(input CheckEmailAvailabilityInput) (bool, error) {
	email := input.Email

	userData, err := s.repository.FindUserByEmail(email)

	if err != nil {
		return false, errors.New("checking email failed.")
	}

	if userData.ID == 0 {
		return true, nil
	}

	return false, nil
}

func (s *service) SaveAvatar(ID int, pathFile string) (User, error) {
	user, err := s.repository.FindUserByID(ID)

	if err != nil {
		return user, err
	}

	user.AvatarFileName = pathFile

	updatedUser, err := s.repository.Update(user)

	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}
