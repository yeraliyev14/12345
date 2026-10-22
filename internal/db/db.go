package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // Драйвер чтения миграций из файлов
	_ "github.com/lib/pq"                                // Драйвер PostgreSQL
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

// RunMigrations автоматически применяет миграции из папки migrations/
func RunMigrations(dbConn *sql.DB, databaseName string) error {
	driver, err := postgres.WithInstance(dbConn, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("ошибка создания драйвера миграций: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations", // Путь к папке с миграциями относительно корня проекта
		databaseName,
		driver,
	)
	if err != nil {
		return fmt.Errorf("ошибка инициализации мигратора: %w", err)
	}

	// Запуск миграций (Up)
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("ошибка применения миграций: %w", err)
	}

	log.Println("Миграции успешно применены!")
	return nil
}
