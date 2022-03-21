package base

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var wg sync.WaitGroup
	var ops uint64
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {
				// 锁上 + 锁 保持有序执行+1
				atomic.AddUint64(&ops, 1)
				// ops++
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("atomic.AddUint64 with sync.WaitGroup (g*50)*1000 =  ", ops)
}
