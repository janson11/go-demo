# go-demo
Go语言学习与研究


## 第一部分：Go 语言快速入门

Go 语言的代码是非常简洁、完整的核心程序，只需要 package、import、func main 这些核心概念就可以实现。


var 变量名 类型 = 表达式

Go 语言中定义的变量必须使用，否则无法编译通过，这也是 Go 语言比较好的特性，防止定义了变量不使用，导致浪费内存的情况；另一方面，在运行程序的时候可以查看变量 i 的结果。

因为 Go 语言具有类型推导功能，所以也可以不去刻意地指定变量的类型，而是让 Go 语言自己推导，比如变量 i 也可以用如下的方式声明


基础类型
常用的有：整型、浮点数、布尔型和字符串

整型
在 Go 语言中，整型分为：

有符号整型：如 int、int8、int16、int32 和 int64。

无符号整型：如 uint、uint8、uint16、uint32 和 uint64。
它们的差别在于，有符号整型表示的数值可以为负数、零和正数，而无符号整型只能为零和正数。

浮点数
浮点数就代表现实中的小数。Go 语言提供了两种精度的浮点数，分别是 float32 和 float64。项目中最常用的是 float64，因为它的精度高，浮点计算的结果相比 float32 误差会更小。

布尔型
一个布尔型的值只有两种：true 和 false，它们代表现实中的“是”和“否”。它们的值会经常被用于一些判断中，比如 if 语句（以后的课时会详细介绍）等。Go 语言中的布尔型使用关键字 bool 定义。

零值
零值其实就是一个变量的默认值，在 Go 语言中，如果我们声明了一个变量，但是没有对其进行初始化，那么 Go 语言会自动初始化其值为对应类型的零值。比如数字类的零值是 0，布尔型的零值是 false，字符串的零值是 "" 空字符串等。

变量简短声明
有没有发现，上面我们演示的示例都有一个 var 关键字，但是这样写代码很烦琐。借助类型推导，Go 语言提供了变量的简短声明 :=，结构如下：

复制代码
变量名:=表达式

指针
在 Go 语言中，指针对应的是变量在内存中的存储位置，也就说指针的值就是变量的内存地址。通过 & 可以获取一个变量的地址，也就是指针。

赋值
在讲变量的时候，我说过变量是可以修改的，那么怎么修改呢？这就是赋值语句要做的事情。最常用也是最简单的赋值语句就是 =

常量
一门编程语言，有变量就有常量，Go 语言也不例外。在程序中，常量的值是指在编译期就确定好的，一旦确定好之后就不能被修改，这样就可以防止在运行期被恶意篡改。

常量的定义
常量的定义和变量类似，只不过它的关键字是 const。

iota
iota 是一个常量生成器，它可以用来初始化相似规则的常量，避免重复的初始化

if 条件语句
if 语句是条件语句，它根据布尔值的表达式来决定选择哪个分支执行：如果表达式的值为 true，则 if 分支被执行；如果表达式的值为 false，则 else 分支被执行

switch 选择语句
if 条件语句比较适合分支较少的情况，如果有很多分支的话，选择 switch 会更方便，在 Go 语言中，switch 的 case 从上到下逐一进行判断，一旦满足条件，
立即执行对应的分支并返回，其余分支不再做判断。也就是说 Go 语言的 switch 在默认情况下，case 最后自带 break。
switch 后的表达式也没有太多限制，是一个合法的表达式即可，也不用一定要求是常量或者整数。

for 循环语句
当需要计算 1 到 100 的数字之和时，如果用代码将一个个数字加起来，会非常复杂，可读性也不好，这就体现出循环语句的存在价值了


Array（数组）
数组存放的是固定长度、相同类型的数据，而且这些存放的元素是连续的。所存放的数据类型没有限制，可以是整型、字符串甚至自定义。

数组循环
使用传统的 for 循环遍历数组，输出对应的索引和对应的值，这种方式很烦琐，一般不使用，大部分情况下，我们使用的是 for range 这种 Go 语言的新型循环，如下面的代码所示：

相比传统的 for 循环，for range 要更简洁，如果返回的值用不到，可以使用 _ 下划线丢弃


Slice（切片）
切片和数组类似，可以把它理解为动态数组。切片是基于数组实现的，它的底层就是一个数组。对数组任意分隔，就可以得到一个切片。现在我们通过一个例子来更好地理解它，同样还是基于上述例子的 array。

切片声明
除了可以从一个数组得到切片外，还可以声明切片，比较简单的是使用 make 函数。

Map（映射）
在 Go 语言中，map 是一个无序的 K-V 键值对集合，结构为 map[K]V。其中 K 对应 Key，V 对应 Value。map 中所有的 Key 必须具有相同的类型，Value 也同样，但 Key 和 Value 的类型可以不同。此外，Key 的类型必须支持 == 比较运算符，这样才可以判断它是否存在，并保证 Key 的唯一。



函数
函数初探
在前面的四节课中，你已经见到了 Go 语言中一个非常重要的函数：main 函数，它是一个 Go 语言程序的入口函数，我在演示代码示例的时候，会一遍遍地使用它。

下面的示例就是一个 main 函数：

复制代码
func main() {
}
它由以下几部分构成：

任何一个函数的定义，都有一个 func 关键字，用于声明一个函数，就像使用 var 关键字声明一个变量一样；

然后紧跟的 main 是函数的名字，命名符合 Go 语言的规范即可，比如不能以数字开头；

