<a href=#1>Day-1 golang语言基础</a>

<a href=#2>Day-2 数据类型与操作符</a>

<a href=#3>Day-3 字符串处理包与函数</a>

<a href=#4>Day-4 数组、切片、map、package介绍</a>

<a href=#5>Day-5 结构体、方法</a>

<a href=#6>Day-6 接口、反射</a>

<a href=#7>Day-7 终端读写</a>>

<a href=#8>Day-8 并发与Chan</a>

<a href=#9>Day-9 网络编程</a>

<a href=#10>Day-10 HTTP编程</a>

<a href=#11>Day-11 日志收集1</a>

<a href=#12>Day-12 日志收集2</a>

<a href=#13>Day-13 日志收集3</a>

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

**重点GOPATH：环境变量 GOPATH 的值可以是一个目录的路径，也可以包含多个目录路径，每个目录都代表 Go 语言的一个工作区（workspace）。这些工作区用于放置 Go 语言的源码文件（source file），以及安装（install）后的归档文件（archive file，也就是以".a"为扩展名的文件）和可执行文件（executable file）。**

### 高并发测试(一个测试)
- 1.[代码位置: ](https://github.com/TianRandai111/buxunxian/blob/master/Day1/LivingExample-1/LivingExample-1/testgo.go)

### 第一个golang程序
- 1.``hello world``[代码位置: ](https://github.com/TianRandai111/buxunxian/blob/master/Day1/LivingExample-2/hello.go)
- 2.``channel``
    - a.管道.类似``unix/linux``中的``pipe``
    - b.多个``goroute``之间通过``channel``进行通信
    - c.支持任何类型
    > 如果管道存放的值数量，多于定义的长度会造成死锁，导致程序无法运行

    [代码位置: Day1/LivingExample-3/pipe.goo](https://github.com/TianRandai111/buxunxian/blob/master/Day1//LivingExample-3/pipe.go)

    ```go
        pipe := make (chan int,3) //初始化一个chan，类型为整形，长度为3。
        pipe <- 1 //将1放到pipe里
        pipe <- 2 //将2放到pipe里
    ```
- 3.返回值
    - [代码位置: Day1/LivingExample-4/return.go](https://github.com/TianRandai111/buxunxian/blob/master/Day1/LivingExample-4/return.go)

### 包的概念
- 1.和``ptyhon``一样，把相同的代码放到一个目录，称之为包
- 2.包可以被其他包引用
- 3.``main``包是用来生成执行文件，每个程序只有一个``main``包
- 4.包的主要用途是提高代码的可复用性


### 包的实战
- 1.
>[main代码位置: Day1/LivingExample-5/main/main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day1/LivingExample-5/main/main.go)

>[add模块代码位置: Day1/LivingExample-5/calc/add.go](https://github.com/TianRandai111/buxunxian/blob/master/Day1/LivingExample-5/calc/add.go)

>[sub模块代码位置: Day1/LivingExample-5/calc/add.go](https://github.com/TianRandai111/buxunxian/blob/master/Day1/LivingExample-5/calc/add.go)

>编译二进制文件``go build -o bin/calc_test.exe  ./main/main.go``

- 2.多进程的``goroute``运算测试，

>[main代码位置: Day1/LivingExample-6/main/main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day1/LivingExample-6/main/main.go)

>[main代码位置: Day1/LivingExample-6/goroute/add.go](https://github.com/TianRandai111/buxunxian/blob/master/Day1/LivingExample-6/goroute/add.go)

作业：使用``fmt``,打印浮点数，二进制，十进制，十六进制

## <a id=2>Day-2笔记</a>

### 第一部分：基本数据类型和操作符

- <a href=#2-1>文件名&关键字&标识符</a>
- <a href=#2-2>Go程序基本结构</a>
- <a href=#2-6>函数的声明和注释</a>
- <a href=#2-3>常量和变量</a>
- <a href=#2-7>值类型和引用类型</a>
- <a href=#2-4>数据类型和操作符</a>
- <a href=#2-5>字符串类型</a>

### 文件名&关键字&标识符

- <a id=2-1>文件名&关键字&标识符</a>
    - 1.所有``go``源码以``.go``结尾
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

### Go程序基本结
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
    >练习1.写一个程序，对于给定一个数字``n``,求出所有亮亮相家等于,的组合。比如：对于``n=3``,所有的组合如下所示
    
    ```
    0+3=3
    1+2=3
    2+1=3
    0+3=3
    ```

    >[main代码位置: Day2/LivingExample-1/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day2/LivingExample-1/main.go)
    
    >练习2.一个程序包包含两个包``add``和``main``,其中add包中有两个变量，``Name``和``age``。请问``main``包中是如何访问``Name``和``Age``的(大写的为公有，小写为私有)

    >[main代码位置: Day2/LivingExample-2/main/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day2/LivingExample-2/main/main.go)

    >[add代码位置: Day2/LivingExample-2/add/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day2/LivingExample-2/add/main.go)

    >练习3.包别名的应用,开发一个程序，使用包别名来访问包中的函数

    >[main代码位置: Day2/LivingExample-3/main/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day2/LivingExample-3/main/main.go)

    >[add代码位置: Day2/LivingExample-3/add/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day2/LivingExample-3/add/main.go)

    >练习4.每个源文件都可以包含一个``init``函数，这个``init``函数自动被``go``运行框架调用。开发一个演示这个功能

    >[代码位置: Day2/LivingExample-4/add/add.go](https://github.com/TianRandai111/buxunxian/tree/master/Day2/LivingExample-4/add/add.go)


### 常量和变量
- <a id=2-3>常量和变量</a>
    - 1.常量使用``const``修饰，代表永远只读，不能修改。
    - 2.``const``只能修饰``boolean``,``number(int相关类型、浮点类型、complex)和string``
    - 3.语法``const identifier [type] = value``,其中``type``可以省略。
    ```go
    const b string = "hello world"
    const b = "hello world"
    const Pi = 3.1415926
    const a = 9/3
    const (
        a = 1
        b = 2
        c = 3
    )
    const (
        a = iota
        b
        c
    )
    ```
    >练习：定义两个常量``Man=1``和``Female=2``,获取当年时间的秒数，如果能被``Femal``e整除，则在终端打印``female``，否则打印``man``
    
    >[代码位置: Day2/LivingExample-5/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day2/LivingExample-5/main.go)

    - 4. 变量
        - 4.1 语法:``var identifier type``
        ```go
        //eg1:
        var a int
        var b string
        var c bool
        var d int = 8
        var e string = "Hello world"
        //eg2:
        var (
            a int  //默认是0
            b string //默认是""
            c bool //默认是false
            d int = 8
            e string = "Hello world"
        )
        //eg3:
        var (
            a int  //默认是0
            b string //默认是""
            c bool //默认是false
            d = 8
            e = "Hello world"
        )
        ```
        >练习写一个程序获取当前运行的系统名称和``PATH``环境变量的值，并打印在终端

        >[代码位置: Day2/LivingExample-6/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day2/LivingExample-6/main.go)

        - 4.2 变量的作用域
            - 4.2.1 在杉树内部声明的变量叫做局部变量，生命周期仅存限于函数内部。
            - 4.2.2 在函数外部生命的变量叫做全局变量，声明周期租用与中个包，如果是大写的则作用于整个程序。
            - 4.2.3 函数内部``{}``大括号中的语句块中定义的变量，只能在语句块中使用。
            >练习下列程序的输出是什么,[代码位置: Day2/LivingExample-10/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day2/LivingExample-10/main.go)

            ```go
            package main

            import (
                "fmt"
            )

            var a = "G"

            func n() {
                fmt.Println(a)
            }

            func m() {
                var a = "O"
                fmt.Println(a)
            }

            func main() {
                n()
                m()
                n()
            }
            ```

            >练习下列程序输出的是什么，[代码位置: Day2/LivingExample-11/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day2/LivingExample-11/main.go)

            ```go
            package main

            import (
                "fmt"
            )

            var a string

            func main() {
                a = "G"
                fmt.Println(a)
                f1()
            }

            func f1() {
                a := "O"
                fmt.Println(a)
                f2()
            }

            func f2() {
                fmt.Println(a)
            }

            ```

- <a id=2-4>数据类型和操作符</a>
- 1.``bool``类型，只能存``true``和``false``
    ```go
    var a bool
    var a bool = true
    var a = true
    ```
- 2.相关操作符
    ```go
    var a bool = true
    var b bool
    ```
    > 请问``!a``、``!b``、``a&&b``、``a||b``的值分别是多少
    
- 3.数字类型，主要有
    - 3.1`` int ``整形(由正负)
        - ``int``、``int8``、``int16``、``int32``、``int68``
    - 3.2 ``uint`` 无符号整形(无负数)
        - ``uint8``、``uint16``、``uint32``、``uint64``
    - 3.3 ``float``浮点数
        - ``float32``、``float64``
        
- 4. 类型转换
    ```go
    var a int8 = 100
    //将int8类型的a转换成int16
    var b int16 = int16(a)
    ```
### 字符串类型
- <a id=2-5>字符串类型</a>
    - 字符串拼装操作[代码位置: Day2/LivingExample-12/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day2/LivingExample-12/main.go)

    - 字符串反转：[代码位置: Day2/LivingExample-13/main/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day2/LivingExample-13/main/main.go)

### 函数的声明和注释
- <a id=2-6>函数的声明和注释</a>
    - 1.函数声明：``func函数名字(参数列表)(返回值列表){}``
        ```go
        举例
        func add(){}
        ```
        ```go
        func add (a int,b int)int {}
        ```
        ```go
        func add (a int,b int)(int,int){}
        ```
    - 2.注释，两种注释，单行注释：``//``多行注释：``/*...*/``
        ```go
        //add计算两个证书的和，并返回结果
        func add(a int, b int) int {}
        ```
        ```go
        /*add计算两个证书的和，并返回结果*/
        func add(a int, b int) int {}
        ```
    
### 值类型和引用类型
- <a id=2-7>值类型和引用类型</a>
    - 1.值类型：变量直接存储值，内存通常在栈中分配
        - a.值类型：基本数据类型``int``、``float``、``bool``、``string``已经数组和``struct.``
    - 2.应用类型：变量存储的是一个地址，这个地址存储最终的值。内存通常在对战分配，通过GC回收
        - a.引用类型：指针、``slice``、``map``、``chan``等都是引用类型。
    > go语言交换值独有的方式[代码位置: Day2/LivingExample-9/main/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day2/LivingExample-9/main/main.go)

- 作业：
    - 判断101和200之间有多少个素数，并输出所有的素数
    >[代码位置: Day2/Assignment/PrimeNumber/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day2/Assignment/PrimeNumber/main.go)

    - 打印出100-999中所有的"水仙花数，所谓的水仙花数是一个三位数，其中苏子的立方和等于该数恩深，例如153是一个水仙花数，因为153=1的三次方+5的三次方+3的三次方
    >[代码位置: Day2/Assignment/PPDI/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day2/Assignment/PPDI/main.go)

    >[代码位置: (使用ASCII码来表示)Day2/Assignment/ASCII/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day2/Assignment/ASCII/main.go)

    - 对于一个数``n``,求··的阶乘之和：，既：``1！+2！+3！+...n``
    >[代码位置: Day2/Assignment/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day2/Assignment/main.go)


    - 变量赋值的区别
    >[代码位置: Day2/Assignment/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day2/Assignment/main.go)

## <a id=3>字符串处理包与函数</a>

- <a href=#3-1>1. strings和strconv的使用</a>
- <a href=#3-2>2. Go中的时间和日期类型</a>
- <a href=#3-3>3. 指针类型</a>
- <a href=#3-6>3. 流程控制</a>
- <a href=#3-4>4. 函数详解</a>
- <a href=#3-5>5. 课后作业</a>


### ``strings``和``strconv``的使
- <a id=3-1>1. ``strings``和``strconv``的使用</a> 
    - 1.1. ``strings.HasPrefix(s string, prefix string) bool``：判断字符串s是否以prefix开头
    >练习1：判断一个url是否以http://开头，如果不是，则加上http://。

    >[代码位置: Day3/LivingExample-1/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day3/LivingExample-1/main.go)

    - 1.2. ``strings.HasSuffix(s string, suffix string) bool``：判断字符串s是否以suffix结尾。
    >练习2：判断一个路径是否以"/"结尾，如果不是，则加上/。

    >[代码位置: Day3/LivingExample-2/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day3/LivingExample-2/main.go)

    - 1.3. ``strings.Index(s string, str string) int``：判断str在s中首次出现的位置，如果没有出现，则返回-1
    >[代码位置: Day3/LivingExample-3/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day3/LivingExample-3/main.go)

    - 1.4. `strings.LastIndex(s string, str string) int`：判断str在s中最后出现的位置，如果没有出现，则返回-1
    >[代码位置: Day3/LivingExample-4/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day3/LivingExample-4/main.go)

    - 1.5. `strings.Replace(str string, old string, new string, n int)`：字符串替换
    - 1.6. `strings.Count(str string, substr string)int`：字符串计数
    - 1.7. `strings.Repeat(str string, count int)string`：重复count次str
    - 1.8. `strings.ToLower(str string)string`：转为小写
    - 1.9. `strings.ToUpper(str string)string`：转为大写
    >练习4：写一个函数分别演示Replace、Count、Repeat、ToLower、ToUpper的用法

    >[代码位置: Day3/LivingExample-5/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day3/LivingExample-5/main.go)

    - 1.10. `strings.TrimSpace(str string)`：去掉字符串首尾空白字符
        - 1.10.1. `strings.Trim(str string, cut string)`：去掉字符串首尾cut字符
        - 1.10.2. `strings.TrimLeft(str string, cut string)`：去掉字符串首cut字符
        - 1.10.3. `strings.TrimRight(str string, cut string)`：去掉字符串首cut字符
    - 1.11.` strings.Field(str string)：返回str空格分隔的所有子串的slice
        - 1.11.1. `strings.Split(str string, split string`)：返回str split分隔的所有子串的slice
    - 1.12. `strings.Join(s1 []string, sep string)`：用sep把s1中的所有元素链接起来
    >练习5：写一个函数分别演示TrimSpace、Trim、TrimLeft、TrimRight、Field、Split、以及Join的用法

    >[代码位置: Day3/LivingExample-6/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day3/LivingExample-6/main.go)

    - 1.13. strconv.Itoa(i int)：把一个整数i转成字符串
    - 1.14. strconv.Atoi(str string)(int, error)：把一个字符串转成整数
    >练习6：写一个函数分别演示Itoa、Atoi的用法
    
    >[代码位置: Day3/LivingExample-7/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day3/LivingExample-7/main.go)


### 2. Go中的时间和日期类型
- <a id=3-2>2. Go中的时间和日期类型</a>
    - 2.1 `time`包
    - 2.2 `time.{Time类型}`，用来表示时间
    - 2.3 获取当前时间``now := time.Now()``(通过这个参数可以获取以下的值)
    - 2.4 通过这个上面参数可以获取以下的值
        - 2.4.1 ``time.Now().Day()``
        - 2.4.2 ``time.Now().Minute()``
        - 2.4.3 ``time.Now().Month()``
        - 2.4.4 ``time.Now().Year()``
    - 2.5 格式化输出,``fmt.Printf("%02d/%02d/%02d %02d:%02d:%02d",now.Year()''')``
        - 2.5.1 格式化输出
        
        ```go
        now := time.Now()
        fmt.Println(now.Format("02/1/2006 15:04"))
        fmt.Println(now.Format("2006/1/02 15:04"))
        fmt.Println(now.Format("2006/1/02"))
        ```
        
        >[代码位置: Day3/LivingExample-8/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day3/LivingExample-8/main.go)

    - 2.6 ``time.Duration``用来表示纳秒
    - 2.7 常量

        ```
        const (
            Nanosecond  Duration = 1
            Microsecond          = 1000 * Nanosecond
            Millisecond          = 1000 * Microsecond
            Second               = 1000 * Millisecond
            Minute               = 60 * Second
            Hour                 = 60 * Minute
        )
        ```
    > 练习6：写一个程序，获取当前时间，并格式化成 2017/06/15 08:05:00形式[代码位置: Day3/LivingExample-9/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day3/LivingExample-9/main.go)

    > 练习7：写一个程序，统计一段代码的执行耗时，单位精确到微秒。[代码位置: Day3/LivingExample-10/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day3/LivingExample-10/main.go)


### 3. 指针类型
- <a id=3-3>3. 指针类型</a>
    - 1. 普通类型，变量存的就是值，也叫值类型
    - 2. 获取变量的地址，用&，比如： ``var a int``, 获取a的地址：``&a``
    - 3. 指针类型，变量存的是一个地址，这个地址存的才是值
    - 4. 获取指针类型所指向的值，使用：``*``，比如：``var *p int``, 使用``*p``获取``p``指向的值
    ```go
    var a int = 5  //a -- >5 变量a指向5
    var p *int = &a //0xefefef -> 5 指针p指向a的地址，a地址对应的值为5
    ```
    >练习8：写一个程序，获取一个变量的地址，并打印到终端。[代码位置: Day3/LivingExample-11/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day3/LivingExample-11/main.go)


    >练习9：写一个函数，传入一个··类型的指针，并在函数中修改所指向的值。在``main``函数中调用这个函数，并把修改前后的值打印到终端，观察结果[代码位置: Day3/LivingExample-12/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day3/LivingExample-12/main.go)

- <a id=3-6>3. 流程控制</a>
    - 1.if/else分钟判断
    ```go
    // no.1
    if condition {

    }

    // no.2
    if condition {
        
    }else {

    }

    // no.3 
    if condition {

    }else if condition {

    }else {

    }
    ```
    >练习10：写一个程序，从终端读取输入，并转成整数，如果转成整数出错，则输出 ``"can not convert to int"``，并返回。否则输出该整数。[代码位置: Day3/LivingExample-13/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day3/LivingExample-13/main.go)

    - 2.``switch case``语句,如果使用``fallthrough``会继续向下执行[代码位置: Day3/LivingExample-14/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day3/LivingExample-14/main.go)
        ```go
        // 写法1
        switch var {
            case var1,var2:
            case var3:
            case var4:
            default:
        }

        //写法2
        var i = 0
        switch {
        condition1:
            fmt.Println("i > 0 and i < 10")
        condition2:
        fmt.Println("i > 10 and i < 20")
        default:
            fmt.Println("def")
        }

        //写法3
        switch i := 0; {
        condition1:
            fmt.Println("i > 0 and i < 10")
        condition2:
        fmt.Println("i > 10 and i < 20")
        default:
            fmt.Println("def")
        }
        ```
        >练习：猜数字，写一个程序，随机生成一个0到100的整数n，然后用户在终端，输入数字，如果和n相等，则提示用户猜对了。如果不相等，则提示用户，大于或小于n。[代码位置: Day3/LivingExample-15 /main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day3/LivingExample-15/main.go)

    - 3.for循环
        ```go
        for 初始化语句;条件判断；变量修改{

        }

        //写法1
        for i := 0 ;i<cont;i++{

        }
        //打印改图形
        A
        AA
        AAA
        AAAA
        AAAAA
        AAAAAA

        //写法2
        for 条件判断{

        }
        //eg:
        for i>0{
            fmt.Println("i>0")
        }
        for true{
            fmt.Println("i>0")
        }
        for {
            fmt.Println("i>0")
        }

        //写法3
        str := "hello world,中国"
        for i ,v := range str{
            ftm.Printf("index:%d,value:%s",i,v)
        }

        //break continue的使用方法
        //Lable语句 字符串加: 配合goto continue使用
        
        ```
        >打印图形[代码位置: Day3/LivingExample-16 /main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day3/LivingExample-16/main.go)

        >``break continue``的使用方法[代码位置: Day3/LivingExample-17 /main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day3/LivingExample-17/main.go)
        
### 4. 函数详解
- <a id=3-4>4. 函数详解</a>
    - 1.语法声明：``func 函数名(参数列表) [(返回值列表)] {}``
    ```golang
    func add(){}
    func add(a int, b int){}
    func add(a int, b int) int {}
    func add(a int, b int) (int, int) {}
    func add(a,b int) (int, int) {}
    ```
    - 2.golang函数特点
        - 2.1 不支持重载，一个包不能有两个一样的函数
        - 2.2 函数是一等公民，函数也是一种类型，函数可以复制给变量[代码位置: Day3/LivingExample-18 /main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day3/LivingExample-18/main.go)
        - 2.3 匿名函数
        - 2.4 多返回值
    - 3.函数参数传递方式
        - 3.1 值传递
        - 3.2 引用传递
    >注意1：无论是值传递，还是引用传递，传递给函数的都是变量的副本，不过，值传递是值的拷贝。引用传递是地址的拷贝，一般来说，地址拷贝更为高效。而值拷贝取决于拷贝的对象大小，对象越大，则性能越低。

    >注意2：``map``、``slice``、``chan``、指针、``interface``默认以引用的方式传递。

    - 4.命名返回值的名字
    ```go
    func add(a, b int) (c int) {
        c = a + b
        return
    }
    func calc(a, b int) (sum int, avg int) {
            sum = a + b
            avg = (a +b)/2
            return
            }
    ```

    - 5.忽略返回值
    ```golang
    func calc (a,b int) (sum int,avg int){
        sum = a + b
        avg = (a + b)/2
        return
    }
    func main() {
        sum, _ := calc(100,200)
    }
    ```

    - 6.可变参数：
    ```golang
    //0个或多个参数,arg属于切片
    func add (arg...int) int {}

    //1个或多个参数
    func add (a int,arg...int) int {}

    //2个或多个参数
    func add (a int,b int,arg...int) int {}
    ```

    >练习14：写一个函数add，支持1个或多个int相加，并返回相加结果

    >练习15：写一个函数concat，支持1个或多个string相拼接，并返回结果

    >[代码位置: Day3/LivingExample-19/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day3/LivingExample-19/main.go)

    - 7.defer的用途：
        - a 当函数返回时，执行defer语句。因此，可以用来做资源清理
        - b 多个defer语句，按先进后出的方式执行
        - c defer语句中的变量，在defer声明时就决定了。

        >[代码位置: Day3/LivingExample-20/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day3/LivingExample-20/main.go)

        ```golang
        func a() {
            i := 0
            defer fmt.Println(i)
            i++
            return
            } 
        func f() {
                for i := 0; i < 5; i++ {
                        defer fmt.Printf("%d ", i)
                } 
            } 
        ```
        - 7.1 锁资源释放

        ```go
        func read() {
        file := open(filename)
        defer file.Close()

        //文件操作
        }
        ```

        - 7.2 锁资源释放

        ```golang
        func read() {
        mc.Lock()
        defer mc.Unlock()
        //其他操作
        }
        ```

        - 7.3 数据库连接释放

        ```golang
        func read() {
        conn := openDatabase()
        defer conn.Close()
        //其他操作
        }
        ```
### 5. 课后作业
- <a id=3-5>5. 课后作业</a>
    - 1.编写程序，在终端输出九九乘法表。[代码位置: Day3/Assignment/one/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day3/Assignment/one/main.go)
    - 2.一个数如果恰好等于它的因子之和，这个数就称为"完数"。例如`6=1＋2＋3`.编程找出1000以内的所有完数。[代码位置: Day3/Assignment/two/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day3/Assignment/two/main.go)
    - 3.输入一个字符串，判断其是否为回文。回文字符串是指从左到右读和从右到左读完全相同的字符串。[代码位置: Day3/Assignment/three/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day3/Assignment/three/main.go)
    - 4.输入一行字符，分别统计出其中英文字母、空格、数字和其它字符的个数。[代码位置: Day3/Assignment/four/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day3/Assignment/four/main.go)
    - 5.计算两个大数相加的和，这两个大数会超过`int64`的表示范围.
    - [代码位置: 自己的版本 Day3/Assignment/five/my-version/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day3/Assignment/five/my-version/main.go)
    - [代码位置: 别人的版本Day3/Assignment/five/teacher-version/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day3/Assignment/five/teacher-version/main.go)

  
## 数组、切片、map、package介
<a id=4>数组、切片、map、package介绍</a>

<a href=#4-1>1.内置函数、递归函数、闭包</a>

<a href=#4-2>2.数组与切片</a>

<a href=#4-3>3.map数组结构</a>

<a href=#4-4>4.package介绍</a>

<a href=#4-5>5.课后作业</a>


### 1.内置函数、递归函数、闭包

<a id=4-1>1.内置函数、递归函数、闭包</a>

- 1.1. 内置函数
    - 1.1.1. `close`：主要用来关闭`channel`
    - 1.1.2. `len`：用来求长度，比如`string`、`array`、`slice`、`map`、`channel`
    - 1.1.3. `new`：用来分配内存，主要用来分配值类型，比如`int`、`struct`。返回的是指针[代码位置: Day4/LivingExample-1/main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day4/LivingExample-1/main.go)
    - 1.1.4. `make`：用来分配内存，主要用来分配引用类型，比如`chan`、`map`、`slice`
    - 1.1.5. `append`：用来追加元素到数组、slice中[代码位置: Day4/LivingExample-2/main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day4/LivingExample-2/main.go)
    - 1.1.6. `panic`和`recover`：用来做错误处理[代码位置: Day4/LivingExample-3/main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day4/LivingExample-3/main.go)

    >new和make的区别

    ![image](https://github.com/TianRandai111/buxunxian/blob/master/Day4/Image/Image-1.png?raw=true)

    [代码位置: Day4/LivingExample-4/main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day4/LivingExample-4/main.go)

- 1.2. 递归函数-自己调用自己
    - 1.2.1 定义好出口条件
    - 1.2.2 一个大问题可以分解成相似的小问题
    - 1.2.3 数字的阶层 [代码位置: Day4/LivingExample-5/main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day4/LivingExample-5/main.go)
    - 1.2.4 斐波那契数列  [代码位置: Day4/LivingExample-6/main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day4/LivingExample-6/main.go)

- 1.3 闭包
    - 1.3.1 一个函数与其相关的引用环境组合成的实体
    - 1.3.2 闭包的函数会保存更改的状态
    - [代码位置: Day4/LivingExample-7/main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day4/LivingExample-7/main.go)
    - [代码位置: Day4/LivingExample-8/main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day4/LivingExample-8/main.go)

### 2.数组与切片
<a id=4-2>2.数组与切片</a>

- 2.1 数组
    - 2.1.1 数组：使用中数据类型的固定长度的序列。
    - 2.1.2 数组定义：`var a [len]int`,比如：`var a[5]int`
    - 2.1.3 长度是数组类型的一部分，因此，var a[5] int和 var a[10] int是不同的类型
    - 2.1.4 数组可以通过下标进行访问，下标是从0开始的，最后一个元素的下标是，len-1
    ```go
    var a [10]int
    for i := 0;i < len(a);i++ {
    }
    for i ,v := range a{
    }
    ```
    - 2.1.5 访问越界：如果下标在数组合法范围之外，则触发访问越界，会panic
    - 2.1.6 数组是值类型，因此改变副本的值，不会改变本身的值[代码位置: Day4/LivingExample-9/main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day4/LivingExample-9/main.go)

    >练习：使用非递归的方式实现斐波那契数列，打印前100个数。 [代码位置: Day4/LivingExample-10/main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day4/LivingExample-10/main.go)

    - 2.1.7 数组初始化
        - a. ``var age0 [5]int = [5]int{1,2,3}``
        - b. ``var age1 = [5]int{1,2,3,4,5,}``
        - c. ``var age2 = [...]int{1,2,3,4,5}``
        - d. ``var age3 = [5]string{3:"hello world",4:"tom"}``

    - 2.1.8 多维数组
        ```
        a ``var age [5][3]int``
        b ``var f[2][3]int = [...][3]int{{1,2,3,},{7,8,9}} //两行三列`` 
        ```

    - 2.1.9 多维数组遍历
        - [代码位置: Day4/LivingExample-11/main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day4/LivingExample-11/main.go)

- 2.2 切片
    - 2.2.1. 切片：切片是数组的一个引用，因此切片是引用类型
    - 2.2.2. 切片的长度可以改变，因此，切片是一个可变的数组
    - 2.2.3. 切片遍历方式和数组一样，可以用``len()``求长度
    - 2.2.4. ``cap//内置函数``可以求出``slice``最大的容量，``0 <= len(slice) <= (array)``，其中``array``是``slice``引用的数组
    - 2.2.5. 切片的定义：``var`` 变量名 []类型，比如`` var str []string  var arr []int``
    - 2.2.6. 切片初始化[代码位置: Day4/LivingExample-12/main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day4/LivingExample-12/main.go)
        - a. 切片初始化：``var slice []int = arr[start:end]``包含``start``到``end``之间的元素，但不包含``end`` 
        - b. ``var slice []int = arr[0:end]``可以简写为`` var slice []int=arr[:end]``
        - c. ``var slice []int = arr[start:len(arr)]`` 可以简写为 ``var slice[]int = arr[start:]``
        - d. ``var slice []int = arr[0, len(arr)]`` 可以简写为`` var slice[]int = arr[:]``
        - e. 如果要切片最后一个元素去掉，可以这么写： ``slice = slice[:len(slice)-1]``
    >练习：写一个程序，演示切片的各个用法[代码位置: Day4/LivingExample-13/main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day4/LivingExample-13/main.go)

    - 2.2.7. 切片的内存布局
        - ![image](https://github.com/TianRandai111/buxunxian/blob/master/Day4/Image/Image-2.png?raw=true)

    - 2.2.8. 通过make来创建切片
        - a. ``var slice []type = make([]type,len)``
        - b. ``slice := make([]type,len)``
        - c. ``slice := make([]type,len,cap)``
        - ![image](https://github.com/TianRandai111/buxunxian/blob/master/Day4/Image/Image-3.png?raw=true)

    - 2.2.9. 用append内置函数操作切片
        ```golang
        slice =append(slice, 10)
        var a = []int{1,2,3}
        var b = []int{4,5,6}
        a= append (a,b)
        ```
    
    - 2.2.10. 遍历切片
        ```go
        for index, val := range slice {}
        ```

    - 2.2.11. 切片resize
        ```golang
        var a = []int{1,3,4,5}
        b := a[1:2]
        b = b[0:3]
        ```
    
    - 2.2.12. 切片拷贝copy()//拷贝不会扩展切片，如果新生成的切片小于被拷贝的切片，只会拷贝到新切片下标的长度。
        ```go
        s1 := []int{1,2,3,4,5}
        s2 := make([]int, 10)
        copy(s2, s1)
        s3 := []int{1,2,3}
        s3 = append(s3, s2…)
        s3 = append(s3, 4,5,6)
        ```
        - [代码位置: Day4/LivingExample-14/main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day4/LivingExample-14/main.go)

    - 2.2.13. string与slice,string底层就是一个byte的数组，因此，也可以进行切片操作
        ```go
        str := "hello world"
        s1 := str[0:5]
        fmt.Println(s1)

        s2 := str[5:]
        fmt.Println(s2)
        ```
        - [代码位置: Day4/LivingExample-15/main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day4/LivingExample-15/main.go)
        - string的底层布局
        - ![image](https://github.com/TianRandai111/buxunxian/blob/master/Day4/Image/Image-4.png?raw=true)
        - 修改string的值[代码位置: Day4/LivingExample-15/main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day4/LivingExample-15/main.go)

    - 2.2.14. 排序和查找
        - a. 排序操作主要都在 ``sort``包中，导入就可以使用了``import("sort")``
        - b. ``sort.Ints``对整数进行排序， ``sort.Strings``对字符串进行排序, ``sort.Float64s``对浮点数进行排序.
        - c. ``sort.SearchInts(a []int, b int)`` 从数组``a``中查找``b``，前提是``a``必须有序
        - d. ``sort.SearchFloats(a []float64, b float64)`` 从数组``a``中查找``b``，前提是``a``必须有序
        - f. ``sort.SearchStrings(a []string, b string) ``从数组``a``中查找``b``，前提是``a``必须有序
        - [代码位置: Day4/LivingExample-17/main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day4/LivingExample-17/main.go)
        
        
### 3.map数组结构
<a id=4-3>3.map数组结构</a>

- 3.1. map简介:key-value的数据结构，又叫字典或关联数组
    - 3.1.1 声明//声明是不会分配内存的，初始化需要make,同样的key:value会覆盖
        ```golang
        var map1 map[keytype]valuetype
        var a map[string]string
        var a map[string]int
        var a map[int]string
        var a map[string]map[string]string
        ```
    - [代码位置: Day4/LivingExample-18/main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day4/LivingExample-18/main.go)

- 3.2. 相关操作
    ```go
    var a map[string]string = map[string]string{"hello": "world",}
    a = make(map[string]string, 10)
    a["hello"] = "world"    //插入和更新
    val, ok := a["hello"]   //查找
    for k, v := range a {   //遍历
        fmt.Println(k,v) 
    }
    delete(a, "hello")      //删除
    len(a)                  //长度
    ```
    - [代码位置: Day4/LivingExample-19/main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day4/LivingExample-19/main.go)

- 3.3. map是引用类型
    ```golang
    func modify(a map[string]int) {
    a[“one”] = 134
    }
    ```

- 3.4. slice of map
    ```golang
    Items := make([]map[int][int], 5)
    For I := 0; I < 5; i++ {
            items[i] = make(map[int][int])
    }
    ```
    - [代码位置: Day4/LivingExample-20/main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day4/LivingExample-20/main.go)

- 3.5. map排序
    - 3.5.1. 先获取所有key，把key进行排序
    - 3.5.2. 按照排序好的key，进行遍历
    - [代码位置: Day4/LivingExample-21/main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day4/LivingExample-21/main.go)

- 3.6. Map反转
    - 3.6.1. 初始化另外一个map，把key、value互换即可
    - [代码位置: Day4/LivingExample-22/main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day4/LivingExample-22/main.go)

### 4.package介绍
<a id=4-4>4.package介绍</a>

- 4.1. golang中的包
    - 4.1.1. golang目前有150个标准包，覆盖了几乎所有的基础库
    - 4.1.2. golang.org有所有包的文档,多多翻阅

- 4.2. 线程同步 //多个线程同时操作一个资源
    - 4.2.1. ``import("sync")``
    - 4.2.2. 互斥锁,``var mu sync.Mutex``
    - 4.2.3. 读写锁,``var mu sync.RWMutex``

- 4.3. go get安装第三方包

### 5.课后作业
<a id=4-5>5.课后作业</a>

- 5.1. 实现一个冒泡排序
    - 自己的版本[代码位置: Day4\Assignment\one\my-version\main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day4/Assignment/one/my-version/main.go)
    - 老师的版本[代码位置: Day4\Assignment\one\my-version\main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day4/Assignment/one/teacher-version/main.go)

- 5.2. 实现一个选择排序
    - 自己的版本[代码位置: Day4\Assignment\two\my-version\main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day4/Assignment/two/my-version/main.go)
    - 老师的版本[代码位置: Day4\Assignment\two\my-version\main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day4/Assignment/two/teacher-version/main.go)
    
- 5.3. 实现一个插入排序
    - 自己的版本[代码位置: Day4\Assignment\three\my-version\main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day4/Assignment/three/my-version/main.go)
    - 老师的版本[代码位置: Day4\Assignment\three\my-version\main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day4/Assignment/three/teacher-version/main.go)
    
- 5.4. 实现一个快速排序
    - 自己的版本[代码位置: Day4\Assignment\four\my-version\main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day4/Assignment/four/my-version/main.go)
    - 老师的版本[代码位置: Day4\Assignment\four\my-version\main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day4/Assignment/four/teacher-version/main.go)

## Day-5 结构体、方法、接口
<a id=5>Day-5 结构体、方法、接口</a>

<a href=#5-1>1. 结构体</a>

<a href=#5-2>2. 方法</a>

<a href=#5-4>3. 作业</a>


### 1.1 结构体
<a id=5-1>1. 结构体</a>

- 1.1 结构体的定义
    - 1.1.1 用来定义复杂的数据结构
    - 1.1.2 struct里面可以包含多个字段（属性）
    - 1.1.3 strunt类型可以定义方法，注意和函数的区分
    - 1.1.4 sturct类型是值类型
    - 1.1.5 struct类型可以嵌套
    - 1.1.6 Go语言没有class类型，只有struct类型

- 1.2 结构体的声明

    ```go
    type 标识符 struct{
        field1 type
        field1 type
    }
    // 例子
    type Student struct{
        Name string 
        age int
        scort int
    }
    ```
    - [代码位置: Day5/LivingExample-1/main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day5/LivingExample-1/main.go)

- 1.3 struct 中字段访问，和其他语言一样，实用点
    ```golang
    //例子
    var stu Student
    stu.Name = "tony"
    stu.Age = 18
    stu.Score = 20

    fmt.Printf("name=%s,age=%d,Score=%d",stu.Name,stu.Age,stu.Score)
    ```

- 1.4 结构体定义的三种形式：

    ```go
    var stu student
    var stu *student = new(Student)
    var stu *student = &Student{}
    ```
    - 1.4.1 其中二、三返回的都是指向结构体的指针，访问形式如下：
    ``stu.Name、stu.Age和stu.Score或者(*stu).Name、(*stu).Age等``

- 1.5 链表定义
    ```go
    type Student struct{
        Name string
        Next *Student
    }
    ```
    - 链表的插入方法
        - 尾部插入法
            - 手动插入链表 [代码位置: Day5/LivingExample-2/main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day5/LivingExample-2/main.go)
            - 循环尾部插入链表 [代码位置: Day5/LivingExample-3/main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day5/LivingExample-3/main.go)
            - 循环头部插入链表，添加、删除链表的一个节点[代码位置: Day5/LivingExample-4/main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day5/LivingExample-4/main.go)

- 1.6双链表定义
    ```golang
    type Student struct{
        Name string
        Next *Student
        Prev *Student
    }
    ```
    >如果有两个指针分别指向前一个节点和后一个节点,我们叫做双链表

- 1.7 二叉树
    ```golang
    type Student struct{
        Name string
        left *Student
        right *Student
    }
    ```
    >如果每个几点有两个指针飞边指向左子树和右子树，我们吧这样的结构叫做二叉树

    >[代码位置: Day5/LivingExample-5/main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day5/LivingExample-5/main.go)

- 1.8 结构体是用户单独定义的类型，不能和其他类型进行强制转换

```golang
type Student struct {
        Number int
}

type Stu Student //alias

//Student与Stu不是同一类型

var a Student
a = Student(30)

var b Stu

//此处赋值会失败
a = b

//需要类型转换
a = Student(b)
```

- 1.9 golang中的struct没有构造函数，一般可以使用工厂模式来解决这个问题
```golang
Package model
type student struct {
       Name stirng
    Age int
}

func NewStudent(name string, age int) *student {
return &student{
       Name:name,
       Age:age,}
}

Package main
S := new (student)
S := model.NewStudent(“tony”, 20)

```

- 1.11 我们可以为struct中的每个字段，写上一个tag。这个tag可以通过反射的机制获取到，最常用的场景就是json序列化和反序列化(描述信息)

```golang
type student struct {
       Name stirng  “this is name field”
    Age int           “this is age field”
}
```
>[代码位置: Day5/LivingExample-6/main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day5/LivingExample-6/main.go)


- 1.12 匿名字段，结构体中字段可以没有名字，即你名字段
```golang
type Car struct {
    Name string
    Age int        
}

type Train struct {
        Car
        int
        Start time.Time
}
```
>[代码位置: Day5/LivingExample-7/main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day5/LivingExample-7/main.go)
### 2. 方法
<a id=5-2>2. 方法</a>

- 2.0 方法 go中的方法是作用在特定类型的变量上，因此自定义类型，都可以有方法，而不仅仅是struct
    - 定义：``func (recevier type) methodName(参数列表)(返回值列表){}``

- 2.1 方法的调用
    - 方法的声明和调用[代码位置: Day5/LivingExample-8/main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day5/LivingExample-8/main.go)
```golang
type A struct {
        a int
}

func (this A) test() {         fmt.Println(this.a)
}

var t A
t.test()
```

- 2.2 方法和函数的区别
    - 2.2.1 函数调用： ``function(variable, 参数列表）``
    - 2.2.2 方法：``variable.function(参数列表）``

      
- 2.3 指针receiver   vs 值receiver
    - 2.3.1 本质上和函数的值传递和地址传递是一样的

- 2.4 方法的访问控制，通过大小写控制
      
- 2.5 继承
    - 2.5.1 如果一个``struct``嵌套了另一个匿名结构体，那么这个结构可以直接访问匿名结构体的方法，从而实现了继承。
    -[代码位置: Day5/LivingExample-9/main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day5/LivingExample-9/main.go)
    
- 2.6 组合和匿名字段
    - 2.6.1 如果一个``struct``嵌套了另一个匿名结构体，那么这个结构可以直接访问匿名结构体的方法，从而实现了继承。
    - 2.6.2 如果一个``struct``嵌套了另一个有名结构体，那么这个模式就叫组合。
    -[代码位置: Day5/LivingExample-9/main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day5/LivingExample-9/main.go)

- 2.7 多重继承
    - 2.7.1 如果一个``struct``嵌套了多个匿名结构体，那么这个结构可以直接访问多个匿名结构体的方法，从而实现了多重继承。

- 2.8 实现``String()``
    - 2.8.1 如果一个变量实现了``String()``这个方法，那么fmt.Println默认会调用这个变量的``String()``进行输出。

### 3. 作业
<a id=5-4>3. 作业</a>

- 1. 实现一个图书管理系统，具有以下功能：
    - a. 书籍录入功能，书籍信息包括书名、副本数、作者、出版日期
    - b. 书籍查询功能，按照书名、作者、出版日期等条件检索
    - c. 学生信息管理功能，管理每个学生的姓名、年级、身份证、性别、借了什么书等信息
    - d. 借书功能，学生可以查询想要的书籍，进行借出
    - e. 书籍管理功能，可以看到每种书被哪些人借出了

## Day-6 接口、反射
<a id=6>Day-6 接口、反射</a>

<a href=#6-1>1.接口</a>

<a href=#6-2>2.反射</a>

<a href=#6-3>3.作业</a>

<a id=6-1>1.接口</a>

### 1. 接口
<a id=6-1>1. 接口</a>

- 1.1. 定义
    - 1.1.1 Interface类型可以定义一组方法，但是这些不需要实现。并且interface不能包含任何变量。
    

- 1.2 接口声明  
    ```go 
    type example interface{

            Method1(参数列表) 返回值列表
            Method2(参数列表) 返回值列表
            …
    }
    ```

- 1.3. interface类型默认是一个指针
```go
type example interface{

        Method1(参数列表) 返回值列表
        Method2(参数列表) 返回值列表
        …
}

var a example
a.Method1()
```

- 1.4. 接口实现
    - 1.5.1. Golang中的接口，不需要显示的实现。只要一个变量，含有接口类型中的所有方法，那么这个变量就实现这个接口。因此，golang中没有implement类似的关键字
    - 1.5.2. 如果一个变量含有了多个interface类型的方法，那么这个变量就实现了多个接口。
    - 1.5.3. 如果一个变量只含有了1个interface的方部分方法，那么这个变量没有实现这个接口。

- 1.5 多态
    - 一种事物的多种形态，都可以按照统一的接口进行操作
    - [代码位置: Day6/LivingExample-1/main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day6/LivingExample-1/main.go)
    - [代码位置: Day6/LivingExample-2/main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day6/LivingExample-2/main.go)


- 1.6. 接口嵌套
    - 一个接口可以嵌套在另外的接口，如下所示：
    ```go
    type ReadWrite interface {
                Read(b Buffer) bool
                Write(b Buffer) bool
    } 
    type Lock interface {
                Lock()
                Unlock() 
    } 
    type File interface {
                ReadWrite
                Lock 
                Close() 
    } 
    ```

- 1.7. 类型断言，由于接口是一般类型，不知道具体类型，如果要转成具体类型
    - 可以采用以下方法进行转换：
    ```go
    var t int
    var x interface{}
    x = t
    y = x.(int)   //转成int

    var t int
    var x interface{}
    x = t
    y, ok = x.(int)   //转成int，带检查
    ```

- 1.8. 练习，写一个函数判断传入参数的类型
```go
 func classifier(items ...interface{}) {
          for i, x := range items { 
                  switch x.(type) {
                    case bool:       fmt.Printf(“param #%d is a bool\n”, i)
                    case float64:    fmt.Printf(“param #%d is a float64\n”, i)
                    case int, int64: fmt.Printf(“param #%d is an int\n”, i)
                    case nil: fmt.Printf(“param #%d is nil\n”, i)
                    case string: fmt.Printf(“param #%d is a string\n”, i)
                    default: fmt.Printf(“param #%d’s type is unknown\n”, i)
            }
}
```

- 1.9. 类型断言，采用type switch方式
```go
switch t := areaIntf.(type) {
case *Square:
    fmt.Printf(“Type Square %T with value %v\n”, t, t) 
case *Circle:
    fmt.Printf(“Type Circle %T with value %v\n”, t, t) 
case float32:
    fmt.Printf(“Type float32 with value %v\n”, t)
case nil:
    fmt.Println(“nil value: nothing to check?”) 
default:
    fmt.Printf(“Unexpected type %T”, t)
}
```
 - [代码位置: Day6/LivingExample-4main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day6/LivingExample-4/main.go)

- 1.10. 16. ``interface{}``，接口中一个方法也没有，所以任何类型都实现了空接口，也就是任何变量都可以赋值给空接口。
    - 1.10.1 空接口没有任何方法，所以所有类型都实现了空接口。
    ```go
    var a int
    var b interface{}
    b  = a
    ```
    - 1.10.2 变量slice和接口slice之间赋值操作，for range
    ```go
    var a []int
    var b []interface{}
    b = a
    ```


- 1.11. 判断一个变量是否实现了指定接口
    - 判断一个变量是否实现了指定接口
    ```go
    type Stringer interface {
            String() string 
    }
    var v MyStruct
    if sv, ok := v.(Stringer); ok {
        fmt.Printf(“v implements String(): %s\n”, sv.String()); 
    } 
    ```
 - [代码位置: Day6/LivingExample-3main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day6/LivingExample-3/main.go)
- 1.12. 指针类型和值类型的区别
      
- 1.13. 实现一个通用的链表类
       - [代码位置: Day6/LivingExample-6main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day6/LivingExample-6/main.go)
- 1.14. 实现一个负载均衡调度算法，支持随机、轮训等算法
      

### 2.反射

<a id=6-2>2.反射</a>

- 2.1. 反射：可以在运行时动态获取变量的相关信息Import (“reflect”)
    - 两个函数：
    - 2.1.1. ``reflect.TypeOf``，获取变量的类型，返回``reflect.Type``类型
    - 2.1.2. ``reflect.ValueOf``，获取变量的值，返回``reflect.Value``类型
    - 2.1.3. ``reflect.Value.Kind``，获取变量的类别，返回一个常量
    - 2.1.4. ``reflect.Value.Interface()``，转换成``interface{}``类型
    - [代码位置: Day6/LivingExample-6main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day6/LivingExample-6/main.go)

```
—————          ————————————          ——————————————
|变量|  <----> |interface{}|  <----> |Reflect.Value|
—————          ————————————          ——————————————
```

- 2.2. reflect.Value.Kind()方法返回的常量
```golang
const (
        Invalid Kind = iota
        Bool
        Int
        Int8
        Int16
        Int32
        Int64
        Uint
        Uint8
        Uint16
        Uint32
        Uint64
        Uintptr
        Float32
        Float64
        Complex64
        Complex128
        Array
        Chan
        Func
        Interface
        Map
        Ptr
        Slice
        String
        Struct
        UnsafePointer
) 
```

- 2.3. 练习：
```golang
package main

import (
	"fmt"
	"reflect"
)

func main() {

	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))
	v := reflect.ValueOf(x)
	fmt.Println("value:", v)
	fmt.Println("type:", v.Type())
	fmt.Println("kind:", v.Kind())
	fmt.Println("value:", v.Float())

	fmt.Println(v.Interface())
	fmt.Printf("value is %5.2e\n", v.Interface())
	y := v.Interface().(float64)
	fmt.Println(y)
}
```



```golang
package main

import (
	"fmt"
	"reflect"
)

func main() {

	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))
	v := reflect.ValueOf(x)
	fmt.Println("value:", v)
	fmt.Println("type:", v.Type())
	fmt.Println("kind:", v.Kind())
	fmt.Println("value:", v.Float())

	fmt.Println(v.Interface())
	fmt.Printf("value is %5.2e\n", v.Interface())
	y := v.Interface().(float64)
	fmt.Println(y)
}
```
- 2.4. 获取变量的值：

``reflect.ValueOf(x).Float() ``

``reflect.ValueOf(x).Int()``

``reflect.ValueOf(x).String()``

``reflect.ValueOf(x).Bool()``
- [代码位置: Day6/LivingExample-7main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day6/LivingExample-7/main.go)


- 2.5. 通过反射的来改变变量的值

``reflect.Value.SetXX``相关方法，比如:

``reflect.Value.SetFloat()``，设置浮点数

``reflect.Value.SetInt()``，设置整数

``reflect.Value.SetString()``，设置字符串
- [代码位置: Day6/LivingExample-8main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day6/LivingExample-8/main.go)

- 2.6. 练习
```golang
package main

import (
	"fmt"
	"reflect"
)

func main() {

	var a float64
	fv := reflect.ValueOf(a)
	fv.SetFloat(3.3)
	fmt.Printf("%v\n", a)
}

```

- 2.7. 崩溃的原因
    - 还是值类型和引用类型的原因
    - ``v := reflect.ValueOf(x) ``
    - v是x的一个拷贝，修改v，x不会修改!

- 2.8. 解决方法，传地址！
```golang
package main

import (
	"fmt"
	"reflect"
)

func main() {

	var a float64
	fv := reflect.ValueOf(&a)
	fv.Elem().SetFloat(3.3)
	fmt.Printf("%v\n", a)
}
```
>其中fv.Elem()用来获取指针指向的变量，相当于： ``var a *int;  *a = 100``

- 2.9. 用反射操作结构体
    - 2.9.1. reflect.Value.NumField()获取结构体中字段的个数
    - 2.9.2. reflect.Value.Method(n).Call来调用结构体中的方法
- [代码位置: Day6/LivingExample-9main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day6/LivingExample-9/main.go)


- 2.10. 练习

```golang
package main

import (
	"fmt"
	“reflect"
)

type NotknownType struct {
	s1 string
	s2 string
	s3 string
}
func (n NotknownType) String() string {
	return n.s1 + "-" + n.s2 + "-" + n.s3
}
var secret interface{} = NotknownType{"Ada", "Go", "Oberon"}

func main() {
	value := reflect.ValueOf(secret) // <main.NotknownType Value>
	typ := reflect.TypeOf(secret)    // main.NotknownType
	fmt.Println(typ)

	knd := value.Kind() // struct
	fmt.Println(knd)
	
	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("Field %d: %v\n", i, value.Field(i))
		//value.Field(i).SetString("C#")
	}
	
	results := value.Method(0).Call(nil)
	fmt.Println(results) // [Ada - Go - Oberon]
}
```

- 2.11. 练习，通过反射操作结构体
```go
package main

import (
	"fmt"
	“reflect"
)

type NotknownType struct {
	s1 string
	s2 string
	s3 string
}
func (n NotknownType) String() string {
	return n.s1 + "-" + n.s2 + "-" + n.s3
}
var secret interface{} = NotknownType{"Ada", "Go", "Oberon"}

func main() {
	value := reflect.ValueOf(secret) // <main.NotknownType Value>
	typ := reflect.TypeOf(secret)    // main.NotknownType
	fmt.Println(typ)

	knd := value.Kind() // struct
	fmt.Println(knd)
	
	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("Field %d: %v\n", i, value.Field(i))
		//value.Field(i).SetString("C#")
	}
	
	results := value.Method(0).Call(nil)
	fmt.Println(results) // [Ada - Go - Oberon]
}
```

- 2.12. 练习2，通过反射修改结构体
```golang
package main

import (
	"fmt"
	"reflect"
)

type T struct {
	A int
	B string
}

func main() {
	t := T{23, "skidoo"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
	s.Field(0).SetInt(77)
	s.Field(1).SetString("Sunset Strip")
	fmt.Println("t is now", t)
}
```

### 3.作业
<a id=6-2>3.作业</a>

- 1. 实现一个图书管理系统v2，具有以下功能：
    - a. 增加用户登录、注册功能
    - b. 增加借书过期的图书界面
    - c. 增加显示热门图书的功能，被借次数最多的top10
    - d. 增加查看某个人的借书记录的功能

## Day-7 终端读写

<a id=#7>Day-7 终端读写</a>

<a href=#7-1>终端读写</a>

<a href=#7-2>文件读写</a>

<a href=#7-3>命令行参数处理</a>

<a href=#7-4>json协议</a>

<a href=#7-5>错误处理</a>

### 7.1 终端读写
<a id=7-1>Day-7 终端读写</a>

- os.stdin

- os.stdout
    - [代码位置: Day7/LivingExample-1main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day7/LivingExample-1/main.go)
      
- os.stderr
    - [代码位置: Day7/LivingExample-2main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day7/LivingExample-2/main.go)
    ```go
    package main

    import (
        "fmt"
    )

    var (
        firstName, lastName, s string
        i                      int
        f                      float32
        input                  = "56.12 / 5212 / Go"
        format                = "%f / %d / %s"
    )

    func main() {
        fmt.Println("Please enter your full name: ")
        fmt.Scanln(&firstName, &lastName)
        // fmt.Scanf("%s %s", &firstName, &lastName)
        fmt.Printf("Hi %s %s!\n", firstName, lastName) // Hi Chris Naegels
        fmt.Sscanf(input, format, &f, &i, &s)
        fmt.Println("From the string we read: ", f, i, s)
    }
    ```
- 缓冲区读写
    - bufio
        - 标准输入读
        >[代码位置: Day7/LivingExample-3main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day7/LivingExample-3/main.go)

        - 文件中读
        >[代码位置: Day7/LivingExample-4main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day7/LivingExample-4/main.go)

        - 练习，从终端读取一行字符串，统计英文、数字、空格以及其他字符的数量。
        >[代码位置: Day7/LivingExample-4main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day7/LivingExample-4/main.go)

### 7.2 文件读
<a id=7-2>7.2 文件读写</a>

- 7.2.1 os.File封装所有文件相关操作，之前讲的 os.Stdin,os.Stdout, os.Stderr都是 os.File
    - a 打开一个文件进行读操作： os.Open(name string) (*File, error)
    - b 关闭一个文件：File.Close()
    >[代码位置: Day7/LivingExample-5main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day7/LivingExample-5/main.go)

    - 复制文件内容(小文件)
    >[代码位置: Day7/LivingExample-6main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day7/LivingExample-6/main.go)

    - 读取*.gz压缩文件中的内容
    >[代码位置: Day7/LivingExample-7main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day7/LivingExample-7/main.go)

- 7.2.2 文件写入
    - ``os.OpenFile(“output.dat”,  os.O_WRONLY|os.O_CREATE, 0666)``
    - 第二个参数：文件打开模式：
        - 1. ``os.O_WRONLY``：只写
        - 2. ``os.O_CREATE``：创建文件
        - 3. ``os.O_RDONLY``：只读
        - 4. ``os.O_RDWR``：读写
        - 5. ``os.O_TRUNC ``：清空
    - 第三个参数：权限控制：
        - ``r ——> 004``
        - ``w——> 002``
        - ``x——> 001``

- 7.2.3 写文件
>[代码位置: Day7/LivingExample-8main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day7/LivingExample-8/main.go)

- 7.2.4 文件拷贝

>[代码位置: Day7/LivingExample-9main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day7/LivingExample-9/main.go)

### 7.3 命令行参数处
<a id=7-3>7.3 命令行参数处理</a>

- 7.3.1 os.Args是一个string的切片，用来存储所有的命令行参数
>[代码位置: Day7/LivingExample-10main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day7/LivingExample-10/main.go)

- 7.3.2  flag包的使用，用来解析命令行参数：
    - ``flag.BoolVar(&test, "b", false, "print on newline")``
	- ``flag.StringVar(&str, "s", "", "print on newline")``
	- ``flag.IntVar(&count, "c", 1001, "print on newline")``
>[代码位置: Day7/LivingExample-11main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day7/LivingExample-11/main.go)

### 7.4 json协议
<a id=7-4>7.4 json协议</a>

(Golang) -序列化-> (Json字符串) -网络传输-> (程序) -反序列化-> (其他语言)

- 1. 导入包：``Import “encoding/json”``
- 2. 序列化: ``json.Marshal(data interface{})``
- 3. 反序列化: ``json.UnMarshal(data []byte,  v  interface{})``
- 练习序列化结构体
>[代码位置: Day7/LivingExample-12main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day7/LivingExample-12/main.go)

- 反序列化
>[代码位置: Day7/LivingExample-13main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day7/LivingExample-13/main.go)

### 7.5 错误处理
<a id=7-5>7.5 错误处理</a>

- 7.5.1. 定义错误

    ```go
    package main

    import (
        "errors"
        "fmt"
    )

    var errNotFound error = errors.New("Not found error")

    func main() {
        fmt.Printf("error: %v", errNotFound)
    }
    ```

- 7.5.2 自定义错误

```go
type error interface { 
        Error() string 
} 
```

- 练习

```go
package main
import (
//	"fmt"
)
type PathError struct {
	Op   string
	Path string
	err string
}
func (e *PathError) Error() string {
	return e.Op + " " + e.Path + ": " + e.Err.Error()
}
func test() error {
	return &PathError{
		Op:   "op",
		Path: "path",
	}
}
func main() {
	test()
}
```
>[代码位置: Day7/LivingExample-14main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day7/LivingExample-14/main.go)

>[代码位置: Day7/LivingExample-15main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day7/LivingExample-15/main.go)


## Day-8 并发与Chan
<a id=8>Day-8 并发与Chan</a>

<a href=#8-1>Day-8.1 Goroute</a>

<a href=#8-2>Day-8.2 Channel</a>

<a href=#8-3>Day-8.3 单元测试</a>

<a href=#8-4>Day-8.4 课后作业</a>

### 8.1 Goroute
<a id=8-1>Day-8.1 Goroute</a>

- 1.进程和线程
    - a.进程是程序在操作系统的一次执行过程，系统进行资源费配合调度的一个独立单位。
    - b.线程时进程的一个执行试题，是CPU调度和分派的基本单位，他是比集成更小的能独立运行的基本单位。
    - c.一个进程何以创建和撤销多个线程；同一个进程中的多个线程之间可以并发执行。

(进程) --> (一个线程) --> (单线程程序)

(进程) --> (多个线程) --> (多线程程序)

- 2.并发和并行
    - a.多线程程序在一个核上的cpu上运行，就是并发
    - b.多线程程序在多个核上的cpu上运行，就是并行

- 3.协程和线程
    - 协程：独立的栈空间，共享的队空间，调度有用户自己控制，本质上有点类似用户级线程，这些用户级线程的调度也是自己实现的
    - 线程: 一个线程可以跑多个协程，协程是轻量级的线程

>[代码位置: Day8/LivingExample-1main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day8/LivingExample-1/main.go)

>[代码位置: Day8/LivingExample-2main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day8/LivingExample-2/main.go)
- 4.gorouteine调度模型
    - M:线程
    - P:上下文
    - G:协程

**调度模型1** 
```
M       M
|       |
P--G    P--G
|  |    |  |
G  G    G  G
   |       |
   G       G
```

**调度模型2** 
```
M1    M0   |   M1       M0
      |    |   |        |
      P--G |   P--G     G0
      |  | |   |  |
      G  G |   G  G
      |    |   |
      G    |   G
```

- 5.如何设置golang占用CPU

```go 

pacakage main

import (
    "fmt"
    "runtime"
)

func main() {
    num := runtime.NumCPU()
    runtime.GOMAXPROC(num)
    fmt.Println(num)
}

```

>[代码位置: Day8/LivingExample-3main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day8/LivingExample-3/main.go)


### 8.2 Channel
<a id=8-2>Day-8.2 Channel</a>

- 1.不同的goroutine之间是如何通信的
    - a.全局变量  [需要加锁]
    >[代码位置: Day8/LivingExample-4main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day8/LivingExample-4/main.go)

    - b.Channel
        - a.类似unix中关掉
        - b.先进先出
        - c.线程安全，多个goroutine同时访问，不需要加锁
        - d.channel是有类型的，一个整数的channel只能存放整数

        - channel声明
            ```go
            var 变量名 chan 类型
            var test chan int
            var test chan string
            var test chan map[string]int
            var test chan interface
            ```
            >[代码位置: Day8/LivingExample-5main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day8/LivingExample-5/main.go)

        - channel阻塞
            - channel的值多余channel的长度时会产生channel阻塞，此时channel的长度只能接受对应长度的值，长于channel的值将会进入阻塞状态
            >[代码位置: Day8/LivingExample-6main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day8/LivingExample-6/main.go)

        - channel同步
            - 检查管道是否被关闭
            >[代码位置: Day8/LivingExample-7main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day8/LivingExample-7/main.go)

            - channel同步 
            >[代码位置: Day8/LivingExample-8main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day8/LivingExample-8/main.go)

        - channel 通信
        >[代码位置: Day8/LivingExample-9main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day8/LivingExample-9/main.go)

        >[代码位置: Day8/LivingExample-10main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day8/LivingExample-10/main.go)

        >[代码位置: Day8/LivingExample-11main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day8/LivingExample-11/main.go) 

        >[代码位置: Day8/LivingExample-12main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day8/LivingExample-12/main.go)

        - channel的只读和只写
            - a. 只读chan的声明
                ```go
                var 变量名 <-chan int
                var readChan <-chan int 
                ```
            - b. 只写chan的声明
                ```go
                var 变量名 chan <- int
                var writeChan chan <- int
                ```

        - channel的switch
            - select


### 8.3 定时器的使用
<a id=8-3>Day-8.3 单元测试</a>

- 一次性定时器
    ```golang
     time.After(time.Second)
    ```

- 超时控制
    ```golang
    time.NewTicker(time.Second)
    ```
    >[代码位置: Day8/LivingExample-13main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day8/LivingExample-13/main.go)


- 协程中捕获异常(recover)
    >[代码位置: Day8/LivingExample-14main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day8/LivingExample-14/main.go)


### 8.4 单元测试
<a id=8-4>Day-8.3 单元测试</a>

- 1.文件名必须以`(文件名)_test.go`结尾

- 2.使用``

    >[代码位置: Day8/LivingExample-15main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day8/LivingExample-15/main.go)

### 8.5 课后作业
<a id=8-4>Day-8.4 课后作业</a>

### Day-9 网络编程

<a id=9>Day-9 TCP编程</a>

<a href=#9>Day-9 TCP编程</a>
<a href=#9>Day-9 TCP变成</a>
<a href=#9>Day-9 TCP变成</a>


<a id=9-1>Day-9 TCP编程</a>

Golang的主要涉及目标之一就是面向大规模后端服务程序，网络通讯这块是服务程序必不可少也是至关重要的一部分

- 9.1网络编程有两种：
    - 1.`TCP socket`变成，是网络编程的主流，之所以叫`TCP socket`编程，是因为底层基于`TCP\IP`协议的，比如`QQ`聊天
    - 2.`b\s`结构的`http`编程，我们的浏览器去访问服务器时，使用的就是`http`协议，而`http`底层依旧是`tcp socket`实现的

- 9.2协议
    - `TCP/IP（Transmission Control Protocol/Internet Pr otocol)`的简写,中文译名为传输控制协议/因特网互联协议，又叫网络通讯协议，这个协议是`Internet` 最基本的协议、`Internet` 国际互联网络的基础，简单地说，就是由网络层的IP 协议和传输层的TCP 协议组成的。

- 9.3`ip` 地址
    - 概述：每个`internet` 上的主机和路由器都有一个

- 9.4端口(port)-介绍
    - 我们这里所指的端口不是指物理意义上的端口，而是特指TCP/IP 协议中的端口，是逻辑意义上的端口。如果把IP 地址比作一间房子，端口就是出入这间房子的门。真正的房子只有几个门，但是一个IP 地址的端口可以有65536（即：256×256）个之多！端口是通过端口号来标记的，端口号只有整数，范围是从0 到65535（256×256-1）
    - 0 号是保留端口.
    - 1-1024 是固定端口(程序员不要使用)
        - 又叫有名端口,即被某些程序固定使用,一般程序员不使用.
            - 22: SSH 远程登录协议23: telnet 使用21: ftp 使用
            - 25: smtp 服务使用80: iis 使用7: echo 服务
    - 1025-65535 是动态端口,这些端口，程序员可以使

- 9.5端口`(port)`-使用注意
	- 1.在计算机(尤其是做服务器)要尽可能的少开端口
	- 2.一个端口只能被一个程序监听
	- 3.如果使用`netstat –an` 可以查看本机有哪些端口在监听
	- 4.可以使用`netstat –anb` 来查看监听端口的pid,在结合任务管理器关闭不安全的端口

- 9.6`tcp socket` 编程的客户端和服务器端
    - 为了授课方法，我们将tcp socket 编程，简称socket 编程.下图为Golang socket 编程中客户端和服

- 9.7`tcp socket` 编程的快速入门
    - 1.服务端的处理流程
        - a.监听端口8888
        - b.接收客户端的tcp 链接，建立客户端和服务器端的链接.
        - c.创建goroutine，处理该链接的请求(通常客户端会通过链接发送请求包)

    - 2.客户端的处理流程
        - a.建立与服务端的链接
        - b.发送请求数据[终端]，接收服务器端返回的结果数据
        - c.关闭链接

    - 3.代码的实现
        - 程序框架图示意图
        - 服务器端功能:
            - a.编写一个服务器端程序，在8888 端口监听
            - b.可以和多个客户端创建链接
            - c.链接成功后，客户端可以发送数据，服务器端接受数据，并显示在终端上.
            - d.先使用telnet 来测试，然后编写客户端程序来测试
        - 服务端的代码:
        >[代码位置: Day9/LivingExample-1/server/main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day9/LivingExample-1/server/main.go)

        - 客户端功能:
            - a.编写一个客户端端程序，能链接到服务器端的8888 端口
            - b.客户端可以发送单行数据，然后就退出
            - c.能通过终端输入数据(输入一行发送一行), 并发送给服务器端[]
            - d.在终端输入exit,表示退出程序.
            - e.代码:
            >[代码位置: Day9/LivingExample-1/client/main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day9/LivingExample-1/client/main.go)
<a id=9-2>Day-9.2 Redis 操作</a>

- [redis文档](http://redisdoc.com/)

- 9.2.1 `Redis` 安装好后，默认有16 个数据库，初始默认使用0 号库, 编号是0...15
   - 1. 添加`key-val [set]`
   - 2. 查看当前redis 的所有`key [keys *]`
   - 3. 获取key 对应的值. `[get key]`
   - 4. 切换redis 数据库`[select index]`
   - 5. 如何查看当前数据库的`key-val` 数量`[dbsize]`
   - 6. 清空当前数据库的`key-val` 和清空所有数据库的`key-val [flushdb flushall]`
- 9.2.2 `Redis` 的五大数据类型:
    - `Redis` 的五大数据类型是: `String(字符串) 、Hash (哈希)、List(列表)、Set(集合)和zset(sorted set：有序集合)`

- 9.2.3 `String`(字符串) -介绍
    - string 是redis 最基本的类型，一个key 对应一个value。
    - string 类型是二进制安全的。除普通的字符串外，也可以存放图片等数据。
    - redis 中字符串value 最大是512M

- 9.2.4 `String`(字符串) -CRUD
    - 举例说明`Redis` 的`String` 字符串的CRUD 操作.
    - `set[如果存在就相当于修改，不存在就是添加]/get/del`

- 9.2.5  `String`(字符串)-使用细节和注意事项
    - `setex(set with expire)`键秒值秒值
    ```
    setex mess01 10 hello.you
    ```
    - 同时设置多个值，同时获取多个值
        - `mset[同时设置一个或多个key-value 对]`
        - `mget[同时获取多个key-val]`

- 9.2.6 Hash` (哈希，类似golang 里的Map)`-介绍
    - 基本的介绍
        - `Redis hash` 是一个键值对集合。`var user1 map[string]string`
        - `Redis hash` 是一个`string` 类型的`field` 和`value` 的映射表，`hash` 特别适合用于存储对象。
    - 举例,存放一个`User` 信息:`(user1)`
        - `user1 name "smith" age 30 job "golang coder"`
        - 说明：
        ```
        key : user1
        name 张三和age 30 就是两对field-value
        ```

    - `Hash``（哈希，类似golang 里的Map）-CRUD`
        - 举例说明Redis 的Hash 的CRUD 的基本操作.
        - `hset/hget/hgetall/hdel`
        - 演示添加`user` 信息的案例`(name,age )`
        ```
        127.0.0.1:6379> hset user1 name "tom"
        127.0.0.1:6379> hset user1 age  19
        127.0.0.1:6379> hset user1 score 100

        127.0.0.1:6379> hget user1 name
        127.0.0.1:6379> hgetall user1

        127.0.0.1:6379> mset user2 name "buxunxian" age "18" score "100"
        127.0.0.1:6379> hmget user1 name age score
        1) "tom"
        2) "19"
        3) "100"

        127.0.0.1:6379> hlen user1
        (integer) 3

        ```
- 9.2.7 List（列表）-介绍
    - 列表是简单的字符串列表，按照插入顺序排序。你可以添加一个元素到列表的头部（左边）或者尾部（右边）。
    - List 本质是个链表, List 的元素是有序的，元素的值可以重复.
    - 举例,存放多个地址信息:
    - `city 北京 天津 上海`
    - 说明：
        - `key : city`
        - 北京 天津 上海就是三个元素
        ```
        127.0.0.1:6379> lpush city beijing shanghai tianjing
        (integer) 3
        127.0.0.1:6379> lrange city 0 -1
        1) "tianjing"
        2) "shanghai"
        3) "beijing"
        127.0.0.1:6379> lrange city 0 -1
        1) "tianjing"
        2) "shanghai"
        3) "beijing"
        4) "hulunbeier"
        127.0.0.1:6379> lpop city
        "tianjing"
        127.0.0.1:6379> lrange city 0 -1
        1) "shanghai"
        2) "beijing"
        3) "hulunbeier"
        127.0.0.1:6379> rpop city
        "hulunbeier"
        127.0.0.1:6379> lrange city 0 -1
        1) "shanghai"
        2) "beijing"
        127.0.0.1:6379> del city
        (integer) 1
        127.0.0.1:6379> lrange city 0 -1
        (empty list or set)
        ```

    - 命令
        - lpush 从list左插入
        - rpush 从list右插入
        - lrange 遍历查询list的值，区间性质查找
        - lpop  弹出一个list中的值,左第一个
        - rpop  弹出一个list中的值,右第一个
        - del   删除一个list
        
- 9.2.8 Set(集合) - 介绍
    - Redis 的Set 是string 类型的无序集合。
    - 底层是HashTable 数据结构, Set 也是存放很多字符串元素，字符串元素是无序的，而且元素的值不能重复
    - 举例,存放多个邮件列表信息:
        - email sgg@sohu.com tom@sohu.com
        - 说明：
        - key : email
        - tn@sohu.com tom@sohu.com 就是二个元素

- 9.2.9 Set(集合)- CRUD
    - 举例说明Redis 的Set 的CRUD 操作.
        ```
        sadd
        smembers[取出所有值]
        sismember[判断值是否是成员]
        srem [删除指定值]
        ```
        ```
        127.0.0.1:6379> sadd emails tom@enn.com jack@suhu.com tianrandai@qq.com
        (integer) 3
        127.0.0.1:6379> smembers emails
        1) "tom@enn.com"
        2) "tianrandai@qq.com"
        3) "jack@suhu.com"
        127.0.0.1:6379> sismember emails tianrandai@qq.com
        (integer) 1
        127.0.0.1:6379> sismember emails bucunzai@qq.com
        (integer) 0
        127.0.0.1:6379> srem emails tianrandai@qq.com
        (integer) 1
        127.0.0.1:6379> del emails
        (integer) 10
        ```
    - 演示添加多个电子邮件信息的案例

- 9.2.10 Golang操作redis
    - `go get github.com/garyburd/redigo/redis`安装支持redis的包
    - 连接redis，并写入数据
    >[代码位置: Day9/LivingExample-2/client/main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day9/LivingExample-2/main.go)

    - 连接redis，并写入hash数据
    >[代码位置: Day9/LivingExample-3/client/main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day9/LivingExample-3/main.go)

- 9.2.11 批量Set/Get 数据
    - 说明: 通过Golang 对Redis 操作，一次操作可以Set / Get 多个key-val 数据
    - 核心代码:
    ```golang
        _, err = c.Do("MSet", "name", "尚硅谷", "address", "北京昌平~")
        r, err := redis.Strings(c.Do("MGet", "name", "address"))
        for _, v := range r {
        fmt.Println(v)
    ```
    >[代码位置: Day9/LivingExample-5/client/main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day9/LivingExample-5/main.go)

- 9.2.12 给数据设置有效时间
    - 说明: 通过Golang 对Redis 操作，给key-value 设置有效时间
    - 核心代码:
    ```golang 
    //给name 数据设置有效时间为10s
    _, err = c.Do("expire", "name", 10)
    19.5.6 操作List
    说明: 通过Golang 对Redis 操作List 数据类型
    核心代码:
    _, err = c.Do("lpush", "heroList", "no1:宋江", 30, "no2:卢俊义", 28)
    r, err := redis.String(c.Do("rpop", "heroList"))
    ```
- 9.2.13 Redis 链接池
    - 说明: 通过Golang 对Redis 操作， 还可以通过Redis 链接池, 流程如下：
        - 1) 事先初始化一定数量的链接，放入到链接池
        - 2) 当Go 需要操作Redis 时，直接从Redis 链接池取出链接即可。
        - 3) 这样可以节省临时获取Redis 链接的时间，从而提高效率.
        ```
        pool = &redis.Pool{
        MaxIdle: 8, //最大空闲链接数
        MaxActive: 0, // 表示和数据库的最大链接数， 0 表示没有限制
        IdleTimeout: 100, // 最大空闲时间
        Dial: func() (redis.Conn, error) { // 初始化链接的代码， 链接哪个ip 的redis
        return redis.Dial("tcp", "localhost:6379")
        }
        c := pool.Get{} //获取一个连接
        pool.Close() //关闭链接就不能再从连接池中拿到连接了
        ```

- 9.2.14 海量用户即时通讯
    >[代码位置: Day9/TX_Projack](https://github.com/TianRandai111/buxunxian/blob/master/Day9/TX_Projack)

## Day-10 HTTP

<a id=10>Day-10 HTTP编程</a>

<a href=#10-1>Day-10.1 HTTP编程</a>

<a href=#10-2>Day-10.2 Mysql使用 </a>

<a href=#10-3>Day-10.3 课后作业</a>

<a id=10-1>Day-10.1 HTTP编程</a>

- 10.1.1 http编程
    -  a. Go原生支持http，import(“net/http”)
    -  b. Go的http服务性能和nginx比较接近
    -  c. 几行代码就可以实现一个web服务
     >[代码位置: Day10/LivingExample-1/client/main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day10/LivingExample-1/main.go)
    

- 10.1.2 http常见请求方法
    - 1）Get请求
     >[代码位置: Day10/LivingExample-2/client/main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day10/LivingExample-2/main.go)
    - 2）Post请求
    - 3）Put请求
    - 4）Delete请求
    - 5）Head请求
      >[代码位置: Day10/LivingExample-3/client/main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day10/LivingExample-3/main.go)

- 10.1.3 表单处理
     >[代码位置: Day10/LivingExample-4/client/main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day10/LivingExample-4/main.go)

- 10.1.4 模板渲染
     >[代码位置: Day10/LivingExample-5/client/main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day10/LivingExample-5/main.go)
    >[代码位置: Day10/LivingExample-6/client/main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day10/LivingExample-6/main.go)
    - 模板操作
    	- not 非 
        ```
        {{if not .condition}}  
        {{end}} 
        ```
        - and 与 
        ```
        {{if and .condition1 .condition2}}  
        {{end}} 
        ```
        - or 或 
        ```
        {{if or .condition1 .condition2}}  
        {{end}} 
        ```
        - eq 等于 
        ```
        {{if eq .var1 .var2}}  
        {{end}} 
        ```
        - ne 不等于 
        ```
        {{if ne .var1 .var2}}  
        {{end}} 
        ```
        - lt 小于 (less than) 
        ```
        {{if lt .var1 .var2}}  
        {{end}} 
        ```
        - le 小于等于 
        ```
        {{if le .var1 .var2}}  
        {{end}} 
        ```
        - gt 大于 
        ```
        {{if gt .var1 .var2}}  
        {{end}} 
        ```
        - ge 大于等于 
        ```
        {{if ge .var1 .var2}}  
        {{end}}
        ```

    - 模板变量
        `{{ .Value }}`
        ```html
        <html>
                <head>
                </head>
                <body>
                        <p>hello, old man, {{.}}</p>
                </body>
        </html>
        ```

    - with
    ```html
    <html>
            <head>
            </head>
            <body>
                    {{with .Name}}
                    //{{.}}=={{with .Name}}
                    <p>hello, old man, {{.}}</p>
                    {{end}}
            </body>
    </html>
    ```

    - range 循环
    ```html
    <html>
        <head>
        </head>
        <body>
                {{range .}}
                    {{if gt .Age 18}}
                    <p>hello, old man, {{.Name}}</p>
                    {{else}}
                    <p>hello,young man, {{.Name}}</p>
                    {{end}}
                {{end}}
        </body>
    </html>
    ```
<a id=10-2>Day-10.2 Mysql使用</a>

- 10.2.1 mysql编程
    - a. 新建测试表
        ```SQL
        CREATE TABLE person (
            user_id int primary key auto_increment,
            username varchar(260),
            sex varchar(260),
            email varchar(260)
        );
        CREATE TABLE place (
            country varchar(200),
            city varchar(200),
            telcode int
        );
        ```

        mysqld的驱动包
        `github.com/go-sql-driver/mysql`
        `github.com/jmoiron/sqlx`

    - b. 链接mysql
    ```golang
    database, err := sqlx.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
    ```
    
    - c. insert操作
    ```golang
    r, err := Db.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu001", "man", "stu01@qq.com")
    ```
    >[代码位置: Day10/LivingExample-7/client/main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day10/LivingExample-7/main.go)

    - d. Select 操作
    ```golang
    err := Db.Select(&person, "select user_id, username, sex, email from person where user_id=?", 1)
    ```
    >[代码位置: Day10/LivingExample-8/client/main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day10/LivingExample-8/main.go)

    - e. update操作
    ```golang
    _, err := Db.Exec("update person set username=? where user_id=?", "stu0001", 1)
    ```
    >[代码位置: Day10/LivingExample-9/client/main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day10/LivingExample-9/main.go)

    - f. Delete 操作
    ```golang
    _, err := Db.Exec("delete from person where user_id=?", 1)
    ```
    >[代码位置: Day10/LivingExample-10/client/main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day10/LivingExample-10/main.go)

<a id=11>Day-11 日志收集1</a>

<a href=#11-1>Day-11.1 日志收集系统设计</a>

<a href=#11-2>Day-11.2 日志客户端开发</a>


<a id=11-1>Day-11.1 日志收集系统设计</a>

- 1.项目背景
    - a. 每个系统都有日志，当系统出现问题时，需要通过日志解决问题
    - b. 当系统机器比较少时，登陆到服务器上查看即可满足
    - c. 当系统机器规模巨大，登陆到机器上查看几乎不现实

- 2.解决方案
    - a. 把机器上的日志实时收集，统一的存储到中心系统
    - b. 然后再对这些日志建立索引，通过搜索即可以找到对应日志
    - c. 通过提供界面友好的web界面，通过web即可以完成日志搜索

- 3.面临的问题
    - a. 实时日志量非常大，每天几十亿条
    - b. 日志准实时收集，延迟控制在分钟级别
    - c. 能够水平可扩展

- 4.业界方案ELK
    ![image](https://github.com/TianRandai111/Projact_Image/blob/master/golang/LivingExample-10/1%20(1).png?raw=true)

- 5.elk方案问题
    - a. 运维成本高，每增加一个日志收集，都需要手动修改配置
    - b. 监控缺失，无法准确获取logstash的状态
    - c. 无法做定制化开发以及维护

- 6.日志收集系统设计

    ![image](https://github.com/TianRandai111/Projact_Image/blob/master/golang/LivingExample-10/1%20(2).png?raw=true)

- 7.组件介绍
    - a. Log Agent，日志收集客户端，用来收集服务器上的日志
    - b. Kafka，高吞吐量的分布式队列，linkin开发，apache顶级开源项目
    - c. ES，elasticsearch，开源的搜索引擎，提供基于http restful的web接口
    - d. adoop，分布式计算框架，能够对大量数据进行分布式处理的平台

    - 7.1 kafka应用场景
        - 1.异步处理, 把非关键流程异步化，提高系统的响应时间和健壮性
        
        ![image](https://github.com/TianRandai111/Projact_Image/blob/master/golang/LivingExample-10/1%20(3).png?raw=true)
        
        ![image](https://github.com/TianRandai111/Projact_Image/blob/master/golang/LivingExample-10/1%20(4).png?raw=true)
        - 2.应用解耦,通过消息队列
        
        ![image](https://github.com/TianRandai111/Projact_Image/blob/master/golang/LivingExample-10/1%20(5).png?raw=true)
        
        ![image](https://github.com/TianRandai111/Projact_Image/blob/master/golang/LivingExample-10/1%20(6).png?raw=true)
        - 3.流量削峰
        
        ![image](https://github.com/TianRandai111/Projact_Image/blob/master/golang/LivingExample-10/1%20(7).png?raw=true)
    - 7.2 zookeeper应用场景
        - 1.服务注册&服务发现
        
        ![image](https://github.com/TianRandai111/Projact_Image/blob/master/golang/LivingExample-10/1%20(8).png?raw=true)
        - 2.配置中心
        
        ![image](https://github.com/TianRandai111/Projact_Image/blob/master/golang/LivingExample-10/1%20(9).png?raw=true)  
        - 3.分布式锁
            - a. Zookeeper是强一致的
            - b. 多个客户端同时在Zookeeper上创建相同znode，只有一个创建成功

- 8.安装kafka
    - a.  安装JDK，从oracle下载最新的SDK安装
    - b.  安装zookeeper3.3.6，下载地址：http://apache.fayea.com/zookeeper/
        - 1）mv conf/zoo_sample.cfg conf/zoo.cfg
        - 2）编辑 conf/zoo.cfg，修改dataDir=D:\zookeeper-3.3.6\data\
        - 3）运行bin/zkServer.cmd
    - c.  安装kafka
        - 1）打开链接：http://kafka.apache.org/downloads.html下载kafka2.1.2
        - 2）打开config目录下的server.properties， 修改log.dirs为D:\kafka_logs,修改advertised.host.name=服务器ip
        - 3）启动kafka ./bin/windows/kafka-server-start.bat ./config/server.preperties

- 9.log agent设计

    ![image](https://github.com/TianRandai111/Projact_Image/blob/master/golang/LivingExample-10/1%20(10).png?raw=true) 

- 10.log agent流程

    ![image](https://github.com/TianRandai111/Projact_Image/blob/master/golang/LivingExample-10/1%20(11).png?raw=true) 

- 11.kafka示例代码

对接kafka包`go get github.com/Shopify/sarama`

>[代码位置: Day11/LivingExample-1/main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day11/LivingExample-1/main.go)

tail包实现tail -f功能`go get github.com/hpcloud/tail`

>[代码位置: Day11/LivingExample-2/main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day11/LivingExample-2/main.go)
    
配置文件包
`go get github.com/astaxie/beego/config`

>[代码位置: Day11/LivingExample-3/main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day11/LivingExample-3/main.go)

```go
String(key string) string
Int(key string) (int, error)
Int64(key string) (int64, error)
Bool(key string) (bool, error)
Float(key string) (float64, error)
```

![image](https://github.com/TianRandai111/Projact_Image/blob/master/golang/LivingExample-11/1.jpg?raw=true) 

![image](https://github.com/TianRandai111/Projact_Image/blob/master/golang/LivingExample-11/2.png?raw=true) 

```go
iniconf, err := config.NewConfig("ini", "testini.conf")
if err != nil {
    t.Fatal(err)
}
```


日志库
`go get github.com/astaxie/beego/logs`

<a id=11-2>Day-11.2 日志客户端开发</a>

- 1. 配置log组件

```go
     config := make(map[string]interface{})
	config["filename"] = "./logs/logcollect.log"
	config["level"] = logs.LevelDebug

	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("marshal failed, err:", err)
		return
	}

```




配置etcd
```golang
gopm get -g "github.com/coreos/etcd/clientv3"
```
     

配置 elasticsearch

```golang
go get "gopkg.in/olivere/elastic.v2"
```


Mysql事物
```golang
_ "github.com/go-sql-driver/mysql" //数据库驱动
"github.com/jmoiron/sqlx" 
2)  Db.Begin()开始事务

3)  Db.Submit()提交事务

4)  Db.Rollback() 回滚事务

```


huanghua-airport_energy-integration 


beego web开发
- 1) 规划好url
- 2) 添加路由
- 3) 开发controller，集成beego.Controller
- 4) 测试