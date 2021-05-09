package main

import (
	"fmt"
	"sync"
)


type Counter struct {
	mu    sync.Mutex
	Count uint64
}

func main() {
	//// 互斥锁的保护计数器
	//var mu sync.Mutex
	//// 计数器的值
	//var count = 0
	var counter Counter
	// 辅助变量，用来确认所有的goroutine都完成
	var wg sync.WaitGroup
	wg.Add(10)

	// 启动10个gourontine
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			// 对变量count执行10次加1
			for j := 0; j < 100; j++ {
				//mu.Lock()
				//count++
				//mu.Unlock()
				counter.mu.Lock()
				counter.Count++
				counter.mu.Unlock()
			}
		}()
	}
	wg.Wait()
	//fmt.Println(count)
	fmt.Println(counter.Count)
}
