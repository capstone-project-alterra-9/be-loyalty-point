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
	PaymentMethod string `gorm:"type:varchar(20);not null" json:"paymentMethod"`
	UserID        string `gorm:"type:varchar(100);not null" json:"userID"`
	ProductID     string `gorm:"type:varchar(100);not null" json:"productID"`
	IdentifierNum string `gorm:"type:varchar(100);not null" json:"identifierNum"`
}

type CreateUserBinding struct {
	Username 	string `json:"username"`
	Email    	string `json:"email"`
	Password 	string `json:"password"`
	Role		string `json:"role"`
}
