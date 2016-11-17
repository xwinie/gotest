package main

import (
	_"github.com/LukeMauldin/lodbc"
	"database/sql"
	"fmt"
)

func main() {
	db, err := sql.Open("lodbc",`Driver={MapR Drill ODBC Driver};
                        Catalog=DRILL; AuthenticationType=No Authentication; ConnectionType=Direct;
                        Host=192.168.10.241;
                        Port=31010;`)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlstr := "select name from mongo.mfm.sys_dict"
	stmt, err := db.Prepare(sqlstr)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	fmt.Println(rows)
	for rows.Next() {
		var name string
		if err := rows.Scan(&name);err==nil {
			fmt.Println(name);

		}

	}
}