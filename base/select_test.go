package base

import (
	"fmt"
	"testing"
	"time"
)

func Test_Select(t *testing.T) {
	// testSelect()
	//
	timeout()
}

// case 必须是 send / received
// 如果default存在 走default
// 否则随机执行 符合条件的case

// 可以理解为 case 下的所有语句都会尝试运行，如果都是等待 则会执行default
// 如果case 都可以执行 那就随机选取
func testSelect() {
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		time.Sleep(10 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()
	// msg1 := <-c1
	// fmt.Println("qqqq:", msg1)
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("recived:", msg1)
		case msg2 := <-c2:
			fmt.Println("recived:", msg2)
			// default:
			// 	time.Sleep(2 * time.Second)
			// 	fmt.Println("暂无消息", i)
		}
	}
}

// 超时设置与处理
func timeout() {

	tch := make(chan string)

	go func() {
		time.Sleep(5 * time.Second)
		tch <- "hi"
	}()

	select {
	case msg := <-tch:
		fmt.Println(msg)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout done!")
		close(tch)
	}
}
