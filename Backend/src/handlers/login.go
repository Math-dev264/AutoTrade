package handlers

import (
	"net/http"

	"github.com/Math-dev264/AutoTrade/Backend/src/auth"
	"github.com/Math-dev264/AutoTrade/Backend/src/utils"
	"github.com/labstack/echo/v4"
)

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
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

	token, err := auth.CreateToken()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Erro ao criar o token")
	}

	// userID := strconv.FormatUint(userDB.ID, 10)
	return c.JSON(http.StatusOK, map[string]string{"token_access": token})
	// c.JSON(http.StatusOK, models.Authentication{ID: userID, Token: token})
}
