# buxunxian
## 001笔记
### Golang语言的特性
- 1.垃圾回收
    - a.内存自动回收，再也不需要开发人员管理内存
    - b.开发人员专注业务实现，降低了心智负担
    - c.只需要new分配内存，不需要释放

- 2.天然并发
    - a.从语言层面支持并发，非常简单
    - b.goroute，轻量级线程，创建成千上万个goroute成为可能
    - c.基于CSP(Communicating Sequential Process)模型实现

```go
func main() {
    go fmt.Println("hello")
}
```

### 高并发测试(一个测试)
- 1.代码位置/001/LivingExample-1/testgo.go

### 第一个golang程序
- 1.hello world代码位置(001\LivingExample-2\hello.go)

- 3.channel
    - a.管道.类似unix/linux中的pipe
    - b.多个goroute之间通过channel进行通信
    - c.支持任何类型
    > 如果管道存放的值数量，多余定义的长度会造成死锁，导致程序无法运行
    代码位置(001\LivingExample-3\pipe.go)
    ```go
        pipe := make (chan int,3) //初始化一个chan，类型为整形，长度为3。
        pipe <- 1 //将1放到pipe里
        pipe <- 2 //将2放到pipe里
    ```
   
![image](le-3)



