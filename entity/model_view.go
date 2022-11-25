package entity

type RegisterView struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
