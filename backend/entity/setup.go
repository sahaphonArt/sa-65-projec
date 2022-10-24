package entity

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {

	return db

}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("sa-65.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	database.AutoMigrate(
		&Admin{},
		&Organization{},
		&TypeFund{},
		&Donator{},
	)

	db = database

	password, err := bcrypt.GenerateFromPassword([]byte("123"), 14)

	var Admins = []Admin{
		{Name: "Sam", Email: "123@gmail.com", Password: string(password)},
		{Name: "Dr", Email: "asdf@gmail.com", Password: string(password)},
	}
	db.CreateInBatches(Admins, 2)

	var TypeFunds = []TypeFund{
		{TypeFund: "aum"},
		{TypeFund: "sum"},
		{TypeFund: "bim"},
	}
	db.CreateInBatches(TypeFunds, 3)

	var Organizations = []Organization{
		{Organization: "a"},
		{Organization: "b"},
		{Organization: "c"},
	}
	db.CreateInBatches(Organizations, 3)

}
