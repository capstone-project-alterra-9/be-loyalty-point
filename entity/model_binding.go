package entity

type LoginBinding struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterBinding struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
