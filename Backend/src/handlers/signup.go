package handlers

import (
	"log"
	"net/http"

	"github.com/Math-dev264/AutoTrade/Backend/src/db"
	"github.com/Math-dev264/AutoTrade/Backend/src/models"
	"github.com/Math-dev264/AutoTrade/Backend/src/security"
	"github.com/Math-dev264/AutoTrade/Backend/src/utils"
	"github.com/badoux/checkmail"
	"github.com/labstack/echo/v4"
)

type RegisterUserRequest struct {
	Name     string `json:"name" `
	Email    string `json:"email"`
	Password string `json:"password"`
}

func post(user models.User) (id uint64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `INSERT INTO users (name, email, password)
						VALUES ($1, $2, $3)
						RETURNING id`

	err = conn.QueryRow(sql, user.Name, user.Email, user.Password).Scan(&id)
	return id, err
}

func validate(user models.User) string {
	if user.Name == "" {
		return "O nome é requirido"
	}
	if user.Email == "" {
		return "O email é requirido"
	}
	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return "O email é inválido"
	}
	duplicateEmail := utils.CheckIfEmailExist(user.Email)
	if duplicateEmail {
		return "O email já está sendo utilizado"
	}
	if user.Password == "" {
		return "A senha é requirida"
	}

	return ""
}

func Signup(c echo.Context) error {
	request := new(RegisterUserRequest)
	if err := c.Bind(request); err != nil {
		log.Println("Erro na bind:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Formato de data inválido"})
	}

	hash, err := security.HashPassword(request.Password)
	if err != nil {
		return c.String(http.StatusNotAcceptable, "Erro ao fazer o hash")
	}

	user := models.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: hash,
	}

	if err := validate(user); err != "" {
		return c.String(http.StatusBadRequest, err)
	}

	id, err := post(user)
	if err != nil {
		log.Println("Erro: ", err)
		return c.String(http.StatusInternalServerError, "Erro ao criar usuário.")
	}

	user.ID = id

	return c.JSON(http.StatusOK, user)
}
