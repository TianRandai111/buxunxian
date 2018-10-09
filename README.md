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

- 3.高并发测试
    file:///001/testgo.go