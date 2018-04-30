package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	name := "go_develop"
	db, err := sql.Open("mysql", "root:Lol384721596@/"+name)
	fmt.Println(err)
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// table_name := "number"
	// _, err = db.Exec("IF NOT EXISTS (SELECT [name] FROM sys.tables WHERE [name] = '" + table_name + "') CREATE TABLE " + table_name + " (number integer(11), squareNumber integer(11))")
	// if err != nil {
	// 	panic(err)
	// }

	_, err = db.Exec("USE " + name)
	if err != nil {
		panic(err)
	}

	// Prepare statement for inserting data
	// stmtIns, err := db.Prepare("INSERT INTO squareNum VALUES( ?, ? )") // ? = placeholder
	// if err != nil {
	// 	panic(err.Error()) // proper error handling instead of panic in your app
	// }
	// defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

	// // Prepare statement for reading data
	// stmtOut, err := db.Prepare("SELECT squareNumber FROM squarenum WHERE number = ?")
	// if err != nil {
	// 	panic(err.Error()) // proper error handling instead of panic in your app
	// }
	// defer stmtOut.Close()

	// // Insert square numbers for 0-24 in the database
	// for i := 0; i < 25; i++ {
	// 	_, err = stmtIns.Exec(i, (i * i)) // Insert tuples (i, i^2)
	// 	if err != nil {
	// 		panic(err.Error()) // proper error handling instead of panic in your app
	// 	}
	// }

	// var squareNum int // we "scan" the result in here

	// // Query the square-number of 13
	// err = stmtOut.QueryRow(13).Scan(&squareNum) // WHERE number = 13
	// if err != nil {
	// 	panic(err.Error()) // proper error handling instead of panic in your app
	// }
	// fmt.Printf("The square number of 13 is: %d", squareNum)

	// // Query another number.. 1 maybe?
	// err = stmtOut.QueryRow(1).Scan(&squareNum) // WHERE number = 1
	// if err != nil {
	// 	panic(err.Error()) // proper error handling instead of panic in your app
	// }
	// fmt.Printf("The square number of 1 is: %d", squareNum)
}
