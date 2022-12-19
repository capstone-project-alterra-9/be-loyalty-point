package config

import (
	"capstone-project/entity"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName,
	)
	var err error
	entity.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return entity.DB
}

func InitDatabaseTest() *gorm.DB {
	dbHostTest := os.Getenv("DB_HOST_TEST")
	dbUserTest := os.Getenv("DB_USER_TEST")
	dbPassTest := os.Getenv("DB_PASS_TEST")
	dbNameTest := os.Getenv("DB_NAME_TEST")
	dbPortTest := os.Getenv("DB_PORT_TEST")
	dsnTest := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		dbHostTest,
		dbUserTest,
		dbPassTest,
		dbNameTest,
		dbPortTest,
	)
	var err error
	entity.DB, err = gorm.Open(postgres.Open(dsnTest), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	entity.DB.Migrator().DropTable(
		&entity.Points{}, &entity.Users{}, &entity.Products{},
		&entity.SerialNumbers{}, &entity.Transactions{},
	)
	entity.DB.AutoMigrate(
		&entity.Points{}, &entity.Users{}, &entity.Products{},
		&entity.SerialNumbers{}, &entity.Transactions{},
	)

	return entity.DB
}
