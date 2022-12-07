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

type TransactionsBinding struct {
	PaymentMethod string `json:"paymentMethod"`
	ProductID     string `json:"productID"`
	IdentifierNum string `json:"identifierNum"`
}

type UpdateTransactionBinding struct {
	Status string `json:"status"`
}

type CreateUserBinding struct {
	Username 	string `json:"username"`
	Email    	string `json:"email"`
	Password 	string `json:"password"`
	Role		string `json:"role"`
}
