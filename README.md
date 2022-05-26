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