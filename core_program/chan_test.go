package core_program

import (
	"fmt"
	"runtime"
	"testing"
)

// chan 可以使用无缓冲的通道实现goroutine之间的同步等待
func TestChan(t *testing.T) {
	c := make(chan struct{})
	go func(i chan struct{}) {
		sum := 0
		for i := 0; i <= 10000; i++ {
			sum += i
		}
		fmt.Println(sum)
		// 写通道
		c <- struct{}{}
	}(c)

	// NumGoroutine 返回当前程序的goroutine数目
	fmt.Println("NumGoroutine=", runtime.NumGoroutine())
	// 读取通道，通过通道进行同步等待
	<-c
}

// 有缓冲的通道，写入的数据不会消失，可以缓冲和适配两个goroutine处理速率不一致的情况
func TestBufferChan(t *testing.T) {
	c := make(chan struct{})
	ci := make(chan int, 100)
	go func(i chan struct{}, j chan int) {
		for i := 0; i < 10; i++ {
			ci <- i
		}
		close(ci)
		// 写通道
		c <- struct{}{}
	}(c,ci)

	// NumGoroutine 返回当前程序的goroutine数目
	fmt.Println("NumGoroutine=", runtime.NumGoroutine())
	// 读取通道，通过通道进行同步等待
	<-c

	// ci通道已经关闭，匿名函数的func 已经关闭
	fmt.Println("NumGoroutine=", runtime.NumGoroutine())
	for v := range ci {
		fmt.Println(v)
	}
}

