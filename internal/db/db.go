package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Драйвер PostgreSQL
)

// ConnectDB инициализирует и проверяет подключение к базе данных
func ConnectDB(connStr string) (*sql.DB, error) {
	database, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("ошибка открытия БД: %w", err)
	}

	// Проверяем реальное соединение
	if err := database.Ping(); err != nil {
		return nil, fmt.Errorf("ошибка проверки соединения (ping): %w", err)
	}

	fmt.Println("Успешное подключение к базе данных!")
	return database, nil
}
