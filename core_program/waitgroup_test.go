package core_program

import (
	"fmt"
	"net/http"
	"sync"
	"testing"
)

var wg sync.WaitGroup

var urls = []string{
	"https://www.golang.net",
	"https://www.qq.com",
	"https://www.baidu.com",
}

func TestWaitGroup(t *testing.T) {
	for _, url := range urls {
		wg.Add(1)

		go func(url string) {
			defer wg.Done()
			resp, err := http.Get(url)
			if err == nil {
				fmt.Println(resp.StatusCode)
			}
		}(url)
	}
	// 等待所有的请求结束
	wg.Wait()
	fmt.Println("Success")
}