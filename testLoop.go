package main

import (
	"runtime"
	"fmt"
)

func doSomething(num int) (sum int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d + %d = %d\n", num, num + i, num + num + i)
		sum = sum + num + i
	}
	return sum
}
func testLoop() {
	// 建立计数器，通道大小为cpu核数
	var NumCPU = runtime.NumCPU()
	fmt.Printf("NumCPU = %d\n", NumCPU)
	sem :=make(chan int, NumCPU);
	//FOR循环体
	data := []int{1, 11, 21, 31, 41, 51, 61, 71, 81, 91}
	for _,v:= range data {
		//建立协程
		go func (v int) {
			fmt.Printf("doSomething(%d)...\n", v)
			sum := doSomething(v);
			//计数
			sem <- sum;
		} (v);
	}
	// 等待循环结束
	var total int = 0
	for i := 0; i < len(data); i++ {
		temp := <- sem
		fmt.Printf("%d <- sem\n", temp)
		total = total + temp
	}
	fmt.Printf("total = %d\n", total)
}
func main() {
	testLoop()
}
