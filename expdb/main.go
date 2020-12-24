package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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
	dbname   = "database_name" //database name here
)

//Create a GORM model
type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"not null;unique_index"`
	Color string
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
	db.AutoMigrate(&User{})

	name, email, color := getInfo()

	u := User{
		Name:  name,
		Email: email,
		Color: color,
	}

	if err := db.Create(&u).Error; err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", u)

	//use gorm to drop the users table
	// db.DropTableIfExists(&User{})
}

func getInfo() (name, email, color string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("What is your name?")
	name, _ = reader.ReadString('\n')

	fmt.Println("What is your email?")
	email, _ = reader.ReadString('\n')

	fmt.Println("What is your favourite color?")
	color, _ = reader.ReadString('\n')

	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)
	color = strings.TrimSpace(color)

	return name, email, color
}
