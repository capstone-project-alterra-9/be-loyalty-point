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
	Status        string `gorm:"type:varchar(100);not null" json:"status"`
}

type CreateUserBinding struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Point    int    `json:"point"`
}

type UpdateUserBinding struct {
	Role       string `json:"role"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Point      int    `json:"point"`
	CostPoints int    `json:"costPoints"`
}
