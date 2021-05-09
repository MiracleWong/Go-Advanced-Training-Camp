package core_program

import (
	"fmt"
	"testing"
)

func TestSelect(t *testing.T) {
	ch := make(chan int, 1)
	go func(chan int) {
		for  {
			// select 关键字，用于多路监听多个通道
			select {
			// 0 or 1 写入是随机的
			case ch <- 0:
			case ch <- 1:
			}
		}
	}(ch)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}
