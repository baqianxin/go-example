package base

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func Test_Q1(t *testing.T) {
	var a uint = 1
	var b uint = 2
	fmt.Println(a - b)
}

// Result:
// 因为 uint为无符号int类型，所以结果肯定不是 -1
// 根据当前系统环境来看 2^32 - 1  || 2^64 - 1
// 原理 a-b = a + (-b)
// === RUN   Test_Q1
// 18446744073709551615
// --- PASS: Test_Q1 (0.00s)
// PASS
// ok      algo/base       0.151s

func cat(a, b chan int, wg *sync.WaitGroup) {
	for {
		select {
		case n := <-a:
			wg.Add(1)
			fmt.Println("cat  - ", n)
			b <- n
			wg.Wait()
		case <-time.After(10 * time.Second):
			return
		}
	}
}
func dog(b, c chan int) {
	for {
		select {
		case n := <-b:
			fmt.Println("dog  - ", n)
			c <- n
		case <-time.After(10 * time.Second):
			close(b)
			return
		}
	}
}

func fish(a, c chan int, wg *sync.WaitGroup) {
	for {
		select {
		case n := <-c:
			fmt.Println("fish - ", n)
			wg.Done()
		case <-time.After(10 * time.Second):
			close(c)
			close(a)
			return
		}
	}
}

func Test_queueMult(t *testing.T) {
	// quitCh := make(chan struct{}) 控对象不占用内存
	var a, b, c chan int = make(chan int), make(chan int), make(chan int)
	wg := &sync.WaitGroup{}
	go cat(a, b, wg)
	go dog(b, c)
	go fish(a, c, wg)
	for i := 0; i < 100; i++ {
		a <- i
	}
	time.Sleep(11 * time.Second)
	wg.Wait()
}
