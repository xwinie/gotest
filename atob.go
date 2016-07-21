package main

import (
	"strconv"
	"fmt"
)
// ParseBool 将字符串转换为布尔值
// 它接受真值：1, t, T, TRUE, true, True
// 它接受假值：0, f, F, FALSE, false, False.
// 其它任何值都返回一个错误
func ParseBool(str string) (value bool, err error)

func main() {
	fmt.Println(strconv.ParseBool("1"))    // true
	fmt.Println(strconv.ParseBool("t"))    // true
	fmt.Println(strconv.ParseBool("T"))    // true
	fmt.Println(strconv.ParseBool("true")) // true
	fmt.Println(strconv.ParseBool("True")) // true
	fmt.Println(strconv.ParseBool("TRUE")) // true
	fmt.Println(strconv.ParseBool("TRue"))
	// false strconv.ParseBool: parsing "TRue": invalid syntax
	fmt.Println(strconv.ParseBool("0"))     // false
	fmt.Println(strconv.ParseBool("f"))     // false
	fmt.Println(strconv.ParseBool("F"))     // false
	fmt.Println(strconv.ParseBool("false")) // false
	fmt.Println(strconv.ParseBool("False")) // false
	fmt.Println(strconv.ParseBool("FALSE")) // false
	fmt.Println(strconv.ParseBool("FALse"))
	fmt.Println(strconv.FormatBool(0 < 1)) // true
	fmt.Println(strconv.FormatBool(0 > 1)) // false
	// false strconv.ParseBool: parsing "FAlse": invalid syntax
}
