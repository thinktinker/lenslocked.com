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

	//use gorm to drop the users table
	// db.DropTableIfExists(&User{})

	//1. gorm chaining example
	var u1 User
	db.Where("id=?", 3).First(&u1)
	//or db.First(&u) or db.Last(&u) or db.First(&u, 4) -- the last query assumes to be asking for id=4
	fmt.Println(u1)

	//2. gorm split the chain
	var u2 User
	newDB := db.Where("id= ? AND color = ?", 5, "red")
	newDB.First(&u2)
	fmt.Println(u2)

	//3. gorm 2nd chaining example
	var u3 User
	db.Where("id>?", 3).
		Where("color=?", "red").
		First(&u3)
	fmt.Println(u3)

	//4. gorm querying user using the model object
	// In this example, this variable is already a pointer to an address
	var u4 *User = &User{
		Color: "red",
		Email: "j_b@gmail.com",
	}
	db.Where(u4).First(u4)
	fmt.Println(u4)

	//5. Querying mulitple records
	var users []User
	db.Find(&users)
	fmt.Println(len(users))
	fmt.Printf("%+v", users)

}
