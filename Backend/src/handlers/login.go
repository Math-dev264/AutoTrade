package handlers

import (
	"log"
	"net/http"

	"github.com/Math-dev264/AutoTrade/Backend/src/auth"
	"github.com/Math-dev264/AutoTrade/Backend/src/db"
	"github.com/Math-dev264/AutoTrade/Backend/src/security"
	"github.com/Math-dev264/AutoTrade/Backend/src/utils"
	"github.com/labstack/echo/v4"
)

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func getPassword(email string) (hash string, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `SELECT password FROM users
						WHERE email = $1`

	err = conn.QueryRow(sql, email).Scan(&hash)
	return hash, err
}

func getUserId(email string) (id uint64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `SELECT id FROM users
						WHERE email = $1`

	err = conn.QueryRow(sql, email).Scan(&id)
	return id, err
}

func postAuthentication(userId uint64, token string) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `INSERT INTO authentication (id_user, token)
						VALUES ($1, $2)`

	_, err = conn.Exec(sql, userId, token)
	if err != nil {
		log.Print("Erro inserir registro de token no banco")
	}
}

func Login(c echo.Context) error {
	bodyRequest := new(loginRequest)
	if err := c.Bind(bodyRequest); err != nil {
		return c.String(http.StatusUnprocessableEntity, "Erro ao bindar")
	}

	emailExist := utils.CheckIfEmailExist(bodyRequest.Email)
	if !emailExist {
		return c.String(http.StatusBadRequest, "Não há nenhuma conta com este e-mail")
	}

	hash, err := getPassword(bodyRequest.Email)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Senha não encontrada")
	}

	equal := security.CheckPasswordHash(bodyRequest.Password, hash)
	if !equal {
		return c.String(402, "Senha e hash são diferentes")
	}

	userId, err := getUserId(bodyRequest.Email)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Id não encontrado")
	}

	token, err := auth.CreateToken(userId, bodyRequest.Email)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Erro ao criar o token")
	}

	postAuthentication(userId, token)

	return c.JSON(http.StatusOK, map[string]string{"token_access": token})
}
