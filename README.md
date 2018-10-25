# buxunxian
<a href=#1>Day-1 golang语言基础</a>

<a href=#2>Day-2 golang语言基础</a>

<a href=#3333>Day-1 golang语言基础</a>

<a id=3333>Day-1笔记</a>

## <a id=1>Day-1笔记</a>
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

- 2.channel
    - a.管道.类似unix/linux中的pipe
    - b.多个goroute之间通过channel进行通信
    - c.支持任何类型
    > 如果管道存放的值数量，多于定义的长度会造成死锁，导致程序无法运行

    代码位置(001\LivingExample-3\pipe.go)
    ```go
        pipe := make (chan int,3) //初始化一个chan，类型为整形，长度为3。
        pipe <- 1 //将1放到pipe里
        pipe <- 2 //将2放到pipe里
    ```
![image](le-3)
- 3.返回值
   代码位置Day1\LivingExample-4\return.go


### 包的概念
- 1.和ptyhon一样，把相同的代码放到一个目录，称之为包
- 2.包可以被其他包引用
- 3.main包是用来生成执行文件，每个程序只有一个main包
- 4.包的主要用途是提高代码的可复用性


### 包的实战
- 1.
>代码位置Day1\LivingExample-5，编译二进制文件
``go build -o bin\calc_test.exe  .\main\main.go``

- 2.多进程的goroute运算测试，
>代码位置Day1\LivingExample-6

作业：使用fmt,打印浮点数，二进制，十进制，十六进制

## <a id=2>Day-2笔记</a>

### 第一部分：基本数据类型和操作符

- <a href=#2-1>文件名&关键字&标识符</a>
- <a href=#2-2>Go程序基本结构</a>
- <a href=#2-3>常量和变量</a>
- <a href=#2-4>数据类型和操作符</a>.
- <a href=#2-5>字符串类型</a>

- <a id=2-1>文件名&关键字&标识符</a>
    - 1.所有go源码以.go结尾
    - 2.标识符以字母或下划线开头，大小写铭感例如

    ``a. boy``  ``b. Boy`` ``c. a+b(错误实例)`` ``d. 0boy(错误实例)`` ``e._boy`` ``f. =_boy(错误实例)`` ``g. _``
    - 3._是特殊标识符，用来忽略结果
    - 4.保留关键字

    |||||
    |-|-|-|-|
    |break|default|func|interface|select|
    |case|defer|go|map|struct|
    |chan|else|goto|package|switch|
    |const|fallthough|if|range|type|
    |continue|for|import|return|var|

- <a id=2-2>Go程序基本结构</a>
    ```go
    pacakge main 

    
    import(
        "fmt"
    )
    //main函数是程序入口
    func main() {
        fmt.Println("hello,world")
    }
    ```
    >练习1.写一个程序，对于给定一个数字n,求出所有亮亮相家等于,的组合。比如：对于n=3,所有的组合如下所示
    
    ```
    0+3=3
    1+2=3
    2+1=3
    0+3=3
    ```

    >代码位置Day2\LivingExample-1\main.go
    
    >练习2.一个程序包包含两个包add和main,其中add包中有两个变量，Name和age。请问main包中是如何访问Name和Age的(大写的为公有，小写为私有)

    >代码位置Day2\LivingExample-2

    >练习3.包别名的应用,开发一个程序，使用包别名来访问包中的函数

    >代码位置Day2\LivingExample-3\main

    >练习4.每个源文件都可以包含一个init函数，这个init函数自动被go运行框架调用。开发一个演示这个功能

    >代码示例Day2\LivingExample-4\add\add.go

- <a id=2-3>常量和变量</a>
- <a id=2-4>数据类型和操作符</a>
- <a id=2-5>字符串类型</a>
