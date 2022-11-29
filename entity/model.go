package entity

import (
	"time"

	"gorm.io/gorm"
)

var DB *gorm.DB

type Points struct {
	ID     string `gorm:"primaryKey"`
	UserID string `gorm:"type:varchar(100);not null" json:"user_id"`
	Points int    `gorm:"not null" json:"points"`
}

type Users struct {
	ID        string         `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"type:timestamp;not null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"type:timestamp;not null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Role      string         `gorm:"type:varchar(10);not null" json:"role"`
	Username  string         `gorm:"type:varchar(20);not null;unique" json:"username"`
	Email     string         `gorm:"type:varchar(100);not null;unique" json:"email"`
	Password  string         `gorm:"type:varchar(100);not null" json:"password"`
}

type Products struct {
	ID          string         `gorm:"primaryKey"`
	CreatedAt   time.Time      `gorm:"type:timestamp;not null" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"type:timestamp;not null" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Category    string         `gorm:"type:varchar(20);not null" json:"category"`
	Name        string         `gorm:"type:varchar(100);not null" json:"name"`
	Description string         `gorm:"type:varchar(200);not null" json:"description"`
	Price       int            `gorm:"type:int;not null" json:"price"`
	Stock       int            `gorm:"type:int;not null" json:"stock"`
	Image       string         `gorm:"type:varchar(200);not null" json:"image"`
}

type SerialNumbers struct {
	ID        string         `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"type:timestamp;not null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"type:timestamp;not null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	ProductID string         `gorm:"type:varchar(100);not null" json:"product_id"`
	Serial    string         `gorm:"type:varchar(100);not null" json:"serial"`
	Status    string         `gorm:"type:varchar(20);not null" json:"status"`
}

type Transactions struct {
	ID            string         `gorm:"primaryKey"`
	CreatedAt     time.Time      `gorm:"type:timestamp;not null" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"type:timestamp;not null" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	PaymentMethod string         `gorm:"type:varchar(20);not null" json:"payment_method"`
	UserID        string         `gorm:"type:varchar(100);not null" json:"user_id"`
	ProductID     string         `gorm:"type:varchar(100);not null" json:"product_id"`
	SerialNumber  string         `gorm:"type:varchar(100);not null" json:"serial_number"`
	IdentifierNum string         `gorm:"type:varchar(100);not null" json:"identifier_num"`
	Price         int            `gorm:"type:int;not null" json:"price"`
	Status        string         `gorm:"type:varchar(20);not null" json:"status"`
}
