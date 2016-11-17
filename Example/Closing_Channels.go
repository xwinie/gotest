package main

import "fmt"
// 这个例子在执行workder的工作时使用jobs管道来监听
// 当jobs关闭了，表示worker执行完毕，开始退出goroutine
// goroutine的退出又使用done来监听
func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()
	// 发送三个jobs worker执行
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")// worker被关闭
	<-done
}// 输出
// sent job 1
// received job 1
// sent job 2
// received job 2
// sent job 3
// received job 3
// sent all jobs
// received all jobs