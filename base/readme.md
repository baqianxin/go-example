# go-example-cn 使用要点

小记与总结

## 重点内容

chan 通道

- 同步特性：阻塞
- 管道顺序： pipie

goroutines 协程

- 并发处理:
- 竞态race: 线程安全与锁
- 限流/限速处理
- 工作池 pool

同步锁/读写锁

- sync.WaitGroup
- chan 实现同步信号锁
- 原子操作 atomic.Add

基础：

- 字符串操作
- list-sort 排序&自定义排序
- 文件/目录操作
- 正则表达式匹配
- 数据结构处理json/xml
- 格式化输出fmt
- 自定义函数(通用函数)
- 接口实现与使用
- context 使用
- http请求/服务/cookie/sission/pwdhash/wensocket
- 缓冲字节数组读写操作
- 中间件使用
- gc 日志检查

其他：

- Signal系统信号+chan
- 启动子进程：exec & syscall
