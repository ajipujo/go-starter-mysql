package user

import "golang.org/x/crypto/bcrypt"

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
}

type service struct {
	repository *repository
}

func NewService(repository *repository) *service {
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

	newUser, err := s.repository.SaveUser(createUser)

	if err != nil {
		return newUser, err
	}

	return newUser, nil

}
