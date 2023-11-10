package models

type User struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Authentication struct {
	Id_user string `json:"id"`
	Token   string `json:"token"`
}
