package main



import (
	"fmt"
	"database/sql"
	_"odbc/driver"
)

func main(){
	conn,err := sql.Open("odbc","DSN=drill");

	if(err!=nil){
		fmt.Println("Connecting Error");
		return;
	}
	defer conn.Close();
	stmt,err := conn.Prepare("select  *  from sys_dict");
	if(err!=nil){
		fmt.Println("Query Error",err);
		return;
	}
	defer stmt.Close();
	row,err := stmt.Query()
	if err!=nil {
		fmt.Println("Query Error",err);
		return;
	}
	defer row.Close();
	for row.Next() {
	  fmt.Println(row.Scan("id"));

	}
	fmt.Printf("%s\n","finish");
	return;
}