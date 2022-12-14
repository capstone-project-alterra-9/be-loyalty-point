package entity

import (
	"time"

	"github.com/google/uuid"
)

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
	Points     int    `json:"points"`
	CostPoints int    `json:"costPoints"`
}

type MidtransTransactionBinding struct {
	PaymentMethod string `json:"paymentMethod"`
	ProductID     string `json:"productID"`
	UserId 		  string `json:"userID"`
}

type TokenBinding struct {
	RefreshToken string `json:"token"`
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type PayloadBinding struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

type UpdateUserByEncryptedIdPayload struct {
	NewPassword        string `json:"newPassword"`
}