学习笔记

## Groutine
* 如果你的 goroutine 在从另一个 goroutine 获得结果之前无法取得进展，那么通常情况下，你自己去做这项工作比委托它( go func() )更简单。
* 每次开始一个goroutine, 问自己
    * 把并发交个调用者，显示让调用者知道 goroutine 的创建； 
    * 搞清楚goroutine什么时候退出，
    * 调用者能够控制goroutine什么时候退出

    * 什么会阻碍它退出；
    * 每个通道需要知道谁创建，谁发送，谁接收；发送者关闭； 
    * timeout control


* log.Fatal 退出不好，因为不能defer；只在 main 函数和 init函数中使用；

* file.Walk()

go serverApp(ctx) 
<-ctx.Done()

* sync.WaitGroup()


## Memory Model
* MESI

## Sync
* CAS


## chan


## context