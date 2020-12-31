package main

import (
	"fmt"

	// using gorm instead to access Go's database/sql package
	// gorm is an Object Relation Mapper or a code library that automates
	// the transfer of data in a RDBMS
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"lenslocked.com/models"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "<username>" // username here uid -UN
	password = ""
	dbname   = "<database_name>" //database name here
)

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

	usr, err := models.NewUserService(psqlinfo)

	if err != nil {
		panic(err)
	}

	defer usr.Close()

	//use this to clear the table's records
	// usr.ResetDB()

	// use the following to create new records
	// u1 := models.User{
	// 	Name:  "King Long",
	// 	Email: "kl@gmail.com",
	// }
	// e1 := usr.CreateRecord(&u1)

	// if e1 != nil {
	// 	panic(err)
	// }

	userByID, err := usr.ByID(2)

	if err != nil {
		panic(err)
	}

	fmt.Println(userByID)

}
