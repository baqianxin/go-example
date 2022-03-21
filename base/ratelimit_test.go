package base

import (
	"fmt"
	"testing"
	"time"
)

// 通过chan+ time.Ticker 定时器来实现 1 / 200ms
// 通过缓冲队列来实现允许 并发（多）请求 + 限速
func Test_ratelimit(t *testing.T) {

	reqCh := make(chan int, 5) //模拟5个请求

	limiter := time.Tick(200 * time.Millisecond)
	for j := 1; j < 6; j++ {
		reqCh <- j
	}
	close(reqCh)
	for req := range reqCh {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// example2: 并发请求数量限制
	fmt.Printf("\n example2: 并发请求数量限制\n")
	burstyRequestCh := make(chan int, 5)
	burstyLimiterCh := make(chan time.Time, 3)
	for j := 0; j < 3; j++ {
		burstyLimiterCh <- time.Now()
	}
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiterCh <- t
		}
	}()
	for j := 1; j < 6; j++ {
		burstyRequestCh <- j
	}
	close(burstyRequestCh)
	for req := range burstyRequestCh {
		<-burstyLimiterCh
		fmt.Println("request", req, time.Now())
	}
}
