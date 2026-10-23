package service

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// UserRepo описывает интерфейс, чтобы сервис мог работать с базой через репозиторий
type UserRepo interface {
	Create(email, passwordHash string) error
}

type AuthService struct {
	userRepo UserRepo
}

func NewAuthService(userRepo UserRepo) *AuthService {
	return &AuthService{userRepo: userRepo}
}

// Register выполняет бизнес-логику регистрации
func (s *AuthService) Register(email, password string) error {
	if email == "" || password == "" {
		return errors.New("email и пароль не должны быть пустыми")
	}

	// 1. Хэшируем пароль перед сохранением
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("ошибка обработки пароля: %w", err)
	}

	// 2. Сохраняем пользователя через репозиторий
	err = s.userRepo.Create(email, string(hashedPassword))
	if err != nil {
		return err
	}

	return nil
}
