package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	// Обратите внимание на путь к вашему модулю (замените "your_module_name" на имя из go.mod, если нужно)
	// Например: myproject/internal/db и myproject/internal/handlers
	"12345/internal/db"
	"12345/internal/handlers"
)

func main() {
	// 1. Строка подключения к PostgreSQL (позже вынесем в .env)
	// Если запускаете локально без докера, укажите свои данные (user, password, dbname)
	connStr := os.Getenv("DB_URL")
	if connStr == "" {
		connStr = "postgres://root:secret@localhost:5432/mydb?sslmode=disable"
	}

	// 2. Подключаемся к базе данных
	database, err := db.ConnectDB(connStr)
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}
	defer func(database *sql.DB) {
		err := database.Close()
		if err != nil {
			log.Printf("Ошибка при закрытии соединения с БД: %v", err)
		}
	}(database)

	// 3. Инициализируем хендеры
	authHandler := handlers.NewAuthHandler(database)

	// 4. Настраиваем роутинг (маршруты)
	mux := http.NewServeMux()

	// Страница регистрации (GET - показать форму, POST - отправить данные)
	mux.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			authHandler.RegisterPage(w, r)
			return
		}
		if r.Method == http.MethodPost {
			authHandler.Register(w, r)
			return
		}
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
	})

	// 5. Запуск веб-сервера
	port := ":8080"
	fmt.Printf("Сервер запущен! Перейдите по ссылке: http://localhost-8080/register (или http://localhost%s/register)\n", port)
	log.Printf("Сервер слушает порт %s...", port)

	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
