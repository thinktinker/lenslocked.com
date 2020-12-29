package main

import (
	"fmt"
	"os"

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
	fmt.Printf("The model: %+v", users)

	// 6. Error Handling in Gorm

	// 6.1
	// The first example chains the query together using DB1
	// This refers to the same object newBD1 and will therefore be able to catch 'no records' error
	var usr1 User
	newDB1 := db.Where("email=?", "blahblah1@gmail.com")
	newDB1 = newDB1.Or("color=?", "red")
	newDB1 = newDB1.First(&usr1)

	if newDB1.Error != nil {
		fmt.Println("There was an error")
		panic(newDB1.Error)
	}

	fmt.Println(usr1)

	// The second example here doesn't chain the database
	// As both statements refer to different objects, any errors won't be caught
	var usr1_1 User
	db.Where("email=? Or color=?", "blahblah@gmail_1.com", "red")
	db.First(&usr1_1)

	if db.Error != nil { // this statement here won't be able to catch the error
		panic(db.Error)
	}

	fmt.Println(usr1_1)

	// 6.2
	// To trap errors using Gorm's GetErrors
	var usr2 User
	newDB2 := db.Where("email=?", "blahblah2@gmail.com")
	newDB2 = newDB2.Or("color=?", "red")
	newDB2 = newDB2.First(&usr2)

	errs := newDB2.GetErrors() //this checks for multiple erros

	if len(errs) > 0 { // if the len of errs is greater than zeor
		fmt.Println(errs) // print the errors
		os.Exit(1)        // and exit the program
	}

	fmt.Println(usr2)

	// 6.3
	// Trapping a RecordNotFound error
	var usr3 User
	newDB3 := db.Where("email=?", "blahblah3@gmail.com")
	newDB3 = newDB3.Or("color=?", "red")
	newDB3 = newDB3.First(&usr3)

	if newDB3.RecordNotFound() {
		fmt.Println("No user found.")
	} else if newDB3.Error != nil {
		panic(newDB3.Error)
	} else {
		fmt.Println(usr3)
	}

	// 6.4
	// Trapping a error using switch statements
	var usr4 User
	newDB4 := db.Where("email=?", "blahblah4@gmail.com")
	newDB4 = newDB4.Where("color=?", "red")
	newDB4 = newDB4.First(&usr4)

	if err := newDB4.Error; err != nil {
		// list of errors found at: https://github.com/jinzhu/gorm/blob/master/errors.go#L25
		// ErrRecordNotFound
		// ErrInvalidSQL
		// ErrInvalidTransaction
		// ErrCantStartTransaction
		// ErrUnaddressable
		switch err {
		case gorm.ErrRecordNotFound:
			fmt.Println("No user found.")
		case gorm.ErrInvalidSQL:
			fmt.Println("Invalid query statement.")
		default:
			panic(err)
		}
	}

	fmt.Println(usr4)

}
