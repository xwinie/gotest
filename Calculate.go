package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

type Sum []int

func (s Sum) Calculate(count, start, end int, flag string, ch chan int) {
	cal := 0

	for i := start; i <= end; i++ {
		//for j := 1; j <= 3000000; j++ {
		//}
		cal += i
	}

	s[count] = cal
	fmt.Println("flag :", flag, ".")
	fmt.Println("cal :", cal, ".")
	ch <- count
}

func (s Sum) LetsGo() {
	// runtime.NumCPU()可以获取CPU核数，我的环境为4核，所以这里就简单起见直接设为4了
	const NCPU = 4
	const RANGE = 10000000
	var ch = make(chan int)

	runtime.GOMAXPROCS(NCPU)
	for i := 0; i < NCPU; i++ {
		go s.Calculate(i, (RANGE / NCPU) * i + 1, (RANGE / NCPU) * (i + 1), strconv.Itoa(i + 1), ch)
	}

	for i := 0; i < NCPU; i++ {
		<-ch
	}
}

func main() {
	var s Sum = make([]int, 4, 4)
	var sum int = 0
	var startTime = time.Now()

	s.LetsGo()

	//for _, v := range s {
	//	sum += v
	//}

	fmt.Println("总数为：", sum, "；所花时间为：",
		(time.Now().Sub(startTime)), "秒。")
}