package entity

import "gorm.io/gorm"

var DB *gorm.DB

type Users struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null;unique" json:"username"`
	Email    string `gorm:"type:varchar(100);not null;unique" json:"email"`
	Password string `gorm:"type:varchar(100);not null" json:"password"`
	Points   int    `gorm:"type:int;not null" json:"points"`
}

type Products struct {
	gorm.Model
	Category    string `gorm:"type:varchar(20);not null" json:"category"`
	Name        string `gorm:"type:varchar(100);not null" json:"name"`
	Description string `gorm:"type:varchar(200);not null" json:"description"`
	Price       int    `gorm:"type:int;not null" json:"price"`
	Stock       int    `gorm:"type:int;not null" json:"stock"`
	Image       string `gorm:"type:varchar(200);not null" json:"image"`
	CreatedBy   string `gorm:"type:varchar(20);not null" json:"created_by"`
	UpdatedBy   string `gorm:"type:varchar(20);not null" json:"updated_by"`

type SerialNumbers struct {
	gorm.Model
	ProductID uint   `gorm:"type:int;not null" json:"product_id"`
	Serial    string `gorm:"type:varchar(100);not null" json:"serial"`
	Status    string `gorm:"type:varchar(20);not null" json:"status"`
	CreatedBy string `gorm:"type:varchar(20);not null" json:"created_by"`
}

type Transactions struct {
	gorm.Model
	UserID        uint   `gorm:"type:int;not null" json:"user_id"`
	ProductID     uint   `gorm:"type:int;not null" json:"product_id"`
	SerialNumber  string `gorm:"type:varchar(100);not null" json:"serial_number"`
	IdentifierNum string `gorm:"type:varchar(100);not null" json:"identifier_num"`
	TotalPrice    int    `gorm:"type:int;not null" json:"total_price"`
	Status        string `gorm:"type:varchar(20);not null" json:"status"`
	LastUpdatedBy string `gorm:"type:varchar(20);not null" json:"last_updated_by"`
}

type Admins struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100);not null" json:"name"`
	Username string `gorm:"type:varchar(20);not null;unique" json:"username"`
	Password string `gorm:"type:varchar(100);not null" json:"password"`
}
