package share_mem

import (
	"sync"
	"testing"
	"time"
)

func TestCounter(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

// 线程安全
func TestCounterThreadSafe(t *testing.T) {
	var mux  sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {
				mux.Unlock()
			}()
			mux.Lock()
			counter++
		}()
	}
	time.Sleep(1 * time.Second)  // OK // 等待所有的协程执行完毕
	//time.Sleep(10 * time.Millisecond) // OK
	//time.Sleep(1 * time.Microsecond)  // 4999
	t.Logf("counter = %d", counter)
}

// WaitGrooup
func TestCounterWaitGroup(t *testing.T) {
	var mux sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mux.Unlock()
			}()
			mux.Lock()
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	t.Logf("counter = %d", counter)
}
