学习笔记

## errors
* 包与包之间的依赖的表面积尽量的少； 用不透明的错误， IsTimeout, IsTemporay()参考
* you should only handle errors once;直接返回error，只在请求的入口打印日志；
* 在出现错误情况，不能对返回值做任何假设；
* pkg/errors;
    1. 引用基础库、外部或者第三方库出错时，需要Wrap, 方便上层定位在哪里发生错误；
    2. 错误只处理一次：后续在程序内部，包括service层，直接透传不处理；直到最外部的用户调用是处理；
    3. package that are resuable across many projects only return root cause value
    4. If the error is not going to be handled, wrap and return up the call stack
    5. Once an error is handled , it is not allowed to be passed up the call stack any longer

* Go 1.13 及 Go 2 错误等；
    * pkg/errors尽量与标准库一致的使用方法；相比于标准库，pkg/errors多了调用栈信息；
    * Is, As等使用；

## Team Action
1. pkg/errors 源码阅读 学习使用，Wrap, WithMessage， WithStack, fmt.Errorf("%w")  及顶层打印等等
2. 实际项目修改和使用：完善当前zjdn101的错误处理及日志；下午先修改下，下周再堆业务代码；
2. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
### 深入学习
0. ppt 最后有reference, 有时间去看看；课程只是指引，原汁原味的味道还要自己品尝
1. defer recover panic 关系及处理: ttps://blog.golang.org/defer-panic-and-recover
2. errors 处理的实践：ppt 第三部分 Handling Error; 用 pkg/errors包；




- Q: 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
- A: 需要，
    1. 引用外部或者第三方库出错时，需要Wrap, 方便上层定位在哪里发生错误；
    2. 错误只处理一次：后续在程序内部，包括service层，直接透传不处理；直到最外部的用户调用是处理；