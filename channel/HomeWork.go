package main

import "time"

const StudentNum = 30

type HomeWork struct {
}

func student(hwChan chan HomeWork) {
	//学生提交作业
	hwChan <- HomeWork{}
}
func teacher(hwChan chan HomeWork) {
	//老师取出30个学生的作业批改
	for i := 0; i < StudentNum; i++ {
		<-hwChan
	}
}
func main() {
	hwChan := make(chan HomeWork, 30)
	//每个学生开启一个goroutine，学生单独做作业，做完作业提交到hwChan中即可
	for i := 0; i < StudentNum; i++ {
		go student(hwChan)
	}
	//老师开启一个goroutine，去批改学生作业
	go teacher(hwChan)
	time.Sleep(5 * time.Second)
}