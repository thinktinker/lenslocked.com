package main

import (
	"database/sql"
	"fmt"

	// pq is a pure Go postgres driver for Go's database/sql package
	_ "github.com/lib/pq" //You need to add the underscore, since it is not used but needed
)

const (
	host     = "localhost"
	port     = 5432
	user     = "<username here>" // username here uid -UN
	password = ""
	dbname   = "<database name here>" //database name here
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

	//open a connection to a DB, postgres in this case
	db, err := sql.Open("postgres", psqlinfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// 1. Ping the database to test the connection
	// ***************************************
	// err = db.Ping()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("successfully connected")
	// ***************************************

	// 2. Insert a sample name and email to the database
	//    db.Exec does not return an id to verify the inserted record
	//    Using back ticks allows you to split the query into two lines
	//    NOTE: Do not use single quotes ('') in a double qoute as that will create vulnerabilities, sql injection
	// ***************************************
	// _, err = db.Exec(`
	// INSERT INTO users(name, email)
	// VALUES($1, $2)`,	"John Doe", "johndoe@gmail.com")
	// if err != nil {
	// 	panic(err)
	// }
	// ***************************************

	// 3. Insert a sample name and email to the database
	//    db.QueryRow returns a value
	// ***************************************
	// var id int
	// row := db.QueryRow(`INSERT INTO users(name, email) VALUES($1, $2) RETURNING id`, "Tesa Law", "tesalaw@gmail.com")

	// err = row.Scan(&id)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("The id is %d\n", id)
	// ***************************************

	// 4. Query a single record from the database
	//    db.QueryRow returns a value
	// ***************************************
	// var id int
	// var name, email string

	// row := db.QueryRow(`SELECT id, name, email FROM users WHERE id=$1`, 3) //try valid and invalid id

	// err = row.Scan(&id, &name, &email)
	// if err == sql.ErrNoRows {
	// 	fmt.Println("No rows found")
	// } else {
	// 	panic(err)
	// }

	// fmt.Printf("The id is %d, name is %s and email is %s\n", id, name, email)
	// ***************************************

	// 5. Query a mulitple record from the database
	//    db.Query returns a value
	// ***************************************

	// type User struct {
	// 	id    int
	// 	name  string
	// 	email string
	// }

	// var users []User

	// rows, err := db.Query(`SELECT * FROM users`) //try valid and invalid id

	// if err != nil {
	// 	panic(err)
	// }

	// defer rows.Close()

	// for rows.Next() {

	// 	var user User
	// 	err = rows.Scan(&user.id, &user.name, &user.email)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	users = append(users, user)
	// }

	// if rows.Err() != nil {
	// 	fmt.Println("No data found!")
	// }

	// fmt.Printf("The results returned is %v\n", users)

	// ***************************************

	// 6. Create orders using a for loop the database
	//    To be used for working on relational data
	// ***************************************

	// for i := 0; i < 6; i++ {
	// 	userId := 1
	// 	if i > 3 {
	// 		userId = 3
	// 	}
	// 	amount := i * 5
	// 	description := fmt.Sprintf("USB Adapter x %d", i)

	// 	_, err := db.Exec(`
	// 		INSERT INTO orders (user_id, amount, description)
	// 		VALUES($1, $2, $3)
	// 	`, userId, amount, description)

	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

	// 7. Use INNER JOIN to query data between related tables
	// ***************************************

	rows, err := db.Query(`SELECT * FROM users INNER JOIN orders ON users.id = orders.user_id`)

	if err != nil {
		panic(err)
	}

	type Order struct {
		userID      int
		name        string
		email       string
		orderID     int
		orderUserID int
		qty         int
		description string
	}

	var orders []Order

	for rows.Next() {
		var order Order
		err = rows.Scan(&order.userID, &order.name, &order.email, &order.orderID, &order.orderUserID, &order.qty, &order.description)
		if err != nil {
			panic(err)
		}

		orders = append(orders, order)

		if rows.Err() != nil {
			fmt.Printf("No results retrieved.")
		}

	}

	fmt.Println(orders)
}
