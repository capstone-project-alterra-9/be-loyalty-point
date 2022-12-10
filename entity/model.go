package entity

import (
	"time"

	"gorm.io/gorm"
)

var DB *gorm.DB

type Points struct {
	ID         string `gorm:"primaryKey"`
	UserID     string `gorm:"type:varchar(100);not null" json:"userID"`
	Points     int    `gorm:"type:int" json:"points"`
	CostPoints int    `gorm:"type:int" json:"costPoints"`
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
	Redeem      bool           `gorm:"type:boolean;not null" json:"redeem"`
	Buy         bool           `gorm:"type:boolean;not null" json:"buy"`
	Name        string         `gorm:"type:varchar(100);not null" json:"name"`
	Description string         `gorm:"type:varchar(200);not null" json:"description"`
	Price       int            `gorm:"type:int;not null" json:"price"`
	Stock       int            `gorm:"type:int" json:"stock"`
	Image       string         `gorm:"type:varchar(200);not null" json:"image"`
}

type SerialNumbers struct {
	ID        string         `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"type:timestamp;not null" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"type:timestamp;not null" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	ProductID string         `gorm:"type:varchar(100);not null" json:"productID"`
	Serial    int64          `gorm:"type:bigint;not null" json:"serial"`
	Status    string         `gorm:"type:varchar(20);not null" json:"status"`
}

type Transactions struct {
	ID            string         `gorm:"primaryKey"`
	CreatedAt     time.Time      `gorm:"type:timestamp;not null" json:"createdAt"`
	UpdatedAt     time.Time      `gorm:"type:timestamp;not null" json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	PaymentMethod string         `gorm:"type:varchar(20);not null" json:"paymentMethod"`
	UserID        string         `gorm:"type:varchar(100);not null" json:"userID"`
	ProductID     string         `gorm:"type:varchar(100);not null" json:"productID"`
	SerialNumber  int64          `gorm:"type:bigint;not null" json:"serialNumber"`
	IdentifierNum string         `gorm:"type:varchar(100);not null" json:"identifierNum"`
	Price         int            `gorm:"type:int;not null" json:"price"`
	Status        string         `gorm:"type:varchar(20);not null" json:"status"`
}

type Pagination struct {
	Limit      int         `json:"limit,omitempty;query:limit"`
	Page       int         `json:"page,omitempty;query:page"`
	Sort       string      `json:"sort,omitempty;query:sort"`
	TotalRows  int64       `json:"total_rows"`
	TotalPages int         `json:"total_pages"`
	Rows       interface{} `json:"rows"`
}