main 函数名字后面的一对括号 () 是不能省略的，括号里可以定义函数使用的参数，这里的 main 函数没有参数，所以是空括号 () ；

括号 () 后还可以有函数的返回值，因为 main 函数没有返回值，所以这里没有定义；

最后就是大括号 {} 函数体了，你可以在函数体里书写代码，写该函数自己的业务逻辑。

多值返回
同有的编程语言不一样，Go 语言的函数可以返回多个值，也就是多值返回。在 Go 语言的标准库中，你可以看到很多这样的函数：第一个值返回函数的结果，第二个值返回函数出错的信息，这种就是多值返回的经典应用。

如果有的函数的返回值不需要，可以使用下划线 _ 丢弃它


可变参数
可变参数，就是函数的参数数量是可变的，比如最常见的 fmt.Println 函数。
只要在参数类型前加三个点 … 


包级函数
不管是自定义的函数 sum、sum1，还是我们使用到的函数 Println，都会从属于一个包，也就是 package。sum 函数属于 main 包，Println 函数属于 fmt 包。

同一个包中的函数哪怕是私有的（函数名称首字母小写）也可以被调用。如果不同包的函数要被调用，那么函数的作用域必须是公有的，也就是函数名称的首字母要大写，比如 Println。

小提示：Go 语言没有用 public、private 这样的修饰符来修饰函数是公有还是私有，而是通过函数名称的大小写来代表，这样省略了烦琐的修饰符，更简洁。

方法
不同于函数的方法
在 Go 语言中，方法和函数是两个概念，但又非常相似，不同点在于方法必须要有一个接收者，这个接收者是一个类型，这样方法就和这个类型绑定在一起，称为这个类型的方法。

结构体
结构体定义
结构体是一种聚合类型，里面可以包含任意类型的值，这些值就是我们定义的结构体的成员，也称为字段。在 Go 语言中，要自定义一个结构体，需要使用 type+struct 关键字组合

结构体的成员字段并不是必需的，也可以一个字段都没有，这种结构体成为空结构体。

根据以上信息，我们可以总结出结构体定义的表达式，如下面的代码所示：

复制代码
type structName struct{
fieldName typeName
....
....
}
其中：

type 和 struct 是 Go 语言的关键字，二者组合就代表要定义一个新的结构体类型。

structName 是结构体类型的名字。

fieldName 是结构体的字段名，而 typeName 是对应的字段类型。

字段可以是零个、一个或者多个。

小提示：结构体也是一种类型，所以以后自定义的结构体，我会称为某结构体或某类型，两者是一个意思。比如 person 结构体和 person 类型其实是一个意思。

结构体声明使用
结构体类型和普通的字符串、整型一样，也可以使用同样的方式声明和初始化。
在下面的例子中，我声明了一个 person 类型的变量 p，因为没有对变量 p 初始化，所以默认会使用结构体里字段的零值。

复制代码
var p person
当然在声明一个结构体变量的时候，也可以通过结构体字面量的方式初始化，如下面的代码所示：

复制代码
p:=person{"飞雪无情",30}

接口
接口的定义
接口是和调用方的一种约定，它是一个高度抽象的类型，不用和具体的实现细节绑定在一起。接口要做的是定义好约定，告诉调用方自己可以做什么，但不用知道它的内部实现，这和我们见到的具体的类型如 int、map、slice 等不一样。

接口的定义和结构体稍微有些差别，虽然都以 type 关键字开始，但接口的关键字是 interface，表示自定义的类型是一个接口

当值类型作为接收者时，person 类型和*person类型都实现了该接口。

当指针类型作为接收者时，只有*person类型实现了该接口。

工厂函数
工厂函数一般用于创建自定义的结构体，便于使用者调用，我们还是以 person 类型为例，用如下代码进行定义：


协程（Goroutine）
Go 语言中没有线程的概念，只有协程，也称为 goroutine。相比线程来说，协程更加轻量，一个程序可以随意启动成千上万个 goroutine。
go function()

Channel
那么如果启动了多个 goroutine，它们之间该如何通信呢？这就是 Go 语言提供的 channel（通道）要解决的问题。
声明一个 channel
在 Go 语言中，声明一个 channel 非常简单，使用内置的 make 函数即可，如下所示：

ch:=make(chan string)
其中 chan 是一个关键字，表示是 channel 类型。后面的 string 表示 channel 里的数据是 string 类型。通过 channel 的声明也可以看到，chan 是一个集合类型。

定义好 chan 后就可以使用了，一个 chan 的操作只有两种：发送和接收。

接收：获取 chan 中的值，操作符为 <- chan。

发送：向 chan 发送值，把值放在 chan 中，操作符为 chan <-。

小技巧：这里注意发送和接收的操作符，都是 <- ，只不过位置不同。接收的 <- 操作符在 chan 的左侧，发送的 <- 操作符在 chan 的右侧。

有缓冲 channel
有缓冲 channel 类似一个可阻塞的队列，内部的元素先进先出。通过 make 函数的第二个参数可以指定 channel 容量的大小，进而创建一个有缓冲 channel，如下面的代码所示：

cacheCh:=make(chan int,5)

数据流动、传递的场景中要优先使用channel，它是并发安全的，性能也不错。

channel 为什么是并发安全的呢？是因为 channel 内部使用了互斥锁来保证并发的安全

小技巧：使用 go build、go run、go test 这些 Go 语言工具链提供的命令时，添加 -race 标识可以帮你检查 Go 语言代码是否存在资源竞争。