package main

import (
	"fmt"
	"sync"
)


type Counter struct {
	sync.Mutex
	Count int
}


func main() {
	// 示例1
	//foo1()

	// 示例2
	//var c Counter
	//c.Lock()
	//defer c.Unlock()
	//c.Count++
	//foo1(c) // 复制锁
	l := &sync.Mutex{}
	foo3(l)
}

// 这里Counter的参数是通过复制的方式传入的
func foo2(c Counter) {
	c.Lock()
	defer c.Unlock()
	fmt.Println("in foo")
}

func foo1() {
	var mu sync.Mutex
	defer mu.Unlock()
	//mu.Lock()
	fmt.Println("hello world!")
}


func foo3(l sync.Locker) {
	fmt.Println("in foo")
	l.Lock()
	bar(l)
	l.Unlock()
}


func bar(l sync.Locker) {
	l.Lock()
	fmt.Println("in bar")
	l.Unlock()
}