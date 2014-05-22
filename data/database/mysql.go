package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// Open mysql database:
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/DB_USER")
	checkErr(err)
	defer db.Close()
	
	// Check connection:
	err = db.Ping()
	checkErr(err)
	
	// Insert data:
	stmt, err := db.Prepare("INSERT INTO `DB_USER`.`TB_USER`(`NAME`,`EMAIL`)VALUES(?,?)")
	checkErr(err)
	res, err := stmt.Exec("Richard Lew", "i@RichardLew.com")
	checkErr(err)
	lastInsertId, err := res.LastInsertId()
	checkErr(err)
	fmt.Printf("Last insert Id: %d\n", lastInsertId)

	// Delete data:
	stmt, err = db.Prepare("DELETE FROM `DB_USER`.`TB_USER` WHERE NAME=?")
	checkErr(err)
	res, err = stmt.Exec("Richard Lew")
	checkErr(err)
	rowsAffectedAfterDelete, err := res.RowsAffected()
	checkErr(err)
	fmt.Printf("Rows affected after delete: %d\n", rowsAffectedAfterDelete)
	
	// Update data:
	stmt, err = db.Prepare("UPDATE `DB_USER`.`TB_USER` SET EMAIL=? WHERE NAME=?")
	checkErr(err)
	res, err = stmt.Exec("i@rincliu.com", "Rinc Liu")
	checkErr(err)
	rowsAffectedAfterUpdate, err := res.RowsAffected()
	checkErr(err)
	fmt.Printf("Rows affected after update: %d\n", rowsAffectedAfterUpdate)

	// Query data:
	rows, err := db.Query("SELECT * FROM `DB_USER`.`TB_USER` WHERE EMAIL LIKE '%_@%'")
	checkErr(err)
	fmt.Println("Query result:")
	for rows.Next() {
		var _id int
		var name string
		var email string
		err = rows.Scan(&_id, &name, &email)
		checkErr(err)
		fmt.Printf("name: %s, email: %s\n", name, email)
	}
}
