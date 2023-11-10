package utils

import (
	"log"

	"github.com/Math-dev264/AutoTrade/Backend/src/db"
)

func CheckIfEmailExist(email string) bool {
	conn, err := db.OpenConnection()
	if err != nil {
		return false
	}
	defer conn.Close()

	var count int
	err = conn.QueryRow("SELECT COUNT(*) FROM users WHERE email = $1", email).Scan(&count)
	if err != nil {
		// Lidar com o erro, por exemplo, logar e retornar false
		log.Println("Erro ao verificar duplicidade de email:", err)
		return false
	}
	return count > 0
}
