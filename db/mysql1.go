package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	conn, err := sql.Open("mysql", "root:@tcp(localhost:3306)/test?charset=utf8")
	if err != nil {
		fmt.Println("数据库连接失败")
		return
	}
	defer conn.Close()
	query, err := conn.Query("select * from foo limit 10")
	if err != nil {
		fmt.Println("数据库查询失败", err.Error())
		return
	}
	defer query.Close()
	cols, _ := query.Columns()
	for i := range cols {
		fmt.Print(cols[i])
		fmt.Print("\t")
	}
	fmt.Println("")
	fmt.Println("=================================")
	values := make([]sql.RawBytes, len(cols))
	scans := make([]interface{}, len(cols))
	for i := range values {
		scans[i] = &values[i]
	}
	results := make(map[int]map[string]string)
	i := 0
	for query.Next() {
		if err := query.Scan(scans...); err != nil {
			fmt.Println("Error")
			return
		}
		row := make(map[string]string)
		for j, v := range values {
			key := cols[j]
			row[key] = string(v)
		}
		results[i] = row
		i++
	}
	// 打印结果
	for i, m := range results {
		fmt.Println(i)
		for k, v := range m {
			fmt.Println(k, " : ", v)
		}
		fmt.Println("========================")
	}
}