package main

import "fmt"

type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}
func (a *Integer) Add(b Integer) {
	*a += b
}

type LessAdder interface {
	Less(b Integer) bool
	Add(b Integer)
}

type Lesser interface {
	Less(b Integer) bool
}

func main() {
	var a Integer = 1
	a.Add(2)
	fmt.Println("a =", a)
}


type ReadWriter interface {
	Read(buf []byte) (n int, err error)
	Write(buf []byte) (n int, err error)
}

type IStream interface {
	Write(buf []byte) (n int, err error)
	Read(buf []byte) (n int, err error)
}