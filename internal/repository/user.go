package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/lib/pq" // Обязательно нужен для проверки ошибок Postgres
)

// Создадим кастомную ошибку, которую сервис сможет легко распознать
var ErrEmailAlreadyExists = errors.New("такой email уже зарегистрирован")

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(email, passwordHash string) error {
	query := `INSERT INTO users (email, password_hash) VALUES ($1, $2)`
	_, err := r.db.Exec(query, email, passwordHash)

	if err != nil {
		// Проверяем, является ли ошибка ошибкой PostgreSQL
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			// Код "23505" означает нарушение уникальности (unique constraint violation)
			if pqErr.Code == "23505" {
				return ErrEmailAlreadyExists
			}
		}
		return fmt.Errorf("ошибка записи в БД: %w", err)
	}

	return nil
}
