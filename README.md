<a href=#1>Day-1 golang语言基础</a>

<a href=#2>Day-2 数据类型与操作符</a>

<a href=#3>Day-3 字符串处理包与函数</a>

<a href=#4>Day-1 数组、切片、map、package介绍</a>


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
- <a href=#3-3>3. 指针类型</a>
- <a href=#3-6>3. 流程控制</a>
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
    - 2.1 time包
    - 2.2 time.{Time类型}，用来表示时间
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

- <a id=3-3>3. 指针类型</a>
    - 1. 普通类型，变量存的就是值，也叫值类型
    - 2. 获取变量的地址，用&，比如： var a int, 获取a的地址：&a
    - 3. 指针类型，变量存的是一个地址，这个地址存的才是值
    - 4. 获取指针类型所指向的值，使用：*，比如：var *p int, 使用*p获取p指向的值
    ```go
    var a int = 5  //a -- >5 变量a指向5
    var p *int = &a //0xefefef -> 5 指针p指向a的地址，a地址对应的值为5
    ```
    >练习8：写一个程序，获取一个变量的地址，并打印到终端。[代码位置: Day3/LivingExample-11/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day3/LivingExample-11/main.go)


    >练习9：写一个函数，传入一个int类型的指针，并在函数中修改所指向的值。在main函数中调用这个函数，并把修改前后的值打印到终端，观察结果[代码位置: Day3/LivingExample-12/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day3/LivingExample-12/main.go)

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
    >练习10：写一个程序，从终端读取输入，并转成整数，如果转成整数出错，则输出 “can not convert to int”，并返回。否则输出该整数。[代码位置: Day3/LivingExample-13/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day3/LivingExample-13/main.go)

    - 2.switch case语句,如果使用``fallthrough``会继续向下执行[代码位置: Day3/LivingExample-14/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day3/LivingExample-14/main.go)
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
            fmt.Println(“i > 0 and i < 10”)
        condition2:
        fmt.Println(“i > 10 and i < 20”)
        default:
            fmt.Println(“def”)
        }

        //写法3
        switch i := 0; {
        condition1:
            fmt.Println(“i > 0 and i < 10”)
        condition2:
        fmt.Println(“i > 10 and i < 20”)
        default:
            fmt.Println(“def”)
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

        >break continue的使用方法[代码位置: Day3/LivingExample-17 /main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day3/LivingExample-17/main.go)

- <a id=3-4>4. 函数详解</a>
    - 1.语法声明：func 函数名(参数列表) [(返回值列表)] {}
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

    >注意2：map、slice、chan、指针、interface默认以引用的方式传递。

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
                        defer fmt.Printf(“%d “, i)
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
    
- <a id=3-5>5. 课后作业</a>
    - 1.编写程序，在终端输出九九乘法表。[Day3/Assignment/one/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day3/Assignment/one/main.go)
    - 2.一个数如果恰好等于它的因子之和，这个数就称为“完数”。例如6=1＋2＋3.编程找出1000以内的所有完数。[Day3/Assignment/two/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day3/Assignment/two/main.go)
    - 3.输入一个字符串，判断其是否为回文。回文字符串是指从左到右读和从右到左读完全相同的字符串。[Day3/Assignment/three/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day3/Assignment/three/main.go)
    - 4.输入一行字符，分别统计出其中英文字母、空格、数字和其它字符的个数。[Day3/Assignment/four/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day3/Assignment/four/main.go)
    - 5.计算两个大数相加的和，这两个大数会超过int64的表示范围.
    - [自己的版本 Day3/Assignment/five/my-version/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day3/Assignment/five/my-version/main.go)
    - [别人的版本Day3/Assignment/five/teacher-version/main.go](https://github.com/TianRandai111/buxunxian/tree/master/Day3/Assignment/five/teacher-version/main.go)

  

<a id=4>Day-1笔记</a>

<a href=#4-1>1.内置函数、递归函数、闭包</a>
<a href=#4-2>2.数组与切片</a>
<a href=#4-3>3.map数组结构</a>
<a href=#4-4>4.package介绍</a>
<a href=#4-5>5.课后作业</a>
<a href=#4-6>Day-1 golang语言基础</a>
<a href=#4-7>Day-1 golang语言基础</a>
<a href=#4-8>Day-1 golang语言基础</a>

<a id=4-1>1.内置函数、递归函数、闭包</a>

- 1.1. 内置函数
    - 1.1.1. close：主要用来关闭channel
    - 1.1.2. len：用来求长度，比如string、array、slice、map、channel
    - 1.1.3. new：用来分配内存，主要用来分配值类型，比如int、struct。返回的是指针
    - 1.1.4. make：用来分配内存，主要用来分配引用类型，比如chan、map、slice
    - 1.1.5. append：用来追加元素到数组、slice中
    - 1.1.6. panic和recover：用来做错误处理
    >new和make的区别
    ![image](https://github.com/TianRandai111/buxunxian/tree/master/Day4/Image/图片1.png)
<a id=4-2>2.数组与切片</a>
<a id=4-3>3.map数组结构</a>
<a id=4-4>4.package介绍</a>
<a id=4-5>5.课后作业</a>




<a id=4-6>Day-1笔记</a>
<a id=4-7>Day-1笔记</a>
<a id=4-8>Day-1笔记</a>
