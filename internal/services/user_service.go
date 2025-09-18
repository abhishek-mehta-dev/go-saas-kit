package services

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/abhishek-mehta-dev/go-saas-kit.git/internal/models"
	"github.com/abhishek-mehta-dev/go-saas-kit.git/internal/repositories"
	"github.com/abhishek-mehta-dev/go-saas-kit.git/pkg/email"
)

type UserService interface {
	Register(user *models.User) error
}

type userService struct {
	repo   repositories.UserRepository
	sender email.EmailSender
}

func NewUserService(repo repositories.UserRepository, sender email.EmailSender) UserService {
	return &userService{repo: repo, sender: sender}
}

func (s *userService) Register(user *models.User) error {
	// Check duplicate
	existing, _ := s.repo.FindByEmail(user.Email)
	if existing != nil {
		return errors.New("email already registered")
	}

	// Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)

	// Save user
	if err := s.repo.Create(user); err != nil {
		return err
	}

	// Send welcome email
	subject := "Welcome to Go SaaS Kit"
	body, err := email.RenderTemplate("welcome.html", email.TemplateData{
	"FirstName": user.FirstName,
	"Title":     "Welcome to Go SaaS Kit",
	})
	if err != nil {
		return err
	}

	if err := s.sender.Send(user.Email, subject, body); err != nil {
	return err
	}

	return nil
}
