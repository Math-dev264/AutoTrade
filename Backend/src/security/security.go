package security

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	const saltRoundsCost = 11
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), saltRoundsCost)
	return string(bytes), err
}
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
