package base

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type Containers struct {
	mu       sync.Mutex
	counters map[string]int
}

// 并发写入map错误 - fatal error: concurrent map writes
// 需要加锁
func (c *Containers) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name] += 1
}

func TestMutex(t *testing.T) {
	c := Containers{
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	wg.Add(3)

	go doIncrement("a", 10000)
	go doIncrement("b", 10000)
	go doIncrement("a", 10000)

	wg.Wait()
	fmt.Println(c.counters)
}

/// //// ////

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

// 通过 chan 读取的阻塞特性 来实现共享数据与同步
// 建立通道
func TestMutexByChan(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	var readOps uint64
	var writeOps uint64

	reads := make(chan readOp)
	writes := make(chan writeOp)

	// 协程 读取读写信号
	go func() {
		var state = make(map[int]int)

		for {
			select {
			case read := <-reads:
				// 发送读取信号
				read.resp <- state[read.key]
			case write := <-writes:
				// 修改状态标记值
				state[write.key] = write.val
				// 发送写入信号
				write.resp <- true

			}
		}
	}()

	// 开始进行读写操作
	// 100个协程 写入 读操作队列，顺序执行，并且累计结果到 readOps 值
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int),
				}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool),
				}

				writes <- write
				// 阻塞-同步等待 === 锁
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	// 一秒执行可以看看总共执行了 多少次 协程的状态管理
	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps: ", readOpsFinal, readOps)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps: ", writeOpsFinal, writeOps)

}
