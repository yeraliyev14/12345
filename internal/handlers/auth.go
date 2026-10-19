package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"

	"golang.org/x/crypto/bcrypt" // Пакет для хэширования паролей
)

// AuthHandler структура для работы с хендерами авторизации/регистрации
type AuthHandler struct {
	DB *sql.DB
}

// NewAuthHandler конструктор
func NewAuthHandler(db *sql.DB) *AuthHandler {
	return &AuthHandler{DB: db}
}

// RegisterPage отображает HTML-форму регистрации
func (h *AuthHandler) RegisterPage(w http.ResponseWriter, r *http.Request) {
	tmplPath := filepath.Join("web", "templates", "register.html")

	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Ошибка загрузки шаблона: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

// Register процесс обработки отправки формы и сохранения в базу
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	if email == "" || password == "" {
		http.Error(w, "Email и пароль не должны быть пустыми", http.StatusBadRequest)
		return
	}

	// 1. Хэшируем пароль перед сохранением (никогда не храните пароли в открытом виде!)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Ошибка обработки пароля", http.StatusInternalServerError)
		return
	}

	// 2. Сохраняем пользователя в базу данных (используем таблицу users из вашей миграции)
	query := `INSERT INTO users (email, password_hash) VALUES ($1, $2)`
	_, err = h.DB.Exec(query, email, string(hashedPassword))
	if err != nil {
		// Здесь может быть ошибка, если email уже существует (уникальный индекс)
		http.Error(w, "Ошибка регистрации (возможно, такой email уже занят): "+err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("Пользователь успешно зарегистрирован: %s\n", email)
	fmt.Fprintln(w, "Регистрация прошла успешно! Пользователь сохранен в базу данных.")
}
