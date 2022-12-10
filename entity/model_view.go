package entity

import "time"

type RegisterView struct {
	ID       string `json: "id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginView struct {
	Username     string `json:"username"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
	Account      string `json:"account"`
}

type CreateUserView struct {
	ID       string `json: "id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Point    int    `json:"point"`
	CostPoints int `json:"costPoints"`
}

type GetUserCountView struct {
	TotalCount int `json:"totalCount"`
}

type GetTransactionsCountView struct {
	TotalTransactions           int `json:"totalTransactions"`
	TotalFailedTransactions     int `json:"totalFailedTransactions"`
	TotalOnProgressTransactions int `json:"totalOnProgressTransactions"`
	TotalSuccessTransactions    int `json:"totalSuccessTransactions"`
}

type TransactionsView struct {
	ID            string    ``
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	PaymentMethod string    `json:"paymentMethod"`
	UserID        string    `json:"userID"`
	Username      string    `json:"username"`
	ProductID     string    `json:"productID"`
	ProductName   string    `json:"productName"`
	SerialNumber  int64     `json:"serialNumber"`
	IdentifierNum string    `json:"identifierNum"`
	Price         int       `json:"price"`
	Status        string    `json:"status"`
}

type UsersView struct {
	ID         string
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	Role       string    `json:"role"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Point      int       `json:"point"`
	CostPoints int       `json:"costPoints"`
}
