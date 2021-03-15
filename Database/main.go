// https://www.bogotobogo.com/GoLang/GoLang_SQLite.php
package main

import (
    "database/sql"
    "fmt"
    _ "strconv"
    "os"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
	database, _ := 
		sql.Open("sqlite3", "./database.db")
        rows, _ := 
		database.Query("SELECT id, firstname, lastname FROM people")
    var id int
    var firstname string
    var lastname string
    for rows.Next() {
        rows.Scan(&id, &firstname, &lastname)
        fmt.Print(firstname);
        fmt.Print(" ");
        fmt.Print(lastname);
        fmt.Println("");
        if firstname == os.Args[1] {
            panic("Nome em uso!");
        } else {
	statement, _ := 
		database.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY AUTOINCREMENT, firstname TEXT, lastname TEXT)")
    statement.Exec()
	statement, _ = 
		database.Prepare("INSERT INTO people (firstname, lastname) VALUES (?, ?)")
    statement.Exec(os.Args[1], os.Args[2]);
        }
	
    
    }
}