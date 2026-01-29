package user

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	CreateUser(email, password string) (*User, error)
	GetUserByID(id string) (*User, error)
	GetAllUsers() ([]*User, error)
	UpdateUser(id string, email, password string) (*User, error)
	DeleteUser(id string) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateUser(email, password string) (*User, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &User{
		ID:       uuid.NewString(),
		Email:    email,
		Password: string(hashedPassword),
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *service) GetUserByID(id string) (*User, error) {
	return s.repo.FindByID(id)
}

func (s *service) GetAllUsers() ([]*User, error) {
	return s.repo.FindAll()
}

func (s *service) UpdateUser(id string, email, password string) (*User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if email != "" {
		user.Email = email
	}

	if password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		user.Password = string(hashedPassword)
	}

	if err := s.repo.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *service) DeleteUser(id string) error {
	return s.repo.Delete(id)
}
