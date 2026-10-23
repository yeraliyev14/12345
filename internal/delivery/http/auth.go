package http

import (
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"

	"12345/internal/repository"
	"12345/internal/service"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
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

// Register процесс обработки отправки формы
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	// Вызываем бизнес-логику из сервиса
	err := h.authService.Register(email, password)
	if err != nil {
		// Проверяем, не занят ли уже этот email
		if errors.Is(err, repository.ErrEmailAlreadyExists) {
			http.Error(w, "Этот email уже занят. Пожалуйста, используйте другой.", http.StatusBadRequest)
			return
		}

		// Для остальных ошибок возвращаем 500 статус
		http.Error(w, "Ошибка регистрации: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("Пользователь успешно зарегистрирован: %s\n", email)
	fmt.Fprintln(w, "Регистрация прошла успешно! Пользователь сохранен в базу данных.")
}
