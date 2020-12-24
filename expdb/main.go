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

//Create a GORM model
type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"not null;unique_index"`
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

	//use Automigrate to gererate the Usertable based on the struct Model created above
	// db.AutoMigrate(&User{})

	//use gorm to drop the users table
	db.DropTableIfExists(&User{})
}
