package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//type Person struct {
	//	Name   string
	//	Age    int
	//	Gender bool
	//}
	////unmarshal to struct
	//var p Person
	var str = `{"Name":"junbin", "Age":21, "Gender":true}`
	//json.Unmarshal([]byte(str), &p)
	////result --> junbin : 21 : true
	//fmt.Println(p.Name, ":", p.Age, ":", p.Gender)
	//
	//// unmarshal to slice-struct
	//var ps []Person
	//var aJson = `[{"Name":"junbin", "Age":21, "Gender":true},
	//			{"Name":"junbin", "Age":21, "Gender":true}]`
	//json.Unmarshal([]byte(aJson), &ps)
	////result --> [{junbin 21 true} {junbin 21 true}] len is 2
	//fmt.Println(ps, "len is", len(ps))

	// unmarshal to map[string]interface{}
	var obj interface{} // var obj map[string]interface{}
	json.Unmarshal([]byte(str), &obj)
	m := obj.(map[string]interface{})
	//result --> junbin : 21 : true
	fmt.Println(m["Name"], ":", m["Age"], ":", m["Gender"])
	//
	////unmarshal to slice
	//var strs string = `["Go", "Java", "C", "Php"]`
	//var aStr []string
	//json.Unmarshal([]byte(strs), &aStr)
	////result --> [Go Java C Php]  len is 4
	//fmt.Println(aStr, " len is", len(aStr))
}