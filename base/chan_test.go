package base

import (
	"fmt"
	"testing"
	"time"
)

// 1. channel 用于协程间的通信，单一进程无需使用，所以不会存在
// 同一进程中的顺序写入读取操作
//  ch <- v    // 发送值v到Channel ch中
//  v := <-ch  // 从Channel ch中接收数据，并将数据赋值给v

// 2. channel 读写都是阻塞的
// 可以通过 <- 特性来等待执行
// 	go sum(s[:len(s)/2], c)
//  go sum(s[len(s)/2:], c)
//  x, y := <-c, <-c // receive from c

func Test_Fib(t *testing.T) {
	// test()
	// test2()
	noBlockChan()
	//
	block()
	//
	rangeoverchan()
	//
	timer()
	//
	timeTicker()
}

func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}

	// close(c)
}

// 简单使用
func test() {
	message := make(chan string)
	go func() { message <- "ping" }()
	msg := <-message
	fmt.Println(msg)
}

func test2() {
	c := make(chan int, 10)
	go func() {
		time.Sleep(3 * time.Second)
		c <- -1

		//如果 chan 未关闭，且无数据写入。range操作会处于阻塞状态，直至超时或数据写入
		close(c)
	}()
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

// 非阻塞channel使用问题
// select的 default子句替换为其他阻塞操作
func noBlockChan() {
	messgaes := make(chan string)
	signals := make(chan bool)
	go func() { messgaes <- "aaas" }()
	// messgaes <- "aaas" // 阻塞操作，需要有对应的消费端
	// a := <-messgaes
	// fmt.Println(a)
	select {
	case msg := <-messgaes:
		fmt.Println("received message:", msg)
	default: // 无缓冲 无阻塞 直接执行default
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case messgaes <- msg:
		fmt.Println("send message", msg)
	default: // 无缓冲 无阻塞 直接执行default
		fmt.Println("no message send")
	}

	select {
	case msg := <-messgaes:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default: // 无缓冲 无阻塞 直接执行default
		fmt.Println("no activity")
	}

}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

// 取值赋值操作是阻塞的
func block() {
	s := []int{1, 2, 3, 4, 5, 6, 7}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}

// range chan需要设置容量
func rangeoverchan() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)
	for elem := range queue {
		fmt.Println(elem)
	}
}

// 定时与超时 time
// case <-time.After(3*time.Second)

// 定时器设置
func timer() {
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(10 * time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer2 fired") // 将不会被触发
	}()
	stop2 := timer2.Stop() // 定时器关闭在前
	if stop2 {
		fmt.Println("Timer2 stopped")
	}
}

// Ticker 打点器与定时Timer相似
// 多了周期触发的特性

func timeTicker() {
	ticker := time.NewTicker(1 * time.Second)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Ticker stopped")
				return
			case t := <-ticker.C:
				fmt.Println("Ticker at ", t)
			}
		}

	}()
	time.Sleep(2 * time.Second) // 这里的耗时操作可以替换为其他业务逻辑
	// 比如： 网络请求

	ticker.Stop()
	done <- true
}
