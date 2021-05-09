package core_program

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestAnonymousFunc(t *testing.T) {
	go func() {
		sum := 0
		for i := 0; i <= 10000; i++ {
			sum += i
		}
		fmt.Println(sum)
		time.Sleep(1 * time.Second)
	}()
	fmt.Println("NumGoroutine=", runtime.NumGoroutine())
	time.Sleep(5 * time.Second)
}
func sum() {
	sum := 0
	for i := 0; i <= 10000; i++ {
		sum += i
	}
	fmt.Println(sum)
	time.Sleep(1 * time.Second)
}

func TestFunc(t *testing.T) {

	go sum()
	// NumGoroutine 返回当前程序的goroutine数目
	fmt.Println("NumGoroutine=", runtime.NumGoroutine())
	time.Sleep(5 * time.Second)
	fmt.Println("GOMAXPROCS=", runtime.GOMAXPROCS(0))
	runtime.GOMAXPROCS(16)
	fmt.Println("GOMAXPROCS=", runtime.GOMAXPROCS(0))
}

