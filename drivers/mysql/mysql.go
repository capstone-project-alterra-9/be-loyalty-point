package mysql_driver

import (

	// "mini-project/util"
	// "errors"

	"capstone-project/drivers/mysql/assignment"
	"capstone-project/drivers/mysql/module"
	"capstone-project/drivers/mysql/student"
	"capstone-project/drivers/mysql/teacher"
	"fmt"
	"log"

	// "golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ConfigDB struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_NAME     string
	DB_HOST     string
	DB_PORT     string
}

func (config *ConfigDB) InitDB() *gorm.DB {
	var err error

	var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB_USERNAME,
		config.DB_PASSWORD,
		config.DB_HOST,
		config.DB_PORT,
		config.DB_NAME,
	)

	// DATABASE_URL='user:pass@tcp(us-cdbr-iron-east-03.cleardb.net:3306)/your_heroku_database'

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("error when connecting to the database: %s", err)
	}

	log.Println("connected to the database")

	return db
}

func DBMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&student.Student{},
		&module.Module{},
		&teacher.Teacher{},
		&assignment.Assignment{},
	)
}

func CloseDB(db *gorm.DB) error {
	database, err := db.DB()

	if err != nil {
		log.Printf("error when getting the database instance: %v", err)
		return err
	}

	if err := database.Close(); err != nil {
		log.Printf("error when closing the database connection: %v", err)
		return err
	}

	log.Println("database connection is closed")

	return nil
}

// func SeedUser(db *gorm.DB) teachers.User {
// 	password, _ := bcrypt.GenerateFromPassword([]byte("123123"), bcrypt.DefaultCost)

// 	fakeUser, _ := util.CreateFaker[teachers.Teacher]()

// 	teacherRecord := teachers.Teacher{
// 		Email:    fakeUser.Email,
// 		Password: string(password),
// 	}

// 	if err := db.Create(&teacherRecord).Error; err != nil {
// 		panic(err)
// 	}

// 	var foundTeacher teachers.Teacher

// 	db.Last(&foundTeacher)

// 	foundTeacher.Password = "123123"

// 	return foundTeacher
// }

// func CleanSeeders(db *gorm.DB) {
// 	db.Exec("SET FOREIGN_KEY_CHECKS = 0")

// 	userResult := db.Exec("DELETE FROM users")

// 	var isFailed bool = userResult.Error != nil

// 	if isFailed {
// 		panic(errors.New("error when cleaning up seeders"))
// 	}

// 	log.Println("Seeders are cleaned up successfully")
// }
