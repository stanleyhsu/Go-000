学习笔记
## brief
* goroutine/chan


## Groutine
* 如果你的 goroutine 在从另一个 goroutine 获得结果之前无法取得进展，那么通常情况下，你自己去做这项工作比委托它( go func() )更简单。
* 每次开始一个goroutine, 问自己
    * 把并发交个调用者，显示让调用者知道 goroutine 的创建； 
    * 搞清楚goroutine什么时候退出，
    * 调用者能够控制goroutine什么时候退出

    * 什么会阻碍它退出；
    * 每个通道需要知道谁创建，谁发送，谁接收；发送者关闭； 
    * timeout control; 防止goroutine泄露；


* log.Fatal 退出不好，因为不能defer；只在 main 函数和 init函数中使用；

* file.Walk()

go serverApp(ctx) 
<-ctx.Done()

* sync.WaitGroup()


## Memory Model
* MESI: MESI协议是一个基于失效的缓存一致性协议;Modified (M) Exclusive (E) Shared (S) Invalid (I)
* 非统一内存访问架构（英语：Non-uniform memory access，简称NUMA）是一种为多处理器的电脑设计的内存架构，内存访问时间取决于内存相对于处理器的位置。在NUMA下，处理器访问它自己的本地内存的速度比非本地内存（内存位于另一个处理器，或者是处理器之间共享的内存）快一些。

## Sync
* CAS
* go tool compile -S *.go
* data race detect: go build -race; go test -race
* 最晚加锁，最早释放； 多个资源锁注意加锁顺序；
* atomic.Value Load()/Store()
* mutex. 锁饥饿；
* COW;
* errorgroup，有人报错全部取消；降级处理；
* 

## chan


## context


# Team Action
1. grpc error and deadline.
2. goroutine cleanup/recover
3. mutex 梳理
4. data race 检测 go build -race