package base

import (
	"sync"
	"testing"
)

// 执行方式：
// GODEBUG=gctrace=1 go test -run ^Test_Gc$ algo/base
// 分析情况：
//  ...
//  gc 33 @0.451s 3%: 0.038+2.9+0.032 ms clock, 0.30+0.31/4.8/7.7+0.26 ms cpu, 5->6->2 MB, 6 MB goal, 8 P
//  格式
//  gc # @#s #%: #+#+# ms clock, #+#/#/#+# ms cpu, #->#-># MB, # MB goal, # P
//  含义
//  gc#：GC 执行次数的编号，每次叠加。
//  @#s：自程序启动后到当前的具体秒数。
//  #%：自程序启动以来在GC中花费的时间百分比。
//  #+...+#：GC 的标记工作共使用的 CPU 时间占总 CPU 时间的百分比。
//  #->#-># MB：分别表示 GC 启动时, GC 结束时, GC 活动时的堆大小(表示被标记对象的大小).
//  #MB goal：下一次触发 GC 的内存占用阈值。
//  #P：当前使用的处理器 P 的数量。

func Test_Gc(t *testing.T) {
	gc()
}

func gc() {
	wg := &sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(wg *sync.WaitGroup) {
			var counter int
			for i := 0; i < 1e8; i++ {
				counter++
			}
			wg.Done()
		}(wg)
	}
	wg.Wait()
}
