<a href=#1>Day-1 golang语言基础</a>

<a href=#2>Day-2 数据类型与操作符</a>

<a href=#3>Day-3 Outline</a>

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

**重点GOPATH：环境变量 GOPATH 的值可以是一个目录的路径，也可以包含多个目录路径，每个目录都代表 Go 语言的一个工作区（workspace）。这些工作区用于放置 Go 语言的源码文件（source file），以及安装（install）后的归档文件（archive file，也就是以“.a”为扩展名的文件）和可执行文件（executable file）。**

### 高并发测试(一个测试)
- 1.[代码位置: ](https://github.com/TianRandai111/buxunxian/blob/master/Day1/LivingExample-1/LivingExample-1/testgo.go)

### 第一个golang程序
- 1.hello world[代码位置: ](https://github.com/TianRandai111/buxunxian/blob/master/Day1/LivingExample-2/hello.go)
- 2.channel
    - a.管道.类似unix/linux中的pipe
    - b.多个goroute之间通过channel进行通信
    - c.支持任何类型
    > 如果管道存放的值数量，多于定义的长度会造成死锁，导致程序无法运行

    [代码位置: Day1/LivingExample-3/pipe.goo](https://github.com/TianRandai111/buxunxian/blob/master/Day1//LivingExample-3/pipe.go)

    ```go
        pipe := make (chan int,3) //初始化一个chan，类型为整形，长度为3。
        pipe <- 1 //将1放到pipe里
        pipe <- 2 //将2放到pipe里
    ```
![image](le-3)
- 3.返回值
    - [代码位置: Day1/LivingExample-4/return.go](https://github.com/TianRandai111/buxunxian/blob/master/Day1/LivingExample-4/return.go)

### 包的概念
- 1.和ptyhon一样，把相同的代码放到一个目录，称之为包
- 2.包可以被其他包引用
- 3.main包是用来生成执行文件，每个程序只有一个main包
- 4.包的主要用途是提高代码的可复用性


### 包的实战
- 1.
>[main代码位置: Day1/LivingExample-5/main/main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day1/LivingExample-5/main/main.go)

>[add模块代码位置: Day1/LivingExample-5/calc/add.go](https://github.com/TianRandai111/buxunxian/blob/master/Day1/LivingExample-5/calc/add.go)

>[sub模块代码位置: Day1/LivingExample-5/calc/add.go](https://github.com/TianRandai111/buxunxian/blob/master/Day1/LivingExample-5/calc/add.go)

>编译二进制文件``go build -o bin/calc_test.exe  ./main/main.go``

- 2.多进程的goroute运算测试，

>[main代码位置: Day1/LivingExample-6/main/main.go](https://github.com/TianRandai111/buxunxian/blob/master/Day1/LivingExample-6/main/main.go)

>[main代码位置: Day1/LivingExample-6/goroute/add.go](https://github.com/TianRandai111/buxunxian/blob/master/Day1/LivingExample-6/goroute/add.go)

作业：使用fmt,打印浮点数，二进制，十进制，十六进制

## <a id=2>Day-2笔记</a>

### 第一部分：基本数据类型和操作符

- <a href=#2-1>文件名&关键字&标识符</a>
- <a href=#2-2>Go程序基本结构</a>
- <a href=#2-6>函数的声明和注释</a>
- <a href=#2-3>常量和变量</a>
- <a href=#2-7>值类型和引用类型</a>
- <a href=#2-4>数据类型和操作符</a>
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

    >[main代码位置: Day2/LivingExample-1/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day2/LivingExample-1/main.go)
    
    >练习2.一个程序包包含两个包add和main,其中add包中有两个变量，Name和age。请问main包中是如何访问Name和Age的(大写的为公有，小写为私有)

    >[main代码位置: Day2/LivingExample-2/main/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day2/LivingExample-2/main/main.go)

    >[add代码位置: Day2/LivingExample-2/add/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day2/LivingExample-2/add/main.go)

    >练习3.包别名的应用,开发一个程序，使用包别名来访问包中的函数

    >[main代码位置: Day2/LivingExample-3/main/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day2/LivingExample-3/main/main.go)

    >[add代码位置: Day2/LivingExample-3/add/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day2/LivingExample-3/add/main.go)

    >练习4.每个源文件都可以包含一个init函数，这个init函数自动被go运行框架调用。开发一个演示这个功能

    >[代码位置: Day2/LivingExample-4/add/add.go](https://github.com/TianRandai111/buxunxian/tree/master/Day2/LivingExample-4/add/add.go)

- <a id=2-6>函数的声明和注释</a>
    - 1.函数声明：func函数名字(参数列表)(返回值列表){}
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
    - 2.注释，两种注释，单行注释：//多行注释：/*...*/
        ```go
        //add计算两个证书的和，并返回结果
        func add(a int, b int) int {}
        ```
        ```go
        /*add计算两个证书的和，并返回结果*/
        func add(a int, b int) int {}
        ```
- <a id=2-3>常量和变量</a>
    - 1.常量使用const修饰，代表永远只读，不能修改。
    - 2.const只能修饰boolean,number(int相关类型、浮点类型、complex)和string
    - 3.语法``const identifier [type] = value``,其中type可以省略。
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
    >练习：定义两个常量Man=1和Female=2,获取当年时间的秒数，如果能被Female整除，则在终端打印female，否则打印man
    
    >[代码位置: Day2/LivingExample-5/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day2/LivingExample-5/main.go)

    - 4. 变量
        - 4.1 语法:var identifier type
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
        >练习写一个程序获取当前运行的系统名称和PATH环境变量的值，并打印在终端

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
- <a id=2-7>值类型和引用类型</a>
    - 1.值类型：变量直接存储值，内存通常在栈中分配
        - a.值类型：基本数据类型int、float、bool、string已经数组和struct.
    - 2.应用类型：变量存储的是一个地址，这个地址存储最终的值。内存通常在对战分配，通过GC回收
        - a.引用类型：指针、slice、map、chan等都是引用类型。
    > go语言交换值独有的方式[代码位置: Day2/LivingExample-9/main/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day2/LivingExample-9/main/main.go)

- <a id=2-4>数据类型和操作符</a>
- 1.bool类型，只能存true和false
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
    > 请问!a、!b、a&&b、a||b的值分别是多少
- 3.数字类型，主要有
    - 3.1 int 整形(由正负)
        - ``int``、``int8``、``int16``、``int32``、``int68``
    - 3.2 uint 无符号整形(无负数)
        - ``uint8``、``uint16``、``uint32``、``uint64``
    - 3.3 float浮点数
        - ``float32``、``float64``
- 4. 类型转换
    ```go
    var a int8 = 100
    //将int8类型的a转换成int16
    var b int16 = int16(a)
    ```
- <a id=2-5>字符串类型</a>
    - 字符串拼装操作[代码位置: Day2/LivingExample-12/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day2/LivingExample-12/main.go)

    - 字符串反转：[代码位置: Day2/LivingExample-13/main/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day2/LivingExample-13/main/main.go)

- 作业：
    - 判断101和200之间有多少个素数，并输出所有的素数
    >[代码位置: Day2/Assignment/PrimeNumber/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day2/Assignment/PrimeNumber/main.go)

    - 打印出100-999中所有的“水仙花数，所谓的水仙花数是一个三位数，其中苏子的立方和等于该数恩深，例如153是一个水仙花数，因为153=1的三次方+5的三次方+3的三次方
    >[代码位置: Day2/Assignment/PPDI/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day2/Assignment/PPDI/main.go)

    >[代码位置: (使用ASCII码来表示)Day2/Assignment/ASCII/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day2/Assignment/ASCII/main.go)

    - 对于一个数n,求n的阶乘之和：，既：1！+2！+3！+...n
    >[代码位置: Day2/Assignment/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day2/Assignment/main.go)


    - 变量赋值的区别
    >[代码位置: Day2/Assignment/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day2/Assignment/main.go)

## <a id=3>Outline</a>

- <a href=#3-1>1. strings和strconv的使用</a>
- <a href=#3-2>2. Go中的时间和日期类型</a>
- <a href=#3-3>3. 流程控制</a>
- <a href=#3-4>4. 函数详解</a>
- <a href=#3-5>5. 课后作业</a>



- <a id=3-1>1. strings和strconv的使用</a> 
    - 1.1. strings.HasPrefix(s string, prefix string) bool：判断字符串s是否以prefix开头
    >练习1：判断一个url是否以http://开头，如果不是，则加上http://。

    >[代码位置: Day3/LivingExample-1/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day3/LivingExample-1/main.go)

    - 1.2. strings.HasSuffix(s string, suffix string) bool：判断字符串s是否以suffix结尾。
    >练习2：判断一个路径是否以“/”结尾，如果不是，则加上/。

    >[代码位置: Day3/LivingExample-2/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day3/LivingExample-2/main.go)

    - 1.3. strings.Index(s string, str string) int：判断str在s中首次出现的位置，如果没有出现，则返回-1
    >[代码位置: Day3/LivingExample-3/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day3/LivingExample-3/main.go)

    - 1.4. strings.LastIndex(s string, str string) int：判断str在s中最后出现的位置，如果没有出现，则返回-1
    >[代码位置: Day3/LivingExample-4/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day3/LivingExample-4/main.go)

    - 1.5. strings.Replace(str string, old string, new string, n int)：字符串替换
    - 1.6. strings.Count(str string, substr string)int：字符串计数
    - 1.7. strings.Repeat(str string, count int)string：重复count次str
    - 1.8. strings.ToLower(str string)string：转为小写
    - 1.9. strings.ToUpper(str string)string：转为大写
    >练习4：写一个函数分别演示Replace、Count、Repeat、ToLower、ToUpper的用法

    >[代码位置: Day3/LivingExample-5/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day3/LivingExample-5/main.go)

    - 1.10. strings.TrimSpace(str string)：去掉字符串首尾空白字符
        - 1.10.1. strings.Trim(str string, cut string)：去掉字符串首尾cut字符
        - 1.10.2. strings.TrimLeft(str string, cut string)：去掉字符串首cut字符
        - 1.10.3. strings.TrimRight(str string, cut string)：去掉字符串首cut字符
    - 1.11. strings.Field(str string)：返回str空格分隔的所有子串的slice
        - 1.11.1. strings.Split(str string, split string)：返回str split分隔的所有子串的slice
    - 1.12. strings.Join(s1 []string, sep string)：用sep把s1中的所有元素链接起来
    >练习5：写一个函数分别演示TrimSpace、Trim、TrimLeft、TrimRight、Field、Split、以及Join的用法

    >[代码位置: Day3/LivingExample-6/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day3/LivingExample-6/main.go)

    - 1.13. strconv.Itoa(i int)：把一个整数i转成字符串
    - 1.14. strconv.Atoi(str string)(int, error)：把一个字符串转成整数
    >练习6：写一个函数分别演示Itoa、Atoi的用法
    
    >[代码位置: Day3/LivingExample-7/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day3/LivingExample-7/main.go)

- <a id=3-2>2. Go中的时间和日期类型</a>
- <a id=3-3>3. 流程控制</a>
- <a id=3-4>4. 函数详解</a>
- <a id=3-5>5. 课后作业</a>


