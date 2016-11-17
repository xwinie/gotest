package main

import (
	"fmt"
	"database/sql"
	_ "github.com/alexbrainman/odbc"
)


func main() {
	connStr := `DSN=drill;`

	db, err := sql.Open("odbc", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sql := "select * from  mysqlplugin.mfm.sys_dict"

	fmt.Println(sql)


	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	for rows.Next() {

		err = rows.Scan("name")
		if err != nil {
			panic(err)
		}
		fmt.Println(rows.Scan("name"))
	}


}