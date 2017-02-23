////panic场景1
////recover场景2
package main
import "fmt"
func  main() {
	defer func(){
		fmt.Println("hello,defer go")
	}()
	panic(11111)
}