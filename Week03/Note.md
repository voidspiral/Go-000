## **Processes and Threads**



空的select语句将永远阻塞（不鼓励，而且我记得现在版本的go更严格了会直接保死锁）

Supervisor 拉起（**记得查询名词**）

>Supervisor是用Python开发的一套通用的进程管理程序，能将一个普通的命令行进程变为后台daemon，并监控进程状态，异常退出时能自动重启。它是通过fork/exec的方式把这些被管理的进程当作supervisor的子进程来启动，这样只要在supervisor的配置文件中，把要管理的进程的可执行文件的路径写进去即可。也实现当子进程挂掉的时候，父进程可以准确获取子进程挂掉的信息的，可以选择是否自己启动和报警。supervisor还提供了一个功能，可以为supervisord或者每个子进程，设置一个非root的user，这个user就可以管理它对应的进程。
>
>
>
>作者：风吹我已散博客
>链接：https://www.jianshu.com/p/0b9054b33db3
>来源：简书
>著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

log.Fatal

>Any time you start a Goroutine you must ask yourself:
>
>•*When will it terminate?*
>
>•*What could prevent it from terminating?*

tips:

​	**无缓冲通道**的特点是，发送的数据需要被读取后，发送才会完成。其阻塞场景有：

1. 通道中无数据，但执行读通道。
2. 通道中无数据，向通道写数据，但无协程读取。

```go
func serverApp(stop chan struct{}) error {
  // goroutine 1
  go func() {
    <-stop 
    http.Shutdown()
  }()
  //caller goroutine 2
  return http.Listen()
}
```

https://github.com/da440dil/go-workgroup/blob/master/README.md

## Never start a grouting without knowing when it will stop

```go
//leak is a buggy function. It launches a goroutine that
//blocks receiving from a channel. Nothing will ever be
//sent on that channel ans the channel is never closed so
//that goroutine will be blocked forever.
func leak() {
  ch := make(chan int)
  
  go func() {
    val := <-ch
    fmt.Println("We received a value:", val)
  }()
}
```

## Incomplete Work

Note:

[【用户行为采集】（一）常见埋点方式及对比](https://www.jianshu.com/p/6f5d60b04b93) 

## 总结：

1. 把并发丢给调用者
2. 搞清楚gorotine什么时候退出， 管控它的生命周期
3. 能够控制gorotine什么时候退出，比如channel or context

##The Go Memory Model

### Introduction

Go内存模型指定了一种条件，在这种条件下，即在一个goroutine中修改一个变量，在另外一个goroutine读取这个变量时需保证读取到修改后的值。

### Advice

修改由多个goroutine同时访问的数据的程序必须序列化此类访问。

要序列化访问，请使用channel操作或其他同步原语（例如sync和sync / atomic包中的原语）保护数据。

如果您必须阅读本文档的其余部分以了解程序的行为，那么您就太聪明了。

别太聪明了。

### Happens Before

在单个goroutine中，读取和写入的行为必须像它们按照程序指定的顺序执行一样。也就是说，仅当重新排序不会改变语言规范所定义的该goroutine中的行为时，编译器和处理器才可以对单个goroutine中执行的读取和写入进行重新排序。因为这种重新排序，一个goroutine观察到的执行顺序可能与另一个goroutine察觉到的执行顺序不同。如果一个goroutine执行a = 1； b = 2，另一个可能会在a的已更新值之前观察b的以更新值。

为了指定读写需求，我们定义了**happens before**， 一种Go程序对内存操作的局部执行顺序。翻译不动了，直接下面的链接吧

https://blog.csdn.net/q191201771/article/details/89472473?utm_medium=distribute.pc_relevant_t0.none-task-blog-BlogCommendFromMachineLearnPai2-1.control&depth_1-utm_source=distribute.pc_relevant_t0.none-task-blog-BlogCommendFromMachineLearnPai2-1.control

https://www.jianshu.com/p/5e44168f47a3



## 总结

这篇文章的意义就在于，告诉你用GO的并发原语来保证数据的一致性。

##Share Memory By Communicating

go tool compile -S file 

##Detecting Race Conditions With Go

注意 康康interface的实现

回头记录一下

go test -bench=.

go help test

## sync.atomic

Copy-On-Write （redis）

## Mutex



## errGroup

>看这个直接：https://github.com/go-kratos/kratos/blob/v2/app.go

