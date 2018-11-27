package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	strSql := "select iduser, login, password from sysuser"
	dbconn := "user=remote dbname=lean_dev host=vps.pilumsoft.ovh password=remote"

	var prepared *sql.Stmt

	db, err := sql.Open("postgres", dbconn)
	if err != nil {
		fmt.Printf("sql.Open failure - err: %v", err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("db.Ping failure - err: %v", err)
	}
	prepared, err = db.Prepare(strSql)
	if err != nil {
		fmt.Printf("db.Prepare failure - err: %v", err)
	}
	defer prepared.Close()

	rows, err := prepared.Query()
	defer rows.Close()

	if err != nil {
		fmt.Printf("prepared.Query() failed - err: %v", err)
	}

	for rows.Next() {
		var login string
		var password string
		if err := rows.Scan(&login, &password); err != nil {
			fmt.Printf("rows.Scan failed - err: %v", err)
		} else {
			//var login string
			//name = brand_name
			fmt.Printf("\tname = %s\n", login)
		}
	}
	if err := rows.Err(); err != nil {
		fmt.Printf("rows.Err is not nil - err: %v", err)
	}
}
