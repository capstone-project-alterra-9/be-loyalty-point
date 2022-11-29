package entity

import (
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

type Points struct {
	ID     string `gorm:"primaryKey"`
	UserID string `gorm:"type:varchar(100);not null" json:"userID"`
	Points int    `gorm:"not null" json:"points"`
}

type Users struct {
	ID        string         `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"type:timestamp;not null" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"type:timestamp;not null" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	Role      string         `gorm:"type:varchar(10);not null" json:"role"`
	Username  string         `gorm:"type:varchar(20);not null;unique" json:"username"`
	Email     string         `gorm:"type:varchar(100);not null;unique" json:"email"`
	Password  string         `gorm:"type:varchar(100);not null" json:"password"`
}

type Products struct {
	ID          string         `gorm:"primaryKey"`
	CreatedAt   time.Time      `gorm:"type:timestamp;not null" json:"createdAt"`
	UpdatedAt   time.Time      `gorm:"type:timestamp;not null" json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	Category    string         `gorm:"type:varchar(20);not null" json:"category"`
	Name        string         `gorm:"type:varchar(100);not null" json:"name"`
	Description string         `gorm:"type:varchar(200);not null" json:"description"`
	Price       int            `gorm:"type:int;not null" json:"price"`
	Stock       int            `gorm:"type:int;not null" json:"stock"`
	Image       string         `gorm:"type:varchar(200);not null" json:"image"`
}

type SerialNumbers struct {
	ID        string         `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"type:timestamp;not null" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"type:timestamp;not null" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	ProductID string         `gorm:"type:varchar(100);not null" json:"productID"`
	Serial    string         `gorm:"type:varchar(100);not null" json:"serial"`
	Status    string         `gorm:"type:varchar(20);not null" json:"status"`
}

type Transactions struct {
	ID            string         `gorm:"primaryKey"`
	CreatedAt     time.Time      `gorm:"type:timestamp;not null" json:"createdAt"`
	UpdatedAt     time.Time      `gorm:"type:timestamp;not null" json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	PaymentMethod string         `gorm:"type:varchar(20);not null" json:"paymentMethod"`
	UserID        string         `gorm:"type:varchar(100);not null" json:"userID"`
	ProductID     string         `gorm:"type:vachar(100);not null" json:"productID"`
	SerialNumber  string         `gorm:"type:varchar(100);not null" json:"serialNumber"`
	IdentifierNum string         `gorm:"type:varchar(100);not null" json:"identifierNum"`
	Price         int            `gorm:"type:int;not null" json:"price"`
	Status        string         `gorm:"type:varchar(20);not null" json:"status"`
}
