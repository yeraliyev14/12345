package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"
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
	// Путь к нашему шаблону, который мы создали ранее
	tmplPath := filepath.Join("web", "templates", "register.html")

	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Ошибка загрузки шаблона: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

// Register процесс обработки отправки формы
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	// Получаем данные из формы
	email := r.FormValue("email")
	password := r.FormValue("password")

	// Простая проверка (в будущем добавим хеширование пароля и запись в БД)
	fmt.Printf("Получена регистрация: Email = %s\n", email)

	// Перенаправляем пользователя после успешной регистрации (например, на главную)
	fmt.Fprintln(w, "Регистрация прошла успешно! (Временный ответ)")
	_ = password // пока не используем до хэширования
}
