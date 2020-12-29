package main

import (
	"fmt"

	// using gorm instead to access Go's database/sql package
	// gorm is an Object Relation Mapper or a code library that automates
	// the transfer of data in a RDBMS
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "<username>" // username here uid -UN
	password = ""
	dbname   = "<database_name>" //database name here
)

// Create a GORM model
// update: add a slice to insert Orders added by the user
type User struct {
	gorm.Model
	Name   string
	Email  string `gorm:"not null;unique_index"`
	Color  string
	Orders []Order
}

// update: unsigned int bcos we don't have -ve values to the id
type Order struct {
	gorm.Model
	UserID      uint
	Amount      int
	Description string
}

func main() {
	// Sprintf formats the value as a string and stores it in the variable psqlinfo
	// Need to differentiate if a password is set or not; otherwise an error would be thrown
	var psqlinfo string

	if password != "" {
		psqlinfo = fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	} else {
		psqlinfo = fmt.Sprintf(
			"host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)
	}

	//open a connection to a DB, postgres in this case using GORM
	db, err := gorm.Open("postgres", psqlinfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	if err := db.DB().Ping(); err != nil {
		panic(err)
	}

	//Run the log mode to see the queries run by Gorm's AutoMigrate
	db.LogMode(true)

	//use Automigrate to gererate the Usertable based on the struct Model created above
	db.AutoMigrate(&User{}, &Order{})

	//use gorm to drop the users table
	// db.DropTableIfExists(&User{})

	// 1. Insert New Orders for user with id 3
	// var u1 User
	// if err := db.Where("id=?", 3).First(&u1).Error; err != nil {
	// 	panic(err)
	// }

	// createOrder(db, u1, 1001, "Desc #1")
	// createOrder(db, u1, 9999, "Desc #2")
	// createOrder(db, u1, 100, "Desc #3")

	//2. Preload the orders for the user with the id of 3
	// var u1 User
	// if err := db.Preload("Orders").Where("id=?", 3).First(&u1).Error; err != nil {
	// 	panic(err)
	// }
	// fmt.Println(u1)
	// fmt.Println(u1.Orders)

	//3. Preload the orders for all users
	var users []User
	if err := db.Preload("Orders").Find(&users).Error; err != nil {
		panic(err)
	}
	fmt.Println(users)

}

func createOrder(db *gorm.DB, user User, amount int, desc string) {
	err := db.Create(&Order{
		UserID:      user.ID,
		Amount:      amount,
		Description: desc,
	}).Error

	if err != nil {
		panic(err)
	}
}
