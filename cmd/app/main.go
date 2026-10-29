package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	db "12345/internal/database"
	deliveryHttp "12345/internal/delivery/http" // <-- Даем псевдоним, чтобы не путать со встроенным net/http
	"12345/internal/repository"
	"12345/internal/service"
)

func main() {
	connStr := os.Getenv("DB_URL")
	if connStr == "" {
		connStr = "postgres://root:secret@localhost:5432/mydb?sslmode=disable"
	}

	database, err := db.ConnectDB(connStr)
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}

	if err := db.RunMigrations(database, "mydb"); err != nil {
		log.Fatalf("Ошибка применения миграций: %v", err)
	}

	defer database.Close()

	// Связываем слои
	userRepo := repository.NewUserRepository(database)
	authService := service.NewAuthService(userRepo)

	// Используем наш алиас deliveryHttp
	authHandler := deliveryHttp.NewAuthHandler(authService)

	mux := http.NewServeMux()

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

	port := ":8080"
	fmt.Printf("Сервер запущен! Перейдите по ссылке: http://localhost%s/register\n", port)

	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
