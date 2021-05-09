package main

import (
	"fmt"
	"math/rand"
	"runtime"
)

// 随机数生成的函数
func GenerateIntA(done chan struct{}) chan int{
	ch := make(chan int)
	go func() {
		Label:
			for {
				select {
				case ch<-rand.Int():
				case <-done:
					break Label
				}
			}
		// 手动通知关闭ch
		close(ch)
	}()
	return ch
}

func main() {
	done := make(chan struct{})
	ch := GenerateIntA(done)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// 发送通知，告诉生产者停止生产
	close(done)

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// NumGoroutine 返回当前程序的goroutine数目
	fmt.Println("NumGoroutine=", runtime.NumGoroutine())
}
