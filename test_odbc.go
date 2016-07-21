package main

import (
	"fmt"
	"database/sql"
	_ "github.com/alexbrainman/odbc"
)


func main() {
	connStr := `Driver=drill;
                        Catalog=DRILL;
                        Schema=hivestg;
                        ConnectionType=Direct;
                        Host=192.168.10.239;
                        Port=31010;
                        AuthenticationType=No Authentication`

	db, err := sql.Open("odbc", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sql := "select PolicyID, LocationID, SomeValue from dfs.`/home/ctownsend/projects/drill-odbc/test_data/`"

	fmt.Println(sql)

	type rslt struct{
		PolID int
		LocID int
		SomeVal int
	}

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var r rslt
		err = rows.Scan(&r.PolID, &r.LocID, &r.SomeVal)
		if err != nil {
			panic(err)
		}
		fmt.Println(r)
	}


}