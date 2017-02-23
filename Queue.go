package main
import (
	"fmt"
)

type Queue struct {
	pBase *[6]int
	pFront,pRear int
}

func initQueue(q *Queue) {
	var arr=new([6]int)
	q.pBase=arr
	q.pFront=0
	q.pRear=0
}

func isEmpty(q *Queue) bool {
	if q.pFront==q.pRear {
		return true
	}else {
		return false
	}
}

func isFull(q *Queue) bool {
	if (q.pRear+1)%6==q.pFront {
		return true
	}else {
		return false
	}
}

func enQueue(q *Queue,val int) bool {
	if isFull(q) {
		return false
	}else {
		q.pBase[q.pRear]=val
		q.pRear=(q.pRear+1)%6
		return true
	}
}

func deQueue(q *Queue) bool {
	if isEmpty(q) {
		return false
	}else {
		q.pFront=(q.pFront+1)%6
		return true
	}
}

func traverse(q *Queue) {
	if isEmpty(q) {
		return
	}
	for i:=q.pFront;i%6!=q.pRear;i++ {
		fmt.Print(q.pBase[i]," ")
	}
	fmt.Println()
}

func main() {
	var q=new(Queue)
	initQueue(q)
	enQueue(q,1)
	enQueue(q,2)
	enQueue(q,3)
	enQueue(q,4)
	enQueue(q,5)
	enQueue(q,6)
	enQueue(q,7)
	traverse(q)
	deQueue(q)
	deQueue(q)
	traverse(q)
	enQueue(q,15)
	traverse(q)
}